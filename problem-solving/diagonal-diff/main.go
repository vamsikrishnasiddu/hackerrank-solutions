package main

import "fmt"

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var sum1, sum2, sum int32
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j {
				sum1 = sum1 + arr[i][j]

			}
			if j == (len(arr) - i - 1) {
				sum2 = sum2 + arr[i][j]
			}
		}
	}

	sum = sum1 - sum2
	return abs(sum)

}

func abs(num int32) int32 {
	var scale int32 = -1
	if num < 0 {
		num = scale * num
	}
	return num
}

func main() {
	var n int32
	fmt.Scanln(&n)
	arr := make([][]int32, n)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			arr[i] = make([]int32, n)
		}

	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Scan(&arr[i][j])
		}

	}

	fmt.Println(diagonalDifference(arr))
}
