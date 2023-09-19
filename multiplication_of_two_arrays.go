package main

import "fmt"

func main() {
	var i int

	var multiArr1 [5]int
	var multiArr2 [5]int
	var MultiplicationArr [5]int

	fmt.Print("Enter the First Array Items for Multiplication = ")
	for i = 0; i < 5; i++ {
		fmt.Scan(&multiArr1[i])
	}

	fmt.Print("Enter the Second Array Items for Multiplication = ")
	for i = 0; i < 5; i++ {
		fmt.Scan(&multiArr2[i])
	}

	fmt.Print("The Multiplication of Two Arrays = ")
	for i = 0; i < len(MultiplicationArr); i++ {
		MultiplicationArr[i] = multiArr1[i] * multiArr2[i]
		fmt.Print(MultiplicationArr[i], "  ")
	}
	fmt.Println()
}
