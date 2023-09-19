package main

import "fmt"

func main() {
	var size, i int

	fmt.Print("Enter the Array Size = ")
	fmt.Scan(&size)

	addArr1 := make([]int, size)
	addArr2 := make([]int, size)
	additionArr := make([]int, size)

	fmt.Print("Enter the First Array Items to Perform Addition = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&addArr1[i])
	}

	fmt.Print("Enter the Second Array Items to Perform Addition = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&addArr2[i])
	}

	fmt.Print("The Sum or Addition of Two Arrays = ")
	for i = 0; i < size; i++ {
		additionArr[i] = addArr1[i] + addArr2[i]
		fmt.Print(additionArr[i], "  ")
	}
	fmt.Println()
}
