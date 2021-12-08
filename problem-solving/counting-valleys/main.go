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
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	stepsArr := strings.Split(path, "")
	arr := make([]int32, len(stepsArr))

	for i := 0; i < len(stepsArr); i++ {
		if stepsArr[i] == "U" {
			arr[i] = 1
		} else if stepsArr[i] == "D" {
			arr[i] = -1
		}
	}

	var sum, count int32
	// for i := 0; i < len(arr); i++ {

	// 	sum = sum + arr[i]

	// 	if sum == 0 {
	// 		if stepsArr[i] == "U" {
	// 			count++
	// 		}
	// 	}
	// }

	for i := 0; i < len(arr); i++ {
		if sum == 0 {
			if stepsArr[i] == "D" {
				count++
			}
		}
		sum = sum + arr[i]
	}
	return count

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	steps := int32(stepsTemp)

	path := readLine(reader)

	result := countingValleys(steps, path)

	fmt.Fprintf(writer, "%d\n", result)

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
