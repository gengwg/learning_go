package main

import (
	"fmt"
)

func main() {
	// better use slices than array
	var colors [3]string	// must be same type in array
	colors[0] = "Red"		// type already defined
	colors[1] = "Green"
	colors[2] = "Blue"		// type already defined
	fmt.Println(colors)
	fmt.Println(colors[1])


	// assingn values as you define array
	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))

}
