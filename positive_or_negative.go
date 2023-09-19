package main

import "fmt"

func main() {

	var num int

	fmt.Print("Please enter any Integer = ")
	fmt.Scanf("%d", &num)

	if num >= 0 {
		fmt.Println("Positive Integer")
	} else {
		fmt.Println("Negative Integer")
	}
}
