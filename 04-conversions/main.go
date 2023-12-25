package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating for our app : ")

	input, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// ParseFloat is something similar to that we use in java
	// to remove '\n' or any whitespace we need to use the TrimSpace function
	// from the strings library.

	if err != nil {
		fmt.Println("Something went wrong, ", err)
	} else {
		fmt.Println("The rating is, ", numRating)
	}

	fmt.Println("Rating after adding 1 : ", numRating+1)
}
