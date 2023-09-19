package main

import (
	"fmt"
)

func main() {
	var strToConvert string

	strToConvert = "Hello World"

	byteString := []byte(strToConvert)
	fmt.Println(byteString)
}
