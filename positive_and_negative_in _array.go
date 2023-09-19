package main

import "fmt"

func main() {
	var size, i int

	fmt.Print("Enter the Positive Negative Array Size = ")
	fmt.Scan(&size)

	posNegArr := make([]int, size)

	fmt.Print("Enter the Positive Negative Array Items  = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&posNegArr[i])
	}

	positiveCount := 0
	negativeCount := 0

	for _, pn := range posNegArr {
		if pn >= 0 {
			positiveCount++
		} else {
			negativeCount++
		}
	}
	fmt.Println("The Total Number of Positive Numbers = ", positiveCount)
	fmt.Println("The Total Number of Negative Numbers = ", negativeCount)
}
