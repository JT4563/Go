package main

import "fmt"

type Player interface {
	play() string
}

type vlc struct{}

type kmp struct{}

func (v vlc) play() string {
	return " i am on"
}
func (k kmp) play() string {
	return " i am on on"
}
func main() {
	pl := []Player{vlc{}, kmp{}}
	for _, p := range pl {
		fmt.Println(p.play())
	}
}
