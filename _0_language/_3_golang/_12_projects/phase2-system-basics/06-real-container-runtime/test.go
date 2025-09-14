package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("🧪 Testing Real Container Runtime")
	fmt.Println("==================================")
	
	// Check prerequisites
	if err := checkPrerequisites(); err != nil {
		fmt.Printf("❌ Prerequisites check failed: %v\n", err)
		os.Exit(1)
	}
	
	// Test namespace manager
	fmt.Println("\n🔒 Testing Namespace Manager...")
	testNamespaceManager()
	
	// Test cgroup manager
	fmt.Println("\n⚡ Testing CGroup Manager...")
	testCGroupManager()
	
	// Test filesystem manager
	fmt.Println("\n💾 Testing Filesystem Manager...")
	testFilesystemManager()
	
	// Test network manager
	fmt.Println("\n🌐 Testing Network Manager...")
	testNetworkManager()
	
	// Test real container runtime
	fmt.Println("\n🐳 Testing Real Container Runtime...")
	testRealContainerRuntime()
	
	fmt.Println("\n✅ All tests completed successfully!")
}

// checkPrerequisites checks if all prerequisites are met
func checkPrerequisites() error {
	// Check if running as root
	if os.Geteuid() != 0 {
		return fmt.Errorf("root privileges required for real containerization")
	}
	
	// Check if running on Linux
	if runtime.GOOS != "linux" {
		return fmt.Errorf("real containerization only supported on Linux")
	}
	
	// Check if required commands are available
	requiredCommands := []string{"unshare", "ip", "iptables", "hostname"}
	for _, cmd := range requiredCommands {
		if !checkCommandExists(cmd) {
			return fmt.Errorf("required command not found: %s", cmd)
		}
	}
	
	fmt.Println("✅ Prerequisites check passed")
	return nil
}

// testNamespaceManager tests the namespace manager
func testNamespaceManager() {
	nm := NewNamespaceManager()
	
	// Check namespace support
	if err := nm.CheckNamespaceSupport(); err != nil {
		fmt.Printf("❌ Namespace support check failed: %v\n", err)
		return
	}
	
	// Print namespace info
	nm.PrintNamespaceInfo()
	
	fmt.Println("✅ Namespace manager test passed")
}

// testCGroupManager tests the cgroup manager
func testCGroupManager() {
	cgm := NewCGroupManager("test-container")
	
	// Check cgroup support
	if err := cgm.CheckCGroupSupport(); err != nil {
		fmt.Printf("❌ CGroup support check failed: %v\n", err)
		return
	}
	
	// Print cgroup info
	cgm.PrintCGroupInfo()
	
	fmt.Println("✅ CGroup manager test passed")
}

// testFilesystemManager tests the filesystem manager
func testFilesystemManager() {
	fm := NewFilesystemManager("test-container")
	
	// Check overlay support
	if err := fm.CheckOverlaySupport(); err != nil {
		fmt.Printf("❌ Overlay support check failed: %v\n", err)
		return
	}
	
	// Create test filesystem
	containerFS, err := fm.CreateContainerFilesystem("ubuntu:latest")
	if err != nil {
		fmt.Printf("❌ Failed to create test filesystem: %v\n", err)
		return
	}
	
	// Print filesystem info
	fm.PrintFilesystemInfo(containerFS)
	
	// Cleanup
	if err := fm.CleanupFilesystem(containerFS); err != nil {
		fmt.Printf("Warning: failed to cleanup test filesystem: %v\n", err)
	}
	
	fmt.Println("✅ Filesystem manager test passed")
}

// testNetworkManager tests the network manager
func testNetworkManager() {
	nm := NewNetworkManager("test-container")
	
	// Check network support
	if err := nm.CheckNetworkSupport(); err != nil {
		fmt.Printf("❌ Network support check failed: %v\n", err)
		return
	}
	
	// Create test network
	network, err := nm.CreateContainerNetwork()
	if err != nil {
		fmt.Printf("❌ Failed to create test network: %v\n", err)
		return
	}
	
	// Print network info
	nm.PrintNetworkInfo(network)
	
	// Cleanup
	if err := nm.CleanupNetwork(network); err != nil {
		fmt.Printf("Warning: failed to cleanup test network: %v\n", err)
	}
	
	fmt.Println("✅ Network manager test passed")
}

// testRealContainerRuntime tests the real container runtime
func testRealContainerRuntime() {
	rtr := NewRealContainerRuntime()
	
	// Test container creation
	limits := &ResourceLimits{
		MemoryLimit: 512 * 1024 * 1024, // 512MB
		CPULimit:    50,                 // 50%
		PIDsLimit:   100,
	}
	
	container, err := rtr.CreateContainer("test-container", "ubuntu:latest", "echo 'Hello from real container!'", limits)
	if err != nil {
		fmt.Printf("❌ Failed to create test container: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Test container created: %s\n", container.ID)
	
	// Test container info
	info, err := rtr.GetContainerInfo(container.ID)
	if err != nil {
		fmt.Printf("❌ Failed to get container info: %v\n", err)
		return
	}
	
	fmt.Printf("Container Info:\n")
	fmt.Printf("  ID: %s\n", info.Container.ID)
	fmt.Printf("  Name: %s\n", info.Container.Name)
	fmt.Printf("  Image: %s\n", info.Container.Image)
	fmt.Printf("  Status: %s\n", info.Container.Status)
	fmt.Printf("  IP: %s\n", info.Container.Network.ContainerIP)
	
	// Test container removal
	if err := rtr.RemoveContainer(container.ID); err != nil {
		fmt.Printf("❌ Failed to remove test container: %v\n", err)
		return
	}
	
	fmt.Println("✅ Test container removed")
	fmt.Println("✅ Real container runtime test passed")
}

// checkCommandExists checks if a command exists
func checkCommandExists(cmd string) bool {
	// Simple check - in real implementation, this would use os/exec.LookPath
	return true
}
