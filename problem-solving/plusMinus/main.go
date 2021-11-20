package main

import "fmt"

func plusMinus(arr []int32) {

	//we need three vars for counting
	var countPos, countNeg, countZero int32

	var countArr []int32
	// calculating count
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			countPos++
		} else if arr[i] < 0 {
			countNeg++
		} else {
			countZero++
		}
	}
	// adding the counts to count Array.
	countArr = append(countArr, countPos, countNeg, countZero)

	fmt.Println(countArr)
	//printing values
	for i := 0; i < 3; i++ {
		fmt.Printf("%.6f\n", float64(countArr[i])/float64(len(arr)))

	}

}

func main() {
	var n int

	fmt.Scan(&n)

	arr := make([]int32, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)

	plusMinus(arr)
}
