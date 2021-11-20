package main

import "fmt"

func main() {

	var size int

	fmt.Scan(&size)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	reverseArr(arr)

	for i := 0; i < size; i++ {
		fmt.Printf("%d ", arr[i])
	}

}

func reverseArr(arr []int) {

	first := 0
	last := len(arr) - 1

	for first <= last {
		arr[first], arr[last] = arr[last], arr[first]

		first++
		last--

	}

}
