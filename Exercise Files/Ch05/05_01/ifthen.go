package main

import "fmt"

func main() {

	// var x float64 = 42
	var result string

	// initial statement w/in conditional block
	// value local to the conditional block
	// memery released exiting block
	if x:= 42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else { // else must be on same line as closing bracket
		result = "Greater than zero"
	}

	fmt.Println("Result:", result)
}
