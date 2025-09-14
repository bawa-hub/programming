package examples

import (
	"fmt"
	"sort"
	"strings"
)

// Student represents a student with grades
type Student struct {
	ID       int
	Name     string
	Grades   []float64
	Average  float64
}

// Course represents a course with students
type Course struct {
	Name     string
	Students []*Student
	Capacity int
}

// GradeCalculator interface for different grading systems
type GradeCalculator interface {
	CalculateGrade(score float64) string
	GetGradePoints(score float64) float64
}

// LetterGradeCalculator implements letter grade system
type LetterGradeCalculator struct {
	Scale map[string]float64
}

// PassFailCalculator implements pass/fail system
type PassFailCalculator struct {
	PassingScore float64
}

// Database represents a simple in-memory database
type Database struct {
	Students map[int]*Student
	Courses  map[string]*Course
	NextID   int
}

// DemonstrateStudentManagement shows a complete student management system
func DemonstrateStudentManagement() {
	fmt.Println("=== STUDENT MANAGEMENT SYSTEM ===")
	
	// 1. Create database
	db := NewDatabase()
	
	// 2. Create students
	alice := &Student{
		ID:     1,
		Name:   "Alice Johnson",
		Grades: []float64{85, 92, 78, 96, 88},
	}
	alice.CalculateAverage()
	
	bob := &Student{
		ID:     2,
		Name:   "Bob Smith",
		Grades: []float64{76, 82, 90, 85, 79},
	}
	bob.CalculateAverage()
	
	charlie := &Student{
		ID:     3,
		Name:   "Charlie Brown",
		Grades: []float64{95, 98, 92, 94, 96},
	}
	charlie.CalculateAverage()
	
	// 3. Add students to database
	db.AddStudent(alice)
	db.AddStudent(bob)
	db.AddStudent(charlie)
	
	// 4. Create course
	course := &Course{
		Name:     "Go Programming",
		Students: []*Student{alice, bob, charlie},
		Capacity: 30,
	}
	
	// 5. Demonstrate various operations
	fmt.Println("\n--- Student Information ---")
	for _, student := range course.Students {
		fmt.Printf("Student: %s (ID: %d)\n", student.Name, student.ID)
		fmt.Printf("  Grades: %v\n", student.Grades)
		fmt.Printf("  Average: %.2f\n", student.Average)
		fmt.Printf("  Letter Grade: %s\n", student.GetLetterGrade())
		fmt.Println()
	}
	
	// 6. Demonstrate sorting
	fmt.Println("--- Students Sorted by Average ---")
	sort.Slice(course.Students, func(i, j int) bool {
		return course.Students[i].Average > course.Students[j].Average
	})
	
	for i, student := range course.Students {
		fmt.Printf("%d. %s: %.2f\n", i+1, student.Name, student.Average)
	}
	
	// 7. Demonstrate grade calculators
	fmt.Println("\n--- Grade Calculators ---")
	letterCalc := NewLetterGradeCalculator()
	passFailCalc := NewPassFailCalculator(70.0)
	
	score := 85.0
	fmt.Printf("Score: %.1f\n", score)
	fmt.Printf("Letter Grade: %s\n", letterCalc.CalculateGrade(score))
	fmt.Printf("Pass/Fail: %s\n", passFailCalc.CalculateGrade(score))
}

// DemonstrateDataStructures shows various data structure implementations
func DemonstrateDataStructures() {
	fmt.Println("\n=== DATA STRUCTURES ===")
	
	// 1. Stack implementation
	fmt.Println("\n--- Stack Implementation ---")
	stack := NewStack()
	
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	fmt.Printf("Stack: %v\n", stack.GetAll())
	fmt.Printf("Top: %d\n", stack.Peek())
	
	popped := stack.Pop()
	fmt.Printf("Popped: %d\n", popped)
	fmt.Printf("Stack after pop: %v\n", stack.GetAll())
	
	// 2. Queue implementation
	fmt.Println("\n--- Queue Implementation ---")
	queue := NewQueue()
	
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")
	
	fmt.Printf("Queue: %v\n", queue.GetAll())
	fmt.Printf("Front: %s\n", queue.Front())
	
	dequeued := queue.Dequeue()
	fmt.Printf("Dequeued: %s\n", dequeued)
	fmt.Printf("Queue after dequeue: %v\n", queue.GetAll())
	
	// 3. Hash table implementation
	fmt.Println("\n--- Hash Table Implementation ---")
	hashTable := NewHashTable()
	
	hashTable.Set("name", "John Doe")
	hashTable.Set("age", 30)
	hashTable.Set("city", "New York")
	
	fmt.Printf("name: %v\n", hashTable.Get("name"))
	fmt.Printf("age: %v\n", hashTable.Get("age"))
	fmt.Printf("city: %v\n", hashTable.Get("city"))
	fmt.Printf("country: %v\n", hashTable.Get("country")) // Not found
	
	// 4. Binary tree implementation
	fmt.Println("\n--- Binary Tree Implementation ---")
	tree := NewBinaryTree()
	
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, value := range values {
		tree.Insert(value)
	}
	
	fmt.Printf("In-order traversal: %v\n", tree.InOrder())
	fmt.Printf("Pre-order traversal: %v\n", tree.PreOrder())
	fmt.Printf("Post-order traversal: %v\n", tree.PostOrder())
}

// DemonstrateAlgorithms shows various algorithm implementations
func DemonstrateAlgorithms() {
	fmt.Println("\n=== ALGORITHMS ===")
	
	// 1. Sorting algorithms
	fmt.Println("\n--- Sorting Algorithms ---")
	
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original: %v\n", numbers)
	
	// Bubble sort
	bubbleSorted := make([]int, len(numbers))
	copy(bubbleSorted, numbers)
	BubbleSort(bubbleSorted)
	fmt.Printf("Bubble Sort: %v\n", bubbleSorted)
	
	// Quick sort
	quickSorted := make([]int, len(numbers))
	copy(quickSorted, numbers)
	QuickSort(quickSorted)
	fmt.Printf("Quick Sort: %v\n", quickSorted)
	
	// 2. Search algorithms
	fmt.Println("\n--- Search Algorithms ---")
	
	sortedNumbers := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7
	
	linearIndex := LinearSearch(sortedNumbers, target)
	fmt.Printf("Linear search for %d: index %d\n", target, linearIndex)
	
	binaryIndex := BinarySearch(sortedNumbers, target)
	fmt.Printf("Binary search for %d: index %d\n", target, binaryIndex)
	
	// 3. String algorithms
	fmt.Println("\n--- String Algorithms ---")
	
	text := "Hello, World!"
	pattern := "World"
	
	index := StringSearch(text, pattern)
	fmt.Printf("String '%s' found in '%s' at index %d\n", pattern, text, index)
	
	// Palindrome check
	palindromes := []string{"racecar", "hello", "madam", "world"}
	for _, word := range palindromes {
		fmt.Printf("'%s' is palindrome: %t\n", word, IsPalindrome(word))
	}
}

// DemonstrateConcurrency shows basic concurrency concepts
func DemonstrateConcurrency() {
	fmt.Println("\n=== CONCURRENCY BASICS ===")
	
	// 1. Goroutines
	fmt.Println("\n--- Goroutines ---")
	
	ch := make(chan string, 3)
	
	go func() {
		ch <- "Hello"
	}()
	
	go func() {
		ch <- "from"
	}()
	
	go func() {
		ch <- "goroutines!"
	}()
	
	// Receive messages
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Printf("Received: %s\n", msg)
	}
	
	// 2. Channels
	fmt.Println("\n--- Channels ---")
	
	numbers := []int{1, 2, 3, 4, 5}
	squares := make(chan int, len(numbers))
	
	// Send squares
	go func() {
		for _, num := range numbers {
			squares <- num * num
		}
		close(squares)
	}()
	
	// Receive squares
	fmt.Printf("Squares of %v: ", numbers)
	for square := range squares {
		fmt.Printf("%d ", square)
	}
	fmt.Println()
	
	// 3. Worker pool
	fmt.Println("\n--- Worker Pool ---")
	
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start workers
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	
	// Send jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	
	// Collect results
	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Printf("Job %d processed, result: %d\n", i, result)
	}
}

// Student methods

func (s *Student) CalculateAverage() {
	if len(s.Grades) == 0 {
		s.Average = 0
		return
	}
	
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade
	}
	s.Average = sum / float64(len(s.Grades))
}

func (s *Student) GetLetterGrade() string {
	if s.Average >= 90 {
		return "A"
	} else if s.Average >= 80 {
		return "B"
	} else if s.Average >= 70 {
		return "C"
	} else if s.Average >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func (s *Student) AddGrade(grade float64) {
	s.Grades = append(s.Grades, grade)
	s.CalculateAverage()
}

// GradeCalculator implementations

func (lgc *LetterGradeCalculator) CalculateGrade(score float64) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func (lgc *LetterGradeCalculator) GetGradePoints(score float64) float64 {
	if score >= 90 {
		return 4.0
	} else if score >= 80 {
		return 3.0
	} else if score >= 70 {
		return 2.0
	} else if score >= 60 {
		return 1.0
	} else {
		return 0.0
	}
}

func (pfc *PassFailCalculator) CalculateGrade(score float64) string {
	if score >= pfc.PassingScore {
		return "PASS"
	} else {
		return "FAIL"
	}
}

func (pfc *PassFailCalculator) GetGradePoints(score float64) float64 {
	if score >= pfc.PassingScore {
		return 1.0
	} else {
		return 0.0
	}
}

// Database methods

func NewDatabase() *Database {
	return &Database{
		Students: make(map[int]*Student),
		Courses:  make(map[string]*Course),
		NextID:   1,
	}
}

func (db *Database) AddStudent(student *Student) {
	db.Students[student.ID] = student
}

func (db *Database) GetStudent(id int) *Student {
	return db.Students[id]
}

func (db *Database) GetAllStudents() []*Student {
	students := make([]*Student, 0, len(db.Students))
	for _, student := range db.Students {
		students = append(students, student)
	}
	return students
}

// Stack implementation

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{items: make([]interface{}, 0)}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) GetAll() []interface{} {
	return s.items
}

// Queue implementation

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{items: make([]interface{}, 0)}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Front() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) GetAll() []interface{} {
	return q.items
}

// Hash table implementation

type HashTable struct {
	buckets []*Bucket
	size    int
}

type Bucket struct {
	key   string
	value interface{}
	next  *Bucket
}

func NewHashTable() *HashTable {
	return &HashTable{
		buckets: make([]*Bucket, 16),
		size:    0,
	}
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash + int(char)) % len(ht.buckets)
	}
	return hash
}

func (ht *HashTable) Set(key string, value interface{}) {
	index := ht.hash(key)
	bucket := ht.buckets[index]
	
	// Check if key already exists
	for bucket != nil {
		if bucket.key == key {
			bucket.value = value
			return
		}
		bucket = bucket.next
	}
	
	// Add new bucket
	newBucket := &Bucket{key: key, value: value, next: ht.buckets[index]}
	ht.buckets[index] = newBucket
	ht.size++
}

func (ht *HashTable) Get(key string) interface{} {
	index := ht.hash(key)
	bucket := ht.buckets[index]
	
	for bucket != nil {
		if bucket.key == key {
			return bucket.value
		}
		bucket = bucket.next
	}
	return nil
}

// Binary tree implementation

type BinaryTree struct {
	root *TreeNode
}

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(value int) {
	bt.root = bt.insertRecursive(bt.root, value)
}

func (bt *BinaryTree) insertRecursive(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{value: value}
	}
	
	if value < node.value {
		node.left = bt.insertRecursive(node.left, value)
	} else if value > node.value {
		node.right = bt.insertRecursive(node.right, value)
	}
	
	return node
}

func (bt *BinaryTree) InOrder() []int {
	var result []int
	bt.inOrderRecursive(bt.root, &result)
	return result
}

func (bt *BinaryTree) inOrderRecursive(node *TreeNode, result *[]int) {
	if node != nil {
		bt.inOrderRecursive(node.left, result)
		*result = append(*result, node.value)
		bt.inOrderRecursive(node.right, result)
	}
}

func (bt *BinaryTree) PreOrder() []int {
	var result []int
	bt.preOrderRecursive(bt.root, &result)
	return result
}

func (bt *BinaryTree) preOrderRecursive(node *TreeNode, result *[]int) {
	if node != nil {
		*result = append(*result, node.value)
		bt.preOrderRecursive(node.left, result)
		bt.preOrderRecursive(node.right, result)
	}
}

func (bt *BinaryTree) PostOrder() []int {
	var result []int
	bt.postOrderRecursive(bt.root, &result)
	return result
}

func (bt *BinaryTree) postOrderRecursive(node *TreeNode, result *[]int) {
	if node != nil {
		bt.postOrderRecursive(node.left, result)
		bt.postOrderRecursive(node.right, result)
		*result = append(*result, node.value)
	}
}

// Sorting algorithms

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	
	left, right := 0, len(arr)-1
	pivot := partition(arr, left, right)
	
	QuickSort(arr[:pivot])
	QuickSort(arr[pivot+1:])
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// Search algorithms

func LinearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}

func StringSearch(text, pattern string) int {
	if len(pattern) > len(text) {
		return -1
	}
	
	for i := 0; i <= len(text)-len(pattern); i++ {
		if text[i:i+len(pattern)] == pattern {
			return i
		}
	}
	
	return -1
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	
	return true
}

// Concurrency functions

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * job
	}
}

// Helper functions

func NewLetterGradeCalculator() *LetterGradeCalculator {
	return &LetterGradeCalculator{
		Scale: map[string]float64{
			"A": 4.0,
			"B": 3.0,
			"C": 2.0,
			"D": 1.0,
			"F": 0.0,
		},
	}
}

func NewPassFailCalculator(passingScore float64) *PassFailCalculator {
	return &PassFailCalculator{
		PassingScore: passingScore,
	}
}

// RunAllExamples runs all practical examples
func RunAllExamples() {
	DemonstrateStudentManagement()
	DemonstrateDataStructures()
	DemonstrateAlgorithms()
	DemonstrateConcurrency()
}
