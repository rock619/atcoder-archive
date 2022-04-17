package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) {
	fmt.Println(N/10000 + N%10000/5000 + N%5000/1000)
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
