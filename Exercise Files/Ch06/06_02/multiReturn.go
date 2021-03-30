package main

import "fmt"

func main() {

	n1, l1 := FullName("Zaphod", "Beeblebox")
	fmt.Printf("Fll name: %v, number of chars: %v\n", n1, l1)

	n2, l2 := FullNameNakedReturn("Arther", "Glurberg")
	fmt.Printf("Fll name: %v, number of chars: %v\n", n2, l2)

}

func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

// same effect as above; define return value in function signature
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
