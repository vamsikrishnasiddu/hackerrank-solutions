package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	birdsArr := make([]int, n)

	for i := 0; i < len(birdsArr); i++ {
		fmt.Scan(&birdsArr[i])
	}

	migratoryBirds(birdsArr)

}

func maxElementArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func migratoryBirds(arr []int) int {
	//birdIds range from 1 to 5
	var birdTypes = []int{1, 2, 3, 4, 5}
	birdTypeCount := make([]int, 5)
	var count int
	for i := 0; i < len(birdTypes); i++ {
		count = 0
		for j := 0; j < len(arr); j++ {
			if arr[j] == birdTypes[i] {
				count++
			}
		}
		// adding the count to birdType array
		birdTypeCount[i] = count
	}
	// calculating the maximum of the frequency
	max := maxElementArr(birdTypeCount)
	var birdId int
	//comparing the maximum with the freqelements
	for i := 0; i < len(birdTypeCount); i++ {
		// if frequency matches maximum element
		//the particular birdType will be returned.
		if birdTypeCount[i] == max {
			birdId = birdTypes[i]
			break
		}
	}

	return birdId

}
