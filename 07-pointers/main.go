package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	myNumber := 100

	// creating a pointer without initializing
	// var ptr *int -> int based pointer

	var ptr = &myNumber

	fmt.Println("Address to myNumber variable : ", ptr)
	fmt.Println("Actual value present at that memory location : ", *ptr)

	// making changes to myNumber variable by making changes to the pointer
	*ptr = *ptr + 2
	// Enters the memory and make changes at that memory location

	fmt.Println("Value for the dereferenced pointer : ", *ptr)
	fmt.Println("New value for myNumber : ", myNumber)
}
