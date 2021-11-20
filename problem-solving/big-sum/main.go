package main

import (
	"fmt"
)

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
	var sum int64
	for i := 0; i < len(ar); i++ {
		sum = sum + ar[i]
	}

	return sum

}

func main() {

	var n int64

	fmt.Scanln(&n)

	ar := make([]int64, n)

	for i := 0; i < len(ar); i++ {
		fmt.Scan(&ar[i])
	}

	fmt.Println(aVeryBigSum(ar))

}
