package main

import (
	"fmt"
	"math"
)

func main() {

	var arr [6][6]int32
	var max int32 = math.MinInt32

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	for l := 0; l < 4; l++ {
		var sum int32
		for k := 0; k < 4; k++ {
			sum = 0
			for i := l; i < 3+l; i++ {
				for j := k; j < k+3; j++ {
					sum = sum + arr[i][j]
				}

			}
			sum = sum - (arr[l+1][k] + arr[l+1][k+2])

			temp := sum

			if temp > max {
				max = temp
			}

		}
	}

	fmt.Println(max)

}
