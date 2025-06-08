package main

import (
	"errors"
	"fmt"
)

func read(filename string) error {
	if filename == "" {
		return errors.New("the file is empty ")
	}
	return nil
}
func file() (string, error) {
	er := read("")
	if er != nil {
		return "", fmt.Errorf("the lol lol \\ %w", er)
	}
	return "all good", nil
}

func main() {
	_, abcd := file()
	if abcd != nil {
		fmt.Println("wow", abcd)
	unw := errors.Unwrap(abcd)
	fmt.Println("unwrapped me :", unw)

	} else {
		fmt.Println("mm", abcd)
	}
}
