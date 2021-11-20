package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func checkLeapYear(year int32) bool {
	var checkYear bool
	if year <= 1917 {
		if year%4 == 0 {
			checkYear = true
		}

	} else if year >= 1919 {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			checkYear = true
		}

	}

	return checkYear

}

func dayOfProgrammer(year int32) string {
	monthsArr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	check := checkLeapYear(year)
	fmt.Println(check)
	if check {
		monthsArr[1]++
	}

	if year == 1918 {
		monthsArr[1] = 15
	}

	var sum, month, day int

	for i := 0; i < len(monthsArr); i++ {
		sum = sum + monthsArr[i]

		if sum+monthsArr[i] > 256 {
			month = i + 2
			day = 256 - sum
			break
		}
	}

	return fmt.Sprintf("%02d.%02d.%d", day, month, year)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
