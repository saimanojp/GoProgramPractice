package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Any String to find ASCII Values = ")
	strData, _ := reader.ReadString('\n')

	for i := 0; i < len(strData); i++ {
		fmt.Printf("The ASCII Value of %c = %d\n", strData[i], strData[i])
	}
}
