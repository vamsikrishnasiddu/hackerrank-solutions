package main

import "fmt"

func minMaxSum(arr []int) {
	var count, sum int
	var brr []int
	//logic to calc sums by skipping one elements at a time
	for i := 0; i < len(arr); i++ {
		sum = 0
		for j := 0; j < len(arr); j++ {
			if j == count {
				continue
			}
			sum = sum + arr[j]

		}
		brr = append(brr, sum)
		count++
	}

	fmt.Println(brr)
	max := brr[0]
	min := brr[0]
	//comparing for the min and max elements
	for i := 1; i < len(brr); i++ {

		if brr[i] > max {
			max = brr[i]
		}

		if brr[i] < min {
			min = brr[i]
		}

	}

	fmt.Println(min, max)
}

func main() {
	var arr = make([]int, 5)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	minMaxSum(arr)
}
