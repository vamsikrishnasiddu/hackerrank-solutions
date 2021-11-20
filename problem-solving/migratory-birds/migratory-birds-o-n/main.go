package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	birdsArr := make([]int, n)

	for i := 0; i < len(birdsArr); i++ {
		fmt.Scan(&birdsArr[i])
	}

	result := migratoryBirds(birdsArr)
	fmt.Println(result)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func migratoryBirds(arr []int) int {
	//birdIds range from 1 to 5
	var counter = make([]int, 6)
	var id, max, result int

	for i := 0; i < len(arr); i++ {
		id = arr[i]
		counter[id]++

		if counter[id] > max {
			max = counter[id]
			result = id
		} else if counter[id] == max {
			result = min(result, id)
		}

	}

	return result

}
