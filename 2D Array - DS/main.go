package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func at(arr [][]int, x, y int) int {
	return arr[y][x]
}

func makeSum(arr [][]int, x, y int) int {
	sum := 0
	sum += at(arr, x, y)
	sum += at(arr, x+1, y)
	sum += at(arr, x+2, y)
	sum += at(arr, x+1, y+1)
	sum += at(arr, x, y+2)
	sum += at(arr, x+1, y+2)
	sum += at(arr, x+2, y+2)
	return sum
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int) int32 {
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -(MaxInt - 1)
	dimension := 6
	max := MinInt

	for y := 0; y < dimension-2; y++ {
		for x := 0; x < dimension-2; x++ {
			sum := makeSum(arr, x, y)
			if sum > max {
				max = sum
			}
		}
	}

	return int32(max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

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
