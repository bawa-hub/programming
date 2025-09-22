package main

import "fmt"

type Command interface { Execute() }

type Light struct{ on bool }
func (l *Light) On() { l.on = true; fmt.Println("Light on") }
func (l *Light) Off() { l.on = false; fmt.Println("Light off") }

type OnCommand struct{ light *Light }
func (c OnCommand) Execute() { c.light.On() }

type OffCommand struct{ light *Light }
func (c OffCommand) Execute() { c.light.Off() }

type Remote struct{ history []Command }
func (r *Remote) Submit(c Command) { c.Execute(); r.history = append(r.history, c) }

func main() {
	l := &Light{}
	r := &Remote{}
	r.Submit(OnCommand{light: l})
	r.Submit(OffCommand{light: l})
}
