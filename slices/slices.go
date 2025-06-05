package main

import "fmt"

func main() {
	t := []byte("Tanmayu")
	fmt.Println(len(t), t) // the t values will be the ascii values of the tanmayu or the string
	
	fmt.Println(t[:3])  
	fmt.Println(t[3:])
}
