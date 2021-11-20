package main

import "fmt"

func main() {

	var x1, v1, x2, v2 int

	fmt.Scan(&x1, &v1, &x2, &v2)

	output := checkLineNumberJumps(x1, v1, x2, v2)
	fmt.Println(output)

}

func checkLineNumberJumps(x1, v1, x2, v2 int) string {

	if (x2 >= x1 && v1 > v2) || (x2 <= x1 && v1 < v2) {
		if (x2-x1)%(v1-v2) == 0 {
			return "YES"
		}
		return "NO"
	}
	return "NO"

}
