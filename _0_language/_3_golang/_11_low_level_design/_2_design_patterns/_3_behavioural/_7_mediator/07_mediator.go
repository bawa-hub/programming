package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// BASIC MEDIATOR PATTERN
// =============================================================================

// Mediator interface
type Mediator interface {
	Notify(sender Colleague, event string)
	AddColleague(colleague Colleague)
	RemoveColleague(colleague Colleague)
}

// Colleague interface
type Colleague interface {
	SetMediator(mediator Mediator)
	Notify(event string)
	GetName() string
}

// Concrete Mediator
type ConcreteMediator struct {
	colleagues []Colleague
	mu         sync.RWMutex
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{
		colleagues: make([]Colleague, 0),
	}
}

func (cm *ConcreteMediator) Notify(sender Colleague, event string) {
	cm.mu.RLock()
	colleagues := make([]Colleague, len(cm.colleagues))
	copy(colleagues, cm.colleagues)
	cm.mu.RUnlock()
	
	fmt.Printf("Mediator: %s sent event '%s'\n", sender.GetName(), event)
	for _, colleague := range colleagues {
		if colleague != sender {
			colleague.Notify(event)
		}
	}
}

func (cm *ConcreteMediator) AddColleague(colleague Colleague) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.colleagues = append(cm.colleagues, colleague)
	colleague.SetMediator(cm)
	fmt.Printf("Mediator: Added colleague %s\n", colleague.GetName())
}

func (cm *ConcreteMediator) RemoveColleague(colleague Colleague) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	
	for i, c := range cm.colleagues {
		if c == colleague {
			cm.colleagues = append(cm.colleagues[:i], cm.colleagues[i+1:]...)
			break
		}
	}
	fmt.Printf("Mediator: Removed colleague %s\n", colleague.GetName())
}

// Concrete Colleagues
type ConcreteColleagueA struct {
	name     string
	mediator Mediator
}

func NewConcreteColleagueA(name string) *ConcreteColleagueA {
	return &ConcreteColleagueA{name: name}
}

func (cca *ConcreteColleagueA) SetMediator(mediator Mediator) {
	cca.mediator = mediator
}

func (cca *ConcreteColleagueA) Notify(event string) {
	fmt.Printf("ColleagueA %s: Received event '%s'\n", cca.name, event)
}

func (cca *ConcreteColleagueA) GetName() string {
	return cca.name
}

func (cca *ConcreteColleagueA) SendEvent(event string) {
	if cca.mediator != nil {
		cca.mediator.Notify(cca, event)
	}
}

type ConcreteColleagueB struct {
	name     string
	mediator Mediator
}

func NewConcreteColleagueB(name string) *ConcreteColleagueB {
	return &ConcreteColleagueB{name: name}
}

func (ccb *ConcreteColleagueB) SetMediator(mediator Mediator) {
	ccb.mediator = mediator
}

func (ccb *ConcreteColleagueB) Notify(event string) {
	fmt.Printf("ColleagueB %s: Received event '%s'\n", ccb.name, event)
}

func (ccb *ConcreteColleagueB) GetName() string {
	return ccb.name
}

func (ccb *ConcreteColleagueB) SendEvent(event string) {
	if ccb.mediator != nil {
		ccb.mediator.Notify(ccb, event)
	}
}

// =============================================================================
// REAL-WORLD EXAMPLES
// =============================================================================

// 1. CHAT ROOM MEDIATOR
type ChatUser interface {
	SetMediator(mediator ChatMediator)
	SendMessage(message string)
	ReceiveMessage(sender string, message string)
	GetName() string
	GetStatus() string
}

type ChatMediator interface {
	AddUser(user ChatUser)
	RemoveUser(user ChatUser)
	SendMessage(sender ChatUser, message string)
	BroadcastMessage(sender ChatUser, message string)
	GetUsers() []ChatUser
}

type ChatRoom struct {
	users []ChatUser
	mu    sync.RWMutex
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		users: make([]ChatUser, 0),
	}
}

func (cr *ChatRoom) AddUser(user ChatUser) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	cr.users = append(cr.users, user)
	user.SetMediator(cr)
	fmt.Printf("Chat Room: %s joined the chat\n", user.GetName())
}

func (cr *ChatRoom) RemoveUser(user ChatUser) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	
	for i, u := range cr.users {
		if u == user {
			cr.users = append(cr.users[:i], cr.users[i+1:]...)
			break
		}
	}
	fmt.Printf("Chat Room: %s left the chat\n", user.GetName())
}

func (cr *ChatRoom) SendMessage(sender ChatUser, message string) {
	cr.mu.RLock()
	users := make([]ChatUser, len(cr.users))
	copy(users, cr.users)
	cr.mu.RUnlock()
	
	fmt.Printf("Chat Room: %s sent: %s\n", sender.GetName(), message)
	for _, user := range users {
		if user != sender {
			user.ReceiveMessage(sender.GetName(), message)
		}
	}
}

func (cr *ChatRoom) BroadcastMessage(sender ChatUser, message string) {
	cr.mu.RLock()
	users := make([]ChatUser, len(cr.users))
	copy(users, cr.users)
	cr.mu.RUnlock()
	
	fmt.Printf("Chat Room: %s broadcasted: %s\n", sender.GetName(), message)
	for _, user := range users {
		if user != sender {
			user.ReceiveMessage(sender.GetName(), message)
		}
	}
}

func (cr *ChatRoom) GetUsers() []ChatUser {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	return cr.users
}

// Concrete Chat Users
type ChatUserImpl struct {
	name   string
	status string
	mediator ChatMediator
}

func NewChatUser(name string) *ChatUserImpl {
	return &ChatUserImpl{
		name:   name,
		status: "online",
	}
}

func (cu *ChatUserImpl) SetMediator(mediator ChatMediator) {
	cu.mediator = mediator
}

func (cu *ChatUserImpl) SendMessage(message string) {
	if cu.mediator != nil {
		cu.mediator.SendMessage(cu, message)
	}
}

func (cu *ChatUserImpl) ReceiveMessage(sender string, message string) {
	fmt.Printf("  %s received from %s: %s\n", cu.name, sender, message)
}

func (cu *ChatUserImpl) GetName() string {
	return cu.name
}

func (cu *ChatUserImpl) GetStatus() string {
	return cu.status
}

func (cu *ChatUserImpl) SetStatus(status string) {
	cu.status = status
}

// 2. AIR TRAFFIC CONTROL MEDIATOR
type Aircraft interface {
	SetMediator(mediator AirTrafficControl)
	RequestLanding()
	RequestTakeoff()
	ReceiveClearance(clearance string)
	GetName() string
	GetStatus() string
}

type AirTrafficControl interface {
	AddAircraft(aircraft Aircraft)
	RemoveAircraft(aircraft Aircraft)
	RequestLanding(aircraft Aircraft)
	RequestTakeoff(aircraft Aircraft)
	GetAircraft() []Aircraft
}

type AirTrafficControlTower struct {
	aircraft []Aircraft
	mu       sync.RWMutex
}

func NewAirTrafficControlTower() *AirTrafficControlTower {
	return &AirTrafficControlTower{
		aircraft: make([]Aircraft, 0),
	}
}

func (atc *AirTrafficControlTower) AddAircraft(aircraft Aircraft) {
	atc.mu.Lock()
	defer atc.mu.Unlock()
	atc.aircraft = append(atc.aircraft, aircraft)
	aircraft.SetMediator(atc)
	fmt.Printf("ATC: %s registered\n", aircraft.GetName())
}

func (atc *AirTrafficControlTower) RemoveAircraft(aircraft Aircraft) {
	atc.mu.Lock()
	defer atc.mu.Unlock()
	
	for i, a := range atc.aircraft {
		if a == aircraft {
			atc.aircraft = append(atc.aircraft[:i], atc.aircraft[i+1:]...)
			break
		}
	}
	fmt.Printf("ATC: %s unregistered\n", aircraft.GetName())
}

func (atc *AirTrafficControlTower) RequestLanding(aircraft Aircraft) {
	atc.mu.RLock()
	aircraftList := make([]Aircraft, len(atc.aircraft))
	copy(aircraftList, atc.aircraft)
	atc.mu.RUnlock()
	
	fmt.Printf("ATC: %s requesting landing\n", aircraft.GetName())
	
	// Check if runway is clear
	runwayClear := true
	for _, a := range aircraftList {
		if a != aircraft && a.GetStatus() == "landing" {
			runwayClear = false
			break
		}
	}
	
	if runwayClear {
		aircraft.ReceiveClearance("Cleared to land")
	} else {
		aircraft.ReceiveClearance("Hold position, runway busy")
	}
}

func (atc *AirTrafficControlTower) RequestTakeoff(aircraft Aircraft) {
	atc.mu.RLock()
	aircraftList := make([]Aircraft, len(atc.aircraft))
	copy(aircraftList, atc.aircraft)
	atc.mu.RUnlock()
	
	fmt.Printf("ATC: %s requesting takeoff\n", aircraft.GetName())
	
	// Check if runway is clear
	runwayClear := true
	for _, a := range aircraftList {
		if a != aircraft && a.GetStatus() == "taking_off" {
			runwayClear = false
			break
		}
	}
	
	if runwayClear {
		aircraft.ReceiveClearance("Cleared for takeoff")
	} else {
		aircraft.ReceiveClearance("Hold position, runway busy")
	}
}

func (atc *AirTrafficControlTower) GetAircraft() []Aircraft {
	atc.mu.RLock()
	defer atc.mu.RUnlock()
	return atc.aircraft
}

// Concrete Aircraft
type AircraftImpl struct {
	name   string
	status string
	mediator AirTrafficControl
}

func NewAircraft(name string) *AircraftImpl {
	return &AircraftImpl{
		name:   name,
		status: "flying",
	}
}

func (a *AircraftImpl) SetMediator(mediator AirTrafficControl) {
	a.mediator = mediator
}

func (a *AircraftImpl) RequestLanding() {
	if a.mediator != nil {
		a.mediator.RequestLanding(a)
	}
}

func (a *AircraftImpl) RequestTakeoff() {
	if a.mediator != nil {
		a.mediator.RequestTakeoff(a)
	}
}

func (a *AircraftImpl) ReceiveClearance(clearance string) {
	fmt.Printf("  %s received clearance: %s\n", a.name, clearance)
	if clearance == "Cleared to land" {
		a.status = "landing"
	} else if clearance == "Cleared for takeoff" {
		a.status = "taking_off"
	}
}

func (a *AircraftImpl) GetName() string {
	return a.name
}

func (a *AircraftImpl) GetStatus() string {
	return a.status
}

func (a *AircraftImpl) SetStatus(status string) {
	a.status = status
}

// 3. GUI COMPONENT MEDIATOR
type GUIComponent interface {
	SetMediator(mediator GUIMediator)
	Click()
	ReceiveEvent(event string)
	GetName() string
	GetType() string
}

type GUIMediator interface {
	AddComponent(component GUIComponent)
	RemoveComponent(component GUIComponent)
	Notify(component GUIComponent, event string)
	GetComponents() []GUIComponent
}

type GUIMediatorImpl struct {
	components []GUIComponent
	mu         sync.RWMutex
}

func NewGUIMediator() *GUIMediatorImpl {
	return &GUIMediatorImpl{
		components: make([]GUIComponent, 0),
	}
}

func (gm *GUIMediatorImpl) AddComponent(component GUIComponent) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.components = append(gm.components, component)
	component.SetMediator(gm)
	fmt.Printf("GUI Mediator: Added %s %s\n", component.GetType(), component.GetName())
}

func (gm *GUIMediatorImpl) RemoveComponent(component GUIComponent) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	
	for i, c := range gm.components {
		if c == component {
			gm.components = append(gm.components[:i], gm.components[i+1:]...)
			break
		}
	}
	fmt.Printf("GUI Mediator: Removed %s %s\n", component.GetType(), component.GetName())
}

func (gm *GUIMediatorImpl) Notify(component GUIComponent, event string) {
	gm.mu.RLock()
	components := make([]GUIComponent, len(gm.components))
	copy(components, gm.components)
	gm.mu.RUnlock()
	
	fmt.Printf("GUI Mediator: %s %s triggered event '%s'\n", 
		component.GetType(), component.GetName(), event)
	
	for _, c := range components {
		if c != component {
			c.ReceiveEvent(event)
		}
	}
}

func (gm *GUIMediatorImpl) GetComponents() []GUIComponent {
	gm.mu.RLock()
	defer gm.mu.RUnlock()
	return gm.components
}

// Concrete GUI Components
type Button struct {
	name     string
	mediator GUIMediator
}

func NewButton(name string) *Button {
	return &Button{name: name}
}

func (b *Button) SetMediator(mediator GUIMediator) {
	b.mediator = mediator
}

func (b *Button) Click() {
	if b.mediator != nil {
		b.mediator.Notify(b, "button_clicked")
	}
}

func (b *Button) ReceiveEvent(event string) {
	fmt.Printf("  Button %s received event: %s\n", b.name, event)
}

func (b *Button) GetName() string {
	return b.name
}

func (b *Button) GetType() string {
	return "Button"
}

type TextField struct {
	name     string
	mediator GUIMediator
}

func NewTextField(name string) *TextField {
	return &TextField{name: name}
}

func (tf *TextField) SetMediator(mediator GUIMediator) {
	tf.mediator = mediator
}

func (tf *TextField) Click() {
	if tf.mediator != nil {
		tf.mediator.Notify(tf, "text_field_focused")
	}
}

func (tf *TextField) ReceiveEvent(event string) {
	fmt.Printf("  TextField %s received event: %s\n", tf.name, event)
}

func (tf *TextField) GetName() string {
	return tf.name
}

func (tf *TextField) GetType() string {
	return "TextField"
}

type Label struct {
	name     string
	mediator GUIMediator
}

func NewLabel(name string) *Label {
	return &Label{name: name}
}

func (l *Label) SetMediator(mediator GUIMediator) {
	l.mediator = mediator
}

func (l *Label) Click() {
	if l.mediator != nil {
		l.mediator.Notify(l, "label_clicked")
	}
}

func (l *Label) ReceiveEvent(event string) {
	fmt.Printf("  Label %s received event: %s\n", l.name, event)
}

func (l *Label) GetName() string {
	return l.name
}

func (l *Label) GetType() string {
	return "Label"
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATION
// =============================================================================

func main() {
	fmt.Println("=== MEDIATOR PATTERN DEMONSTRATION ===\n")

	// 1. BASIC MEDIATOR
	fmt.Println("1. BASIC MEDIATOR:")
	mediator := NewConcreteMediator()
	
	colleagueA1 := NewConcreteColleagueA("A1")
	colleagueA2 := NewConcreteColleagueA("A2")
	colleagueB1 := NewConcreteColleagueB("B1")
	
	mediator.AddColleague(colleagueA1)
	mediator.AddColleague(colleagueA2)
	mediator.AddColleague(colleagueB1)
	
	colleagueA1.SendEvent("Hello from A1")
	colleagueB1.SendEvent("Hello from B1")
	
	mediator.RemoveColleague(colleagueA2)
	colleagueA1.SendEvent("A2 is gone")
	fmt.Println()

	// 2. REAL-WORLD EXAMPLES
	fmt.Println("2. REAL-WORLD EXAMPLES:")

	// Chat Room Mediator
	fmt.Println("Chat Room Mediator:")
	chatRoom := NewChatRoom()
	
	user1 := NewChatUser("Alice")
	user2 := NewChatUser("Bob")
	user3 := NewChatUser("Charlie")
	
	chatRoom.AddUser(user1)
	chatRoom.AddUser(user2)
	chatRoom.AddUser(user3)
	
	user1.SendMessage("Hello everyone!")
	user2.SendMessage("Hi Alice!")
	user3.SendMessage("Hey guys!")
	
	chatRoom.RemoveUser(user2)
	user1.SendMessage("Bob left the chat")
	fmt.Println()

	// Air Traffic Control Mediator
	fmt.Println("Air Traffic Control Mediator:")
	atc := NewAirTrafficControlTower()
	
	aircraft1 := NewAircraft("Flight 123")
	aircraft2 := NewAircraft("Flight 456")
	aircraft3 := NewAircraft("Flight 789")
	
	atc.AddAircraft(aircraft1)
	atc.AddAircraft(aircraft2)
	atc.AddAircraft(aircraft3)
	
	aircraft1.RequestLanding()
	aircraft2.RequestLanding()
	aircraft3.RequestTakeoff()
	
	aircraft1.SetStatus("landed")
	aircraft2.RequestLanding()
	fmt.Println()

	// GUI Component Mediator
	fmt.Println("GUI Component Mediator:")
	guiMediator := NewGUIMediator()
	
	button1 := NewButton("Submit")
	button2 := NewButton("Cancel")
	textField1 := NewTextField("Username")
	textField2 := NewTextField("Password")
	label1 := NewLabel("Status")
	
	guiMediator.AddComponent(button1)
	guiMediator.AddComponent(button2)
	guiMediator.AddComponent(textField1)
	guiMediator.AddComponent(textField2)
	guiMediator.AddComponent(label1)
	
	button1.Click()
	textField1.Click()
	button2.Click()
	fmt.Println()

	fmt.Println("=== END OF DEMONSTRATION ===")
}
