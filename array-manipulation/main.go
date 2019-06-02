package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// We loop over the rows in the query, and then sub-loop over the elements of
// the array than need summation. This approach worksfine , but it will not pass
// (in an acceptable amount of time) the higher tests in when n is very
// large - O(n**2)
func arrayManipulation_too_slow_solution(n int32, queries [][]int32) int64 {
	var maxValue int64
	arr := make([]int32, n)

	for i := 0; i < len(queries); i++ {
		query := queries[i]
		// fmt.Println(query)
		for j := query[0] - 1; j < query[1]; j++ {
			arr[j] += query[2]
			if int64(arr[j]) > maxValue {
				maxValue = int64(arr[j])
			}
		}
		// fmt.Println(arr)
	}
	fmt.Println(maxValue)
	return maxValue
}

// This is better - O(N*M)
func arrayManipulation(n int32, queries [][]int32) int64 {
	var maxValue int64
	var itt int64
	arr := make([]int32, n)

	for _, v := range queries {
		arr[v[0]-1] += v[2]
		if v[1] != int32(len(arr)) {
			arr[v[1]] -= v[2]
		}
	}
	fmt.Println(arr)

	for _, q := range arr {
		itt += int64(q)
		if itt > maxValue {
			maxValue = itt
		}
	}

	fmt.Println(maxValue)
	return maxValue
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

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
