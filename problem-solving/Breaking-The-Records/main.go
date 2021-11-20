package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	//input for the array

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	min, max := minMaxRecords(arr)
	fmt.Printf("min:%d, max:%d", max, min)

}

func minMaxRecords(arr []int) (int, int) {
	var max, min, countMax, countMin int

	max = arr[0]
	min = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			countMin++
		} else if arr[i] > max {
			max = arr[i]
			countMax++
		}
	}

	return countMax, countMin
}
