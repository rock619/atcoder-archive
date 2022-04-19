package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(A int64, B int64, C int64, X int64) {
	count := 0
	for i := int64(0); i <= A; i++ {
		for j := int64(0); j <= B; j++ {
			for k := int64(0); k <= C; k++ {
				if i*500+100*j+50*k == X {
					count++
				}
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
	scanner.Scan()
	A, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	B, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	C, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	X, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	solve(A, B, C, X)
}
