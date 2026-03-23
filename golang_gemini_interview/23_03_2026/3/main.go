package main

import (
	"fmt"
	"unsafe"
)

type A struct {
	a bool
	b int64
	c bool
}

type B struct {
	a bool
	c bool
	b int64
}

func main() {
	fmt.Println(unsafe.Sizeof(A{}))
	fmt.Println(unsafe.Sizeof(B{}))
}