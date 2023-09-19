package main

import "fmt"

func main() {
	var a, b, temp int
	fmt.Print("Enter the First  = ")
	fmt.Scanln(&a)

	fmt.Print("Enter the Second  = ")
	fmt.Scanln(&b)

	temp = a
	a = b
	b = temp

	fmt.Println("The First Number after  = ", a)
	fmt.Println("The Second Number after = ", b)
}
