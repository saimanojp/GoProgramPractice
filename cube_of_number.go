package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter the number")
	fmt.Scanln(&num)

	cube := num * num * num

	fmt.Print("The cube of ", num, " is ", cube)

}
