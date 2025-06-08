package main

import (
	"errors"
	"fmt"
)

func check(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negative number")
	}
	return "positive number found", nil
}

func main() {
	msg, err := check(-1)
	if err != nil {
		fmt.Println("no",  err)
	} else {
		fmt.Println("yes", msg)
	}
}
