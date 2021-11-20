package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	grades := make([]int, n)

	for i := 0; i < len(grades); i++ {
		fmt.Scanln(&grades[i])
	}
	newGrades := roundGrades(grades)

	for i := 0; i < len(newGrades); i++ {
		fmt.Println(newGrades[i])
	}

}

func roundGrades(grades []int) []int {
	var rem int
	var nextMultiple int
	for i := 0; i < len(grades); i++ {
		rem = grades[i] % 5
		if rem != 0 && grades[i] >= 38 {
			nextMultiple = grades[i] + (5 - rem)
			if (nextMultiple - grades[i]) < 3 {
				grades[i] = nextMultiple
			}
		}

	}
	return grades

}
