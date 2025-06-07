package main

import "fmt"

type payproces interface {
	payment() string
}

type paytm struct{}

func makepayment(p payproces) {
	fmt.Println(p.payment())
}

func (p paytm) payment() string {
	return "payment is successful"
}
func main() {
	pay := paytm{}
	makepayment(pay)
}
