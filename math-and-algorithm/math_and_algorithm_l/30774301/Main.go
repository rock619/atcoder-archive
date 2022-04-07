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

func solve(N int64) {
	if isPrime(N) {
		fmt.Println(YES)
	} else {
		fmt.Println(NO)
	}
}

func isPrime(N int64) bool {
	for i := int64(2); i*i <= N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
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
	solve(N)
}
