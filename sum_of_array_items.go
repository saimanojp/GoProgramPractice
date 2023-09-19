package main

import "fmt"

func main() {

	numarray := []int{10, 20, 30, 40, 50}

	arrSum := 0

	for i := 0; i < len(numarray); i++ {
		arrSum = arrSum + numarray[i]
	}
	fmt.Println("The Sum of Array Items = ", arrSum)
}
