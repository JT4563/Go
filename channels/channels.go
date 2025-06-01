package main

import (
	"fmt"
)

func main() {
	messchan := make(chan string, 1) // Buffered channel with capacity 1 //creating a channel 
	messchan <- "lol"
	tj := <-messchan
	fmt.Println(tj)
	
}
