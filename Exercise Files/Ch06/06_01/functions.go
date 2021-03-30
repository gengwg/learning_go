package main

import "fmt"

func main() {
	doSomething()

	sum := addValues(23, 54)
	fmt.Println("sum: ", sum)

	sum = addAllValues(12, 54, 79, 28)
	fmt.Println("New sum: ", sum)

}

func doSomething() {
	fmt.Println("Doing something!")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

// go has no function overloading; you can't have the same functions with same name in the same package
// lower initial means private func; not accessible to other applications
func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
