package main

import (
	"fmt"
)

func main() {
	notes := [8]int{500, 100, 50, 20, 10, 5, 2, 1}
	var amount int
	fmt.Print(" Enter the total Amount of Cash ")
	fmt.Scanln(&amount)

	temp := amount

	for i := 0; i < 8; i++ {
		fmt.Println(notes[i], "notes = ", temp/notes[i])
		temp = temp % notes[i]
	}
}
