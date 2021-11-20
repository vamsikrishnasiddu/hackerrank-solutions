package main

import "fmt"

func main() {
	var a, b, m, n, s, t int32

	fmt.Scan(&s)
	fmt.Scan(&t)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&m)
	fmt.Scan(&n)
	apples := make([]int32, m)
	oranges := make([]int32, n)

	for i := 0; i < len(apples); i++ {
		fmt.Scan(&apples[i])
	}

	for i := 0; i < len(oranges); i++ {
		fmt.Scan(&oranges[i])
	}

	count1, count2 := countApplesAndOranges(s, t, a, b, apples, oranges)
	fmt.Println(count1)
	fmt.Println(count2)

}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (int32, int32) {

	var countApples, countOranges int32

	for i := 0; i < len(apples); i++ {
		if a+apples[i] >= s && a+apples[i] <= t {
			countApples++
		}
	}

	for i := 0; i < len(oranges); i++ {
		if b+oranges[i] >= s && b+oranges[i] <= t {
			countOranges++
		}
	}

	return countApples, countOranges

}
