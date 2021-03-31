package main

import "fmt"

func main() {

	var p *int

	if p != nil {
		fmt.Println("value of p: ", *p)		// * references a pointer
	} else {
		fmt.Println("p is nil")
	}

	var v int = 42
	p = &v	// & connects a variable and pointer
	if p != nil {
		fmt.Println("value of p: ", *p)
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1:", *pointer1)
	fmt.Println("Value 1:", value1)
}
