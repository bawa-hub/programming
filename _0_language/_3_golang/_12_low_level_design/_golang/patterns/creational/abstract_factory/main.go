package main

import "fmt"

type Button interface { Paint() }

type Checkbox interface { Paint() }

type WinButton struct{}
func (WinButton) Paint() { fmt.Println("Render a Windows button") }

type MacButton struct{}
func (MacButton) Paint() { fmt.Println("Render a macOS button") }

type WinCheckbox struct{}
func (WinCheckbox) Paint() { fmt.Println("Render a Windows checkbox") }

type MacCheckbox struct{}
func (MacCheckbox) Paint() { fmt.Println("Render a macOS checkbox") }

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WinFactory struct{}
func (WinFactory) CreateButton() Button   { return WinButton{} }
func (WinFactory) CreateCheckbox() Checkbox { return WinCheckbox{} }

type MacFactory struct{}
func (MacFactory) CreateButton() Button   { return MacButton{} }
func (MacFactory) CreateCheckbox() Checkbox { return MacCheckbox{} }

type Application struct{ factory GUIFactory }

func NewApplication(factory GUIFactory) *Application { return &Application{factory: factory} }

func (a *Application) Render() {
	btn := a.factory.CreateButton()
	chk := a.factory.CreateCheckbox()
	btn.Paint()
	chk.Paint()
}

func main() {
	app := NewApplication(WinFactory{})
	app.Render()

	app = NewApplication(MacFactory{})
	app.Render()
}
