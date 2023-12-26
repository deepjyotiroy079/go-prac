package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// the syntax
	var fruitList = []string{}
	fmt.Printf("Type of fruit List : %T\n", fruitList)

	// initializing
	var initializedFruitList = []string{"Apple", "Banana", "Grapes"}
	fmt.Println("Initialized Fruit List : ", initializedFruitList)
	fmt.Println("Length of initialized fruit list : ", len(initializedFruitList))

	// appending to the slice
	initializedFruitList = append(initializedFruitList, "Mango", "Guava")
	fmt.Println("Appended initialized Fruit List : ", initializedFruitList)

	// slicing

	initializedFruitList = append(initializedFruitList[:3])
	fmt.Println("Sliced initialized fruit list : ", initializedFruitList)

	// using the make function to allocate memory
	highestScores := make([]int, 4)

	highestScores[0] = 100
	highestScores[1] = 600
	highestScores[2] = 200
	highestScores[3] = 400

	fmt.Println("Highest Scores : ", highestScores)

	// this will throw out of index error
	// highestScores[4] = 500

	// appending to the slice
	highestScores = append(highestScores, 600, 900, 101)
	fmt.Println("Appended Highest Scores : ", highestScores)
	fmt.Println("is Sorted? ", sort.IntsAreSorted(highestScores))

	// sorting the slice
	sort.Ints(highestScores)
	fmt.Println("Sorted highest Scores : ", highestScores)
	fmt.Println("is Sorted? ", sort.IntsAreSorted(highestScores))
}
