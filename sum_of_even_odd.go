package main

import "fmt"

func main() {

	var eonum, eventotal, oddtotal int

	fmt.Print("Enter the Number to find Even and Odd Sum = ")
	fmt.Scanln(&eonum)

	eventotal = 0
	oddtotal = 0

	for x := 1; x <= eonum; x++ {
		if x%2 == 0 {
			eventotal = eventotal + x
		} else {
			oddtotal = oddtotal + x
		}
	}
	fmt.Println("\nSum of Even Numbers from 1 to ", eonum, " = ", eventotal)
	fmt.Println("\nSum of Odd Numbers from 1 to ", eonum, "  = ", oddtotal)
}
