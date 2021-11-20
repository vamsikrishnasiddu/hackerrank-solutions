package main

import "fmt"

func main() {
	var n, m, d int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&d, &m)

	fmt.Println(subArrayDivison(arr, d, m))

}

func subArrayDivison(arr []int, d, m int) int {
	var sum, count int
	for i := 0; i <= len(arr)-m; i++ {
		sum = 0
		for j := i; j <= m+i-1; j++ {
			sum = sum + arr[j]
		}
		fmt.Println("sum:", sum)

		if sum == d {
			count++
		}

	}

	return count
}
