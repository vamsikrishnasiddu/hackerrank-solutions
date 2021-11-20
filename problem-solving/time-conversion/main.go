package main

import (
	"fmt"
	"strings"
)

func main() {
	var hh, mm, ss int
	var time string
	fmt.Scanf("%d:%d:%d%s", &hh, &mm, &ss, &time)

	if strings.Contains(time, "AM") && hh == 12 {
		hh = 0
	}

	if strings.Contains(time, "PM") && hh != 12 {
		hh += 12
	}

	fmt.Printf("%02d:%02d:%02d", hh, mm, ss)
}
