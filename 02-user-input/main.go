package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "User"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a ratting for our course: ")

	// ok comma syntex || ok error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("The type of the rating is, %T\n", input)

}
