package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC PROTOTYPE PATTERN
// =============================================================================

// Prototype interface
type Prototype interface {
	Clone() Prototype
	GetType() string
	GetInfo() string
}

// Concrete Prototype 1 - User
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
	Settings  map[string]interface{}
	Roles     []string
}

func (u *User) Clone() Prototype {
	// Deep copy implementation
	clonedUser := &User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		Settings:  make(map[string]interface{}),
		Roles:     make([]string, len(u.Roles)),
	}
	
	// Deep copy settings map
	for k, v := range u.Settings {
		clonedUser.Settings[k] = v
	}
	
	// Deep copy roles slice
	copy(clonedUser.Roles, u.Roles)
	
	return clonedUser
}

func (u *User) GetType() string {
	return "User"
}

func (u *User) GetInfo() string {
	return fmt.Sprintf("User{ID: %d, Name: %s, Email: %s, Roles: %v}", 
		u.ID, u.Name, u.Email, u.Roles)
}

// Concrete Prototype 2 - Product
type Product struct {
	ID          int
	Name        string
	Price       float64
	Category    string
	Tags        []string
	Attributes  map[string]string
	CreatedAt   time.Time
}

func (p *Product) Clone() Prototype {
	// Deep copy implementation
	clonedProduct := &Product{
		ID:         p.ID,
		Name:       p.Name,
		Price:      p.Price,
		Category:   p.Category,
		Tags:       make([]string, len(p.Tags)),
		Attributes: make(map[string]string),
		CreatedAt:  p.CreatedAt,
	}
	
	// Deep copy tags slice
	copy(clonedProduct.Tags, p.Tags)
	
	// Deep copy attributes map
	for k, v := range p.Attributes {
		clonedProduct.Attributes[k] = v
	}
	
	return clonedProduct
}

func (p *Product) GetType() string {
	return "Product"
}

func (p *Product) GetInfo() string {
	return fmt.Sprintf("Product{ID: %d, Name: %s, Price: $%.2f, Category: %s}", 
		p.ID, p.Name, p.Price, p.Category)
}

// =============================================================================
// PROTOTYPE REGISTRY PATTERN
// =============================================================================

// Prototype Registry
type PrototypeRegistry struct {
	prototypes map[string]Prototype
}

func NewPrototypeRegistry() *PrototypeRegistry {
	return &PrototypeRegistry{
		prototypes: make(map[string]Prototype),
	}
}

func (pr *PrototypeRegistry) Register(name string, prototype Prototype) {
	pr.prototypes[name] = prototype
}

func (pr *PrototypeRegistry) Unregister(name string) {
	delete(pr.prototypes, name)
}

func (pr *PrototypeRegistry) Get(name string) Prototype {
	if prototype, exists := pr.prototypes[name]; exists {
		return prototype.Clone()
	}
	return nil
}

func (pr *PrototypeRegistry) List() []string {
	var names []string
	for name := range pr.prototypes {
		names = append(names, name)
	}
	return names
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. DATABASE RECORD PROTOTYPE
type DatabaseRecord struct {
	TableName string
	Fields    map[string]interface{}
	Indexes   []string
	CreatedAt time.Time
}

func (dr *DatabaseRecord) Clone() Prototype {
	clonedRecord := &DatabaseRecord{
		TableName: dr.TableName,
		Fields:    make(map[string]interface{}),
		Indexes:   make([]string, len(dr.Indexes)),
		CreatedAt: dr.CreatedAt,
	}
	
	// Deep copy fields
	for k, v := range dr.Fields {
		clonedRecord.Fields[k] = v
	}
	
	// Deep copy indexes
	copy(clonedRecord.Indexes, dr.Indexes)
	
	return clonedRecord
}

func (dr *DatabaseRecord) GetType() string {
	return "DatabaseRecord"
}

func (dr *DatabaseRecord) GetInfo() string {
	return fmt.Sprintf("DatabaseRecord{Table: %s, Fields: %v}", dr.TableName, dr.Fields)
}

// 2. UI COMPONENT PROTOTYPE
type UIComponent struct {
	Type       string
	Properties map[string]interface{}
	Children   []*UIComponent
	Style      map[string]string
}

func (uc *UIComponent) Clone() Prototype {
	clonedComponent := &UIComponent{
		Type:       uc.Type,
		Properties: make(map[string]interface{}),
		Children:   make([]*UIComponent, len(uc.Children)),
		Style:      make(map[string]string),
	}
	
	// Deep copy properties
	for k, v := range uc.Properties {
		clonedComponent.Properties[k] = v
	}
	
	// Deep copy children (recursive)
	for i, child := range uc.Children {
		clonedComponent.Children[i] = child.Clone().(*UIComponent)
	}
	
	// Deep copy style
	for k, v := range uc.Style {
		clonedComponent.Style[k] = v
	}
	
	return clonedComponent
}

func (uc *UIComponent) GetType() string {
	return "UIComponent"
}

func (uc *UIComponent) GetInfo() string {
	return fmt.Sprintf("UIComponent{Type: %s, Properties: %v}", uc.Type, uc.Properties)
}

// 3. GAME OBJECT PROTOTYPE
type GameObject struct {
	ID          string
	Type        string
	Position    struct{ X, Y, Z float64 }
	Rotation    struct{ X, Y, Z float64 }
	Scale       struct{ X, Y, Z float64 }
	Components  map[string]interface{}
	Children    []*GameObject
}

func (go *GameObject) Clone() Prototype {
	clonedObject := &GameObject{
		ID:         go.ID,
		Type:       go.Type,
		Position:   go.Position,
		Rotation:   go.Rotation,
		Scale:      go.Scale,
		Components: make(map[string]interface{}),
		Children:   make([]*GameObject, len(go.Children)),
	}
	
	// Deep copy components
	for k, v := range go.Components {
		clonedObject.Components[k] = v
	}
	
	// Deep copy children (recursive)
	for i, child := range go.Children {
		clonedObject.Children[i] = child.Clone().(*GameObject)
	}
	
	return clonedObject
}

func (go *GameObject) GetType() string {
	return "GameObject"
}

func (go *GameObject) GetInfo() string {
	return fmt.Sprintf("GameObject{ID: %s, Type: %s, Position: (%.2f, %.2f, %.2f)}", 
		go.ID, go.Type, go.Position.X, go.Position.Y, go.Position.Z)
}

// =============================================================================
// SHALLOW VS DEEP COPY DEMONSTRATION
// =============================================================================

type ShallowCopyExample struct {
	Name    string
	Numbers []int
	Data    map[string]string
}

func (sce *ShallowCopyExample) ShallowClone() *ShallowCopyExample {
	return &ShallowCopyExample{
		Name:    sce.Name,
		Numbers: sce.Numbers, // Shared reference
		Data:    sce.Data,    // Shared reference
	}
}

func (sce *ShallowCopyExample) DeepClone() *ShallowCopyExample {
	cloned := &ShallowCopyExample{
		Name:    sce.Name,
		Numbers: make([]int, len(sce.Numbers)),
		Data:    make(map[string]string),
	}
	
	// Deep copy slice
	copy(cloned.Numbers, sce.Numbers)
	
	// Deep copy map
	for k, v := range sce.Data {
		cloned.Data[k] = v
	}
	
	return cloned
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== PROTOTYPE PATTERN DEMONSTRATION ===\n")

	// 1. BASIC PROTOTYPE
	fmt.Println("1. BASIC PROTOTYPE:")
	
	// Create original user
	originalUser := &User{
		ID:   1,
		Name: "John Doe",
		Email: "john@example.com",
		CreatedAt: time.Now(),
		Settings: map[string]interface{}{
			"theme": "dark",
			"lang":  "en",
		},
		Roles: []string{"admin", "user"},
	}
	
	fmt.Printf("Original User: %s\n", originalUser.GetInfo())
	
	// Clone user
	clonedUser := originalUser.Clone().(*User)
	clonedUser.ID = 2
	clonedUser.Name = "Jane Doe"
	clonedUser.Email = "jane@example.com"
	clonedUser.Roles = append(clonedUser.Roles, "moderator")
	
	fmt.Printf("Cloned User: %s\n", clonedUser.GetInfo())
	fmt.Printf("Original User (unchanged): %s\n", originalUser.GetInfo())
	fmt.Println()

	// 2. PROTOTYPE REGISTRY
	fmt.Println("2. PROTOTYPE REGISTRY:")
	
	registry := NewPrototypeRegistry()
	
	// Register prototypes
	registry.Register("admin_user", &User{
		ID:    0,
		Name:  "Admin Template",
		Email: "admin@template.com",
		CreatedAt: time.Now(),
		Settings: map[string]interface{}{
			"theme": "dark",
			"lang":  "en",
		},
		Roles: []string{"admin"},
	})
	
	registry.Register("product_template", &Product{
		ID:       0,
		Name:     "Product Template",
		Price:    0.0,
		Category: "General",
		Tags:     []string{"template"},
		Attributes: map[string]string{
			"status": "active",
		},
		CreatedAt: time.Now(),
	})
	
	// Create objects from registry
	adminUser := registry.Get("admin_user").(*User)
	adminUser.ID = 100
	adminUser.Name = "New Admin"
	adminUser.Email = "newadmin@example.com"
	
	product := registry.Get("product_template").(*Product)
	product.ID = 200
	product.Name = "New Product"
	product.Price = 99.99
	product.Category = "Electronics"
	
	fmt.Printf("Created from registry - Admin: %s\n", adminUser.GetInfo())
	fmt.Printf("Created from registry - Product: %s\n", product.GetInfo())
	fmt.Printf("Available prototypes: %v\n", registry.List())
	fmt.Println()

	// 3. REAL-WORLD EXAMPLES
	fmt.Println("3. REAL-WORLD EXAMPLES:")

	// Database Record
	fmt.Println("Database Record Prototype:")
	userRecord := &DatabaseRecord{
		TableName: "users",
		Fields: map[string]interface{}{
			"id":    "INT PRIMARY KEY",
			"name":  "VARCHAR(100)",
			"email": "VARCHAR(255)",
		},
		Indexes: []string{"idx_email", "idx_name"},
		CreatedAt: time.Now(),
	}
	
	clonedRecord := userRecord.Clone().(*DatabaseRecord)
	clonedRecord.TableName = "customers"
	clonedRecord.Fields["phone"] = "VARCHAR(20)"
	
	fmt.Printf("Original Record: %s\n", userRecord.GetInfo())
	fmt.Printf("Cloned Record: %s\n", clonedRecord.GetInfo())
	fmt.Println()

	// UI Component
	fmt.Println("UI Component Prototype:")
	buttonComponent := &UIComponent{
		Type: "Button",
		Properties: map[string]interface{}{
			"text":    "Click Me",
			"onClick": "handleClick",
		},
		Style: map[string]string{
			"backgroundColor": "blue",
			"color":           "white",
		},
		Children: []*UIComponent{},
	}
	
	clonedButton := buttonComponent.Clone().(*UIComponent)
	clonedButton.Properties["text"] = "Submit"
	clonedButton.Style["backgroundColor"] = "green"
	
	fmt.Printf("Original Button: %s\n", buttonComponent.GetInfo())
	fmt.Printf("Cloned Button: %s\n", clonedButton.GetInfo())
	fmt.Println()

	// 4. SHALLOW VS DEEP COPY
	fmt.Println("4. SHALLOW VS DEEP COPY:")
	
	original := &ShallowCopyExample{
		Name:    "Original",
		Numbers: []int{1, 2, 3},
		Data:    map[string]string{"key": "value"},
	}
	
	shallowCopy := original.ShallowClone()
	deepCopy := original.DeepClone()
	
	// Modify original
	original.Numbers[0] = 999
	original.Data["key"] = "modified"
	
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Shallow Copy: %+v\n", shallowCopy)
	fmt.Printf("Deep Copy: %+v\n", deepCopy)
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
