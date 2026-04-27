package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 12.5
	
	t := reflect.TypeOf(x)
	
	fmt.Println(t)
}