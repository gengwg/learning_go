package main

import (
	"fmt"
)

type Dog struct {
	Breed string
	Weight int
}

func main() {

	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)	// print details of struct
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight) // access fields of struct


}
