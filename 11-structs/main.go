package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	fmt.Println("Welcome to structs")
	// there is no inheritance in go lang; no classes and objects

	// creating a struct type variable
	user1 := User{"User1", "user1@gmail.com", 18}

	fmt.Printf("User is %v\n", user1)
	fmt.Printf("User details are %+v\n", user1)

	fmt.Printf("The name of the user is %v, the email of the user is %v\n", user1.Name, user1.Email)
}
