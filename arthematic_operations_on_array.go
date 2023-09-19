package main

import "fmt"

func main() {
	var size, i int

	fmt.Print("Enter the Array Size = ")
	fmt.Scan(&size)

	garr1 := make([]int, size)
	garr2 := make([]int, size)

	fmt.Print("Enter the First Array Items  = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&garr1[i])
	}

	fmt.Print("Enter the Second Array Items = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&garr2[i])
	}

	fmt.Println("Add\tSub\tMul\tDiv\tMod")
	for i = 0; i < size; i++ {
		fmt.Print("\n", garr1[i]+garr2[i], "\t")
		fmt.Print(garr1[i]-garr2[i], "\t")
		fmt.Print(garr1[i]*garr2[i], "\t")
		fmt.Print(garr1[i]/garr2[i], "\t")
		fmt.Print(garr1[i]%garr2[i], "\t")
	}
	fmt.Println()
}
