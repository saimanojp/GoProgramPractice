package main

import (
	"fmt"
)

func main() {

	byteArray := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	var strToConvert string

	strToConvert = string(byteArray)

	fmt.Println(strToConvert)
}
