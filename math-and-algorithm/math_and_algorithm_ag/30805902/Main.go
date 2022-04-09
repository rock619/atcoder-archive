package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(x []int64, y []int64, r []int64) {
	dist := math.Sqrt(float64((x[0]-x[1])*(x[0]-x[1]) + (y[0]-y[1])*(y[0]-y[1])))

	if r[0] < r[1] {
		r[0], r[1] = r[1], r[0]
	}

	fmt.Println(pattern(float64(r[0]), float64(r[1]), dist))
}

func pattern(r1, r2, dist float64) int {
	switch {
	case r1 > r2+dist:
		return 1
	case r1 == r2+dist:
		return 2
	case dist == r1+r2:
		return 4
	case dist > r1+r2:
		return 5
	default:
		return 3
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	x := make([]int64, 2)
	y := make([]int64, 2)
	r := make([]int64, 2)
	for i := int64(0); i < 2; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		r[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(x, y, r)
}
