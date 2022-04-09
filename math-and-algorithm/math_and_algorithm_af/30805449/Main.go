package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64, x []int64, y []int64) {
	min := int64(math.MaxInt64)
	for i := int64(0); i < N; i++ {
		for j := int64(i + 1); j < N; j++ {
			ds := (x[i]-x[j])*(x[i]-x[j]) + (y[i]-y[j])*(y[i]-y[j])
			if ds < min {
				min = ds
			}
		}
	}

	fmt.Println(math.Sqrt(float64(min)))
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
	x := make([]int64, N)
	y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, x, y)
}
