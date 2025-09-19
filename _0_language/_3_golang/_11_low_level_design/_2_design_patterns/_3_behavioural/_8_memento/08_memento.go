package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC MEMENTO PATTERN
// =============================================================================

// Memento interface
type Memento interface {
	GetState() interface{}
	GetTimestamp() time.Time
}

// Originator
type Originator struct {
	state string
}

func NewOriginator(initialState string) *Originator {
	return &Originator{state: initialState}
}

func (o *Originator) SetState(state string) {
	o.state = state
	fmt.Printf("Originator: State set to '%s'\n", state)
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) CreateMemento() Memento {
	return &ConcreteMemento{
		state:     o.state,
		timestamp: time.Now(),
	}
}

func (o *Originator) RestoreMemento(memento Memento) {
	o.state = memento.GetState().(string)
	fmt.Printf("Originator: State restored to '%s' (from %s)\n", 
		o.state, memento.GetTimestamp().Format("15:04:05"))
}

// Concrete Memento
type ConcreteMemento struct {
	state     interface{}
	timestamp time.Time
}

func (cm *ConcreteMemento) GetState() interface{} {
	return cm.state
}

func (cm *ConcreteMemento) GetTimestamp() time.Time {
	return cm.timestamp
}

// Caretaker
type Caretaker struct {
	mementos []Memento
}

func NewCaretaker() *Caretaker {
	return &Caretaker{
		mementos: make([]Memento, 0),
	}
}

func (c *Caretaker) AddMemento(memento Memento) {
	c.mementos = append(c.mementos, memento)
	fmt.Printf("Caretaker: Added memento from %s\n", 
		memento.GetTimestamp().Format("15:04:05"))
}

func (c *Caretaker) GetMemento(index int) Memento {
	if index >= 0 && index < len(c.mementos) {
		return c.mementos[index]
	}
	return nil
}

func (c *Caretaker) GetMementoCount() int {
	return len(c.mementos)
}

func (c *Caretaker) GetLastMemento() Memento {
	if len(c.mementos) > 0 {
		return c.mementos[len(c.mementos)-1]
	}
	return nil
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. TEXT EDITOR MEMENTO
type TextEditorMemento interface {
	GetContent() string
	GetCursorPosition() int
	GetTimestamp() time.Time
}

type TextEditor struct {
	content        string
	cursorPosition int
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		content:        "",
		cursorPosition: 0,
	}
}

func (te *TextEditor) InsertText(text string) {
	te.content = te.content[:te.cursorPosition] + text + te.content[te.cursorPosition:]
	te.cursorPosition += len(text)
	fmt.Printf("TextEditor: Inserted '%s' at position %d\n", text, te.cursorPosition)
}

func (te *TextEditor) DeleteText(length int) {
	if te.cursorPosition >= length {
		te.content = te.content[:te.cursorPosition-length] + te.content[te.cursorPosition:]
		te.cursorPosition -= length
		fmt.Printf("TextEditor: Deleted %d characters at position %d\n", length, te.cursorPosition)
	}
}

func (te *TextEditor) MoveCursor(position int) {
	if position >= 0 && position <= len(te.content) {
		te.cursorPosition = position
		fmt.Printf("TextEditor: Moved cursor to position %d\n", te.cursorPosition)
	}
}

func (te *TextEditor) GetContent() string {
	return te.content
}

func (te *TextEditor) GetCursorPosition() int {
	return te.cursorPosition
}

func (te *TextEditor) CreateMemento() TextEditorMemento {
	return &TextEditorMementoImpl{
		content:        te.content,
		cursorPosition: te.cursorPosition,
		timestamp:      time.Now(),
	}
}

func (te *TextEditor) RestoreMemento(memento TextEditorMemento) {
	te.content = memento.GetContent()
	te.cursorPosition = memento.GetCursorPosition()
	fmt.Printf("TextEditor: Restored to '%s' (cursor at %d) from %s\n", 
		te.content, te.cursorPosition, memento.GetTimestamp().Format("15:04:05"))
}

type TextEditorMementoImpl struct {
	content        string
	cursorPosition int
	timestamp      time.Time
}

func (tem *TextEditorMementoImpl) GetContent() string {
	return tem.content
}

func (tem *TextEditorMementoImpl) GetCursorPosition() int {
	return tem.cursorPosition
}

func (tem *TextEditorMementoImpl) GetTimestamp() time.Time {
	return tem.timestamp
}

type TextEditorCaretaker struct {
	mementos []TextEditorMemento
}

func NewTextEditorCaretaker() *TextEditorCaretaker {
	return &TextEditorCaretaker{
		mementos: make([]TextEditorMemento, 0),
	}
}

func (tec *TextEditorCaretaker) SaveMemento(memento TextEditorMemento) {
	tec.mementos = append(tec.mementos, memento)
	fmt.Printf("TextEditorCaretaker: Saved memento from %s\n", 
		memento.GetTimestamp().Format("15:04:05"))
}

func (tec *TextEditorCaretaker) Undo() TextEditorMemento {
	if len(tec.mementos) > 1 {
		tec.mementos = tec.mementos[:len(tec.mementos)-1]
		return tec.mementos[len(tec.mementos)-1]
	}
	return nil
}

func (tec *TextEditorCaretaker) GetMementoCount() int {
	return len(tec.mementos)
}

// 2. GAME STATE MEMENTO
type GameStateMemento interface {
	GetLevel() int
	GetScore() int
	GetHealth() int
	GetPosition() (int, int)
	GetTimestamp() time.Time
}

type GameState struct {
	level    int
	score    int
	health   int
	positionX int
	positionY int
}

func NewGameState(level, score, health, x, y int) *GameState {
	return &GameState{
		level:     level,
		score:     score,
		health:    health,
		positionX: x,
		positionY: y,
	}
}

func (gs *GameState) UpdateScore(points int) {
	gs.score += points
	fmt.Printf("GameState: Score updated to %d\n", gs.score)
}

func (gs *GameState) TakeDamage(damage int) {
	gs.health -= damage
	if gs.health < 0 {
		gs.health = 0
	}
	fmt.Printf("GameState: Health reduced to %d\n", gs.health)
}

func (gs *GameState) MoveTo(x, y int) {
	gs.positionX = x
	gs.positionY = y
	fmt.Printf("GameState: Moved to position (%d, %d)\n", x, y)
}

func (gs *GameState) LevelUp() {
	gs.level++
	fmt.Printf("GameState: Leveled up to %d\n", gs.level)
}

func (gs *GameState) GetLevel() int {
	return gs.level
}

func (gs *GameState) GetScore() int {
	return gs.score
}

func (gs *GameState) GetHealth() int {
	return gs.health
}

func (gs *GameState) GetPosition() (int, int) {
	return gs.positionX, gs.positionY
}

func (gs *GameState) CreateMemento() GameStateMemento {
	return &GameStateMementoImpl{
		level:     gs.level,
		score:     gs.score,
		health:    gs.health,
		positionX: gs.positionX,
		positionY: gs.positionY,
		timestamp: time.Now(),
	}
}

func (gs *GameState) RestoreMemento(memento GameStateMemento) {
	gs.level = memento.GetLevel()
	gs.score = memento.GetScore()
	gs.health = memento.GetHealth()
	gs.positionX, gs.positionY = memento.GetPosition()
	fmt.Printf("GameState: Restored to level %d, score %d, health %d, position (%d, %d) from %s\n", 
		gs.level, gs.score, gs.health, gs.positionX, gs.positionY, 
		memento.GetTimestamp().Format("15:04:05"))
}

type GameStateMementoImpl struct {
	level     int
	score     int
	health    int
	positionX int
	positionY int
	timestamp time.Time
}

func (gsm *GameStateMementoImpl) GetLevel() int {
	return gsm.level
}

func (gsm *GameStateMementoImpl) GetScore() int {
	return gsm.score
}

func (gsm *GameStateMementoImpl) GetHealth() int {
	return gsm.health
}

func (gsm *GameStateMementoImpl) GetPosition() (int, int) {
	return gsm.positionX, gsm.positionY
}

func (gsm *GameStateMementoImpl) GetTimestamp() time.Time {
	return gsm.timestamp
}

type GameSaveManager struct {
	saves []GameStateMemento
}

func NewGameSaveManager() *GameSaveManager {
	return &GameSaveManager{
		saves: make([]GameStateMemento, 0),
	}
}

func (gsm *GameSaveManager) SaveGame(memento GameStateMemento) {
	gsm.saves = append(gsm.saves, memento)
	fmt.Printf("GameSaveManager: Saved game state from %s\n", 
		memento.GetTimestamp().Format("15:04:05"))
}

func (gsm *GameSaveManager) LoadGame(index int) GameStateMemento {
	if index >= 0 && index < len(gsm.saves) {
		return gsm.saves[index]
	}
	return nil
}

func (gsm *GameSaveManager) GetSaveCount() int {
	return len(gsm.saves)
}

func (gsm *GameSaveManager) ListSaves() {
	fmt.Printf("GameSaveManager: %d saves available:\n", len(gsm.saves))
	for i, save := range gsm.saves {
		fmt.Printf("  Save %d: Level %d, Score %d, Health %d, Position (%d, %d) - %s\n", 
			i, save.GetLevel(), save.GetScore(), save.GetHealth(), 
			save.GetPosition(), save.GetTimestamp().Format("15:04:05"))
	}
}

// 3. CONFIGURATION MEMENTO
type ConfigurationMemento interface {
	GetSettings() map[string]interface{}
	GetTimestamp() time.Time
}

type Configuration struct {
	settings map[string]interface{}
}

func NewConfiguration() *Configuration {
	return &Configuration{
		settings: make(map[string]interface{}),
	}
}

func (c *Configuration) SetSetting(key string, value interface{}) {
	c.settings[key] = value
	fmt.Printf("Configuration: Set %s = %v\n", key, value)
}

func (c *Configuration) GetSetting(key string) interface{} {
	return c.settings[key]
}

func (c *Configuration) GetSettings() map[string]interface{} {
	return c.settings
}

func (c *Configuration) CreateMemento() ConfigurationMemento {
	// Deep copy settings
	settingsCopy := make(map[string]interface{})
	for k, v := range c.settings {
		settingsCopy[k] = v
	}
	
	return &ConfigurationMementoImpl{
		settings:  settingsCopy,
		timestamp: time.Now(),
	}
}

func (c *Configuration) RestoreMemento(memento ConfigurationMemento) {
	c.settings = memento.GetSettings()
	fmt.Printf("Configuration: Restored settings from %s\n", 
		memento.GetTimestamp().Format("15:04:05"))
}

type ConfigurationMementoImpl struct {
	settings  map[string]interface{}
	timestamp time.Time
}

func (cm *ConfigurationMementoImpl) GetSettings() map[string]interface{} {
	return cm.settings
}

func (cm *ConfigurationMementoImpl) GetTimestamp() time.Time {
	return cm.timestamp
}

type ConfigurationManager struct {
	configurations []ConfigurationMemento
}

func NewConfigurationManager() *ConfigurationManager {
	return &ConfigurationManager{
		configurations: make([]ConfigurationMemento, 0),
	}
}

func (cm *ConfigurationManager) SaveConfiguration(memento ConfigurationMemento) {
	cm.configurations = append(cm.configurations, memento)
	fmt.Printf("ConfigurationManager: Saved configuration from %s\n", 
		memento.GetTimestamp().Format("15:04:05"))
}

func (cm *ConfigurationManager) RestoreConfiguration(index int) ConfigurationMemento {
	if index >= 0 && index < len(cm.configurations) {
		return cm.configurations[index]
	}
	return nil
}

func (cm *ConfigurationManager) GetConfigurationCount() int {
	return len(cm.configurations)
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== MEMENTO PATTERN DEMONSTRATION ===\n")

	// 1. BASIC MEMENTO
	fmt.Println("1. BASIC MEMENTO:")
	originator := NewOriginator("Initial State")
	caretaker := NewCaretaker()
	
	// Save initial state
	caretaker.AddMemento(originator.CreateMemento())
	
	// Change state
	originator.SetState("Modified State")
	caretaker.AddMemento(originator.CreateMemento())
	
	// Change state again
	originator.SetState("Another State")
	caretaker.AddMemento(originator.CreateMemento())
	
	// Restore to previous state
	originator.RestoreMemento(caretaker.GetMemento(1))
	
	// Restore to initial state
	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Text Editor Memento
	fmt.Println("Text Editor Memento:")
	textEditor := NewTextEditor()
	textCaretaker := NewTextEditorCaretaker()
	
	// Save initial state
	textCaretaker.SaveMemento(textEditor.CreateMemento())
	
	// Make changes
	textEditor.InsertText("Hello, ")
	textCaretaker.SaveMemento(textEditor.CreateMemento())
	
	textEditor.InsertText("World!")
	textCaretaker.SaveMemento(textEditor.CreateMemento())
	
	textEditor.MoveCursor(6)
	textEditor.InsertText("Beautiful ")
	textCaretaker.SaveMemento(textEditor.CreateMemento())
	
	// Undo changes
	textEditor.RestoreMemento(textCaretaker.Undo())
	textEditor.RestoreMemento(textCaretaker.Undo())
	textEditor.RestoreMemento(textCaretaker.Undo())
	fmt.Println()

	// Game State Memento
	fmt.Println("Game State Memento:")
	gameState := NewGameState(1, 0, 100, 0, 0)
	gameSaveManager := NewGameSaveManager()
	
	// Save initial state
	gameSaveManager.SaveGame(gameState.CreateMemento())
	
	// Play the game
	gameState.UpdateScore(100)
	gameState.MoveTo(10, 20)
	gameSaveManager.SaveGame(gameState.CreateMemento())
	
	gameState.UpdateScore(50)
	gameState.TakeDamage(20)
	gameState.LevelUp()
	gameState.MoveTo(30, 40)
	gameSaveManager.SaveGame(gameState.CreateMemento())
	
	gameState.UpdateScore(200)
	gameState.TakeDamage(50)
	gameState.MoveTo(50, 60)
	gameSaveManager.SaveGame(gameState.CreateMemento())
	
	// List saves
	gameSaveManager.ListSaves()
	
	// Load previous save
	gameState.RestoreMemento(gameSaveManager.LoadGame(2))
	fmt.Println()

	// Configuration Memento
	fmt.Println("Configuration Memento:")
	configuration := NewConfiguration()
	configManager := NewConfigurationManager()
	
	// Save initial configuration
	configManager.SaveConfiguration(configuration.CreateMemento())
	
	// Modify configuration
	configuration.SetSetting("theme", "dark")
	configuration.SetSetting("language", "en")
	configuration.SetSetting("font_size", 14)
	configManager.SaveConfiguration(configuration.CreateMemento())
	
	configuration.SetSetting("theme", "light")
	configuration.SetSetting("font_size", 16)
	configuration.SetSetting("auto_save", true)
	configManager.SaveConfiguration(configuration.CreateMemento())
	
	// Restore previous configuration
	configuration.RestoreMemento(configManager.RestoreConfiguration(1))
	fmt.Printf("Current configuration: %v\n", configuration.GetSettings())
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
