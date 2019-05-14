package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int) {
	bribe := 0
	chaotic := false
	n := len(q)
	for i := 0; i < n; i++ {
		if q[i]-(i+1) > 2 {
			chaotic = true
			break
		}
		max := q[i] - 1 - 1
		if max <= 0 {
			max = 0
		}
		for j := max; j < i; j++ {
			if q[j] > q[i] {
				bribe++
			}
		}
	}
	if chaotic {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(bribe)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
