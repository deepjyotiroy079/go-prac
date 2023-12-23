package main

import "fmt"

// the Capital L in LoginToken (Capital letter) represents public variable and
// this can be accessed from anywhere in the module.

const LoginToken string = "asdfasdfasdfasdfasdf" // public variable

func main() {

	/*
		REFER THE BELOW DOCS FOR DETAILS
		https://go.dev/ref/spec#Types
	*/
	// string
	var username string = "deepjyoti"
	fmt.Println(username)
	fmt.Printf("The type of the variable is: %T \n", username)

	// int
	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("The type of the variable is: %T \n", smallValue)

	// floats
	var userCounts float64 = 1000.234234234234
	fmt.Println(userCounts)
	fmt.Printf("The type of the variable is: %T \n", userCounts)

	// boolean
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of the variable is: %T \n", isLoggedIn)

	// implicit declaration
	var anotherVariable = "deepjyotiroy079.github.io"
	fmt.Println(anotherVariable)
	fmt.Printf("The type of the variable is: %T \n", anotherVariable)

	// no var style
	// we cannot use this style outside the function
	jwtToken := "asfasdfasdfb"
	fmt.Println(jwtToken)
	fmt.Printf("The type of the variable is: %T \n", jwtToken)

	// consts
	// go to the top of the file
	fmt.Println(LoginToken)
	fmt.Printf("The type of the variable is: %T \n", LoginToken)
}
