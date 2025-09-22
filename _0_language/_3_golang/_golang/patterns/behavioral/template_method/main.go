package main

import "fmt"

type Game interface {
	initialize()
	startPlay()
	endPlay()
}

type GameTemplate struct{ Game }

func (t GameTemplate) Play() {
	t.initialize()
	t.startPlay()
	t.endPlay()
}

type Cricket struct{}
func (Cricket) initialize() { fmt.Println("Cricket Game Initialized!") }
func (Cricket) startPlay() { fmt.Println("Cricket Game Started. Enjoy!") }
func (Cricket) endPlay() { fmt.Println("Cricket Game Finished!") }

func main() {
	g := GameTemplate{Game: Cricket{}}
	g.Play()
}
