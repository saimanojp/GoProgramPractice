package main

import "fmt"

func main() {

	var num, i int

	fmt.Print("\nEnter the Last Natural Number = ")
	fmt.Scanln(&num)

	fmt.Println("\nNatural Numbers from ", num, " to 1 are = ")
	for i = num; i >= 1; i = i - 1 {
		fmt.Print(i, "\t")
	}
	fmt.Println()
}
