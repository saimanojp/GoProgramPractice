package main

import "fmt"

func main() {

	var perfNum, perfsum int
	perfsum = 0

	fmt.Print("Enter the Number to find the Perfect = ")
	fmt.Scanln(&perfNum)

	for i := 1; i < perfNum; i++ {
		if perfNum%i == 0 {
			perfsum = perfsum + i
		}
	}

	if perfNum == perfsum {
		fmt.Println(perfNum, " is a Perfect Number")
	} else {
		fmt.Println(perfNum, " is Not a Perfect Number")
	}
}
