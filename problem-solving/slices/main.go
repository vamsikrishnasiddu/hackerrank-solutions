package main

import "fmt"

func main() {
	var s []int = []int{1, 2, 3, 4, 5}
	var m []int = make([]int, 5)
	m = s
	fmt.Println(s)
	fmt.Println(m)

	s[0] = 5

	fmt.Println(s)
	fmt.Println(m)
}
