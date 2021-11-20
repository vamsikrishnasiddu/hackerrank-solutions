package main

import "fmt"

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var sum int32
	for i := 0; i < len(ar); i++ {
		sum += ar[i]

	}

	return sum

}

func main() {

	var n int32
	fmt.Scanln(&n)

	var arr []int32
	arr = make([]int32, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	sum := simpleArraySum(arr)

	fmt.Println(sum)

}
