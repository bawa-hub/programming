package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// =============================================================================
// BASIC COMPOSITE PATTERN
// =============================================================================

// Component interface - defines common operations for both leaf and composite objects
type Component interface {
	Operation() string
	Add(component Component)
	Remove(component Component)
	GetChild(index int) Component
	GetName() string
	GetSize() int
}

// Leaf - represents individual objects that don't have children
type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) Operation() string {
	return f.name
}

func (f *File) Add(component Component) {
	// Files don't have children
	fmt.Printf("Cannot add component to file: %s\n", f.name)
}

func (f *File) Remove(component Component) {
	// Files don't have children
	fmt.Printf("Cannot remove component from file: %s\n", f.name)
}

func (f *File) GetChild(index int) Component {
	// Files don't have children
	return nil
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

// Composite - represents objects that have children
type Folder struct {
	name     string
	children []Component
}

func NewFolder(name string) *Folder {
	return &Folder{
		name:     name,
		children: make([]Component, 0),
	}
}

func (f *Folder) Operation() string {
	result := f.name + " ["
	for i, child := range f.children {
		if i > 0 {
			result += ", "
		}
		result += child.Operation()
	}
	result += "]"
	return result
}

func (f *Folder) Add(component Component) {
	f.children = append(f.children, component)
}

func (f *Folder) Remove(component Component) {
	for i, child := range f.children {
		if child == component {
			f.children = append(f.children[:i], f.children[i+1:]...)
			break
		}
	}
}

func (f *Folder) GetChild(index int) Component {
	if index >= 0 && index < len(f.children) {
		return f.children[index]
	}
	return nil
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) GetSize() int {
	totalSize := 0
	for _, child := range f.children {
		totalSize += child.GetSize()
	}
	return totalSize
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. FILE SYSTEM COMPOSITE
type FileSystemComponent interface {
	GetName() string
	GetSize() int
	GetType() string
	Display(indent int)
	Search(name string) []FileSystemComponent
}

type FileSystemFile struct {
	name string
	size int
	ext  string
}

func NewFileSystemFile(name string, size int, ext string) *FileSystemFile {
	return &FileSystemFile{name: name, size: size, ext: ext}
}

func (fsf *FileSystemFile) GetName() string {
	return fsf.name
}

func (fsf *FileSystemFile) GetSize() int {
	return fsf.size
}

func (fsf *FileSystemFile) GetType() string {
	return "File"
}

func (fsf *FileSystemFile) Display(indent int) {
	indentStr := strings.Repeat("  ", indent)
	fmt.Printf("%s📄 %s (%d bytes)\n", indentStr, fsf.name, fsf.size)
}

func (fsf *FileSystemFile) Search(name string) []FileSystemComponent {
	if strings.Contains(fsf.name, name) {
		return []FileSystemComponent{fsf}
	}
	return []FileSystemComponent{}
}

type FileSystemFolder struct {
	name     string
	children []FileSystemComponent
}

func NewFileSystemFolder(name string) *FileSystemFolder {
	return &FileSystemFolder{
		name:     name,
		children: make([]FileSystemComponent, 0),
	}
}

func (fsf *FileSystemFolder) GetName() string {
	return fsf.name
}

func (fsf *FileSystemFolder) GetSize() int {
	totalSize := 0
	for _, child := range fsf.children {
		totalSize += child.GetSize()
	}
	return totalSize
}

func (fsf *FileSystemFolder) GetType() string {
	return "Folder"
}

func (fsf *FileSystemFolder) Display(indent int) {
	indentStr := strings.Repeat("  ", indent)
	fmt.Printf("%s📁 %s (%d bytes)\n", indentStr, fsf.name, fsf.GetSize())
	for _, child := range fsf.children {
		child.Display(indent + 1)
	}
}

func (fsf *FileSystemFolder) Search(name string) []FileSystemComponent {
	var results []FileSystemComponent
	if strings.Contains(fsf.name, name) {
		results = append(results, fsf)
	}
	for _, child := range fsf.children {
		results = append(results, child.Search(name)...)
	}
	return results
}

func (fsf *FileSystemFolder) Add(component FileSystemComponent) {
	fsf.children = append(fsf.children, component)
}

func (fsf *FileSystemFolder) Remove(component FileSystemComponent) {
	for i, child := range fsf.children {
		if child == component {
			fsf.children = append(fsf.children[:i], fsf.children[i+1:]...)
			break
		}
	}
}

// 2. UI COMPONENT COMPOSITE
type UIComponent interface {
	Render() string
	GetName() string
	GetType() string
	Add(component UIComponent)
	Remove(component UIComponent)
	GetChild(index int) UIComponent
	GetChildren() []UIComponent
}

type UIButton struct {
	name    string
	text    string
	onClick string
}

func NewUIButton(name, text, onClick string) *UIButton {
	return &UIButton{name: name, text: text, onClick: onClick}
}

func (ub *UIButton) Render() string {
	return fmt.Sprintf("<button name='%s' onclick='%s'>%s</button>", ub.name, ub.onClick, ub.text)
}

func (ub *UIButton) GetName() string {
	return ub.name
}

func (ub *UIButton) GetType() string {
	return "Button"
}

func (ub *UIButton) Add(component UIComponent) {
	fmt.Printf("Cannot add component to button: %s\n", ub.name)
}

func (ub *UIButton) Remove(component UIComponent) {
	fmt.Printf("Cannot remove component from button: %s\n", ub.name)
}

func (ub *UIButton) GetChild(index int) UIComponent {
	return nil
}

func (ub *UIButton) GetChildren() []UIComponent {
	return []UIComponent{}
}

type UIPanel struct {
	name     string
	children []UIComponent
}

func NewUIPanel(name string) *UIPanel {
	return &UIPanel{
		name:     name,
		children: make([]UIComponent, 0),
	}
}

func (up *UIPanel) Render() string {
	result := fmt.Sprintf("<div name='%s'>", up.name)
	for _, child := range up.children {
		result += child.Render()
	}
	result += "</div>"
	return result
}

func (up *UIPanel) GetName() string {
	return up.name
}

func (up *UIPanel) GetType() string {
	return "Panel"
}

func (up *UIPanel) Add(component UIComponent) {
	up.children = append(up.children, component)
}

func (up *UIPanel) Remove(component UIComponent) {
	for i, child := range up.children {
		if child == component {
			up.children = append(up.children[:i], up.children[i+1:]...)
			break
		}
	}
}

func (up *UIPanel) GetChild(index int) UIComponent {
	if index >= 0 && index < len(up.children) {
		return up.children[index]
	}
	return nil
}

func (up *UIPanel) GetChildren() []UIComponent {
	return up.children
}

type UIWindow struct {
	name     string
	title    string
	children []UIComponent
}

func NewUIWindow(name, title string) *UIWindow {
	return &UIWindow{
		name:     name,
		title:    title,
		children: make([]UIComponent, 0),
	}
}

func (uw *UIWindow) Render() string {
	result := fmt.Sprintf("<window name='%s' title='%s'>", uw.name, uw.title)
	for _, child := range uw.children {
		result += child.Render()
	}
	result += "</window>"
	return result
}

func (uw *UIWindow) GetName() string {
	return uw.name
}

func (uw *UIWindow) GetType() string {
	return "Window"
}

func (uw *UIWindow) Add(component UIComponent) {
	uw.children = append(uw.children, component)
}

func (uw *UIWindow) Remove(component UIComponent) {
	for i, child := range uw.children {
		if child == component {
			uw.children = append(uw.children[:i], uw.children[i+1:]...)
			break
		}
	}
}

func (uw *UIWindow) GetChild(index int) UIComponent {
	if index >= 0 && index < len(uw.children) {
		return uw.children[index]
	}
	return nil
}

func (uw *UIWindow) GetChildren() []UIComponent {
	return uw.children
}

// 3. ORGANIZATION STRUCTURE COMPOSITE
type Employee interface {
	GetName() string
	GetPosition() string
	GetSalary() float64
	GetDepartment() string
	Add(employee Employee)
	Remove(employee Employee)
	GetSubordinates() []Employee
	GetTotalSalary() float64
}

type IndividualEmployee struct {
	name       string
	position   string
	salary     float64
	department string
}

func NewIndividualEmployee(name, position, department string, salary float64) *IndividualEmployee {
	return &IndividualEmployee{
		name:       name,
		position:   position,
		salary:     salary,
		department: department,
	}
}

func (ie *IndividualEmployee) GetName() string {
	return ie.name
}

func (ie *IndividualEmployee) GetPosition() string {
	return ie.position
}

func (ie *IndividualEmployee) GetSalary() float64 {
	return ie.salary
}

func (ie *IndividualEmployee) GetDepartment() string {
	return ie.department
}

func (ie *IndividualEmployee) Add(employee Employee) {
	fmt.Printf("Cannot add employee to individual: %s\n", ie.name)
}

func (ie *IndividualEmployee) Remove(employee Employee) {
	fmt.Printf("Cannot remove employee from individual: %s\n", ie.name)
}

func (ie *IndividualEmployee) GetSubordinates() []Employee {
	return []Employee{}
}

func (ie *IndividualEmployee) GetTotalSalary() float64 {
	return ie.salary
}

type Department struct {
	name        string
	employees   []Employee
	manager     string
	description string
}

func NewDepartment(name, manager, description string) *Department {
	return &Department{
		name:        name,
		manager:     manager,
		description: description,
		employees:   make([]Employee, 0),
	}
}

func (d *Department) GetName() string {
	return d.name
}

func (d *Department) GetPosition() string {
	return "Department"
}

func (d *Department) GetSalary() float64 {
	return 0 // Departments don't have salaries
}

func (d *Department) GetDepartment() string {
	return d.name
}

func (d *Department) Add(employee Employee) {
	d.employees = append(d.employees, employee)
}

func (d *Department) Remove(employee Employee) {
	for i, emp := range d.employees {
		if emp == employee {
			d.employees = append(d.employees[:i], d.employees[i+1:]...)
			break
		}
	}
}

func (d *Department) GetSubordinates() []Employee {
	return d.employees
}

func (d *Department) GetTotalSalary() float64 {
	totalSalary := 0.0
	for _, employee := range d.employees {
		totalSalary += employee.GetTotalSalary()
	}
	return totalSalary
}

// =============================================================================
// COMPOSITE WITH VISITOR PATTERN
// =============================================================================

// Visitor interface for operations on composite structures
type Visitor interface {
	VisitFile(file *File)
	VisitFolder(folder *Folder)
}

// Concrete visitor for calculating total size
type SizeVisitor struct {
	totalSize int
}

func (sv *SizeVisitor) VisitFile(file *File) {
	sv.totalSize += file.GetSize()
}

func (sv *SizeVisitor) VisitFolder(folder *Folder) {
	for _, child := range folder.children {
		switch c := child.(type) {
		case *File:
			sv.VisitFile(c)
		case *Folder:
			sv.VisitFolder(c)
		}
	}
}

func (sv *SizeVisitor) GetTotalSize() int {
	return sv.totalSize
}

// Concrete visitor for counting files
type FileCountVisitor struct {
	fileCount int
}

func (fcv *FileCountVisitor) VisitFile(file *File) {
	fcv.fileCount++
}

func (fcv *FileCountVisitor) VisitFolder(folder *Folder) {
	for _, child := range folder.children {
		switch c := child.(type) {
		case *File:
			fcv.VisitFile(c)
		case *Folder:
			fcv.VisitFolder(c)
		}
	}
}

func (fcv *FileCountVisitor) GetFileCount() int {
	return fcv.fileCount
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== COMPOSITE PATTERN DEMONSTRATION ===\n")

	// 1. BASIC COMPOSITE
	fmt.Println("1. BASIC COMPOSITE:")
	
	// Create files
	file1 := NewFile("document.txt", 1024)
	file2 := NewFile("image.jpg", 2048)
	file3 := NewFile("video.mp4", 10240)
	
	// Create folders
	folder1 := NewFolder("Documents")
	folder2 := NewFolder("Media")
	rootFolder := NewFolder("Root")
	
	// Add files to folders
	folder1.Add(file1)
	folder2.Add(file2)
	folder2.Add(file3)
	
	// Add folders to root
	rootFolder.Add(folder1)
	rootFolder.Add(folder2)
	
	// Display structure
	fmt.Printf("File structure: %s\n", rootFolder.Operation())
	fmt.Printf("Total size: %d bytes\n", rootFolder.GetSize())
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// File System Composite
	fmt.Println("File System Composite:")
	rootDir := NewFileSystemFolder("Root")
	documentsDir := NewFileSystemFolder("Documents")
	imagesDir := NewFileSystemFolder("Images")
	
	// Add files
	documentsDir.Add(NewFileSystemFile("report.pdf", 1024, "pdf"))
	documentsDir.Add(NewFileSystemFile("notes.txt", 512, "txt"))
	imagesDir.Add(NewFileSystemFile("photo1.jpg", 2048, "jpg"))
	imagesDir.Add(NewFileSystemFile("photo2.png", 1536, "png"))
	
	// Add folders to root
	rootDir.Add(documentsDir)
	rootDir.Add(imagesDir)
	
	// Display file system
	rootDir.Display(0)
	fmt.Printf("Total size: %d bytes\n", rootDir.GetSize())
	
	// Search for files
	results := rootDir.Search("photo")
	fmt.Printf("Search results for 'photo': %d files found\n", len(results))
	for _, result := range results {
		fmt.Printf("  - %s\n", result.GetName())
	}
	fmt.Println()

	// UI Component Composite
	fmt.Println("UI Component Composite:")
	mainWindow := NewUIWindow("main", "My Application")
	headerPanel := NewUIPanel("header")
	contentPanel := NewUIPanel("content")
	footerPanel := NewUIPanel("footer")
	
	// Add buttons
	headerPanel.Add(NewUIButton("home", "Home", "goHome()"))
	headerPanel.Add(NewUIButton("about", "About", "showAbout()"))
	contentPanel.Add(NewUIButton("save", "Save", "saveData()"))
	contentPanel.Add(NewUIButton("cancel", "Cancel", "cancelAction()"))
	footerPanel.Add(NewUIButton("help", "Help", "showHelp()"))
	
	// Add panels to window
	mainWindow.Add(headerPanel)
	mainWindow.Add(contentPanel)
	mainWindow.Add(footerPanel)
	
	// Render UI
	fmt.Printf("UI Structure:\n%s\n", mainWindow.Render())
	fmt.Println()

	// Organization Structure Composite
	fmt.Println("Organization Structure Composite:")
	engineeringDept := NewDepartment("Engineering", "John Smith", "Software development")
	marketingDept := NewDepartment("Marketing", "Jane Doe", "Product marketing")
	
	// Add employees
	engineeringDept.Add(NewIndividualEmployee("Alice", "Senior Developer", "Engineering", 80000))
	engineeringDept.Add(NewIndividualEmployee("Bob", "Junior Developer", "Engineering", 60000))
	engineeringDept.Add(NewIndividualEmployee("Charlie", "DevOps Engineer", "Engineering", 75000))
	
	marketingDept.Add(NewIndividualEmployee("David", "Marketing Manager", "Marketing", 70000))
	marketingDept.Add(NewIndividualEmployee("Eve", "Content Writer", "Marketing", 50000))
	
	// Display organization
	fmt.Printf("Engineering Department:\n")
	fmt.Printf("  Manager: %s\n", engineeringDept.manager)
	fmt.Printf("  Total Salary: $%.2f\n", engineeringDept.GetTotalSalary())
	fmt.Printf("  Employees: %d\n", len(engineeringDept.GetSubordinates()))
	
	fmt.Printf("Marketing Department:\n")
	fmt.Printf("  Manager: %s\n", marketingDept.manager)
	fmt.Printf("  Total Salary: $%.2f\n", marketingDept.GetTotalSalary())
	fmt.Printf("  Employees: %d\n", len(marketingDept.GetSubordinates()))
	fmt.Println()

	// 3. COMPOSITE WITH VISITOR
	fmt.Println("3. COMPOSITE WITH VISITOR:")
	
	// Create a complex file structure
	complexRoot := NewFolder("ComplexRoot")
	subFolder1 := NewFolder("SubFolder1")
	subFolder2 := NewFolder("SubFolder2")
	
	subFolder1.Add(NewFile("file1.txt", 100))
	subFolder1.Add(NewFile("file2.txt", 200))
	subFolder2.Add(NewFile("file3.txt", 300))
	subFolder2.Add(NewFile("file4.txt", 400))
	
	complexRoot.Add(subFolder1)
	complexRoot.Add(subFolder2)
	complexRoot.Add(NewFile("rootFile.txt", 500))
	
	// Use visitors
	sizeVisitor := &SizeVisitor{}
	sizeVisitor.VisitFolder(complexRoot)
	fmt.Printf("Total size (visitor): %d bytes\n", sizeVisitor.GetTotalSize())
	
	fileCountVisitor := &FileCountVisitor{}
	fileCountVisitor.VisitFolder(complexRoot)
	fmt.Printf("Total files (visitor): %d\n", fileCountVisitor.GetFileCount())
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
