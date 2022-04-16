package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(A int64, B int64) {
	max := int64(1)
	for i := int64(1); i <= B; i++ {
		if TwoFactorsIn(A, B, i) {
			max = i
		}
	}

	fmt.Println(max)
}

func TwoFactorsIn(lower, upper, i int64) bool {
	return upper/i-(lower+i-1)/i >= 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var A int64
	scanner.Scan()
	A, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B)
}
