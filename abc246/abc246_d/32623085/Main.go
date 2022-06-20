package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	min := 1 << 62
	for i := 0; i <= 1_000_000; i++ {
		for left, right := -1, 1_000_000; right-left > 1; {
			center := (left + right) / 2
			if x := F(i, center); x < N {
				left = center
			} else {
				UpdateMin(&min, x)
				right = center
			}
		}
	}

	fmt.Fprintln(w, min)
}

func F(a, b int) int {
	return a*a*a + a*a*b + a*b*b + b*b*b
}

func UpdateMin(min *int, v int) {
	if v < *min {
		*min = v
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	solve(N)
}
