package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func minimumSwaps(arr []int) int {
	var swaps int

	// Create an map of pairs where key
	// is array element and value
	// is position of first element
	pairs := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		pairs[arr[i]] = i
	}
	fmt.Println(pairs)

	// Sort the keys in the map to
	// get right position of every element as second
	// element of pair.
	var keys []int
	for k := range pairs {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	// To keep track of visited elements. Initialize
	// all elements as false
	visitedElements := make(map[int]bool)
	for _, v := range keys {
		visitedElements[v] = false
	}
	fmt.Println(visitedElements)

	for i := 0; i < len(arr); i++ {
		if visitedElements[i] || pairs[keys[i]] == i {
			continue
		}

		cycleSize := 0
		j := i
		for !visitedElements[j] {
			visitedElements[j] = true

			// move to next node
			j = pairs[keys[j]]
			cycleSize++
		}

		// Update answer by adding current cycle.
		if cycleSize > 0 {
			swaps += (cycleSize - 1)
		}
	}

	return swaps
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
