package main

import "fmt"

var pl = fmt.Println

type customer struct {
	name    string
	address string
	bal     float64
}

// Customer passed as values
func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

type contact struct {
	fName string
	lName string
	phone string
}

// Struct composition : Putting a struct in another
type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}

// ----- DEFINED TYPES -----
// I'll define different cooking measurement types
type Tsp float64
type TBs float64
type ML float64

// Convert with functions (Bad Way)
func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

// Associate method with types
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}





func main() {

	// Structs allow you to store values with many data types

	// Add values
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 Main St"
	tS.bal = 234.56

	// Pass to function as values
	getCustInfo(tS)
	// or as reference
	newCustAdd(&tS, "123 South st")
	pl("Address :", tS.address)

	// Create a struct literal
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name :", sS.name)

	// Structs with functions
	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area :", rect1.Area())

	// Go doesn't support inheritance, but it does support composition by embedding a struct in another
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}

	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}

	bus1.info()

	// ----- DEFINED TYPES -----
	// You can use them also to enhance the quality of other data types
	// We'll create them for different measurements

	// Convert from tsp to mL
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f mL\n", ml1)

	// Convert from TBs to mL
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f mL\n", ml2)

	// You can use arithmetic and comparison
	// operators
	pl("2 tsp + 4 tsp =", Tsp(2), Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))

	// We can convert with functions
	// Bad Way
	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))
	fmt.Printf("3 TBs = %.2f mL\n", TBToML(3))

	// We can solve this by using methods which
	// are functions associated with a type
	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())

	// ----- PROTECTING DATA -----
	// We want to protect our data from receiving bad values by moving our date struct to another package using encapsulation
}


// A struct is a collection of fields, like a lightweight object.
type User struct {
    ID    int
    Name  string
    Email string
}

// Struct Methods (Value vs Pointer Receivers)
// Methods can be attached to structs.
// Value receiver → works on a copy.
// Pointer receiver → works on original.
type Counter struct {
    value int
}

func (c Counter) IncrementByCopy() {
    c.value++ // modifies copy
}

func (c *Counter) IncrementByPointer() {
    c.value++ // modifies original
}

// “When should you use pointer vs value receiver?”
// Use pointer receiver if:
// Method modifies receiver
// Struct is large (avoid copying)
// Use value receiver if:
// Struct is small and immutable (e.g., time.Time, string wrappers)

// Struct Embedding (Composition > Inheritance)
// Go doesn’t have inheritance, but embedding lets you compose behaviors.	

type Person struct {
    Name string
}

type Employee struct {
    Person  // embedded
    Role    string
}

func main() {

	u1 := User{ID: 1, Name: "Alice", Email: "a@example.com"} // full init
	u2 := User{Name: "Bob"} // partial init
	u3 := new(User)         // pointer to zero-value struct
	u4 := &User{ID: 2}      // pointer to struct literal
	fmt.Println(u1);
	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println(u4)

	c := Counter{}
    c.IncrementByCopy()
    fmt.Println(c.value) // 0
    c.IncrementByPointer()
    fmt.Println(c.value) // 1

	e := Employee{Person: Person{Name: "Alice"}, Role: "Engineer"}
    fmt.Println(e.Name)  // promoted field → "Alice"

}

// Person represents a basic struct
type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

// Address represents an embedded struct
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

// Employee represents a struct with methods
type Employee struct {
	Person
	ID        int
	Position  string
	Salary    float64
	StartDate time.Time
	IsActive  bool
}

// Rectangle represents a struct for demonstrating methods
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle represents another shape
type Circle struct {
	Radius float64
}

// Shape interface for polymorphism
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// DemonstrateBasicStructs shows basic struct operations
func DemonstrateBasicStructs() {
	fmt.Println("=== BASIC STRUCTS ===")
	
	// 1. Struct Declaration and Initialization
	fmt.Println("\n--- Struct Declaration ---")
	
	// Zero-initialized struct
	var person1 Person
	fmt.Printf("Zero-initialized person: %+v\n", person1)
	
	// Struct literal
	person2 := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
			ZipCode: "10001",
		},
	}
	fmt.Printf("Initialized person: %+v\n", person2)
	
	// Partial initialization
	person3 := Person{
		Name: "Jane Smith",
		Age:  25,
		// Email and Address will be zero values
	}
	fmt.Printf("Partially initialized person: %+v\n", person3)
	
	// 2. Field Access and Modification
	fmt.Println("\n--- Field Access ---")
	
	person := Person{
		Name:  "Alice Johnson",
		Age:   28,
		Email: "alice@example.com",
	}
	
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)
	
	// Modify fields
	person.Age = 29
	person.Email = "alice.johnson@example.com"
	fmt.Printf("After modification: %+v\n", person)
	
	// 3. Anonymous Structs
	fmt.Println("\n--- Anonymous Structs ---")
	
	// Inline struct definition
	config := struct {
		DatabaseURL string
		Port        int
		Debug       bool
	}{
		DatabaseURL: "localhost:5432",
		Port:        8080,
		Debug:       true,
	}
	
	fmt.Printf("Config: %+v\n", config)
}

// DemonstrateStructMethods shows struct methods
func DemonstrateStructMethods() {
	fmt.Println("\n=== STRUCT METHODS ===")
	
	// 1. Value Receiver Methods
	fmt.Println("\n--- Value Receiver Methods ---")
	
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
	fmt.Printf("String: %s\n", rect.String())
	
	// 2. Pointer Receiver Methods
	fmt.Println("\n--- Pointer Receiver Methods ---")
	
	employee := Employee{
		Person: Person{
			Name:  "Bob Wilson",
			Age:   35,
			Email: "bob@company.com",
		},
		ID:        1001,
		Position:  "Software Engineer",
		Salary:    75000,
		StartDate: time.Now().AddDate(-2, 0, 0), // 2 years ago
		IsActive:  true,
	}
	
	fmt.Printf("Employee before promotion: %+v\n", employee)
	employee.Promote("Senior Software Engineer", 85000)
	fmt.Printf("Employee after promotion: %+v\n", employee)
	
	// 3. Method Chaining
	fmt.Println("\n--- Method Chaining ---")
	
	rect2 := Rectangle{Width: 3, Height: 4}
	rect2.Scale(2).Scale(1.5)
	fmt.Printf("Scaled rectangle: %+v\n", rect2)
}

// DemonstrateStructEmbedding shows struct embedding (composition)
func DemonstrateStructEmbedding() {
	fmt.Println("\n=== STRUCT EMBEDDING ===")
	
	// 1. Basic Embedding
	fmt.Println("\n--- Basic Embedding ---")
	
	employee := Employee{
		Person: Person{
			Name:  "Charlie Brown",
			Age:   40,
			Email: "charlie@company.com",
			Address: Address{
				Street:  "456 Oak Ave",
				City:    "San Francisco",
				Country: "USA",
				ZipCode: "94102",
			},
		},
		ID:        1002,
		Position:  "Manager",
		Salary:    95000,
		StartDate: time.Now().AddDate(-5, 0, 0),
		IsActive:  true,
	}
	
	// Access embedded struct fields directly
	fmt.Printf("Employee Name: %s\n", employee.Name)        // From Person
	fmt.Printf("Employee Age: %d\n", employee.Age)          // From Person
	fmt.Printf("Employee ID: %d\n", employee.ID)            // From Employee
	fmt.Printf("Employee Position: %s\n", employee.Position) // From Employee
	
	// Access embedded struct fields explicitly
	fmt.Printf("Employee City: %s\n", employee.Person.Address.City)
	
	// 2. Method Promotion
	fmt.Println("\n--- Method Promotion ---")
	
	// Methods from embedded structs are promoted
	fmt.Printf("Employee Info: %s\n", employee.GetInfo())
	
	// 3. Field Shadowing
	fmt.Println("\n--- Field Shadowing ---")
	
	type Manager struct {
		Person
		Name string // This shadows Person.Name
	}
	
	manager := Manager{
		Person: Person{Name: "John Manager"},
		Name:   "Jane Manager",
	}
	
	fmt.Printf("Manager.Person.Name: %s\n", manager.Person.Name)
	fmt.Printf("Manager.Name: %s\n", manager.Name)
}

// DemonstrateStructTags shows struct tags for metadata
func DemonstrateStructTags() {
	fmt.Println("\n=== STRUCT TAGS ===")
	
	// Struct with JSON tags
	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email,omitempty"`
		Password string `json:"-"` // Ignore in JSON
		Age      int    `json:"age,string"` // Convert to string in JSON
	}
	
	user := User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@example.com",
		Password: "secret123",
		Age:      30,
	}
	
	// Convert to JSON
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	
	fmt.Printf("User as JSON:\n%s\n", string(jsonData))
	
	// Convert from JSON
	jsonString := `{"id":2,"username":"janedoe","email":"jane@example.com","age":"25"}`
	var user2 User
	err = json.Unmarshal([]byte(jsonString), &user2)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}
	
	fmt.Printf("User from JSON: %+v\n", user2)
}

// DemonstrateStructComparison shows struct comparison
func DemonstrateStructComparison() {
	fmt.Println("\n=== STRUCT COMPARISON ===")
	
	// Structs with comparable fields can be compared
	point1 := struct {
		X, Y int
	}{1, 2}
	
	point2 := struct {
		X, Y int
	}{1, 2}
	
	point3 := struct {
		X, Y int
	}{2, 3}
	
	fmt.Printf("point1 == point2: %t\n", point1 == point2)
	fmt.Printf("point1 == point3: %t\n", point1 == point3)
	
	// Structs with non-comparable fields cannot be compared
	// (This would cause a compile error)
	// person1 := Person{Name: "John"}
	// person2 := Person{Name: "John"}
	// fmt.Printf("person1 == person2: %t\n", person1 == person2) // Error!
}

// Rectangle methods (Value receivers)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Width: %.2f, Height: %.2f}", r.Width, r.Height)
}

func (r Rectangle) Scale(factor float64) Rectangle {
	return Rectangle{
		Width:  r.Width * factor,
		Height: r.Height * factor,
	}
}

// Circle methods
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// Employee methods (Pointer receivers)
func (e *Employee) Promote(newPosition string, newSalary float64) {
	e.Position = newPosition
	e.Salary = newSalary
}

func (e *Employee) GetInfo() string {
	return fmt.Sprintf("%s (%s) - %s", e.Name, e.Position, e.Email)
}

// Person methods
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s, %d years old", p.Name, p.Age)
}

// DemonstratePolymorphism shows interface implementation
func DemonstratePolymorphism() {
	fmt.Println("\n=== POLYMORPHISM WITH INTERFACES ===")
	
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 4},
		Circle{Radius: 2.5},
	}
	
	for i, shape := range shapes {
		fmt.Printf("Shape %d: %s\n", i+1, shape.String())
		fmt.Printf("  Area: %.2f\n", shape.Area())
		fmt.Printf("  Perimeter: %.2f\n", shape.Perimeter())
		fmt.Println()
	}
}

// RunAllStructExamples runs all struct examples
func RunAllStructExamples() {
	DemonstrateBasicStructs()
	DemonstrateStructMethods()
	DemonstrateStructEmbedding()
	DemonstrateStructTags()
	DemonstrateStructComparison()
	DemonstratePolymorphism()
}


// Basic struct types
type Person struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Age       int       `json:"age" db:"age"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Embedded structs
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type Contact struct {
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
}

// Anonymous structs
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Contact  struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
	} `json:"contact"`
}

// Nested structs
type Company struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CEO         Person    `json:"ceo"`
	Employees   []Person  `json:"employees"`
	Departments []Dept    `json:"departments"`
	Founded     time.Time `json:"founded"`
}

type Dept struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Manager     Person   `json:"manager"`
	Members     []Person `json:"members"`
	Budget      float64  `json:"budget"`
	IsActive    bool     `json:"is_active"`
}

// Struct with methods
type BankAccount struct {
	AccountNumber string    `json:"account_number"`
	Holder        Person    `json:"holder"`
	Balance       float64   `json:"balance"`
	Currency      string    `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
	Transactions  []BankTransaction `json:"transactions"`
}

type BankTransaction struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"` // "deposit", "withdrawal", "transfer"
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

// Struct with pointer fields
type Node struct {
	Value    int    `json:"value"`
	Data     string `json:"data"`
	Parent   *Node  `json:"parent,omitempty"`
	Children []*Node `json:"children,omitempty"`
}

// Generic struct (Go 1.18+)
type Container[T any] struct {
	Items []T `json:"items"`
	Count int `json:"count"`
}

// Struct with validation
type Product struct {
	ID          int     `json:"id" validate:"required,min=1"`
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Description string  `json:"description" validate:"max=500"`
	InStock     bool    `json:"in_stock"`
	Category    string  `json:"category" validate:"required"`
	Tags        []string `json:"tags"`
}

// StructManager handles CRUD operations for structs
type StructManager struct {
	People      []Person      `json:"people"`
	Companies   []Company     `json:"companies"`
	Accounts    []BankAccount `json:"accounts"`
	Products    []Product     `json:"products"`
	Nodes       []*Node       `json:"nodes"`
	Containers  map[string]interface{} `json:"containers"`
}

// NewStructManager creates a new struct manager
func NewStructManager() *StructManager {
	return &StructManager{
		People:     make([]Person, 0),
		Companies:  make([]Company, 0),
		Accounts:   make([]BankAccount, 0),
		Products:   make([]Product, 0),
		Nodes:      make([]*Node, 0),
		Containers: make(map[string]interface{}),
	}
}

// CRUD Operations for Structs

// Create - Initialize and create struct instances
func (sm *StructManager) Create() {
	fmt.Println("🔧 Creating struct instances...")
	
	// Create Person instances
	person1 := Person{
		ID:        1,
		Name:      "Alice Johnson",
		Email:     "alice@example.com",
		Age:       30,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	person2 := Person{
		ID:        2,
		Name:      "Bob Smith",
		Email:     "bob@example.com",
		Age:       25,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Create with pointer
	person3 := &Person{
		ID:        3,
		Name:      "Charlie Brown",
		Email:     "charlie@example.com",
		Age:       35,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	sm.People = append(sm.People, person1, person2, *person3)
	
	// Create Company with nested structs
	company := Company{
		ID:   1,
		Name: "TechCorp Inc.",
		CEO:  person1,
		Employees: []Person{person1, person2, *person3},
		Departments: []Dept{
			{
				ID:       1,
				Name:     "Engineering",
				Manager:  person1,
				Members:  []Person{person2, *person3},
				Budget:   1000000.0,
				IsActive: true,
			},
		},
		Founded: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	
	sm.Companies = append(sm.Companies, company)
	
	// Create BankAccount with transactions
	account := BankAccount{
		AccountNumber: "ACC-001",
		Holder:        person1,
		Balance:       5000.0,
		Currency:      "USD",
		CreatedAt:     time.Now(),
		Transactions: []BankTransaction{
			{
				ID:          "TXN-001",
				Type:        "deposit",
				Amount:      5000.0,
				Description: "Initial deposit",
				Timestamp:   time.Now(),
			},
		},
	}
	
	sm.Accounts = append(sm.Accounts, account)
	
	// Create Product
	product := Product{
		ID:          1,
		Name:        "Laptop Pro",
		Price:       1299.99,
		Description: "High-performance laptop for professionals",
		InStock:     true,
		Category:    "Electronics",
		Tags:        []string{"laptop", "computer", "professional"},
	}
	
	sm.Products = append(sm.Products, product)
	
	// Create Node tree
	root := &Node{
		Value: 1,
		Data:  "Root",
	}
	
	child1 := &Node{
		Value:  2,
		Data:   "Child 1",
		Parent: root,
	}
	
	child2 := &Node{
		Value:  3,
		Data:   "Child 2",
		Parent: root,
	}
	
	root.Children = []*Node{child1, child2}
	sm.Nodes = append(sm.Nodes, root)
	
	// Create generic containers
	intContainer := Container[int]{
		Items: []int{1, 2, 3, 4, 5},
		Count: 5,
	}
	
	stringContainer := Container[string]{
		Items: []string{"apple", "banana", "cherry"},
		Count: 3,
	}
	
	sm.Containers["integers"] = intContainer
	sm.Containers["strings"] = stringContainer
	
	fmt.Println("✅ Struct instances created successfully")
}

// Read - Display struct information
func (sm *StructManager) Read() {
	fmt.Println("\n📖 READING STRUCT DATA:")
	fmt.Println("=======================")
	
	// Read People
	fmt.Printf("People (%d):\n", len(sm.People))
	for i, person := range sm.People {
		fmt.Printf("  %d: %+v\n", i+1, person)
	}
	
	// Read Companies
	fmt.Printf("\nCompanies (%d):\n", len(sm.Companies))
	for i, company := range sm.Companies {
		fmt.Printf("  %d: %s (CEO: %s, Employees: %d)\n", 
			i+1, company.Name, company.CEO.Name, len(company.Employees))
	}
	
	// Read Accounts
	fmt.Printf("\nBank Accounts (%d):\n", len(sm.Accounts))
	for i, account := range sm.Accounts {
		fmt.Printf("  %d: %s (Holder: %s, Balance: %.2f %s)\n", 
			i+1, account.AccountNumber, account.Holder.Name, account.Balance, account.Currency)
	}
	
	// Read Products
	fmt.Printf("\nProducts (%d):\n", len(sm.Products))
	for i, product := range sm.Products {
		fmt.Printf("  %d: %s ($%.2f, In Stock: %t)\n", 
			i+1, product.Name, product.Price, product.InStock)
	}
	
	// Read Nodes (tree structure)
	fmt.Printf("\nNode Tree:\n")
	for i, node := range sm.Nodes {
		sm.printNode(node, 0, i+1)
	}
	
	// Read Containers
	fmt.Printf("\nContainers:\n")
	for name, container := range sm.Containers {
		fmt.Printf("  %s: %+v\n", name, container)
	}
}

// printNode recursively prints node tree
func (sm *StructManager) printNode(node *Node, depth int, index int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	
	fmt.Printf("%s%d. %s (Value: %d)\n", indent, index, node.Data, node.Value)
	
	for i, child := range node.Children {
		sm.printNode(child, depth+1, i+1)
	}
}

// Update - Modify struct instances
func (sm *StructManager) Update() {
	fmt.Println("\n🔄 UPDATING STRUCT DATA:")
	fmt.Println("========================")
	
	// Update Person
	if len(sm.People) > 0 {
		sm.People[0].Name = "Alice Johnson-Updated"
		sm.People[0].Age = 31
		sm.People[0].UpdatedAt = time.Now()
		fmt.Printf("Updated person: %+v\n", sm.People[0])
	}
	
	// Update Company
	if len(sm.Companies) > 0 {
		sm.Companies[0].Name = "TechCorp Inc. - Updated"
		sm.Companies[0].Departments[0].Budget = 1500000.0
		fmt.Printf("Updated company: %s\n", sm.Companies[0].Name)
	}
	
	// Update BankAccount
	if len(sm.Accounts) > 0 {
		sm.Accounts[0].Balance += 1000.0
		sm.Accounts[0].Transactions = append(sm.Accounts[0].Transactions, BankTransaction{
			ID:          "TXN-002",
			Type:        "deposit",
			Amount:      1000.0,
			Description: "Salary deposit",
			Timestamp:   time.Now(),
		})
		fmt.Printf("Updated account balance: %.2f\n", sm.Accounts[0].Balance)
	}
	
	// Update Product
	if len(sm.Products) > 0 {
		sm.Products[0].Price *= 0.9 // 10% discount
		sm.Products[0].Tags = append(sm.Products[0].Tags, "sale")
		fmt.Printf("Updated product price: %.2f\n", sm.Products[0].Price)
	}
	
	// Update Node
	if len(sm.Nodes) > 0 && len(sm.Nodes[0].Children) > 0 {
		sm.Nodes[0].Children[0].Value = 99
		sm.Nodes[0].Children[0].Data = "Updated Child 1"
		fmt.Printf("Updated node: %+v\n", sm.Nodes[0].Children[0])
	}
	
	fmt.Println("✅ Struct data updated successfully")
}

// Delete - Remove struct instances
func (sm *StructManager) Delete() {
	fmt.Println("\n🗑️  DELETING STRUCT DATA:")
	fmt.Println("=========================")
	
	// Delete Person (remove last person)
	if len(sm.People) > 0 {
		deleted := sm.People[len(sm.People)-1]
		sm.People = sm.People[:len(sm.People)-1]
		fmt.Printf("Deleted person: %s\n", deleted.Name)
	}
	
	// Delete Company
	if len(sm.Companies) > 0 {
		deleted := sm.Companies[0]
		sm.Companies = sm.Companies[1:]
		fmt.Printf("Deleted company: %s\n", deleted.Name)
	}
	
	// Delete Account
	if len(sm.Accounts) > 0 {
		deleted := sm.Accounts[0]
		sm.Accounts = sm.Accounts[1:]
		fmt.Printf("Deleted account: %s\n", deleted.AccountNumber)
	}
	
	// Delete Product
	if len(sm.Products) > 0 {
		deleted := sm.Products[0]
		sm.Products = sm.Products[1:]
		fmt.Printf("Deleted product: %s\n", deleted.Name)
	}
	
	// Delete Node
	if len(sm.Nodes) > 0 {
		deleted := sm.Nodes[0]
		sm.Nodes = sm.Nodes[1:]
		fmt.Printf("Deleted node: %s\n", deleted.Data)
	}
	
	// Clear containers
	sm.Containers = make(map[string]interface{})
	fmt.Println("Cleared all containers")
	
	fmt.Println("✅ Struct data deleted successfully")
}

// Struct Methods and Behaviors

// DemonstrateStructMethods shows struct methods
func (sm *StructManager) DemonstrateStructMethods() {
	fmt.Println("\n🔧 STRUCT METHODS DEMONSTRATION:")
	fmt.Println("===============================")
	
	// Create a bank account for demonstration
	account := BankAccount{
		AccountNumber: "DEMO-001",
		Holder: Person{
			ID:    1,
			Name:  "Demo User",
			Email: "demo@example.com",
			Age:   25,
		},
		Balance:     1000.0,
		Currency:    "USD",
		CreatedAt:   time.Now(),
		Transactions: []BankTransaction{},
	}
	
	// Demonstrate methods
	fmt.Printf("Initial balance: %.2f\n", account.GetBalance())
	
	account.Deposit(500.0, "Demo deposit")
	fmt.Printf("After deposit: %.2f\n", account.GetBalance())
	
	account.Withdraw(200.0, "Demo withdrawal")
	fmt.Printf("After withdrawal: %.2f\n", account.GetBalance())
	
	account.Transfer(100.0, "Demo transfer")
	fmt.Printf("After transfer: %.2f\n", account.GetBalance())
	
	// Show transaction history
	fmt.Printf("Transaction history (%d transactions):\n", len(account.Transactions))
	for i, txn := range account.Transactions {
		fmt.Printf("  %d: %s - %.2f (%s)\n", i+1, txn.Type, txn.Amount, txn.Description)
	}
}

// BankAccount methods
func (ba *BankAccount) GetBalance() float64 {
	return ba.Balance
}

func (ba *BankAccount) Deposit(amount float64, description string) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	
	ba.Balance += amount
	ba.addTransaction("deposit", amount, description)
	return nil
}

func (ba *BankAccount) Withdraw(amount float64, description string) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	
	if amount > ba.Balance {
		return fmt.Errorf("insufficient funds")
	}
	
	ba.Balance -= amount
	ba.addTransaction("withdrawal", amount, description)
	return nil
}

func (ba *BankAccount) Transfer(amount float64, description string) error {
	if amount <= 0 {
		return fmt.Errorf("transfer amount must be positive")
	}
	
	if amount > ba.Balance {
		return fmt.Errorf("insufficient funds for transfer")
	}
	
	ba.Balance -= amount
	ba.addTransaction("transfer", amount, description)
	return nil
}

func (ba *BankAccount) addTransaction(txnType string, amount float64, description string) {
	txn := BankTransaction{
		ID:          fmt.Sprintf("TXN-%d", len(ba.Transactions)+1),
		Type:        txnType,
		Amount:      amount,
		Description: description,
		Timestamp:   time.Now(),
	}
	ba.Transactions = append(ba.Transactions, txn)
}

// DemonstrateStructComposition shows struct composition
func (sm *StructManager) DemonstrateStructComposition() {
	fmt.Println("\n🧩 STRUCT COMPOSITION DEMONSTRATION:")
	fmt.Println("===================================")
	
	// Create a person with contact information
	person := Person{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}
	
	// Create contact with embedded address
	contact := Contact{
		Phone: "+1-555-123-4567",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			State:   "NY",
			ZipCode: "10001",
			Country: "USA",
		},
	}
	
	// Create a user with anonymous struct
	user := User{
		ID:       1,
		Username: "johndoe",
		Contact: struct {
			Email string `json:"email"`
			Phone string `json:"phone"`
		}{
			Email: "john@example.com",
			Phone: "+1-555-123-4567",
		},
	}
	
	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Contact: %+v\n", contact)
	fmt.Printf("User: %+v\n", user)
	
	// Demonstrate field access
	fmt.Printf("Person email: %s\n", person.Email)
	fmt.Printf("Contact phone: %s\n", contact.Phone)
	fmt.Printf("Contact address: %s, %s, %s %s\n", 
		contact.Address.Street, contact.Address.City, 
		contact.Address.State, contact.Address.ZipCode)
	fmt.Printf("User contact email: %s\n", user.Contact.Email)
}

// DemonstrateStructTags shows struct tags usage
func (sm *StructManager) DemonstrateStructTags() {
	fmt.Println("\n🏷️  STRUCT TAGS DEMONSTRATION:")
	fmt.Println("=============================")
	
	// Create a person with tags
	person := Person{
		ID:        1,
		Name:      "Alice Johnson",
		Email:     "alice@example.com",
		Age:       30,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	// Demonstrate JSON marshaling
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Printf("JSON marshaling error: %v\n", err)
		return
	}
	
	fmt.Printf("JSON representation:\n%s\n", string(jsonData))
	
	// Demonstrate JSON unmarshaling
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Printf("JSON unmarshaling error: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled person: %+v\n", newPerson)
	
	// Demonstrate reflection on struct tags
	sm.demonstrateReflection(person)
}

// demonstrateReflection shows reflection on struct tags
func (sm *StructManager) demonstrateReflection(person Person) {
	fmt.Println("\nReflection on struct tags:")
	
	t := reflect.TypeOf(person)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field: %s\n", field.Name)
		fmt.Printf("  JSON tag: %s\n", field.Tag.Get("json"))
		fmt.Printf("  DB tag: %s\n", field.Tag.Get("db"))
		fmt.Printf("  Type: %s\n", field.Type)
		fmt.Println()
	}
}

// DemonstrateStructValidation shows struct validation
func (sm *StructManager) DemonstrateStructValidation() {
	fmt.Println("\n✅ STRUCT VALIDATION DEMONSTRATION:")
	fmt.Println("===================================")
	
	// Valid product
	validProduct := Product{
		ID:          1,
		Name:        "Valid Product",
		Price:       99.99,
		Description: "This is a valid product",
		InStock:     true,
		Category:    "Electronics",
		Tags:        []string{"valid", "test"},
	}
	
	// Invalid product (negative price)
	invalidProduct := Product{
		ID:          2,
		Name:        "Invalid Product",
		Price:       -10.0, // Invalid: negative price
		Description: "This product has invalid data",
		InStock:     true,
		Category:    "Electronics",
		Tags:        []string{"invalid", "test"},
	}
	
	// Validate products
	fmt.Printf("Valid product validation: %t\n", sm.validateProduct(validProduct))
	fmt.Printf("Invalid product validation: %t\n", sm.validateProduct(invalidProduct))
	
	// Demonstrate custom validation
	fmt.Printf("Valid product custom validation: %t\n", sm.customValidateProduct(validProduct))
	fmt.Printf("Invalid product custom validation: %t\n", sm.customValidateProduct(invalidProduct))
}

// validateProduct performs basic validation
func (sm *StructManager) validateProduct(product Product) bool {
	if product.ID <= 0 {
		return false
	}
	if len(product.Name) < 2 || len(product.Name) > 100 {
		return false
	}
	if product.Price < 0 {
		return false
	}
	if len(product.Category) == 0 {
		return false
	}
	return true
}

// customValidateProduct performs custom validation
func (sm *StructManager) customValidateProduct(product Product) bool {
	// Check if product name contains only valid characters
	for _, char := range product.Name {
		if !((char >= 'a' && char <= 'z') || 
			 (char >= 'A' && char <= 'Z') || 
			 (char >= '0' && char <= '9') || 
			 char == ' ' || char == '-') {
			return false
		}
	}
	
	// Check if price is reasonable
	if product.Price > 1000000 {
		return false
	}
	
	// Check if category is valid
	validCategories := []string{"Electronics", "Clothing", "Books", "Home", "Sports"}
	validCategory := false
	for _, cat := range validCategories {
		if product.Category == cat {
			validCategory = true
			break
		}
	}
	
	return validCategory
}

// Base struct for embedding demonstration
type Base struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Add method to Base struct
func (b *Base) UpdateTimestamp() {
	b.UpdatedAt = time.Now()
}

// DemonstrateStructEmbedding shows struct embedding
func (sm *StructManager) DemonstrateStructEmbedding() {
	fmt.Println("\n🔗 STRUCT EMBEDDING DEMONSTRATION:")
	fmt.Println("==================================")
	
	type User struct {
		Base
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	type Admin struct {
		User
		Permissions []string `json:"permissions"`
		IsActive    bool     `json:"is_active"`
	}
	
	// Create instances
	user := User{
		Base: Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:  "Regular User",
		Email: "user@example.com",
	}
	
	admin := Admin{
		User: User{
			Base: Base{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:  "Admin User",
			Email: "admin@example.com",
		},
		Permissions: []string{"read", "write", "delete"},
		IsActive:    true,
	}
	
	// Demonstrate field access
	fmt.Printf("User: %+v\n", user)
	fmt.Printf("User ID: %d\n", user.ID) // Access embedded field
	fmt.Printf("User CreatedAt: %s\n", user.CreatedAt)
	
	fmt.Printf("\nAdmin: %+v\n", admin)
	fmt.Printf("Admin ID: %d\n", admin.ID) // Access embedded field
	fmt.Printf("Admin Name: %s\n", admin.Name) // Access embedded field
	fmt.Printf("Admin Permissions: %v\n", admin.Permissions)
	
	// Demonstrate method promotion
	user.UpdateTimestamp()
	admin.UpdateTimestamp()
	
	fmt.Printf("User after update: %s\n", user.UpdatedAt.Format(time.RFC3339))
	fmt.Printf("Admin after update: %s\n", admin.UpdatedAt.Format(time.RFC3339))
}