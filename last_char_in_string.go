package main

import "fmt"

func main() {

	var strLastChar string

	strLastChar = "Welcome Mj"

	fmt.Println(strLastChar)

	LastChar := strLastChar[len(strLastChar)-1]

	fmt.Printf("The Last Character in this String = %c\n", LastChar)
}
