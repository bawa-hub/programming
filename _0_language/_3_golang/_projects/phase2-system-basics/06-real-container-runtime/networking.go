package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

// NetworkManager manages container networking
type NetworkManager struct {
	bridgeName    string
	bridgeIP      string
	containerID   string
	networkMode   string
}

// ContainerNetwork represents a container's network configuration
type ContainerNetwork struct {
	ContainerID   string            `json:"container_id"`
	NetworkMode   string            `json:"network_mode"`
	BridgeName    string            `json:"bridge_name"`
	ContainerIP   string            `json:"container_ip"`
	GatewayIP     string            `json:"gateway_ip"`
	PortMappings  []PortMapping     `json:"port_mappings"`
	DNS           []string          `json:"dns"`
	Hostname      string            `json:"hostname"`
	VethHost      string            `json:"veth_host"`
	VethContainer string            `json:"veth_container"`
}

// PortMapping represents port mapping configuration
type PortMapping struct {
	HostPort      string `json:"host_port"`
	ContainerPort string `json:"container_port"`
	Protocol      string `json:"protocol"`
	HostIP        string `json:"host_ip"`
}

// NewNetworkManager creates a new network manager
func NewNetworkManager(containerID string) *NetworkManager {
	return &NetworkManager{
		bridgeName:  "br-" + containerID,
		bridgeIP:    "172.17.0.1/16",
		containerID: containerID,
		networkMode: "bridge",
	}
}

// CreateContainerNetwork creates a network for the container
func (nm *NetworkManager) CreateContainerNetwork() (*ContainerNetwork, error) {
	fmt.Printf("🌐 Creating network for container: %s\n", nm.containerID)
	
	network := &ContainerNetwork{
		ContainerID:   nm.containerID,
		NetworkMode:   nm.networkMode,
		BridgeName:    nm.bridgeName,
		ContainerIP:   nm.generateContainerIP(),
		GatewayIP:     "172.17.0.1",
		PortMappings:  make([]PortMapping, 0),
		DNS:           []string{"8.8.8.8", "8.8.4.4"},
		Hostname:      nm.containerID,
		VethHost:      "veth-" + nm.containerID + "-host",
		VethContainer: "veth-" + nm.containerID + "-container",
	}
	
	// Create bridge network
	if err := nm.createBridge(network); err != nil {
		return nil, fmt.Errorf("failed to create bridge: %w", err)
	}
	
	// Create veth pair
	if err := nm.createVethPair(network); err != nil {
		return nil, fmt.Errorf("failed to create veth pair: %w", err)
	}
	
	// Configure container network namespace
	if err := nm.configureContainerNamespace(network); err != nil {
		return nil, fmt.Errorf("failed to configure container namespace: %w", err)
	}
	
	// Set up iptables rules
	if err := nm.setupIptablesRules(network); err != nil {
		return nil, fmt.Errorf("failed to setup iptables rules: %w", err)
	}
	
	fmt.Printf("✅ Container network created: %s\n", network.ContainerIP)
	return network, nil
}

// createBridge creates a bridge network
func (nm *NetworkManager) createBridge(network *ContainerNetwork) error {
	fmt.Println("  Creating bridge network...")
	
	// Create bridge using ip command
	cmd := exec.Command("ip", "link", "add", "name", network.BridgeName, "type", "bridge")
	if err := cmd.Run(); err != nil {
		// Bridge might already exist, continue
		fmt.Printf("    Bridge %s might already exist\n", network.BridgeName)
	}
	
	// Set bridge IP
	cmd = exec.Command("ip", "addr", "add", network.GatewayIP, "dev", network.BridgeName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set bridge IP: %w", err)
	}
	
	// Bring bridge up
	cmd = exec.Command("ip", "link", "set", "dev", network.BridgeName, "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to bring bridge up: %w", err)
	}
	
	fmt.Printf("    Bridge %s created with IP %s\n", network.BridgeName, network.GatewayIP)
	return nil
}

// createVethPair creates a veth pair for container networking
func (nm *NetworkManager) createVethPair(network *ContainerNetwork) error {
	fmt.Println("  Creating veth pair...")
	
	// Create veth pair
	cmd := exec.Command("ip", "link", "add", network.VethHost, "type", "veth", "peer", "name", network.VethContainer)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create veth pair: %w", err)
	}
	
	// Add host end to bridge
	cmd = exec.Command("ip", "link", "set", network.VethHost, "master", network.BridgeName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add veth to bridge: %w", err)
	}
	
	// Bring host end up
	cmd = exec.Command("ip", "link", "set", "dev", network.VethHost, "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to bring host veth up: %w", err)
	}
	
	fmt.Printf("    Veth pair created: %s <-> %s\n", network.VethHost, network.VethContainer)
	return nil
}

// configureContainerNamespace configures the container network namespace
func (nm *NetworkManager) configureContainerNamespace(network *ContainerNetwork) error {
	fmt.Println("  Configuring container network namespace...")
	
	// Create network namespace
	nsPath := fmt.Sprintf("/var/run/netns/%s", nm.containerID)
	if err := os.MkdirAll(filepath.Dir(nsPath), 0755); err != nil {
		return fmt.Errorf("failed to create netns directory: %w", err)
	}
	
	// Create namespace file
	file, err := os.Create(nsPath)
	if err != nil {
		return fmt.Errorf("failed to create namespace file: %w", err)
	}
	file.Close()
	
	// Move container veth to namespace
	cmd := exec.Command("ip", "link", "set", network.VethContainer, "netns", nm.containerID)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to move veth to namespace: %w", err)
	}
	
	// Configure container interface
	cmd = exec.Command("ip", "netns", "exec", nm.containerID, "ip", "addr", "add", network.ContainerIP+"/16", "dev", network.VethContainer)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set container IP: %w", err)
	}
	
	// Bring container interface up
	cmd = exec.Command("ip", "netns", "exec", nm.containerID, "ip", "link", "set", "dev", network.VethContainer, "up")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to bring container interface up: %w", err)
	}
	
	// Set default route
	cmd = exec.Command("ip", "netns", "exec", nm.containerID, "ip", "route", "add", "default", "via", network.GatewayIP)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set default route: %w", err)
	}
	
	// Configure DNS
	if err := nm.configureDNS(network); err != nil {
		return fmt.Errorf("failed to configure DNS: %w", err)
	}
	
	fmt.Printf("    Container network configured: %s\n", network.ContainerIP)
	return nil
}

// configureDNS configures DNS for the container
func (nm *NetworkManager) configureDNS(network *ContainerNetwork) error {
	// Create resolv.conf
	resolvConf := fmt.Sprintf("/var/run/netns/%s/resolv.conf", nm.containerID)
	
	var dnsContent strings.Builder
	dnsContent.WriteString("# Container DNS configuration\n")
	for _, dns := range network.DNS {
		dnsContent.WriteString(fmt.Sprintf("nameserver %s\n", dns))
	}
	
	if err := os.WriteFile(resolvConf, []byte(dnsContent.String()), 0644); err != nil {
		return fmt.Errorf("failed to write resolv.conf: %w", err)
	}
	
	return nil
}

// setupIptablesRules sets up iptables rules for container networking
func (nm *NetworkManager) setupIptablesRules(network *ContainerNetwork) error {
	fmt.Println("  Setting up iptables rules...")
	
	// Enable IP forwarding
	cmd := exec.Command("sysctl", "-w", "net.ipv4.ip_forward=1")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to enable IP forwarding: %w", err)
	}
	
	// Add NAT rule for container traffic
	cmd = exec.Command("iptables", "-t", "nat", "-A", "POSTROUTING", "-s", network.ContainerIP+"/16", "!", "-o", network.BridgeName, "-j", "MASQUERADE")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add NAT rule: %w", err)
	}
	
	// Add forward rule
	cmd = exec.Command("iptables", "-A", "FORWARD", "-i", network.BridgeName, "!", "-o", network.BridgeName, "-j", "ACCEPT")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add forward rule: %w", err)
	}
	
	// Add return forward rule
	cmd = exec.Command("iptables", "-A", "FORWARD", "-i", network.BridgeName, "-o", network.BridgeName, "-j", "ACCEPT")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add return forward rule: %w", err)
	}
	
	fmt.Println("    Iptables rules configured")
	return nil
}

// AddPortMapping adds a port mapping to the container
func (nm *NetworkManager) AddPortMapping(network *ContainerNetwork, hostPort, containerPort, protocol string) error {
	fmt.Printf("🌐 Adding port mapping: %s:%s/%s\n", hostPort, containerPort, protocol)
	
	// Add port mapping to network
	portMapping := PortMapping{
		HostPort:      hostPort,
		ContainerPort: containerPort,
		Protocol:      protocol,
		HostIP:        "0.0.0.0",
	}
	network.PortMappings = append(network.PortMappings, portMapping)
	
	// Add iptables rule for port forwarding
	cmd := exec.Command("iptables", "-t", "nat", "-A", "PREROUTING", "-p", protocol, "--dport", hostPort, "-j", "DNAT", "--to-destination", network.ContainerIP+":"+containerPort)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add port forwarding rule: %w", err)
	}
	
	// Add iptables rule for local port forwarding
	cmd = exec.Command("iptables", "-t", "nat", "-A", "OUTPUT", "-p", protocol, "--dport", hostPort, "-j", "DNAT", "--to-destination", network.ContainerIP+":"+containerPort)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add local port forwarding rule: %w", err)
	}
	
	fmt.Printf("✅ Port mapping added: %s:%s/%s -> %s:%s\n", hostPort, protocol, network.ContainerIP, containerPort)
	return nil
}

// RemovePortMapping removes a port mapping
func (nm *NetworkManager) RemovePortMapping(network *ContainerNetwork, hostPort, containerPort, protocol string) error {
	fmt.Printf("🌐 Removing port mapping: %s:%s/%s\n", hostPort, containerPort, protocol)
	
	// Remove iptables rule for port forwarding
	cmd := exec.Command("iptables", "-t", "nat", "-D", "PREROUTING", "-p", protocol, "--dport", hostPort, "-j", "DNAT", "--to-destination", network.ContainerIP+":"+containerPort)
	cmd.Run() // Ignore error, rule might not exist
	
	// Remove iptables rule for local port forwarding
	cmd = exec.Command("iptables", "-t", "nat", "-D", "OUTPUT", "-p", protocol, "--dport", hostPort, "-j", "DNAT", "--to-destination", network.ContainerIP+":"+containerPort)
	cmd.Run() // Ignore error, rule might not exist
	
	// Remove from network configuration
	for i, mapping := range network.PortMappings {
		if mapping.HostPort == hostPort && mapping.ContainerPort == containerPort && mapping.Protocol == protocol {
			network.PortMappings = append(network.PortMappings[:i], network.PortMappings[i+1:]...)
			break
		}
	}
	
	fmt.Printf("✅ Port mapping removed: %s:%s/%s\n", hostPort, containerPort, protocol)
	return nil
}

// ExecuteInNamespace executes a command in the container network namespace
func (nm *NetworkManager) ExecuteInNamespace(network *ContainerNetwork, command string) error {
	fmt.Printf("🌐 Executing command in network namespace: %s\n", command)
	
	// Execute command in network namespace
	cmd := exec.Command("ip", "netns", "exec", nm.containerID, "sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute command in namespace: %w", err)
	}
	
	return nil
}

// GetNetworkInfo returns network information
func (nm *NetworkManager) GetNetworkInfo(network *ContainerNetwork) (*NetworkInfo, error) {
	info := &NetworkInfo{
		ContainerID:   network.ContainerID,
		NetworkMode:   network.NetworkMode,
		BridgeName:    network.BridgeName,
		ContainerIP:   network.ContainerIP,
		GatewayIP:     network.GatewayIP,
		PortMappings:  network.PortMappings,
		DNS:           network.DNS,
		Hostname:      network.Hostname,
	}
	
	// Get interface information
	if interfaces, err := net.Interfaces(); err == nil {
		for _, iface := range interfaces {
			if iface.Name == network.VethHost {
				info.HostInterface = iface.Name
				info.HostInterfaceIndex = iface.Index
				break
			}
		}
	}
	
	return info, nil
}

// CleanupNetwork cleans up the container network
func (nm *NetworkManager) CleanupNetwork(network *ContainerNetwork) error {
	fmt.Printf("🧹 Cleaning up network for container: %s\n", nm.containerID)
	
	// Remove port mappings
	for _, mapping := range network.PortMappings {
		nm.RemovePortMapping(network, mapping.HostPort, mapping.ContainerPort, mapping.Protocol)
	}
	
	// Remove iptables rules
	nm.cleanupIptablesRules(network)
	
	// Remove veth pair
	cmd := exec.Command("ip", "link", "delete", network.VethHost)
	cmd.Run() // Ignore error, interface might not exist
	
	// Remove network namespace
	nsPath := fmt.Sprintf("/var/run/netns/%s", nm.containerID)
	os.Remove(nsPath)
	
	// Remove bridge (only if no other containers are using it)
	cmd = exec.Command("ip", "link", "delete", network.BridgeName)
	cmd.Run() // Ignore error, bridge might be in use
	
	fmt.Printf("✅ Network cleaned up for container: %s\n", nm.containerID)
	return nil
}

// cleanupIptablesRules removes iptables rules
func (nm *NetworkManager) cleanupIptablesRules(network *ContainerNetwork) {
	// Remove NAT rule
	cmd := exec.Command("iptables", "-t", "nat", "-D", "POSTROUTING", "-s", network.ContainerIP+"/16", "!", "-o", network.BridgeName, "-j", "MASQUERADE")
	cmd.Run()
	
	// Remove forward rules
	cmd = exec.Command("iptables", "-D", "FORWARD", "-i", network.BridgeName, "!", "-o", network.BridgeName, "-j", "ACCEPT")
	cmd.Run()
	
	cmd = exec.Command("iptables", "-D", "FORWARD", "-i", network.BridgeName, "-o", network.BridgeName, "-j", "ACCEPT")
	cmd.Run()
}

// generateContainerIP generates a unique IP for the container
func (nm *NetworkManager) generateContainerIP() string {
	// Simple IP generation - in real implementation, this would be more sophisticated
	containerNum := len(nm.containerID) % 254
	return fmt.Sprintf("172.17.0.%d", containerNum+2)
}

// CheckNetworkSupport checks if networking features are supported
func (nm *NetworkManager) CheckNetworkSupport() error {
	fmt.Println("🔍 Checking network support...")
	
	// Check if ip command is available
	if _, err := exec.LookPath("ip"); err != nil {
		return fmt.Errorf("ip command not found: %w", err)
	}
	
	// Check if iptables is available
	if _, err := exec.LookPath("iptables"); err != nil {
		return fmt.Errorf("iptables command not found: %w", err)
	}
	
	// Check if bridge module is loaded
	if _, err := os.Stat("/sys/class/net/br0"); err != nil {
		// Bridge might not exist yet, check if module is available
		if _, err := os.Stat("/sys/module/bridge"); err != nil {
			return fmt.Errorf("bridge module not loaded")
		}
	}
	
	// Check if veth module is loaded
	if _, err := os.Stat("/sys/module/veth"); err != nil {
		return fmt.Errorf("veth module not loaded")
	}
	
	fmt.Println("✅ Network support available")
	return nil
}

// PrintNetworkInfo prints network information
func (nm *NetworkManager) PrintNetworkInfo(network *ContainerNetwork) {
	fmt.Println("🌐 Network Information")
	fmt.Println("====================")
	
	// Check network support
	if err := nm.CheckNetworkSupport(); err != nil {
		fmt.Printf("❌ Network support check failed: %v\n", err)
		return
	}
	
	// Get network info
	info, err := nm.GetNetworkInfo(network)
	if err != nil {
		fmt.Printf("❌ Failed to get network info: %v\n", err)
		return
	}
	
	fmt.Printf("Container ID: %s\n", info.ContainerID)
	fmt.Printf("Network Mode: %s\n", info.NetworkMode)
	fmt.Printf("Bridge Name: %s\n", info.BridgeName)
	fmt.Printf("Container IP: %s\n", info.ContainerIP)
	fmt.Printf("Gateway IP: %s\n", info.GatewayIP)
	fmt.Printf("Hostname: %s\n", info.Hostname)
	fmt.Printf("Host Interface: %s (index: %d)\n", info.HostInterface, info.HostInterfaceIndex)
	
	fmt.Println("\nDNS Servers:")
	for _, dns := range info.DNS {
		fmt.Printf("  - %s\n", dns)
	}
	
	fmt.Println("\nPort Mappings:")
	for _, mapping := range info.PortMappings {
		fmt.Printf("  - %s:%s/%s -> %s:%s\n", mapping.HostIP, mapping.HostPort, mapping.Protocol, info.ContainerIP, mapping.ContainerPort)
	}
}

// NetworkInfo represents network information
type NetworkInfo struct {
	ContainerID        string        `json:"container_id"`
	NetworkMode        string        `json:"network_mode"`
	BridgeName         string        `json:"bridge_name"`
	ContainerIP        string        `json:"container_ip"`
	GatewayIP          string        `json:"gateway_ip"`
	PortMappings       []PortMapping `json:"port_mappings"`
	DNS                []string      `json:"dns"`
	Hostname           string        `json:"hostname"`
	HostInterface      string        `json:"host_interface"`
	HostInterfaceIndex int           `json:"host_interface_index"`
}
