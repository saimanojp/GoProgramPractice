package main

import "fmt"

func main() {

	var strFirstChar string

	strFirstChar = "Welcome Mj"

	fmt.Println(strFirstChar)

	firstChar := strFirstChar[0]

	fmt.Printf("The First Character in this String = %c\n", firstChar)
}
