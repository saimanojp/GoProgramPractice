package main

import (
	"fmt"
	"math"
)

func main() {

	var pnum, expo float64

	fmt.Print("\nEnter the Number to find the Power = ")
	fmt.Scanln(&pnum)

	fmt.Print("Enter the Exponent Value = ")
	fmt.Scanln(&expo)

	power := math.Pow(pnum, expo)

	fmt.Println(pnum, " Power ", expo, " = ", power)
}
