package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time world")

	currentTime := time.Now()
	fmt.Println("Current Time is : ", currentTime)

	// formating this 2023-12-25 14:06:49.974627705 +0530 IST m=+0.000281637
	formatedTime := currentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("Formated Time : ", formatedTime)

	// creating a time
	createdDate := time.Date(2023, time.February, 21, 23, 0, 0, 0, time.UTC)
	fmt.Println("New Created Time : ", createdDate.Format("01-02-2006 15:04:05 Monday"))
}
