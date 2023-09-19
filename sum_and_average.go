package main

import "fmt"

func main() {

	var num, i int

	fmt.Print("\nEnter the Maximum Natural Number = ")
	fmt.Scanln(&num)

	total := 0

	for i = 1; i <= num; i++ {
		total = total + i
	}
	average := total / num

	fmt.Println("The Sum of Natural Numbers from 1 to ", num, " = ", total)
	fmt.Println("The Average of Natural Numbers from 1 to ", num, " = ", average)
}
