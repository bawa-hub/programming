package main

import "fmt"

type CPU struct{}
func (CPU) Freeze() { fmt.Println("CPU freeze") }
func (CPU) Jump(addr int) { fmt.Println("CPU jump", addr) }
func (CPU) Execute() { fmt.Println("CPU execute") }

type Memory struct{}
func (Memory) Load(addr int, data string) { fmt.Println("Memory load", addr, data) }

type Disk struct{}
func (Disk) Read(lba int, size int) string { fmt.Println("Disk read", lba, size); return "data" }

// Facade
type Computer struct { cpu CPU; mem Memory; disk Disk }
func (c Computer) Start() {
	c.cpu.Freeze()
	data := c.disk.Read(0, 1024)
	c.mem.Load(0, data)
	c.cpu.Jump(0)
	c.cpu.Execute()
}

func main() { (Computer{}).Start() }
