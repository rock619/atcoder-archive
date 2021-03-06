package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	YES = "Yes"
	NO  = "No"
)

func solve(N int64, S string) {
	count := 0
	for _, c := range S {
		if c == '(' {
			count++
		} else {
			count--
		}
		if count < 0 {
			fmt.Println(NO)
			return
		}
	}
	fmt.Println(YES)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var S string
	scanner.Scan()
	S = scanner.Text()
	solve(N, S)
}
