package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Custom types for demonstration
type Server struct {
	Port    int
	Handler func(net.Conn)
}

type Client struct {
	Host string
	Port int
}

func (c Client) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func main() {
	fmt.Println("🚀 Go net Package Mastery Examples")
	fmt.Println("===================================")

	// 1. Basic IP Address Operations
	fmt.Println("\n1. Basic IP Address Operations:")
	
	// Parse IP addresses
	ipv4 := net.ParseIP("192.168.1.1")
	ipv6 := net.ParseIP("2001:db8::1")
	invalidIP := net.ParseIP("invalid")
	
	fmt.Printf("IPv4: %s (valid: %t)\n", ipv4, ipv4 != nil)
	fmt.Printf("IPv6: %s (valid: %t)\n", ipv6, ipv6 != nil)
	fmt.Printf("Invalid: %s (valid: %t)\n", invalidIP, invalidIP != nil)
	
	// IP address properties
	if ipv4 != nil {
		fmt.Printf("IPv4 properties:\n")
		fmt.Printf("  Is loopback: %t\n", ipv4.IsLoopback())
		fmt.Printf("  Is multicast: %t\n", ipv4.IsMulticast())
		fmt.Printf("  Is unspecified: %t\n", ipv4.IsUnspecified())
		fmt.Printf("  Is private: %t\n", ipv4.IsPrivate())
		fmt.Printf("  Is global unicast: %t\n", ipv4.IsGlobalUnicast())
	}
	
	// Create IP addresses
	localhost := net.IPv4(127, 0, 0, 1)
	fmt.Printf("Localhost: %s\n", localhost.String())

	// 2. Network Interface Operations
	fmt.Println("\n2. Network Interface Operations:")
	
	// Get network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("Error getting interfaces: %v", err)
	} else {
		fmt.Printf("Found %d network interfaces:\n", len(interfaces))
		for i, iface := range interfaces {
			fmt.Printf("  [%d] %s (Index: %d, MTU: %d)\n", i, iface.Name, iface.Index, iface.MTU)
			
			// Get addresses for this interface
			addrs, err := iface.Addrs()
			if err == nil {
				for _, addr := range addrs {
					fmt.Printf("    Address: %s\n", addr.String())
				}
			}
			
			// Get multicast addresses
			multicastAddrs, err := iface.MulticastAddrs()
			if err == nil && len(multicastAddrs) > 0 {
				for _, addr := range multicastAddrs {
					fmt.Printf("    Multicast: %s\n", addr.String())
				}
			}
		}
	}

	// 3. DNS Resolution
	fmt.Println("\n3. DNS Resolution:")
	
	// Lookup hostname
	hostname := "google.com"
	ips, err := net.LookupHost(hostname)
	if err != nil {
		log.Printf("Error looking up %s: %v", hostname, err)
	} else {
		fmt.Printf("IP addresses for %s:\n", hostname)
		for _, ip := range ips {
			fmt.Printf("  %s\n", ip)
		}
	}
	
	// Lookup IP addresses
	ipAddrs, err := net.LookupIP(hostname)
	if err != nil {
		log.Printf("Error looking up IP for %s: %v", hostname, err)
	} else {
		fmt.Printf("IP addresses (net.IP) for %s:\n", hostname)
		for _, ip := range ipAddrs {
			fmt.Printf("  %s (IPv4: %t, IPv6: %t)\n", ip.String(), ip.To4() != nil, ip.To16() != nil)
		}
	}
	
	// Lookup CNAME
	cname, err := net.LookupCNAME(hostname)
	if err != nil {
		log.Printf("Error looking up CNAME for %s: %v", hostname, err)
	} else {
		fmt.Printf("CNAME for %s: %s\n", hostname, cname)
	}
	
	// Lookup MX records
	mxRecords, err := net.LookupMX(hostname)
	if err != nil {
		log.Printf("Error looking up MX for %s: %v", hostname, err)
	} else {
		fmt.Printf("MX records for %s:\n", hostname)
		for _, mx := range mxRecords {
			fmt.Printf("  %s (priority: %d)\n", mx.Host, mx.Pref)
		}
	}

	// 4. TCP Client Operations
	fmt.Println("\n4. TCP Client Operations:")
	
	// Test TCP connection to a public service
	testTCPConnection := func(host, port string) {
		address := net.JoinHostPort(host, port)
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			fmt.Printf("  Failed to connect to %s: %v\n", address, err)
			return
		}
		defer conn.Close()
		
		fmt.Printf("  Successfully connected to %s\n", address)
		
		// Set read deadline
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		
		// Try to read some data
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("  Read error: %v\n", err)
		} else {
			fmt.Printf("  Read %d bytes: %s\n", n, string(buffer[:n]))
		}
	}
	
	// Test connections to common services
	testTCPConnection("httpbin.org", "80")
	testTCPConnection("google.com", "80")

	// 5. TCP Server Operations
	fmt.Println("\n5. TCP Server Operations:")
	
	// Start a simple TCP server
	startTCPServer := func(port int) {
		address := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", address)
		if err != nil {
			log.Printf("Error starting TCP server on port %d: %v", port, err)
			return
		}
		defer listener.Close()
		
		fmt.Printf("TCP server listening on %s\n", address)
		
		// Accept one connection for demonstration
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			return
		}
		defer conn.Close()
		
		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())
		
		// Echo server
		io.Copy(conn, conn)
	}
	
	// Start server in a goroutine
	go startTCPServer(8080)
	time.Sleep(100 * time.Millisecond) // Give server time to start
	
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Printf("Error connecting to server: %v", err)
	} else {
		defer conn.Close()
		
		// Send data
		message := "Hello TCP Server!"
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error writing to server: %v", err)
		} else {
			fmt.Printf("Sent: %s\n", message)
		}
		
		// Read response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading from server: %v", err)
		} else {
			fmt.Printf("Received: %s\n", string(buffer[:n]))
		}
	}

	// 6. UDP Operations
	fmt.Println("\n6. UDP Operations:")
	
	// Start UDP server
	startUDPServer := func(port int) {
		address := fmt.Sprintf(":%d", port)
		addr, err := net.ResolveUDPAddr("udp", address)
		if err != nil {
			log.Printf("Error resolving UDP address: %v", err)
			return
		}
		
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			log.Printf("Error starting UDP server: %v", err)
			return
		}
		defer conn.Close()
		
		fmt.Printf("UDP server listening on %s\n", address)
		
		// Read one packet
		buffer := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading UDP packet: %v", err)
			return
		}
		
		fmt.Printf("Received %d bytes from %s: %s\n", n, clientAddr, string(buffer[:n]))
		
		// Echo back
		_, err = conn.WriteToUDP(buffer[:n], clientAddr)
		if err != nil {
			log.Printf("Error writing UDP response: %v", err)
		}
	}
	
	// Start UDP server in goroutine
	go startUDPServer(8081)
	time.Sleep(100 * time.Millisecond)
	
	// Send UDP packet
	conn, err = net.Dial("udp", "localhost:8081")
	if err != nil {
		log.Printf("Error connecting to UDP server: %v", err)
	} else {
		defer conn.Close()
		
		message := "Hello UDP Server!"
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error writing to UDP server: %v", err)
		} else {
			fmt.Printf("Sent UDP: %s\n", message)
		}
		
		// Read response
		buffer := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading UDP response: %v", err)
		} else {
			fmt.Printf("Received UDP: %s\n", string(buffer[:n]))
		}
	}

	// 7. Advanced TCP Operations
	fmt.Println("\n7. Advanced TCP Operations:")
	
	// Custom dialer with timeout
	dialer := &net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	
	conn, err = dialer.Dial("tcp", "httpbin.org:80")
	if err != nil {
		log.Printf("Error with custom dialer: %v", err)
	} else {
		defer conn.Close()
		fmt.Printf("Connected using custom dialer to %s\n", conn.RemoteAddr())
		
		// Set connection options
		if tcpConn, ok := conn.(*net.TCPConn); ok {
			tcpConn.SetKeepAlive(true)
			tcpConn.SetKeepAlivePeriod(30 * time.Second)
			fmt.Printf("Set TCP keep-alive options\n")
		}
	}

	// 8. Context with Network Operations
	fmt.Println("\n8. Context with Network Operations:")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Dial with context
	conn, err = (&net.Dialer{}).DialContext(ctx, "tcp", "httpbin.org:80")
	if err != nil {
		log.Printf("Error with context dial: %v", err)
	} else {
		defer conn.Close()
		fmt.Printf("Connected with context to %s\n", conn.RemoteAddr())
	}

	// 9. Network Address Parsing
	fmt.Println("\n9. Network Address Parsing:")
	
	// Parse network addresses
	addresses := []string{
		"192.168.1.1:8080",
		"[2001:db8::1]:8080",
		"localhost:3000",
		"example.com:443",
	}
	
	for _, addrStr := range addresses {
		host, port, err := net.SplitHostPort(addrStr)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", addrStr, err)
			continue
		}
		
		fmt.Printf("Address: %s\n", addrStr)
		fmt.Printf("  Host: %s\n", host)
		fmt.Printf("  Port: %s\n", port)
		
		// Resolve address
		addr, err := net.ResolveTCPAddr("tcp", addrStr)
		if err != nil {
			fmt.Printf("  Error resolving: %v\n", err)
		} else {
			fmt.Printf("  Resolved: %s\n", addr.String())
		}
	}

	// 10. Unix Domain Sockets
	fmt.Println("\n10. Unix Domain Sockets:")
	
	// Create Unix socket server
	socketPath := "/tmp/test.sock"
	
	// Remove existing socket file
	net.Dial("unix", socketPath) // This will fail, but that's okay
	
	// Start Unix socket server
	go func() {
		addr, err := net.ResolveUnixAddr("unix", socketPath)
		if err != nil {
			log.Printf("Error resolving Unix address: %v", err)
			return
		}
		
		listener, err := net.ListenUnix("unix", addr)
		if err != nil {
			log.Printf("Error starting Unix server: %v", err)
			return
		}
		defer listener.Close()
		
		fmt.Printf("Unix socket server listening on %s\n", socketPath)
		
		conn, err := listener.AcceptUnix()
		if err != nil {
			log.Printf("Error accepting Unix connection: %v", err)
			return
		}
		defer conn.Close()
		
		fmt.Printf("Accepted Unix connection from %s\n", conn.RemoteAddr())
		
		// Echo server
		io.Copy(conn, conn)
	}()
	
	time.Sleep(100 * time.Millisecond)
	
	// Connect to Unix socket
	conn, err = net.Dial("unix", socketPath)
	if err != nil {
		log.Printf("Error connecting to Unix socket: %v", err)
	} else {
		defer conn.Close()
		
		message := "Hello Unix Socket!"
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error writing to Unix socket: %v", err)
		} else {
			fmt.Printf("Sent to Unix socket: %s\n", message)
		}
		
		// Read response
		buffer := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading from Unix socket: %v", err)
		} else {
			fmt.Printf("Received from Unix socket: %s\n", string(buffer[:n]))
		}
	}

	// 11. Network Error Handling
	fmt.Println("\n11. Network Error Handling:")
	
	// Test various network error conditions
	testNetworkErrors := func() {
		// Test connection to non-existent host
		_, err := net.DialTimeout("tcp", "nonexistent.example.com:80", 2*time.Second)
		if err != nil {
			fmt.Printf("Expected error for non-existent host: %v\n", err)
		}
		
		// Test connection to closed port
		_, err = net.DialTimeout("tcp", "localhost:99999", 2*time.Second)
		if err != nil {
			fmt.Printf("Expected error for invalid port: %v\n", err)
		}
		
		// Test connection with invalid protocol
		_, err = net.Dial("invalid", "localhost:80")
		if err != nil {
			fmt.Printf("Expected error for invalid protocol: %v\n", err)
		}
	}
	
	testNetworkErrors()

	// 12. Network Interface Statistics
	fmt.Println("\n12. Network Interface Statistics:")
	
	// Get interface statistics
	interfaces, err = net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			fmt.Printf("Interface: %s\n", iface.Name)
			fmt.Printf("  Index: %d\n", iface.Index)
			fmt.Printf("  MTU: %d\n", iface.MTU)
			fmt.Printf("  Flags: %s\n", iface.Flags.String())
			fmt.Printf("  Hardware Address: %s\n", iface.HardwareAddr.String())
		}
	}

	// 13. Port Scanning
	fmt.Println("\n13. Port Scanning:")
	
	// Simple port scanner
	scanPort := func(host string, port int) bool {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			return false
		}
		conn.Close()
		return true
	}
	
	host := "localhost"
	ports := []int{22, 80, 443, 8080, 3306, 5432}
	
	fmt.Printf("Scanning ports on %s:\n", host)
	for _, port := range ports {
		open := scanPort(host, port)
		status := "closed"
		if open {
			status = "open"
		}
		fmt.Printf("  Port %d: %s\n", port, status)
	}

	// 14. Network Timeout Handling
	fmt.Println("\n14. Network Timeout Handling:")
	
	// Test different timeout scenarios
	testTimeouts := func() {
		// Short timeout
		conn, err := net.DialTimeout("tcp", "httpbin.org:80", 100*time.Millisecond)
		if err != nil {
			fmt.Printf("Short timeout error: %v\n", err)
		} else {
			conn.Close()
			fmt.Printf("Short timeout succeeded\n")
		}
		
		// Long timeout
		conn, err = net.DialTimeout("tcp", "httpbin.org:80", 10*time.Second)
		if err != nil {
			fmt.Printf("Long timeout error: %v\n", err)
		} else {
			conn.Close()
			fmt.Printf("Long timeout succeeded\n")
		}
	}
	
	testTimeouts()

	// 15. Advanced Network Operations
	fmt.Println("\n15. Advanced Network Operations:")
	
	// Create a simple network utility
	networkUtility := func(host, port string) {
		address := net.JoinHostPort(host, port)
		
		// Test TCP connection
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			fmt.Printf("TCP connection to %s failed: %v\n", address, err)
		} else {
			defer conn.Close()
			fmt.Printf("TCP connection to %s successful\n", address)
			
			// Get connection info
			if tcpConn, ok := conn.(*net.TCPConn); ok {
				localAddr := tcpConn.LocalAddr()
				remoteAddr := tcpConn.RemoteAddr()
				fmt.Printf("  Local: %s\n", localAddr)
				fmt.Printf("  Remote: %s\n", remoteAddr)
			}
		}
		
		// Test UDP connection
		conn, err = net.DialTimeout("udp", address, 5*time.Second)
		if err != nil {
			fmt.Printf("UDP connection to %s failed: %v\n", address, err)
		} else {
			defer conn.Close()
			fmt.Printf("UDP connection to %s successful\n", address)
		}
	}
	
	networkUtility("httpbin.org", "80")
	networkUtility("google.com", "443")

	fmt.Println("\n🎉 net Package Mastery Complete!")
}
