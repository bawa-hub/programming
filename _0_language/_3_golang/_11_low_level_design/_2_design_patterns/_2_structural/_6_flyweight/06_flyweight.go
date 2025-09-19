package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// BASIC FLYWEIGHT PATTERN
// =============================================================================

// Flyweight interface
type Flyweight interface {
	Operation(extrinsicState string)
}

// Concrete Flyweight - stores intrinsic state
type ConcreteFlyweight struct {
	intrinsicState string
}

func NewConcreteFlyweight(intrinsicState string) *ConcreteFlyweight {
	return &ConcreteFlyweight{intrinsicState: intrinsicState}
}

func (cf *ConcreteFlyweight) Operation(extrinsicState string) {
	fmt.Printf("ConcreteFlyweight: Intrinsic state = %s, Extrinsic state = %s\n", 
		cf.intrinsicState, extrinsicState)
}

// Flyweight Factory - manages flyweight instances
type FlyweightFactory struct {
	flyweights map[string]Flyweight
	mu         sync.RWMutex
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

func (ff *FlyweightFactory) GetFlyweight(key string) Flyweight {
	ff.mu.RLock()
	if flyweight, exists := ff.flyweights[key]; exists {
		ff.mu.RUnlock()
		return flyweight
	}
	ff.mu.RUnlock()
	
	ff.mu.Lock()
	defer ff.mu.Unlock()
	
	// Double-check locking
	if flyweight, exists := ff.flyweights[key]; exists {
		return flyweight
	}
	
	// Create new flyweight
	flyweight := NewConcreteFlyweight(key)
	ff.flyweights[key] = flyweight
	return flyweight
}

func (ff *FlyweightFactory) GetFlyweightCount() int {
	ff.mu.RLock()
	defer ff.mu.RUnlock()
	return len(ff.flyweights)
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. TEXT EDITOR FLYWEIGHT
type CharacterFlyweight interface {
	Display(font string, size int, color string)
}

type Character struct {
	char rune
}

func NewCharacter(char rune) *Character {
	return &Character{char: char}
}

func (c *Character) Display(font string, size int, color string) {
	fmt.Printf("Character '%c' with font=%s, size=%d, color=%s\n", 
		c.char, font, size, color)
}

type CharacterFactory struct {
	characters map[rune]CharacterFlyweight
	mu         sync.RWMutex
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{
		characters: make(map[rune]CharacterFlyweight),
	}
}

func (cf *CharacterFactory) GetCharacter(char rune) CharacterFlyweight {
	cf.mu.RLock()
	if character, exists := cf.characters[char]; exists {
		cf.mu.RUnlock()
		return character
	}
	cf.mu.RUnlock()
	
	cf.mu.Lock()
	defer cf.mu.Unlock()
	
	// Double-check locking
	if character, exists := cf.characters[char]; exists {
		return character
	}
	
	// Create new character
	character := NewCharacter(char)
	cf.characters[char] = character
	return character
}

func (cf *CharacterFactory) GetCharacterCount() int {
	cf.mu.RLock()
	defer cf.mu.RUnlock()
	return len(cf.characters)
}

// 2. GAME DEVELOPMENT FLYWEIGHT
type TreeFlyweight interface {
	Render(x, y int, season string)
}

type TreeType struct {
	name        string
	texture     string
	mesh        string
	height      float64
	width       float64
}

func NewTreeType(name, texture, mesh string, height, width float64) *TreeType {
	return &TreeType{
		name:    name,
		texture: texture,
		mesh:    mesh,
		height:  height,
		width:   width,
	}
}

func (tt *TreeType) Render(x, y int, season string) {
	fmt.Printf("Rendering %s tree at (%d, %d) in %s season\n", 
		tt.name, x, y, season)
	fmt.Printf("  Texture: %s, Mesh: %s, Size: %.1fx%.1f\n", 
		tt.texture, tt.mesh, tt.width, tt.height)
}

type TreeFactory struct {
	treeTypes map[string]TreeFlyweight
	mu        sync.RWMutex
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: make(map[string]TreeFlyweight),
	}
}

func (tf *TreeFactory) GetTreeType(name string) TreeFlyweight {
	tf.mu.RLock()
	if treeType, exists := tf.treeTypes[name]; exists {
		tf.mu.RUnlock()
		return treeType
	}
	tf.mu.RUnlock()
	
	tf.mu.Lock()
	defer tf.mu.Unlock()
	
	// Double-check locking
	if treeType, exists := tf.treeTypes[name]; exists {
		return treeType
	}
	
	// Create new tree type based on name
	var treeType TreeFlyweight
	switch name {
	case "oak":
		treeType = NewTreeType("Oak", "oak_texture.png", "oak_mesh.obj", 15.0, 8.0)
	case "pine":
		treeType = NewTreeType("Pine", "pine_texture.png", "pine_mesh.obj", 20.0, 6.0)
	case "maple":
		treeType = NewTreeType("Maple", "maple_texture.png", "maple_mesh.obj", 12.0, 10.0)
	default:
		treeType = NewTreeType("Generic", "generic_texture.png", "generic_mesh.obj", 10.0, 5.0)
	}
	
	tf.treeTypes[name] = treeType
	return treeType
}

func (tf *TreeFactory) GetTreeTypeCount() int {
	tf.mu.RLock()
	defer tf.mu.RUnlock()
	return len(tf.treeTypes)
}

// Tree instance with extrinsic state
type Tree struct {
	treeType TreeFlyweight
	x        int
	y        int
	season   string
}

func NewTree(treeType TreeFlyweight, x, y int, season string) *Tree {
	return &Tree{
		treeType: treeType,
		x:        x,
		y:        y,
		season:   season,
	}
}

func (t *Tree) Render() {
	t.treeType.Render(t.x, t.y, t.season)
}

// 3. DATABASE CONNECTION POOL FLYWEIGHT
type ConnectionFlyweight interface {
	Execute(query string) ([]map[string]interface{}, error)
	Close() error
	GetConnectionID() string
}

type DatabaseConnection struct {
	connectionID string
	host         string
	port         int
	database     string
	username     string
	password     string
	isConnected  bool
}

func NewDatabaseConnection(connectionID, host string, port int, database, username, password string) *DatabaseConnection {
	return &DatabaseConnection{
		connectionID: connectionID,
		host:         host,
		port:         port,
		database:     database,
		username:     username,
		password:     password,
		isConnected:  false,
	}
}

func (dc *DatabaseConnection) Execute(query string) ([]map[string]interface{}, error) {
	if !dc.isConnected {
		dc.isConnected = true
		fmt.Printf("Connection %s: Connecting to database\n", dc.connectionID)
	}
	fmt.Printf("Connection %s: Executing query: %s\n", dc.connectionID, query)
	return []map[string]interface{}{{"result": "data"}}, nil
}

func (dc *DatabaseConnection) Close() error {
	if dc.isConnected {
		dc.isConnected = false
		fmt.Printf("Connection %s: Closing connection\n", dc.connectionID)
	}
	return nil
}

func (dc *DatabaseConnection) GetConnectionID() string {
	return dc.connectionID
}

type ConnectionPool struct {
	connections map[string]ConnectionFlyweight
	mu          sync.RWMutex
	maxSize     int
}

func NewConnectionPool(maxSize int) *ConnectionPool {
	return &ConnectionPool{
		connections: make(map[string]ConnectionFlyweight),
		maxSize:     maxSize,
	}
}

func (cp *ConnectionPool) GetConnection(host string, port int, database string) ConnectionFlyweight {
	key := fmt.Sprintf("%s:%d/%s", host, port, database)
	
	cp.mu.RLock()
	if connection, exists := cp.connections[key]; exists {
		cp.mu.RUnlock()
		return connection
	}
	cp.mu.RUnlock()
	
	cp.mu.Lock()
	defer cp.mu.Unlock()
	
	// Double-check locking
	if connection, exists := cp.connections[key]; exists {
		return connection
	}
	
	// Check pool size
	if len(cp.connections) >= cp.maxSize {
		// Remove oldest connection (simple implementation)
		for k := range cp.connections {
			delete(cp.connections, k)
			break
		}
	}
	
	// Create new connection
	connectionID := fmt.Sprintf("conn_%d", time.Now().UnixNano())
	connection := NewDatabaseConnection(connectionID, host, port, database, "user", "pass")
	cp.connections[key] = connection
	return connection
}

func (cp *ConnectionPool) GetConnectionCount() int {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return len(cp.connections)
}

func (cp *ConnectionPool) CloseAll() {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	
	for _, connection := range cp.connections {
		connection.Close()
	}
	cp.connections = make(map[string]ConnectionFlyweight)
}

// 4. GUI ICON FLYWEIGHT
type IconFlyweight interface {
	Render(x, y int, size int, color string)
}

type Icon struct {
	name     string
	imageData []byte
	width    int
	height   int
}

func NewIcon(name string, width, height int) *Icon {
	// Simulate image data
	imageData := make([]byte, width*height*4) // RGBA
	for i := range imageData {
		imageData[i] = byte(i % 256)
	}
	
	return &Icon{
		name:      name,
		imageData: imageData,
		width:     width,
		height:    height,
	}
}

func (i *Icon) Render(x, y int, size int, color string) {
	fmt.Printf("Rendering %s icon at (%d, %d) with size %d and color %s\n", 
		i.name, x, y, size, color)
	fmt.Printf("  Original size: %dx%d, Data size: %d bytes\n", 
		i.width, i.height, len(i.imageData))
}

type IconFactory struct {
	icons map[string]IconFlyweight
	mu    sync.RWMutex
}

func NewIconFactory() *IconFactory {
	return &IconFactory{
		icons: make(map[string]IconFlyweight),
	}
}

func (if_ *IconFactory) GetIcon(name string) IconFlyweight {
	if_.mu.RLock()
	if icon, exists := if_.icons[name]; exists {
		if_.mu.RUnlock()
		return icon
	}
	if_.mu.RUnlock()
	
	if_.mu.Lock()
	defer if_.mu.Unlock()
	
	// Double-check locking
	if icon, exists := if_.icons[name]; exists {
		return icon
	}
	
	// Create new icon based on name
	var icon IconFlyweight
	switch name {
	case "home":
		icon = NewIcon("Home", 32, 32)
	case "settings":
		icon = NewIcon("Settings", 32, 32)
	case "user":
		icon = NewIcon("User", 32, 32)
	case "search":
		icon = NewIcon("Search", 32, 32)
	default:
		icon = NewIcon("Default", 32, 32)
	}
	
	if_.icons[name] = icon
	return icon
}

func (if_ *IconFactory) GetIconCount() int {
	if_.mu.RLock()
	defer if_.mu.RUnlock()
	return len(if_.icons)
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== FLYWEIGHT PATTERN DEMONSTRATION ===\n")

	// 1. BASIC FLYWEIGHT
	fmt.Println("1. BASIC FLYWEIGHT:")
	factory := NewFlyweightFactory()
	
	// Create multiple flyweights with same intrinsic state
	flyweight1 := factory.GetFlyweight("shared_state")
	flyweight2 := factory.GetFlyweight("shared_state")
	flyweight3 := factory.GetFlyweight("different_state")
	
	// Use flyweights with different extrinsic state
	flyweight1.Operation("extrinsic_state_1")
	flyweight2.Operation("extrinsic_state_2")
	flyweight3.Operation("extrinsic_state_3")
	
	fmt.Printf("Total flyweights created: %d\n", factory.GetFlyweightCount())
	fmt.Printf("flyweight1 == flyweight2: %t\n", flyweight1 == flyweight2)
	fmt.Printf("flyweight1 == flyweight3: %t\n", flyweight1 == flyweight3)
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Text Editor Flyweight
	fmt.Println("Text Editor Flyweight:")
	charFactory := NewCharacterFactory()
	
	// Create text with repeated characters
	text := "Hello, World!"
	for _, char := range text {
		character := charFactory.GetCharacter(char)
		character.Display("Arial", 12, "black")
	}
	
	fmt.Printf("Total unique characters: %d\n", charFactory.GetCharacterCount())
	fmt.Println()

	// Game Development Flyweight
	fmt.Println("Game Development Flyweight:")
	treeFactory := NewTreeFactory()
	
	// Create forest with many trees
	trees := []*Tree{
		NewTree(treeFactory.GetTreeType("oak"), 10, 20, "spring"),
		NewTree(treeFactory.GetTreeType("oak"), 15, 25, "summer"),
		NewTree(treeFactory.GetTreeType("pine"), 20, 30, "winter"),
		NewTree(treeFactory.GetTreeType("pine"), 25, 35, "autumn"),
		NewTree(treeFactory.GetTreeType("maple"), 30, 40, "spring"),
		NewTree(treeFactory.GetTreeType("oak"), 35, 45, "summer"),
	}
	
	for _, tree := range trees {
		tree.Render()
	}
	
	fmt.Printf("Total unique tree types: %d\n", treeFactory.GetTreeTypeCount())
	fmt.Println()

	// Database Connection Pool Flyweight
	fmt.Println("Database Connection Pool Flyweight:")
	connectionPool := NewConnectionPool(3)
	
	// Use connections
	conn1 := connectionPool.GetConnection("localhost", 5432, "database1")
	conn2 := connectionPool.GetConnection("localhost", 5432, "database1")
	conn3 := connectionPool.GetConnection("localhost", 5432, "database2")
	conn4 := connectionPool.GetConnection("localhost", 5432, "database1")
	
	conn1.Execute("SELECT * FROM users")
	conn2.Execute("SELECT * FROM products")
	conn3.Execute("SELECT * FROM orders")
	conn4.Execute("SELECT * FROM categories")
	
	fmt.Printf("Total connections in pool: %d\n", connectionPool.GetConnectionCount())
	connectionPool.CloseAll()
	fmt.Println()

	// GUI Icon Flyweight
	fmt.Println("GUI Icon Flyweight:")
	iconFactory := NewIconFactory()
	
	// Create UI with many icons
	icons := []struct {
		name  string
		x, y  int
		size  int
		color string
	}{
		{"home", 10, 10, 24, "blue"},
		{"home", 50, 10, 24, "red"},
		{"settings", 90, 10, 24, "green"},
		{"user", 130, 10, 24, "purple"},
		{"search", 170, 10, 24, "orange"},
		{"home", 210, 10, 32, "blue"},
		{"settings", 250, 10, 32, "green"},
	}
	
	for _, iconData := range icons {
		icon := iconFactory.GetIcon(iconData.name)
		icon.Render(iconData.x, iconData.y, iconData.size, iconData.color)
	}
	
	fmt.Printf("Total unique icons: %d\n", iconFactory.GetIconCount())
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
