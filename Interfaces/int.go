import "fmt"

type robot struct {
	battery int
}

func (r robot) talk() string {
	if r.battery > 0 {
		return "beep boop"
	}
	return "died omg"
}

func main() {
	r := robot{battery: 5}
	fmt.Println(r.talk())
}