package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	sum := solveMeFirst(a, b)

	fmt.Println(sum)
}

func solveMeFirst(a int, b int) int {
	return a + b
}
