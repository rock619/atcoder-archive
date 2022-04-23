package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(L int64, R int64) {
	for d := R - L; d > 0; d-- {
		for x := L; x+d <= R; x++ {
			y := x + d
			if GCD(x, y) == 1 {
				fmt.Println(d)
				return
			}
		}
	}
}

func GCD(a, b int64) int64 {
	for a > 0 && b > 0 {
		if a < b {
			b = b % a
		} else {
			a = a % b
		}
	}

	if a > 0 {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	L, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	scanner.Scan()
	R, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	solve(L, R)
}
