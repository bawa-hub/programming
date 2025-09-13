package main

import (
	"fmt"
	"sort"
)

// 🔧 GENERICS MASTERY
// Understanding Go generics and type parameters (Go 1.18+)

func main() {
	fmt.Println("🔧 GENERICS MASTERY")
	fmt.Println("===================")

	// 1. Basic Generics
	fmt.Println("\n1. Basic Generics:")
	basicGenerics()

	// 2. Type Constraints
	fmt.Println("\n2. Type Constraints:")
	typeConstraints()

	// 3. Generic Functions
	fmt.Println("\n3. Generic Functions:")
	genericFunctions()

	// 4. Generic Types
	fmt.Println("\n4. Generic Types:")
	genericTypes()

	// 5. Generic Interfaces
	fmt.Println("\n5. Generic Interfaces:")
	genericInterfaces()

	// 6. Generic Methods
	fmt.Println("\n6. Generic Methods:")
	genericMethods()

	// 7. Advanced Generic Patterns
	fmt.Println("\n7. Advanced Generic Patterns:")
	advancedGenericPatterns()

	// 8. Generic Best Practices
	fmt.Println("\n8. Generic Best Practices:")
	genericBestPractices()
}

// BASIC GENERICS: Understanding basic generics
func basicGenerics() {
	fmt.Println("Understanding basic generics...")
	
	// Generic function with type parameter
	result1 := Max(10, 20)
	fmt.Printf("  📊 Max(10, 20) = %d\n", result1)
	
	result2 := Max(3.14, 2.71)
	fmt.Printf("  📊 Max(3.14, 2.71) = %.2f\n", result2)
	
	result3 := Max("hello", "world")
	fmt.Printf("  📊 Max(\"hello\", \"world\") = %s\n", result3)
	
	// Generic function with multiple type parameters
	pair := Pair[string, int]{First: "hello", Second: 42}
	fmt.Printf("  📊 Pair: %+v\n", pair)
	
	// Generic slice operations
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(x int) int { return x * 2 })
	fmt.Printf("  📊 Doubled: %v\n", doubled)
	
	strings := []string{"hello", "world", "golang"}
	lengths := Map(strings, func(s string) int { return len(s) })
	fmt.Printf("  📊 Lengths: %v\n", lengths)
}

// TYPE CONSTRAINTS: Understanding type constraints
func typeConstraints() {
	fmt.Println("Understanding type constraints...")
	
	// Numeric constraint
	result1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Printf("  📊 Sum of ints: %d\n", result1)
	
	result2 := Sum([]float64{1.1, 2.2, 3.3, 4.4, 5.5})
	fmt.Printf("  📊 Sum of floats: %.2f\n", result2)
	
	// Comparable constraint
	index1 := Find([]int{1, 2, 3, 4, 5}, 3)
	fmt.Printf("  📊 Find 3 in ints: index %d\n", index1)
	
	index2 := Find([]string{"a", "b", "c", "d", "e"}, "c")
	fmt.Printf("  📊 Find \"c\" in strings: index %d\n", index2)
	
	// Ordered constraint
	max1 := MaxOrdered(10, 20, 30, 5, 15)
	fmt.Printf("  📊 Max ordered: %d\n", max1)
	
	max2 := MaxOrdered("apple", "banana", "cherry")
	fmt.Printf("  📊 Max string: %s\n", max2)
}

// GENERIC FUNCTIONS: Understanding generic functions
func genericFunctions() {
	fmt.Println("Understanding generic functions...")
	
	// Map function
	numbers := []int{1, 2, 3, 4, 5}
	squared := Map(numbers, func(x int) int { return x * x })
	fmt.Printf("  📊 Squared: %v\n", squared)
	
	// Filter function
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("  📊 Evens: %v\n", evens)
	
	// Reduce function
	sum := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("  📊 Sum: %d\n", sum)
	
	// Generic sort
	names := []string{"charlie", "alice", "bob"}
	Sort(names)
	fmt.Printf("  📊 Sorted names: %v\n", names)
	
	// Generic reverse
	Reverse(numbers)
	fmt.Printf("  📊 Reversed numbers: %v\n", numbers)
}

// GENERIC TYPES: Understanding generic types
func genericTypes() {
	fmt.Println("Understanding generic types...")
	
	// Generic stack
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	fmt.Printf("  📊 Stack: %v\n", stack.items)
	
	popped, _ := stack.Pop()
	fmt.Printf("  📊 Popped: %d\n", popped)
	fmt.Printf("  📊 Stack after pop: %v\n", stack.items)
	
	// Generic queue
	queue := NewQueue[string]()
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	
	fmt.Printf("  📊 Queue: %v\n", queue.items)
	
	dequeued, _ := queue.Dequeue()
	fmt.Printf("  📊 Dequeued: %s\n", dequeued)
	fmt.Printf("  📊 Queue after dequeue: %v\n", queue.items)
	
	// Generic map
	genericMap := NewGenericMap[string, int]()
	genericMap.Set("one", 1)
	genericMap.Set("two", 2)
	genericMap.Set("three", 3)
	
	fmt.Printf("  📊 Generic map: %v\n", genericMap.data)
	
	value, exists := genericMap.Get("two")
	fmt.Printf("  📊 Get \"two\": %d (exists: %t)\n", value, exists)
}

// GENERIC INTERFACES: Understanding generic interfaces
func genericInterfaces() {
	fmt.Println("Understanding generic interfaces...")
	
	// Generic repository
	userRepo := NewUserRepository()
	userRepo.Save(&User{ID: "1", Name: "John"})
	userRepo.Save(&User{ID: "2", Name: "Jane"})
	
	user, _ := userRepo.FindByID("1")
	fmt.Printf("  📊 Found user: %s\n", user.Name)
	
	users := userRepo.FindAll()
	fmt.Printf("  📊 All users: %v\n", users)
	
	// Generic service
	service := NewUserService(userRepo)
	service.ProcessUser("1")
	
	// Generic validator
	validator := NewValidator[User]()
	validator.AddRule(func(u *User) error {
		if u.Name == "" {
			return fmt.Errorf("name is required")
		}
		return nil
	})
	
	validUser := &User{ID: "1", Name: "John"}
	invalidUser := &User{ID: "2", Name: ""}
	
	err1 := validator.Validate(validUser)
	fmt.Printf("  📊 Valid user: %v\n", err1)
	
	err2 := validator.Validate(invalidUser)
	fmt.Printf("  📊 Invalid user: %v\n", err2)
}

// GENERIC METHODS: Understanding generic methods
func genericMethods() {
	fmt.Println("Understanding generic methods...")
	
	// Generic container
	container := NewContainer[int]()
	container.Add(1)
	container.Add(2)
	container.Add(3)
	
	fmt.Printf("  📊 Container: %v\n", container.items)
	
	// Generic methods
	exists := container.Contains(2)
	fmt.Printf("  📊 Contains 2: %t\n", exists)
	
	count := container.Count()
	fmt.Printf("  📊 Count: %d\n", count)
	
	// Generic transformation
	transformed := container.Transform(func(x int) interface{} {
		return fmt.Sprintf("item-%d", x)
	})
	fmt.Printf("  📊 Transformed: %v\n", transformed)
}

// ADVANCED GENERIC PATTERNS: Understanding advanced patterns
func advancedGenericPatterns() {
	fmt.Println("Understanding advanced generic patterns...")
	
	// Pattern 1: Generic builder
	builder := NewBuilder[Person]()
	person := builder.
		Set("Name", "John").
		Set("Age", 30).
		Set("Email", "john@example.com").
		Build()
	
	fmt.Printf("  📊 Built person: %+v\n", person)
	
	// Pattern 2: Generic factory
	factory := NewFactory[Shape]()
	factory.Register("circle", func() Shape { return &Circle{Radius: 5.0} })
	factory.Register("rectangle", func() Shape { return &Rectangle{Width: 10, Height: 8} })
	
	circle, _ := factory.Create("circle")
	rectangle, _ := factory.Create("rectangle")
	
	fmt.Printf("  📊 Circle area: %.2f\n", circle.Area())
	fmt.Printf("  📊 Rectangle area: %.2f\n", rectangle.Area())
	
	// Pattern 3: Generic pipeline
	pipeline := NewPipeline[int]()
	result := pipeline.
		Add(func(x int) int { return x * 2 }).
		Add(func(x int) int { return x + 1 }).
		Add(func(x int) int { return x * 3 }).
		Process(5)
	
	fmt.Printf("  📊 Pipeline result: %d\n", result)
}

// GENERIC BEST PRACTICES: Following best practices
func genericBestPractices() {
	fmt.Println("Understanding generic best practices...")
	
	// 1. Use descriptive type parameter names
	fmt.Println("  📝 Best Practice 1: Use descriptive type parameter names")
	descriptiveTypeParameters()
	
	// 2. Use constraints when possible
	fmt.Println("  📝 Best Practice 2: Use constraints when possible")
	useConstraints()
	
	// 3. Avoid over-genericization
	fmt.Println("  📝 Best Practice 3: Avoid over-genericization")
	avoidOverGenericization()
	
	// 4. Use type inference
	fmt.Println("  📝 Best Practice 4: Use type inference")
	useTypeInference()
}

func descriptiveTypeParameters() {
	// Good: Descriptive type parameter names
	result := ProcessItems([]int{1, 2, 3}, func(item int) string {
		return fmt.Sprintf("item-%d", item)
	})
	fmt.Printf("    ✅ Processed: %v\n", result)
}

func useConstraints() {
	// Good: Use constraints
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Printf("    ✅ Sum: %d\n", sum)
}

func avoidOverGenericization() {
	// Good: Specific when appropriate
	numbers := []int{1, 2, 3, 4, 5}
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("    ✅ Evens: %v\n", evens)
}

func useTypeInference() {
	// Good: Let Go infer types
	result := Max(10, 20) // Type inferred as int
	fmt.Printf("    ✅ Max: %d\n", result)
}

// GENERIC FUNCTION IMPLEMENTATIONS

// Basic generic function
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Generic function with type constraint
func Sum[T Numeric](items []T) T {
	var sum T
	for _, item := range items {
		sum += item
	}
	return sum
}

// Generic function with comparable constraint
func Find[T comparable](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

// Generic function with ordered constraint
func MaxOrdered[T Ordered](items ...T) T {
	if len(items) == 0 {
		var zero T
		return zero
	}
	
	max := items[0]
	for _, item := range items[1:] {
		if item > max {
			max = item
		}
	}
	return max
}

// Generic map function
func Map[T, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}

// Generic filter function
func Filter[T any](items []T, fn func(T) bool) []T {
	var result []T
	for _, item := range items {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// Generic reduce function
func Reduce[T, U any](items []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, item := range items {
		result = fn(result, item)
	}
	return result
}

// Generic sort function
func Sort[T Ordered](items []T) {
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})
}

// Generic reverse function
func Reverse[T any](items []T) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

// Generic process function
func ProcessItems[T, U any](items []T, fn func(T) U) []U {
	return Map(items, fn)
}

// GENERIC TYPE IMPLEMENTATIONS

// Generic stack
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Generic queue
type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Generic map
type GenericMap[K comparable, V any] struct {
	data map[K]V
}

func NewGenericMap[K comparable, V any]() *GenericMap[K, V] {
	return &GenericMap[K, V]{data: make(map[K]V)}
}

func (m *GenericMap[K, V]) Set(key K, value V) {
	m.data[key] = value
}

func (m *GenericMap[K, V]) Get(key K) (V, bool) {
	value, exists := m.data[key]
	return value, exists
}

// Generic container
type Container[T any] struct {
	items []T
}

func NewContainer[T any]() *Container[T] {
	return &Container[T]{items: make([]T, 0)}
}

func (c *Container[T]) Add(item T) {
	c.items = append(c.items, item)
}

func (c *Container[T]) Contains(item T) bool {
	for _, i := range c.items {
		if any(i) == any(item) {
			return true
		}
	}
	return false
}

func (c *Container[T]) Count() int {
	return len(c.items)
}

func (c *Container[T]) Transform(fn func(T) interface{}) []interface{} {
	result := make([]interface{}, len(c.items))
	for i, item := range c.items {
		result[i] = fn(item)
	}
	return result
}

// GENERIC INTERFACE IMPLEMENTATIONS

// Generic repository
type Repository[T any] interface {
	Save(entity *T) error
	FindByID(id string) (*T, error)
	FindAll() []*T
}

type UserRepository struct {
	users map[string]*User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]*User)}
}

func (r *UserRepository) Save(user *User) error {
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(id string) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *UserRepository) FindAll() []*User {
	users := make([]*User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users
}

// Generic service
type UserService struct {
	repo Repository[User]
}

func NewUserService(repo Repository[User]) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ProcessUser(id string) {
	user, _ := s.repo.FindByID(id)
	fmt.Printf("  📊 Processing user: %s\n", user.Name)
}

// Generic validator
type Validator[T any] struct {
	rules []func(*T) error
}

func NewValidator[T any]() *Validator[T] {
	return &Validator[T]{rules: make([]func(*T) error, 0)}
}

func (v *Validator[T]) AddRule(rule func(*T) error) {
	v.rules = append(v.rules, rule)
}

func (v *Validator[T]) Validate(entity *T) error {
	for _, rule := range v.rules {
		if err := rule(entity); err != nil {
			return err
		}
	}
	return nil
}

// ADVANCED PATTERN IMPLEMENTATIONS

// Generic builder
type Builder[T any] struct {
	fields map[string]interface{}
}

func NewBuilder[T any]() *Builder[T] {
	return &Builder[T]{fields: make(map[string]interface{})}
}

func (b *Builder[T]) Set(field string, value interface{}) *Builder[T] {
	b.fields[field] = value
	return b
}

func (b *Builder[T]) Build() *T {
	// Simplified implementation
	var result T
	// In a real implementation, you would use reflection
	// to set the fields on the struct
	return &result
}

// Generic factory
type Factory[T any] struct {
	creators map[string]func() T
}

func NewFactory[T any]() *Factory[T] {
	return &Factory[T]{creators: make(map[string]func() T)}
}

func (f *Factory[T]) Register(name string, creator func() T) {
	f.creators[name] = creator
}

func (f *Factory[T]) Create(name string) (T, error) {
	creator, exists := f.creators[name]
	if !exists {
		var zero T
		return zero, fmt.Errorf("unknown type: %s", name)
	}
	return creator(), nil
}

// Generic pipeline
type Pipeline[T any] struct {
	steps []func(T) T
}

func NewPipeline[T any]() *Pipeline[T] {
	return &Pipeline[T]{steps: make([]func(T) T, 0)}
}

func (p *Pipeline[T]) Add(step func(T) T) *Pipeline[T] {
	p.steps = append(p.steps, step)
	return p
}

func (p *Pipeline[T]) Process(input T) T {
	result := input
	for _, step := range p.steps {
		result = step(result)
	}
	return result
}

// TYPE CONSTRAINTS

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// DATA STRUCTURES

type Pair[T, U any] struct {
	First  T
	Second U
}

type User struct {
	ID    string
	Name  string
	Email string
}

type Person struct {
	Name  string
	Age   int
	Email string
}

// SHAPE INTERFACES

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
