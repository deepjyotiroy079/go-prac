package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")

	countries := make(map[string]string)

	countries["IND"] = "India"
	countries["ENG"] = "England"
	countries["AUS"] = "Australia"
	countries["SAF"] = "South Africa"

	fmt.Println("List of all the countries are : ", countries)

	fmt.Println("IND is for : ", countries["IND"])

	// DELETING BY KEY
	delete(countries, "ENG")

	fmt.Println("List of all the countries are : ", countries)

	// LOOPING THROUGH THE MAP
	for key, value := range countries {
		fmt.Printf("The Key is %v, and the value is %v\n", key, value)
	}

}
