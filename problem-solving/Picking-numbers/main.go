package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func pickingNumbers(arr []int32) int32 {
	// Write your code here
	var longSlice []int32 = []int32{}
	var max int32
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println(arr)

	list := removeDuplicates(arr)
	fmt.Println(list)

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(arr); j++ {
			if !(arr[j] < list[i]) {
				if math.Abs(float64(list[i])-float64(arr[j])) <= 1 {
					longSlice = append(longSlice, arr[j])
				}
			}
		}
		fmt.Println(longSlice)

		if int32(len(longSlice)) > max {
			max = int32(len(longSlice))
		}
		longSlice = []int32{}
	}

	return max

}

func removeDuplicates(arr []int32) []int32 {
	keys := make(map[int32]bool)
	var list []int32 = []int32{}

	for _, elem := range arr {

		if _, value := keys[elem]; !value {
			keys[elem] = true
			list = append(list, elem)
		}
	}
	return list
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

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

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
