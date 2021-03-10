package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	// if have a value that declared, you must address it in the code
	stringLength, _ := fmt.Println(str1, str2, str3)

	//if err == nil {
		// Println function will add space between
		fmt.Println("String length:", stringLength)
	//}

	// print verbals
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))

	// print types
	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n",
		str1, str2, str3, aNumber, isTrue)


	// assign printf output to a string
	myString :=	fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T\n",
		str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)
}
