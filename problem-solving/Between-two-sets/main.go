package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	var arr = make([]int, m)
	var brr = make([]int, n)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(brr); i++ {
		fmt.Scan(&brr[i])
	}

	fmt.Printf("lcm of arr %v is : %d\n", arr, lcmofArr(arr))
	fmt.Printf("gcd of arr %v is : %d\n", brr, gcdOfArr(brr))
	fmt.Println("getTotalCount:", getTotalCount(arr, brr))

}

func getTotalCount(arr, brr []int) int {
	a := lcmofArr(arr)
	b := gcdOfArr(brr)
	i := 1
	count := 0
	temp := a
	for temp <= b {
		temp = a * i
		if b%temp == 0 {
			count++
		}

		i++
	}
	return count

}

func gcd(a, b int) int {
	for b%a != 0 {
		temp := a
		a = b % a
		b = temp
	}

	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmofArr(arr []int) int {
	x := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		x = (x * arr[i+1]) / gcd(x, arr[i+1])

	}

	return x
}

func gcdOfArr(arr []int) int {
	x := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		x = gcd(x, arr[i+1])
	}

	return x
}
