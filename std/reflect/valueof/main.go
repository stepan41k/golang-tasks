package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x float64 = 10.0
	
	v := reflect.ValueOf(x)
	
	fmt.Println(v)
}
