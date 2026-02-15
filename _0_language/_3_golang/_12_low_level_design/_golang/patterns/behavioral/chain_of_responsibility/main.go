package main

import "fmt"

type Handler interface {
	SetNext(h Handler)
	Handle(req int)
}

type BaseHandler struct{ next Handler }

func (b *BaseHandler) SetNext(h Handler) { b.next = h }
func (b *BaseHandler) pass(req int) { if b.next != nil { b.next.Handle(req) } }

type LowHandler struct{ BaseHandler }
func (h *LowHandler) Handle(req int) {
	if req < 10 { fmt.Println("Low handled", req); return }
	h.pass(req)
}

type MidHandler struct{ BaseHandler }
func (h *MidHandler) Handle(req int) {
	if req < 100 { fmt.Println("Mid handled", req); return }
	h.pass(req)
}

type HighHandler struct{ BaseHandler }
func (h *HighHandler) Handle(req int) {
	fmt.Println("High handled", req)
}

func main() {
	l, m, hi := &LowHandler{}, &MidHandler{}, &HighHandler{}
	l.SetNext(m); m.SetNext(hi)
	l.Handle(5); l.Handle(50); l.Handle(500)
}
