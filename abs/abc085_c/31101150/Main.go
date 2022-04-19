package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, Y int64) {
	for i := int64(0); i <= N; i++ {
		for j := int64(0); j <= N; j++ {
			if k := N - (i + j); k >= 0 && i*10000+j*5000+k*1000 == Y {
				fmt.Printf("%d %d %d\n", i, j, k)
				return
			}
		}
	}

	fmt.Printf("%d %d %d\n", -1, -1, -1)
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
	Y, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, Y)
}
