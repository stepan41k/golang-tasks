package main

import "fmt"

type MyError struct{}

func (m *MyError) Error() string {
	return "my error"
}

func getError() error {
	var err *MyError
	return err
}

func main() {
	err := getError()
	if err != nil {
		fmt.Println("Error is NOT nil")
	} else {
		fmt.Println("Error IS nil")
	}
}