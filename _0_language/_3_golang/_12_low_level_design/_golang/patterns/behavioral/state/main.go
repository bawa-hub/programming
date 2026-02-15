package main

import "fmt"

type State interface { Handle(ctx *Context) }

type Context struct{ state State }
func (c *Context) SetState(s State) { c.state = s }
func (c *Context) Request() { c.state.Handle(c) }

type Locked struct{}
func (Locked) Handle(c *Context) { fmt.Println("Locked -> unlocking"); c.SetState(Unlocked{}) }

type Unlocked struct{}
func (Unlocked) Handle(c *Context) { fmt.Println("Unlocked -> locking"); c.SetState(Locked{}) }

func main() {
	ctx := &Context{state: Locked{}}
	ctx.Request()
	ctx.Request()
}
