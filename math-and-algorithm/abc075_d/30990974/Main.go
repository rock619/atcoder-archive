package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	X, Y int64
}

func solve(N int64, K int64, x []int64, y []int64) {
	points := make([]Point, N)
	for i := range points {
		points[i] = Point{X: x[i], Y: y[i]}
	}

	min := int64(1 << 62)
	for _, left := range points {
		for _, right := range points {
			for _, bottom := range points {
				for _, top := range points {
					if ContainedPointsCount(points, left, right, bottom, top) < K {
						continue
					}

					if a := Area(left, right, bottom, top); a < min {
						min = a
					}
				}
			}
		}
	}

	fmt.Println(min)
}

func ContainedPointsCount(points []Point, left, right, bottom, top Point) int64 {
	result := int64(0)
	for _, p := range points {
		if p.X >= left.X && p.X <= right.X && p.Y >= bottom.Y && p.Y <= top.Y {
			result++
		}
	}
	return result
}

func Area(left, right, bottom, top Point) int64 {
	return (right.X - left.X) * (top.Y - bottom.Y)
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
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	x := make([]int64, N)
	y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		x[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		y[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, K, x, y)
}
