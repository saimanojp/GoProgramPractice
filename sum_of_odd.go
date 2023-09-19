package main

import "fmt"

func main() {

	var oddnum, Oddtotal int

	fmt.Print("Enter the Number to find Odd Sum = ")
	fmt.Scanln(&oddnum)

	Oddtotal = 0
	fmt.Println("List of Odd Numbers from 1 to ", oddnum, " are = ")
	for x := 1; x <= oddnum; x++ {
		if x%2 != 0 {
			fmt.Print(x, "\t")
			Oddtotal = Oddtotal + x
		}
	}
	fmt.Println("\nSum of Odd Numbers from 1 to ", oddnum, " = ", Oddtotal)
}
