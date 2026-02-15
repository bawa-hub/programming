package main

import "fmt"

type Player interface { Play(file string) }

type MP3Player struct{}
func (MP3Player) Play(file string) { fmt.Println("Playing mp3:", file) }

// Adaptee with incompatible interface
type VLC struct{}
func (VLC) PlayVLC(file string) { fmt.Println("Playing vlc:", file) }

// Adapter
type VLCAdapter struct{ vlc VLC }
func (a VLCAdapter) Play(file string) { a.vlc.PlayVLC(file) }

func main() {
	var p Player = MP3Player{}
	p.Play("song.mp3")

	p = VLCAdapter{vlc: VLC{}}
	p.Play("movie.vlc")
}
