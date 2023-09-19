package main

import "fmt"

func main() {

	var evenNum int

	fmt.Print("Enter the Number to Print Even's = ")
	fmt.Scanln(&evenNum)

	fmt.Println("Even Numbers from 1 to ", evenNum, " are = ")
	for i := 1; i <= evenNum; i++ {
		if i%2 == 0 {
			fmt.Print(i, "\t")
		}
	}
	fmt.Println()
}
