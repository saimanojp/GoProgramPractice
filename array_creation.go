package main

import "fmt"

func main() {
	var array1 = [3]int{1, 2, 3}
	array2 := [3]int{2, 3, 4}

	fmt.Println(array1)
	fmt.Println(array2)

	// Arrays with indefiente lengths

	var array3 = [...]int{2, 3, 5}
	array4 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(array3)
	fmt.Println(array4)

	// Initialzing selectively

	array5 := [...]int{1: 10, 3: 2}

	fmt.Println(array5)

}
