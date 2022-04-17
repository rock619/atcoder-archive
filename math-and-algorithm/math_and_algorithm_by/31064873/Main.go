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

func solve(a int64, b int64, c int64) {
	if c == 1 {
		fmt.Println(NO)
		return
	}

	v := int64(1)
	for i := int64(1); i <= b; i++ {
		if a/c < v {
			fmt.Println(YES)
			return
		}
		v *= c
	}

	fmt.Println(NO)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var a int64
	scanner.Scan()
	a, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var b int64
	scanner.Scan()
	b, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var c int64
	scanner.Scan()
	c, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(a, b, c)
}
