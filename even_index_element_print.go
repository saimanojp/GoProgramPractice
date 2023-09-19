package main

import "fmt"

func main() {

	var size int

	fmt.Print("Enter the Even Odd Array Size = ")
	fmt.Scan(&size)

	numarray := make([]int, size)

	fmt.Print("Enter the Even Odd Array Items  = ")
	for i := 0; i < size; i++ {
		fmt.Scan(&numarray[i])
	}

	fmt.Println("The List of Array Items in Even Index Position = ")
	for i := 0; i < len(numarray); i += 2 {
		fmt.Println(numarray[i])
	}
}
