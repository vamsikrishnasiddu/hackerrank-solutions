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
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func removeDuplicates(ar []int32) []int32 {
	keys := make(map[int32]bool)
	list := []int32{}

	for _, entry := range ar {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}

	}
	return list
}

func sockMerchant(n int32, ar []int32) int32 {
	counter := []int32{}
	list := removeDuplicates(ar)
	var count int32

	for i := 0; i < len(list); i++ {
		count = 0
		for j := 0; j < len(ar); j++ {
			if list[i] == ar[j] {
				count++
			}
		}
		counter = append(counter, count)
	}
	var sum int32
	for i := 0; i < len(counter); i++ {
		sum = sum + (counter[i] / 2)
	}
	return sum

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

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
