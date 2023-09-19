package main

import (
	"fmt"
	"math"
)

func main() {

	var sqrtnum, sqrtOut float64

	fmt.Print("\nEnter the Number to find Square Root = ")
	fmt.Scanln(&sqrtnum)

	sqrtOut = math.Sqrt(sqrtnum)

	fmt.Println("\nThe Square of a Given Number  = ", sqrtOut)
}
