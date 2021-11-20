package main

import "fmt"

func main() {

	var n, k int

	fmt.Scan(&n, &k)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(divSumPairs(arr, k))

}

func divSumPairs(arr []int, k int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if (i < j) && ((arr[i]+arr[j])%k) == 0 {
				count++

			}
		}
	}
	return count
}
