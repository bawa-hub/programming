package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC VISITOR PATTERN
// =============================================================================

// Visitor interface
type Visitor interface {
	VisitElementA(element *ElementA)
	VisitElementB(element *ElementB)
	VisitElementC(element *ElementC)
}

// Element interface
type Element interface {
	Accept(visitor Visitor)
	GetName() string
}

// Concrete Elements
type ElementA struct {
	name string
}

func NewElementA(name string) *ElementA {
	return &ElementA{name: name}
}

func (ea *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(ea)
}

func (ea *ElementA) GetName() string {
	return ea.name
}

type ElementB struct {
	name string
}

func NewElementB(name string) *ElementB {
	return &ElementB{name: name}
}

func (eb *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(eb)
}

func (eb *ElementB) GetName() string {
	return eb.name
}

type ElementC struct {
	name string
}

func NewElementC(name string) *ElementC {
	return &ElementC{name: name}
}

func (ec *ElementC) Accept(visitor Visitor) {
	visitor.VisitElementC(ec)
}

func (ec *ElementC) GetName() string {
	return ec.name
}

// Concrete Visitors
type ConcreteVisitor1 struct{}

func (cv1 *ConcreteVisitor1) VisitElementA(element *ElementA) {
	fmt.Printf("Visitor1: Processing ElementA - %s\n", element.GetName())
}

func (cv1 *ConcreteVisitor1) VisitElementB(element *ElementB) {
	fmt.Printf("Visitor1: Processing ElementB - %s\n", element.GetName())
}

func (cv1 *ConcreteVisitor1) VisitElementC(element *ElementC) {
	fmt.Printf("Visitor1: Processing ElementC - %s\n", element.GetName())
}

type ConcreteVisitor2 struct{}

func (cv2 *ConcreteVisitor2) VisitElementA(element *ElementA) {
	fmt.Printf("Visitor2: Analyzing ElementA - %s\n", element.GetName())
}

func (cv2 *ConcreteVisitor2) VisitElementB(element *ElementB) {
	fmt.Printf("Visitor2: Analyzing ElementB - %s\n", element.GetName())
}

func (cv2 *ConcreteVisitor2) VisitElementC(element *ElementC) {
	fmt.Printf("Visitor2: Analyzing ElementC - %s\n", element.GetName())
}

// Object Structure
type ObjectStructure struct {
	elements []Element
}

func NewObjectStructure() *ObjectStructure {
	return &ObjectStructure{
		elements: make([]Element, 0),
	}
}

func (os *ObjectStructure) AddElement(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range os.elements {
		element.Accept(visitor)
	}
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. SHOPPING CART VISITOR
type ShoppingCartItem interface {
	Accept(visitor ShoppingCartVisitor)
	GetName() string
	GetPrice() float64
	GetQuantity() int
}

type ShoppingCartVisitor interface {
	VisitBook(book *Book)
	VisitElectronics(electronics *Electronics)
	VisitClothing(clothing *Clothing)
	GetTotal() float64
	GetName() string
}

// Concrete Items
type Book struct {
	name     string
	price    float64
	quantity int
	author   string
}

func NewBook(name, author string, price float64, quantity int) *Book {
	return &Book{
		name:     name,
		author:   author,
		price:    price,
		quantity: quantity,
	}
}

func (b *Book) Accept(visitor ShoppingCartVisitor) {
	visitor.VisitBook(b)
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) GetPrice() float64 {
	return b.price
}

func (b *Book) GetQuantity() int {
	return b.quantity
}

func (b *Book) GetAuthor() string {
	return b.author
}

type Electronics struct {
	name     string
	price    float64
	quantity int
	brand    string
}

func NewElectronics(name, brand string, price float64, quantity int) *Electronics {
	return &Electronics{
		name:     name,
		brand:    brand,
		price:    price,
		quantity: quantity,
	}
}

func (e *Electronics) Accept(visitor ShoppingCartVisitor) {
	visitor.VisitElectronics(e)
}

func (e *Electronics) GetName() string {
	return e.name
}

func (e *Electronics) GetPrice() float64 {
	return e.price
}

func (e *Electronics) GetQuantity() int {
	return e.quantity
}

func (e *Electronics) GetBrand() string {
	return e.brand
}

type Clothing struct {
	name     string
	price    float64
	quantity int
	size     string
}

func NewClothing(name, size string, price float64, quantity int) *Clothing {
	return &Clothing{
		name:     name,
		size:     size,
		price:    price,
		quantity: quantity,
	}
}

func (c *Clothing) Accept(visitor ShoppingCartVisitor) {
	visitor.VisitClothing(c)
}

func (c *Clothing) GetName() string {
	return c.name
}

func (c *Clothing) GetPrice() float64 {
	return c.price
}

func (c *Clothing) GetQuantity() int {
	return c.quantity
}

func (c *Clothing) GetSize() string {
	return c.size
}

// Concrete Visitors
type PriceCalculator struct {
	total float64
}

func NewPriceCalculator() *PriceCalculator {
	return &PriceCalculator{total: 0}
}

func (pc *PriceCalculator) VisitBook(book *Book) {
	itemTotal := book.GetPrice() * float64(book.GetQuantity())
	pc.total += itemTotal
	fmt.Printf("Book: %s by %s - $%.2f x %d = $%.2f\n", 
		book.GetName(), book.GetAuthor(), book.GetPrice(), book.GetQuantity(), itemTotal)
}

func (pc *PriceCalculator) VisitElectronics(electronics *Electronics) {
	itemTotal := electronics.GetPrice() * float64(electronics.GetQuantity())
	pc.total += itemTotal
	fmt.Printf("Electronics: %s by %s - $%.2f x %d = $%.2f\n", 
		electronics.GetName(), electronics.GetBrand(), electronics.GetPrice(), electronics.GetQuantity(), itemTotal)
}

func (pc *PriceCalculator) VisitClothing(clothing *Clothing) {
	itemTotal := clothing.GetPrice() * float64(clothing.GetQuantity())
	pc.total += itemTotal
	fmt.Printf("Clothing: %s (Size: %s) - $%.2f x %d = $%.2f\n", 
		clothing.GetName(), clothing.GetSize(), clothing.GetPrice(), clothing.GetQuantity(), itemTotal)
}

func (pc *PriceCalculator) GetTotal() float64 {
	return pc.total
}

func (pc *PriceCalculator) GetName() string {
	return "Price Calculator"
}

type TaxCalculator struct {
	total float64
	taxRate float64
}

func NewTaxCalculator(taxRate float64) *TaxCalculator {
	return &TaxCalculator{
		total:    0,
		taxRate:  taxRate,
	}
}

func (tc *TaxCalculator) VisitBook(book *Book) {
	itemTotal := book.GetPrice() * float64(book.GetQuantity())
	tc.total += itemTotal
	fmt.Printf("Book: %s - $%.2f (tax included)\n", book.GetName(), itemTotal)
}

func (tc *TaxCalculator) VisitElectronics(electronics *Electronics) {
	itemTotal := electronics.GetPrice() * float64(electronics.GetQuantity())
	taxAmount := itemTotal * tc.taxRate
	tc.total += itemTotal + taxAmount
	fmt.Printf("Electronics: %s - $%.2f + $%.2f tax = $%.2f\n", 
		electronics.GetName(), itemTotal, taxAmount, itemTotal + taxAmount)
}

func (tc *TaxCalculator) VisitClothing(clothing *Clothing) {
	itemTotal := clothing.GetPrice() * float64(clothing.GetQuantity())
	taxAmount := itemTotal * tc.taxRate
	tc.total += itemTotal + taxAmount
	fmt.Printf("Clothing: %s - $%.2f + $%.2f tax = $%.2f\n", 
		clothing.GetName(), itemTotal, taxAmount, itemTotal + taxAmount)
}

func (tc *TaxCalculator) GetTotal() float64 {
	return tc.total
}

func (tc *TaxCalculator) GetName() string {
	return "Tax Calculator"
}

type InventoryManager struct {
	items map[string]int
}

func NewInventoryManager() *InventoryManager {
	return &InventoryManager{
		items: make(map[string]int),
	}
}

func (im *InventoryManager) VisitBook(book *Book) {
	im.items[book.GetName()] += book.GetQuantity()
	fmt.Printf("Inventory: Added %d copies of '%s' by %s\n", 
		book.GetQuantity(), book.GetName(), book.GetAuthor())
}

func (im *InventoryManager) VisitElectronics(electronics *Electronics) {
	im.items[electronics.GetName()] += electronics.GetQuantity()
	fmt.Printf("Inventory: Added %d units of '%s' by %s\n", 
		electronics.GetQuantity(), electronics.GetName(), electronics.GetBrand())
}

func (im *InventoryManager) VisitClothing(clothing *Clothing) {
	im.items[clothing.GetName()] += clothing.GetQuantity()
	fmt.Printf("Inventory: Added %d units of '%s' (Size: %s)\n", 
		clothing.GetQuantity(), clothing.GetName(), clothing.GetSize())
}

func (im *InventoryManager) GetTotal() float64 {
	return 0 // Not applicable for inventory
}

func (im *InventoryManager) GetName() string {
	return "Inventory Manager"
}

func (im *InventoryManager) GetInventory() map[string]int {
	return im.items
}

// Shopping Cart
type ShoppingCart struct {
	items []ShoppingCartItem
}

func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		items: make([]ShoppingCartItem, 0),
	}
}

func (sc *ShoppingCart) AddItem(item ShoppingCartItem) {
	sc.items = append(sc.items, item)
}

func (sc *ShoppingCart) Accept(visitor ShoppingCartVisitor) {
	for _, item := range sc.items {
		item.Accept(visitor)
	}
}

// 2. DOCUMENT PROCESSING VISITOR
type DocumentElement interface {
	Accept(visitor DocumentVisitor)
	GetType() string
	GetContent() string
}

type DocumentVisitor interface {
	VisitText(text *Text)
	VisitImage(image *Image)
	VisitTable(table *Table)
	GetResult() string
	GetName() string
}

// Concrete Document Elements
type Text struct {
	content string
	font    string
	size    int
}

func NewText(content, font string, size int) *Text {
	return &Text{
		content: content,
		font:    font,
		size:    size,
	}
}

func (t *Text) Accept(visitor DocumentVisitor) {
	visitor.VisitText(t)
}

func (t *Text) GetType() string {
	return "Text"
}

func (t *Text) GetContent() string {
	return t.content
}

func (t *Text) GetFont() string {
	return t.font
}

func (t *Text) GetSize() int {
	return t.size
}

type Image struct {
	src    string
	width  int
	height int
	alt    string
}

func NewImage(src, alt string, width, height int) *Image {
	return &Image{
		src:    src,
		alt:    alt,
		width:  width,
		height: height,
	}
}

func (i *Image) Accept(visitor DocumentVisitor) {
	visitor.VisitImage(i)
}

func (i *Image) GetType() string {
	return "Image"
}

func (i *Image) GetContent() string {
	return i.src
}

func (i *Image) GetAlt() string {
	return i.alt
}

func (i *Image) GetWidth() int {
	return i.width
}

func (i *Image) GetHeight() int {
	return i.height
}

type Table struct {
	rows    int
	columns int
	data    [][]string
}

func NewTable(rows, columns int, data [][]string) *Table {
	return &Table{
		rows:    rows,
		columns: columns,
		data:    data,
	}
}

func (t *Table) Accept(visitor DocumentVisitor) {
	visitor.VisitTable(t)
}

func (t *Table) GetType() string {
	return "Table"
}

func (t *Table) GetContent() string {
	return fmt.Sprintf("Table with %d rows and %d columns", t.rows, t.columns)
}

func (t *Table) GetRows() int {
	return t.rows
}

func (t *Table) GetColumns() int {
	return t.columns
}

func (t *Table) GetData() [][]string {
	return t.data
}

// Concrete Document Visitors
type HTMLExporter struct {
	result string
}

func NewHTMLExporter() *HTMLExporter {
	return &HTMLExporter{result: ""}
}

func (he *HTMLExporter) VisitText(text *Text) {
	he.result += fmt.Sprintf("<p style=\"font-family: %s; font-size: %dpx;\">%s</p>\n", 
		text.GetFont(), text.GetSize(), text.GetContent())
}

func (he *HTMLExporter) VisitImage(image *Image) {
	he.result += fmt.Sprintf("<img src=\"%s\" alt=\"%s\" width=\"%d\" height=\"%d\">\n", 
		image.GetContent(), image.GetAlt(), image.GetWidth(), image.GetHeight())
}

func (he *HTMLExporter) VisitTable(table *Table) {
	he.result += "<table>\n"
	for _, row := range table.GetData() {
		he.result += "  <tr>\n"
		for _, cell := range row {
			he.result += fmt.Sprintf("    <td>%s</td>\n", cell)
		}
		he.result += "  </tr>\n"
	}
	he.result += "</table>\n"
}

func (he *HTMLExporter) GetResult() string {
	return he.result
}

func (he *HTMLExporter) GetName() string {
	return "HTML Exporter"
}

type MarkdownExporter struct {
	result string
}

func NewMarkdownExporter() *MarkdownExporter {
	return &MarkdownExporter{result: ""}
}

func (me *MarkdownExporter) VisitText(text *Text) {
	me.result += fmt.Sprintf("%s\n\n", text.GetContent())
}

func (me *MarkdownExporter) VisitImage(image *Image) {
	me.result += fmt.Sprintf("![%s](%s)\n\n", image.GetAlt(), image.GetContent())
}

func (me *MarkdownExporter) VisitTable(table *Table) {
	me.result += "| "
	for i := 0; i < table.GetColumns(); i++ {
		me.result += fmt.Sprintf("Column %d | ", i+1)
	}
	me.result += "\n| "
	for i := 0; i < table.GetColumns(); i++ {
		me.result += "--- | "
	}
	me.result += "\n"
	
	for _, row := range table.GetData() {
		me.result += "| "
		for _, cell := range row {
			me.result += fmt.Sprintf("%s | ", cell)
		}
		me.result += "\n"
	}
	me.result += "\n"
}

func (me *MarkdownExporter) GetResult() string {
	return me.result
}

func (me *MarkdownExporter) GetName() string {
	return "Markdown Exporter"
}

// Document
type Document struct {
	elements []DocumentElement
}

func NewDocument() *Document {
	return &Document{
		elements: make([]DocumentElement, 0),
	}
}

func (d *Document) AddElement(element DocumentElement) {
	d.elements = append(d.elements, element)
}

func (d *Document) Accept(visitor DocumentVisitor) {
	for _, element := range d.elements {
		element.Accept(visitor)
	}
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== VISITOR PATTERN DEMONSTRATION ===\n")

	// 1. BASIC VISITOR
	fmt.Println("1. BASIC VISITOR:")
	objectStructure := NewObjectStructure()
	objectStructure.AddElement(NewElementA("Element A1"))
	objectStructure.AddElement(NewElementB("Element B1"))
	objectStructure.AddElement(NewElementC("Element C1"))
	
	visitor1 := &ConcreteVisitor1{}
	visitor2 := &ConcreteVisitor2{}
	
	objectStructure.Accept(visitor1)
	objectStructure.Accept(visitor2)
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Shopping Cart Visitor
	fmt.Println("Shopping Cart Visitor:")
	cart := NewShoppingCart()
	cart.AddItem(NewBook("The Go Programming Language", "Alan Donovan", 45.99, 2))
	cart.AddItem(NewElectronics("MacBook Pro", "Apple", 2499.99, 1))
	cart.AddItem(NewClothing("T-Shirt", "L", 19.99, 3))
	
	// Price Calculator
	priceCalculator := NewPriceCalculator()
	cart.Accept(priceCalculator)
	fmt.Printf("Total Price: $%.2f\n", priceCalculator.GetTotal())
	fmt.Println()
	
	// Tax Calculator
	taxCalculator := NewTaxCalculator(0.08) // 8% tax
	cart.Accept(taxCalculator)
	fmt.Printf("Total with Tax: $%.2f\n", taxCalculator.GetTotal())
	fmt.Println()
	
	// Inventory Manager
	inventoryManager := NewInventoryManager()
	cart.Accept(inventoryManager)
	fmt.Printf("Inventory: %v\n", inventoryManager.GetInventory())
	fmt.Println()

	// Document Processing Visitor
	fmt.Println("Document Processing Visitor:")
	document := NewDocument()
	document.AddElement(NewText("Hello, World!", "Arial", 12))
	document.AddElement(NewImage("image.jpg", "Sample Image", 800, 600))
	document.AddElement(NewTable(2, 3, [][]string{
		{"Name", "Age", "City"},
		{"John", "25", "New York"},
		{"Jane", "30", "Los Angeles"},
	}))
	
	// HTML Exporter
	htmlExporter := NewHTMLExporter()
	document.Accept(htmlExporter)
	fmt.Printf("HTML Export:\n%s\n", htmlExporter.GetResult())
	
	// Markdown Exporter
	markdownExporter := NewMarkdownExporter()
	document.Accept(markdownExporter)
	fmt.Printf("Markdown Export:\n%s\n", markdownExporter.GetResult())
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
