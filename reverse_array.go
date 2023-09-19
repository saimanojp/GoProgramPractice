import "fmt"

func main() {
	var actsize, i, j int

	fmt.Print("Enter the Even Array Size = ")
	fmt.Scan(&actsize)

	actArr := make([]int, actsize)
	revArr := make([]int, actsize)

	fmt.Print("Enter the Even Array Items  = ")
	for i = 0; i < actsize; i++ {
		fmt.Scan(&actArr[i])
	}
	j = 0
	for i = actsize - 1; i >= 0; i-- {
		revArr[j] = actArr[i]
		j++
	}

	fmt.Println("\nThe Result of a Reverse Array = ", revArr)
}