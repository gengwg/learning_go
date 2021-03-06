package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// more readable using range
	for i := range colors {
		fmt.Println(colors[i])
	}

	// for loop simulating while loop
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
		if sum > 200 {
			goto endofprogram
		}
		if sum > 500 {
			break // break jumps out current block
			// continue goes back to beginning and reevaluate
		}
	}

	// go to label
	endofprogram : fmt.Println("end of program")

}
