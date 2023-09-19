package main

import "fmt"

func main() {
	var arrTofindAvg [5]int

	fmt.Print("Enter the Array Items =  ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&arrTofindAvg[i])
	}

	fmt.Println(arrTofindAvg)

	arraySum := 0

	for _, arr := range arrTofindAvg {
		arraySum += arr
	}

	arrayAvg := arraySum / 5
	fmt.Println("The Average of Array Items = ", arrayAvg)
	fmt.Println("The sum of Array Items     = ", arraySum)
}
