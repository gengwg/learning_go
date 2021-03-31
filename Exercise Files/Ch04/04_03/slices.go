package main

import (
	"fmt"
	"sort"
)

func main() {

	// without numbers in [] it's slice; with number its array
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// slices are resizable
	colors = append(colors, "Purple")
	fmt.Println(colors)

	// remove 1st item
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// remove last item
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// type, initial size, capcity
	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156
	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))	// output 10!

	sort.Ints(numbers)
	fmt.Println(numbers)
}
