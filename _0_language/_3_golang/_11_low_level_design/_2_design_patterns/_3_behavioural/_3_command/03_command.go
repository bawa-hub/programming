package main

import (
	"fmt"
	"time"
)

// =============================================================================
// BASIC COMMAND PATTERN
// =============================================================================

// Command interface
type Command interface {
	Execute()
	Undo()
	GetName() string
}

// Receiver - knows how to perform operations
type Receiver struct {
	state string
}

func NewReceiver() *Receiver {
	return &Receiver{state: "Initial State"}
}

func (r *Receiver) Action1() {
	fmt.Printf("Receiver: Performing Action1, current state: %s\n", r.state)
	r.state = "Action1 Completed"
}

func (r *Receiver) Action2() {
	fmt.Printf("Receiver: Performing Action2, current state: %s\n", r.state)
	r.state = "Action2 Completed"
}

func (r *Receiver) Action3() {
	fmt.Printf("Receiver: Performing Action3, current state: %s\n", r.state)
	r.state = "Action3 Completed"
}

func (r *Receiver) GetState() string {
	return r.state
}

// Concrete Commands
type ConcreteCommand1 struct {
	receiver *Receiver
}

func NewConcreteCommand1(receiver *Receiver) *ConcreteCommand1 {
	return &ConcreteCommand1{receiver: receiver}
}

func (cc1 *ConcreteCommand1) Execute() {
	cc1.receiver.Action1()
}

func (cc1 *ConcreteCommand1) Undo() {
	fmt.Printf("Command1: Undoing Action1\n")
	cc1.receiver.state = "Action1 Undone"
}

func (cc1 *ConcreteCommand1) GetName() string {
	return "Command1"
}

type ConcreteCommand2 struct {
	receiver *Receiver
}

func NewConcreteCommand2(receiver *Receiver) *ConcreteCommand2 {
	return &ConcreteCommand2{receiver: receiver}
}

func (cc2 *ConcreteCommand2) Execute() {
	cc2.receiver.Action2()
}

func (cc2 *ConcreteCommand2) Undo() {
	fmt.Printf("Command2: Undoing Action2\n")
	cc2.receiver.state = "Action2 Undone"
}

func (cc2 *ConcreteCommand2) GetName() string {
	return "Command2"
}

// Invoker - invokes commands
type Invoker struct {
	commands []Command
}

func NewInvoker() *Invoker {
	return &Invoker{
		commands: make([]Command, 0),
	}
}

func (i *Invoker) ExecuteCommand(command Command) {
	command.Execute()
	i.commands = append(i.commands, command)
	fmt.Printf("Invoker: Executed %s\n", command.GetName())
}

func (i *Invoker) UndoLastCommand() {
	if len(i.commands) > 0 {
		lastCommand := i.commands[len(i.commands)-1]
		lastCommand.Undo()
		i.commands = i.commands[:len(i.commands)-1]
		fmt.Printf("Invoker: Undone %s\n", lastCommand.GetName())
	} else {
		fmt.Println("Invoker: No commands to undo")
	}
}

func (i *Invoker) GetCommandHistory() []string {
	history := make([]string, len(i.commands))
	for i, cmd := range i.commands {
		history[i] = cmd.GetName()
	}
	return history
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. TEXT EDITOR COMMANDS
type TextEditor struct {
	content string
	cursor  int
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		content: "",
		cursor:  0,
	}
}

func (te *TextEditor) InsertText(text string) {
	te.content = te.content[:te.cursor] + text + te.content[te.cursor:]
	te.cursor += len(text)
	fmt.Printf("TextEditor: Inserted '%s', cursor at %d\n", text, te.cursor)
}

func (te *TextEditor) DeleteText(length int) {
	if te.cursor >= length {
		te.content = te.content[:te.cursor-length] + te.content[te.cursor:]
		te.cursor -= length
		fmt.Printf("TextEditor: Deleted %d characters, cursor at %d\n", length, te.cursor)
	}
}

func (te *TextEditor) MoveCursor(position int) {
	if position >= 0 && position <= len(te.content) {
		te.cursor = position
		fmt.Printf("TextEditor: Moved cursor to %d\n", te.cursor)
	}
}

func (te *TextEditor) GetContent() string {
	return te.content
}

func (te *TextEditor) GetCursor() int {
	return te.cursor
}

type TextCommand interface {
	Execute()
	Undo()
	GetName() string
}

type InsertTextCommand struct {
	editor    *TextEditor
	text      string
	position  int
	executed  bool
}

func NewInsertTextCommand(editor *TextEditor, text string, position int) *InsertTextCommand {
	return &InsertTextCommand{
		editor:   editor,
		text:     text,
		position: position,
		executed: false,
	}
}

func (itc *InsertTextCommand) Execute() {
	if !itc.executed {
		itc.editor.MoveCursor(itc.position)
		itc.editor.InsertText(itc.text)
		itc.executed = true
	}
}

func (itc *InsertTextCommand) Undo() {
	if itc.executed {
		itc.editor.MoveCursor(itc.position)
		itc.editor.DeleteText(len(itc.text))
		itc.executed = false
	}
}

func (itc *InsertTextCommand) GetName() string {
	return fmt.Sprintf("InsertText('%s' at %d)", itc.text, itc.position)
}

type DeleteTextCommand struct {
	editor    *TextEditor
	text      string
	position  int
	executed  bool
}

func NewDeleteTextCommand(editor *TextEditor, position int, length int) *DeleteTextCommand {
	text := editor.content[position : position+length]
	return &DeleteTextCommand{
		editor:   editor,
		text:     text,
		position: position,
		executed: false,
	}
}

func (dtc *DeleteTextCommand) Execute() {
	if !dtc.executed {
		dtc.editor.MoveCursor(dtc.position)
		dtc.editor.DeleteText(len(dtc.text))
		dtc.executed = true
	}
}

func (dtc *DeleteTextCommand) Undo() {
	if dtc.executed {
		dtc.editor.MoveCursor(dtc.position)
		dtc.editor.InsertText(dtc.text)
		dtc.executed = false
	}
}

func (dtc *DeleteTextCommand) GetName() string {
	return fmt.Sprintf("DeleteText('%s' at %d)", dtc.text, dtc.position)
}

// 2. REMOTE CONTROL COMMANDS
type Device interface {
	On()
	Off()
	GetName() string
	GetState() string
}

type Light struct {
	name  string
	state string
}

func NewLight(name string) *Light {
	return &Light{
		name:  name,
		state: "off",
	}
}

func (l *Light) On() {
	l.state = "on"
	fmt.Printf("Light %s is now ON\n", l.name)
}

func (l *Light) Off() {
	l.state = "off"
	fmt.Printf("Light %s is now OFF\n", l.name)
}

func (l *Light) GetName() string {
	return l.name
}

func (l *Light) GetState() string {
	return l.state
}

type TV struct {
	name  string
	state string
}

func NewTV(name string) *TV {
	return &TV{
		name:  name,
		state: "off",
	}
}

func (t *TV) On() {
	t.state = "on"
	fmt.Printf("TV %s is now ON\n", t.name)
}

func (t *TV) Off() {
	t.state = "off"
	fmt.Printf("TV %s is now OFF\n", t.name)
}

func (t *TV) GetName() string {
	return t.name
}

func (t *TV) GetState() string {
	return t.state
}

type DeviceCommand interface {
	Execute()
	Undo()
	GetName() string
}

type TurnOnCommand struct {
	device   Device
	executed bool
}

func NewTurnOnCommand(device Device) *TurnOnCommand {
	return &TurnOnCommand{
		device:   device,
		executed: false,
	}
}

func (toc *TurnOnCommand) Execute() {
	if !toc.executed {
		toc.device.On()
		toc.executed = true
	}
}

func (toc *TurnOnCommand) Undo() {
	if toc.executed {
		toc.device.Off()
		toc.executed = false
	}
}

func (toc *TurnOnCommand) GetName() string {
	return fmt.Sprintf("TurnOn(%s)", toc.device.GetName())
}

type TurnOffCommand struct {
	device   Device
	executed bool
}

func NewTurnOffCommand(device Device) *TurnOffCommand {
	return &TurnOffCommand{
		device:   device,
		executed: false,
	}
}

func (tofc *TurnOffCommand) Execute() {
	if !tofc.executed {
		tofc.device.Off()
		tofc.executed = true
	}
}

func (tofc *TurnOffCommand) Undo() {
	if tofc.executed {
		tofc.device.On()
		tofc.executed = false
	}
}

func (tofc *TurnOffCommand) GetName() string {
	return fmt.Sprintf("TurnOff(%s)", tofc.device.GetName())
}

// Macro Command - composite command
type MacroCommand struct {
	commands []DeviceCommand
	name     string
}

func NewMacroCommand(name string, commands ...DeviceCommand) *MacroCommand {
	return &MacroCommand{
		commands: commands,
		name:     name,
	}
}

func (mc *MacroCommand) Execute() {
	fmt.Printf("Executing macro: %s\n", mc.name)
	for _, cmd := range mc.commands {
		cmd.Execute()
	}
}

func (mc *MacroCommand) Undo() {
	fmt.Printf("Undoing macro: %s\n", mc.name)
	// Undo in reverse order
	for i := len(mc.commands) - 1; i >= 0; i-- {
		mc.commands[i].Undo()
	}
}

func (mc *MacroCommand) GetName() string {
	return mc.name
}

// Remote Control
type RemoteControl struct {
	commands []DeviceCommand
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		commands: make([]DeviceCommand, 0),
	}
}

func (rc *RemoteControl) PressButton(command DeviceCommand) {
	command.Execute()
	rc.commands = append(rc.commands, command)
}

func (rc *RemoteControl) UndoLastCommand() {
	if len(rc.commands) > 0 {
		lastCommand := rc.commands[len(rc.commands)-1]
		lastCommand.Undo()
		rc.commands = rc.commands[:len(rc.commands)-1]
	}
}

// 3. DATABASE TRANSACTION COMMANDS
type Database struct {
	data map[string]interface{}
}

func NewDatabase() *Database {
	return &Database{
		data: make(map[string]interface{}),
	}
}

func (db *Database) Insert(key string, value interface{}) {
	db.data[key] = value
	fmt.Printf("Database: Inserted %s = %v\n", key, value)
}

func (db *Database) Update(key string, value interface{}) {
	if _, exists := db.data[key]; exists {
		db.data[key] = value
		fmt.Printf("Database: Updated %s = %v\n", key, value)
	}
}

func (db *Database) Delete(key string) {
	if _, exists := db.data[key]; exists {
		delete(db.data, key)
		fmt.Printf("Database: Deleted %s\n", key)
	}
}

func (db *Database) Get(key string) interface{} {
	return db.data[key]
}

type DatabaseCommand interface {
	Execute()
	Undo()
	GetName() string
}

type InsertCommand struct {
	database *Database
	key      string
	value    interface{}
	executed bool
}

func NewInsertCommand(database *Database, key string, value interface{}) *InsertCommand {
	return &InsertCommand{
		database: database,
		key:      key,
		value:    value,
		executed: false,
	}
}

func (ic *InsertCommand) Execute() {
	if !ic.executed {
		ic.database.Insert(ic.key, ic.value)
		ic.executed = true
	}
}

func (ic *InsertCommand) Undo() {
	if ic.executed {
		ic.database.Delete(ic.key)
		ic.executed = false
	}
}

func (ic *InsertCommand) GetName() string {
	return fmt.Sprintf("Insert(%s, %v)", ic.key, ic.value)
}

type UpdateCommand struct {
	database *Database
	key      string
	value    interface{}
	oldValue interface{}
	executed bool
}

func NewUpdateCommand(database *Database, key string, value interface{}) *UpdateCommand {
	oldValue := database.Get(key)
	return &UpdateCommand{
		database: database,
		key:      key,
		value:    value,
		oldValue: oldValue,
		executed: false,
	}
}

func (uc *UpdateCommand) Execute() {
	if !uc.executed {
		uc.database.Update(uc.key, uc.value)
		uc.executed = true
	}
}

func (uc *UpdateCommand) Undo() {
	if uc.executed {
		uc.database.Update(uc.key, uc.oldValue)
		uc.executed = false
	}
}

func (uc *UpdateCommand) GetName() string {
	return fmt.Sprintf("Update(%s, %v)", uc.key, uc.value)
}

type DeleteCommand struct {
	database *Database
	key      string
	value    interface{}
	executed bool
}

func NewDeleteCommand(database *Database, key string) *DeleteCommand {
	value := database.Get(key)
	return &DeleteCommand{
		database: database,
		key:      key,
		value:    value,
		executed: false,
	}
}

func (dc *DeleteCommand) Execute() {
	if !dc.executed {
		dc.database.Delete(dc.key)
		dc.executed = true
	}
}

func (dc *DeleteCommand) Undo() {
	if dc.executed {
		dc.database.Insert(dc.key, dc.value)
		dc.executed = false
	}
}

func (dc *DeleteCommand) GetName() string {
	return fmt.Sprintf("Delete(%s)", dc.key)
}

// Transaction Manager
type TransactionManager struct {
	commands []DatabaseCommand
	database *Database
}

func NewTransactionManager(database *Database) *TransactionManager {
	return &TransactionManager{
		commands: make([]DatabaseCommand, 0),
		database: database,
	}
}

func (tm *TransactionManager) AddCommand(command DatabaseCommand) {
	tm.commands = append(tm.commands, command)
}

func (tm *TransactionManager) ExecuteTransaction() {
	fmt.Println("Executing transaction...")
	for _, cmd := range tm.commands {
		cmd.Execute()
	}
}

func (tm *TransactionManager) RollbackTransaction() {
	fmt.Println("Rolling back transaction...")
	for i := len(tm.commands) - 1; i >= 0; i-- {
		tm.commands[i].Undo()
	}
}

func (tm *TransactionManager) ClearTransaction() {
	tm.commands = make([]DatabaseCommand, 0)
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== COMMAND PATTERN DEMONSTRATION ===\n")

	// 1. BASIC COMMAND
	fmt.Println("1. BASIC COMMAND:")
	receiver := NewReceiver()
	invoker := NewInvoker()
	
	command1 := NewConcreteCommand1(receiver)
	command2 := NewConcreteCommand2(receiver)
	
	invoker.ExecuteCommand(command1)
	invoker.ExecuteCommand(command2)
	
	fmt.Printf("Receiver state: %s\n", receiver.GetState())
	fmt.Printf("Command history: %v\n", invoker.GetCommandHistory())
	
	invoker.UndoLastCommand()
	invoker.UndoLastCommand()
	invoker.UndoLastCommand() // No more commands to undo
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Text Editor Commands
	fmt.Println("Text Editor Commands:")
	editor := NewTextEditor()
	
	insertCmd1 := NewInsertTextCommand(editor, "Hello", 0)
	insertCmd2 := NewInsertTextCommand(editor, " World", 5)
	insertCmd3 := NewInsertTextCommand(editor, "!", 11)
	
	insertCmd1.Execute()
	insertCmd2.Execute()
	insertCmd3.Execute()
	
	fmt.Printf("Content: '%s'\n", editor.GetContent())
	
	insertCmd3.Undo()
	insertCmd2.Undo()
	insertCmd1.Undo()
	
	fmt.Printf("Content after undo: '%s'\n", editor.GetContent())
	fmt.Println()

	// Remote Control Commands
	fmt.Println("Remote Control Commands:")
	light1 := NewLight("Living Room")
	light2 := NewLight("Bedroom")
	tv := NewTV("Living Room")
	
	remote := NewRemoteControl()
	
	// Turn on devices
	remote.PressButton(NewTurnOnCommand(light1))
	remote.PressButton(NewTurnOnCommand(tv))
	remote.PressButton(NewTurnOnCommand(light2))
	
	// Undo last command
	remote.UndoLastCommand()
	
	// Create macro command
	macroCmd := NewMacroCommand("Movie Night", 
		NewTurnOnCommand(light1),
		NewTurnOnCommand(tv),
		NewTurnOffCommand(light2),
	)
	
	remote.PressButton(macroCmd)
	remote.UndoLastCommand()
	fmt.Println()

	// Database Transaction Commands
	fmt.Println("Database Transaction Commands:")
	database := NewDatabase()
	transactionManager := NewTransactionManager(database)
	
	// Add commands to transaction
	transactionManager.AddCommand(NewInsertCommand(database, "user1", "John"))
	transactionManager.AddCommand(NewInsertCommand(database, "user2", "Jane"))
	transactionManager.AddCommand(NewUpdateCommand(database, "user1", "John Doe"))
	transactionManager.AddCommand(NewDeleteCommand(database, "user2"))
	
	// Execute transaction
	transactionManager.ExecuteTransaction()
	
	// Rollback transaction
	transactionManager.RollbackTransaction()
	
	// Clear transaction
	transactionManager.ClearTransaction()
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
