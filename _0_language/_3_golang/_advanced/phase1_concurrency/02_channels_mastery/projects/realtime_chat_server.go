package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// 💬 REAL-TIME CHAT SERVER PROJECT
// A WebSocket-like chat server using channels for communication

type Client struct {
	ID       string
	Conn     net.Conn
	Messages chan string
	Quit     chan struct{}
}

type ChatServer struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan string
	quit       chan struct{}
	mutex      sync.RWMutex
}

func main() {
	fmt.Println("💬 REAL-TIME CHAT SERVER")
	fmt.Println("========================")

	// Create chat server
	server := NewChatServer()
	defer server.Close()

	// Start server
	go server.Run()

	// Start listening for connections
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer listener.Close()

	fmt.Println("🚀 Server started on :8081")
	fmt.Println("📝 Connect with: telnet localhost 8081")
	fmt.Println("🛑 Press Ctrl+C to stop")

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// Create client and register
		client := NewClient(conn)
		server.RegisterClient(client)
	}
}

// NewChatServer creates a new chat server
func NewChatServer() *ChatServer {
	return &ChatServer{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan string),
		quit:       make(chan struct{}),
	}
}

// NewClient creates a new client
func NewClient(conn net.Conn) *Client {
	return &Client{
		ID:       generateID(),
		Conn:     conn,
		Messages: make(chan string, 10),
		Quit:     make(chan struct{}),
	}
}

// Run starts the chat server
func (s *ChatServer) Run() {
	for {
		select {
		case client := <-s.register:
			s.registerClient(client)
		case client := <-s.unregister:
			s.unregisterClient(client)
		case message := <-s.broadcast:
			s.broadcastMessage(message)
		case <-s.quit:
			return
		}
	}
}

// RegisterClient registers a new client
func (s *ChatServer) RegisterClient(client *Client) {
	s.register <- client
}

// registerClient handles client registration
func (s *ChatServer) registerClient(client *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.clients[client.ID] = client
	fmt.Printf("👤 Client %s connected (Total: %d)\n", client.ID, len(s.clients))

	// Start client handler
	go s.handleClient(client)
}

// unregisterClient handles client disconnection
func (s *ChatServer) unregisterClient(client *Client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.clients[client.ID]; exists {
		delete(s.clients, client.ID)
		close(client.Messages)
		client.Conn.Close()
		fmt.Printf("👋 Client %s disconnected (Total: %d)\n", client.ID, len(s.clients))
	}
}

// broadcastMessage broadcasts a message to all clients
func (s *ChatServer) broadcastMessage(message string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, client := range s.clients {
		select {
		case client.Messages <- message:
		default:
			// Client's message channel is full, skip
		}
	}
}

// handleClient handles a client connection
func (s *ChatServer) handleClient(client *Client) {
	defer func() {
		s.unregister <- client
	}()

	// Send welcome message
	welcome := fmt.Sprintf("Welcome to the chat! Your ID: %s\n", client.ID)
	client.Conn.Write([]byte(welcome))

	// Start message writer
	go s.writeMessages(client)

	// Start message reader
	s.readMessages(client)
}

// writeMessages writes messages to the client
func (s *ChatServer) writeMessages(client *Client) {
	for {
		select {
		case message := <-client.Messages:
			client.Conn.Write([]byte(message + "\n"))
		case <-client.Quit:
			return
		}
	}
}

// readMessages reads messages from the client
func (s *ChatServer) readMessages(client *Client) {
	scanner := bufio.NewScanner(client.Conn)
	
	for scanner.Scan() {
		message := strings.TrimSpace(scanner.Text())
		
		if message == "" {
			continue
		}

		// Handle special commands
		if strings.HasPrefix(message, "/") {
			s.handleCommand(client, message)
			continue
		}

		// Broadcast message
		formattedMessage := fmt.Sprintf("[%s] %s: %s", 
			time.Now().Format("15:04:05"), client.ID, message)
		s.broadcast <- formattedMessage
	}

	// Client disconnected
	close(client.Quit)
}

// handleCommand handles special commands
func (s *ChatServer) handleCommand(client *Client, command string) {
	parts := strings.Fields(command)
	
	switch parts[0] {
	case "/help":
		help := "Available commands:\n" +
			"  /help - Show this help\n" +
			"  /list - List connected clients\n" +
			"  /quit - Disconnect\n"
		client.Messages <- help
		
	case "/list":
		s.mutex.RLock()
		clientList := "Connected clients:\n"
		for id := range s.clients {
			clientList += fmt.Sprintf("  - %s\n", id)
		}
		s.mutex.RUnlock()
		client.Messages <- clientList
		
	case "/quit":
		client.Messages <- "Goodbye!"
		close(client.Quit)
		
	default:
		client.Messages <- "Unknown command. Type /help for available commands."
	}
}

// Close gracefully shuts down the server
func (s *ChatServer) Close() {
	close(s.quit)
	
	// Close all client connections
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	for _, client := range s.clients {
		close(client.Quit)
		client.Conn.Close()
	}
}

// generateID generates a unique client ID
func generateID() string {
	return fmt.Sprintf("client_%d", time.Now().UnixNano()%10000)
}
