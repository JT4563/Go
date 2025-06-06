package main
import "fmt"

func main() {
    // Our toy box that only holds toys that can talk
    toyBox := []TalkingToy{
        TeddyBear{}, // Allowed because it has Talk()
        Robot{},     // Allowed because it has Talk()
    }
    
    // Make all toys talk!
    for _, toy := range toyBox {
        fmt.Println(toy.Talk())
    }
}