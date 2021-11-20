package main

import "fmt"

func main() {
	var size int

	fmt.Scan(&size)

	A := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&A[i])
	}

	revArr := reverseArray(A)

	for i := 0; i < len(revArr); i++ {
		fmt.Printf("%d ", revArr[i])
	}

}

func reverseArray(A []int) []int {
	temp := make([]int, len(A))

	j := len(temp) - 1

	for i := 0; i < len(A); i++ {
		temp[j] = A[i]
		j--
	}

	return temp
}
