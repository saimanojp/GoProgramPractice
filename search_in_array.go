package main

import "fmt"

func main() {
	var sersize, i, search int

	fmt.Print("Enter the Even Array Size = ")
	fmt.Scan(&sersize)

	serArr := make([]int, sersize)

	fmt.Print("Enter the Even Array Items  = ")
	for i = 0; i < sersize; i++ {
		fmt.Scan(&serArr[i])
	}
	fmt.Print("Enter the Array Search Item  = ")
	fmt.Scan(&search)
	flag := 0
	for i = 0; i < sersize; i++ {
		if serArr[i] == search {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Println("We Found the Search Item ", search, " at position = ", i)
	} else {
		fmt.Println("We haven't Found the Search Item ")
	}
}
