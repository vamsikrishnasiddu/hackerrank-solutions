package main

import "fmt"

func compareTriplets(a []int32, b []int32) {
	countA := 0
	countB := 0

	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			countA++
		} else if a[i] < b[i] {
			countB++
		}
	}
	fmt.Println(countA, countB)

}

func main() {
	alice := make([]int32, 3)
	bob := make([]int32, 3)

	fmt.Println("Provide input for alice challenges")

	for i := 0; i < 3; i++ {
		fmt.Scan(&alice[i])
	}

	fmt.Println("Provide input for bobs challenges")

	for i := 0; i < 3; i++ {
		fmt.Scan(&bob[i])
	}

	compareTriplets(alice, bob)
}
