package main

import "fmt"

type Renderer interface { RenderCircle(radius float64) }

type VectorRenderer struct{}
func (VectorRenderer) RenderCircle(radius float64) { fmt.Println("Vector circle of radius", radius) }

type RasterRenderer struct{}
func (RasterRenderer) RenderCircle(radius float64) { fmt.Println("Raster circle of radius", radius) }

type Circle struct {
	r Renderer
	radius float64
}

func NewCircle(r Renderer, radius float64) *Circle { return &Circle{r: r, radius: radius} }

func (c *Circle) Draw() { c.r.RenderCircle(c.radius) }

func main() {
	NewCircle(VectorRenderer{}, 5).Draw()
	NewCircle(RasterRenderer{}, 10).Draw()
}
