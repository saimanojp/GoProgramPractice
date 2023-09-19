package main

import "fmt"

func main() {

	var digiNum, digiSum, digiReminder int

	fmt.Print("Enter the Number to find the Sum of Digits = ")
	fmt.Scanln(&digiNum)

	for digiSum = 0; digiNum > 0; digiNum = digiNum / 10 {
		digiReminder = digiNum % 10
		digiSum = digiSum + digiReminder
	}
	fmt.Println("The Sum of Digits in this Number = ", digiSum)
}
