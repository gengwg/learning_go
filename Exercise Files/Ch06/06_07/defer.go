package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("Close the file!")
	fmt.Println("Open the file!")

	// stack; lifo
	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")

	myFunc()

	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	fmt.Println("Undeferred Statement")

	x := 1000
	// evaluated and saved at the moment of deferring
	defer fmt.Println("Value of x:", x)
	x++
	fmt.Println("Value of x after incrementing: ", x)


}

// defer only works within a function
func myFunc() {

	defer fmt.Println("Deferred in the function")

	fmt.Println("Not Deferred in the function")
}
