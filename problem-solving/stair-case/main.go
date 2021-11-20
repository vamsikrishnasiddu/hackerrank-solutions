package main

import "fmt"

func stairCase(n int) {
	//main logic

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < n-i-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {

	var n int
	fmt.Scan(&n)

	stairCase(n)

}
