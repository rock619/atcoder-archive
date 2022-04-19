package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A int64, B int64) {
	sum := int64(0)
	for i := int64(1); i <= N; i++ {
		if s := DigitsSum(i); s >= A && s <= B {
			sum += i
		}
	}
	fmt.Println(sum)
}

func DigitsSum(n int64) int64 {
	sum := int64(0)
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	A, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	B, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, A, B)
}
