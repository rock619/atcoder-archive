package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, X int64) {
	count := 0
	for a := int64(1); a <= N; a++ {
		for b := a + 1; b <= N; b++ {
			if c := X - (a + b); c > b && c <= N {
				count++
			}
		}
	}
	fmt.Println(count)
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
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, X)
}
