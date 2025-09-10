# net Package - Network Operations 🌐

The `net` package provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets. It's essential for building distributed systems and network applications.

## 🎯 Key Concepts

### 1. **Network Interfaces**
- `Interface` - Network interface representation
- `Addrs()` - Get interface addresses
- `Flags()` - Get interface flags
- `HardwareAddr` - MAC address
- `MulticastAddrs()` - Get multicast addresses

### 2. **IP Addresses**
- `IP` - IP address type
- `ParseIP()` - Parse IP address string
- `IPv4()` - Create IPv4 address
- `IPv6()` - Create IPv6 address
- `IsLoopback()` - Check if loopback
- `IsMulticast()` - Check if multicast
- `IsUnspecified()` - Check if unspecified

### 3. **TCP Operations**
- `TCPAddr` - TCP address
- `TCPConn` - TCP connection
- `DialTCP()` - Connect to TCP server
- `ListenTCP()` - Listen for TCP connections
- `AcceptTCP()` - Accept TCP connection
- `Read()` - Read from connection
- `Write()` - Write to connection

### 4. **UDP Operations**
- `UDPAddr` - UDP address
- `UDPConn` - UDP connection
- `DialUDP()` - Connect to UDP server
- `ListenUDP()` - Listen for UDP packets
- `ReadFromUDP()` - Read UDP packet
- `WriteToUDP()` - Write UDP packet

### 5. **Domain Name Resolution**
- `LookupHost()` - Lookup hostname
- `LookupIP()` - Lookup IP addresses
- `LookupCNAME()` - Lookup canonical name
- `LookupMX()` - Lookup mail exchange
- `LookupTXT()` - Lookup text records
- `LookupSRV()` - Lookup service records

### 6. **Unix Domain Sockets**
- `UnixAddr` - Unix socket address
- `UnixConn` - Unix socket connection
- `DialUnix()` - Connect to Unix socket
- `ListenUnix()` - Listen on Unix socket

## 🚀 Common Patterns

### Basic TCP Client
```go
conn, err := net.Dial("tcp", "localhost:8080")
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

// Send data
conn.Write([]byte("Hello Server"))

// Read response
buffer := make([]byte, 1024)
n, err := conn.Read(buffer)
```

### Basic TCP Server
```go
listener, err := net.Listen("tcp", ":8080")
if err != nil {
    log.Fatal(err)
}
defer listener.Close()

for {
    conn, err := listener.Accept()
    if err != nil {
        log.Println(err)
        continue
    }
    go handleConnection(conn)
}
```

### UDP Communication
```go
// UDP client
conn, err := net.Dial("udp", "localhost:8080")
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

conn.Write([]byte("Hello UDP"))

// UDP server
addr, _ := net.ResolveUDPAddr("udp", ":8080")
conn, err := net.ListenUDP("udp", addr)
if err != nil {
    log.Fatal(err)
}
defer conn.Close()
```

## ⚠️ Common Pitfalls

1. **Not closing connections** - Always close connections to prevent leaks
2. **Blocking operations** - Use timeouts for network operations
3. **Buffer overflow** - Check buffer sizes for reads
4. **Address resolution** - Handle DNS resolution errors
5. **Concurrent access** - Protect shared network resources

## 🎯 Best Practices

1. **Use timeouts** - Set timeouts for all network operations
2. **Handle errors** - Always check and handle network errors
3. **Close connections** - Use defer to ensure connections are closed
4. **Use context** - Use context for cancellation and timeouts
5. **Validate addresses** - Validate IP addresses and ports

## 🔍 Advanced Features

### Custom Dialer
```go
dialer := &net.Dialer{
    Timeout:   30 * time.Second,
    KeepAlive: 30 * time.Second,
}
conn, err := dialer.Dial("tcp", "example.com:80")
```

### Network Interface Management
```go
interfaces, err := net.Interfaces()
if err != nil {
    log.Fatal(err)
}

for _, iface := range interfaces {
    fmt.Printf("Interface: %s\n", iface.Name)
    addrs, _ := iface.Addrs()
    for _, addr := range addrs {
        fmt.Printf("  Address: %s\n", addr.String())
    }
}
```

### IP Address Manipulation
```go
ip := net.ParseIP("192.168.1.1")
if ip != nil {
    fmt.Printf("IP: %s\n", ip.String())
    fmt.Printf("Is IPv4: %t\n", ip.To4() != nil)
    fmt.Printf("Is loopback: %t\n", ip.IsLoopback())
}
```

## 📚 Real-world Applications

1. **Web Servers** - HTTP server implementation
2. **Database Connections** - Database client connections
3. **Microservices** - Service-to-service communication
4. **Load Balancers** - Traffic distribution
5. **Monitoring** - Network health checks

## 🧠 Memory Tips

- **net** = **N**etwork **E**ngine **T**oolkit
- **TCP** = **T**ransmission **C**ontrol **P**rotocol
- **UDP** = **U**ser **D**atagram **P**rotocol
- **IP** = **I**nternet **P**rotocol
- **DNS** = **D**omain **N**ame **S**ystem
- **Dial** = **D**ial connection
- **Listen** = **L**isten for connections
- **Accept** = **A**ccept connection

Remember: The net package is your gateway to network programming in Go! 🎯
