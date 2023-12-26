package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var fruitList [3]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	// fruitList[0] = "Apple"

	fmt.Println("Fruit List elements are : ", fruitList)
	fmt.Println("Length of the fruit list : ", len(fruitList))

	var vegList = [3]string{"Potato", "Onion", "Carrots"}

	fmt.Println("Veg List elements are : ", vegList)
	fmt.Println("Length of the veg list : ", len(vegList))
}
