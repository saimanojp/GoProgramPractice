package main

import (
	"fmt"
	"math"
)

func main() {
	var pAmount, intrestRate, timePeriod, ciFuture, compoundI float64

	fmt.Print("Enter Total Amount = ")
	fmt.Scanln(&pAmount)

	fmt.Print("Enter Intrest Rate = ")
	fmt.Scanln(&intrestRate)

	fmt.Print("Enter Total No.of Years = ")
	fmt.Scanln(&timePeriod)

	ciFuture = pAmount * (math.Pow((1 + intrestRate/100), timePeriod))
	compoundI = ciFuture - pAmount

	fmt.Println(" The Compound Intrest = ", compoundI)
	fmt.Println(" The Future Compound Intrest = ", ciFuture)

}
