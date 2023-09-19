package main

import "fmt"

func main() {
	var ngsize, i int

	fmt.Print("Enter the Negative Array Size = ")
	fmt.Scan(&ngsize)

	negArr := make([]int, ngsize)

	fmt.Print("Enter the Negative Array Items  = ")
	for i = 0; i < ngsize; i++ {
		fmt.Scan(&negArr[i])
	}

	fmt.Print("\nThe Negative Numbers in this negArr = ")
	for i = 0; i < ngsize; i++ {
		if negArr[i] < 0 {
			fmt.Print(negArr[i], " ")
		}
	}
	fmt.Println()
}
