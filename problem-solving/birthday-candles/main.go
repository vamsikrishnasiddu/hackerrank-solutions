package main

import "fmt"

func birthdayCandles(arr []int) {
	max := arr[0]
	count := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if max == arr[i] {
			count++
		}
	}

	fmt.Println(count)
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	birthdayCandles(arr)
}
