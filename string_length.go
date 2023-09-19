package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Any String to find Length = ")
	str, _ := reader.ReadString('\n')
	length := 0
	for _, l := range str {
		fmt.Printf("%c  ", l)
		length++
	}
	fmt.Println("\nThe Length of a Given String = ", length-1)
}
