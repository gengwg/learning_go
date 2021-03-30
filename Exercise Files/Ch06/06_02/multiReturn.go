package main

import "fmt"

func main() {
	n1, l1 := FullName("Zaphod", "Beeblebox")
	fmt.Printf("Fll name: %v, number of chars: %v", n1, l1)
}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}
