package main

import "fmt"

func main() {
	var lgsmsize, i, minPosition, maxPosition int

	fmt.Print("Enter the Array Size to find Smallest & Largest = ")
	fmt.Scan(&lgsmsize)

	lgsmArr := make([]int, lgsmsize)

	fmt.Print("Enter the Array Items  = ")
	for i = 0; i < lgsmsize; i++ {
		fmt.Scan(&lgsmArr[i])
	}
	largest := lgsmArr[0]
	smallest := lgsmArr[0]

	for i = 0; i < lgsmsize; i++ {
		if largest < lgsmArr[i] {
			largest = lgsmArr[i]
			maxPosition = i
		}
		if smallest > lgsmArr[i] {
			smallest = lgsmArr[i]
			minPosition = i
		}
	}
	fmt.Println("\nThe Largest Number in this lgsmArr    = ", largest)
	fmt.Println("The Index Position of Largest Number = ", maxPosition)
	fmt.Println("\nThe Smallest Number in this lgsmArr    = ", smallest)
	fmt.Println("The Index Position of Smallest Number = ", minPosition)
}
