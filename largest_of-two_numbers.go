package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Enter the First Number to find Largest of Two  = ")
	fmt.Scanln(&num1)

	fmt.Print("Enter the Second Number to find Largest of Two = ")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Println("The Largest amoung two  = ", num1)
	} else {
		fmt.Println("The Largest amoung two  = ", num2)
	}
}
