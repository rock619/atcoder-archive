package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	X, Y int64
}

func NewPoint(x, y int64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (p Point) Angle() float64 {
	a := math.Atan2(float64(p.Y), float64(p.X)) * 180 / math.Pi
	if a < 0 {
		return 360.0 + a
	}
	return a
}

func (p Point) Sub(p2 Point) Point {
	return NewPoint(p.X-p2.X, p.Y-p2.Y)
}

func Angle2(a, b float64) float64 {
	abs := math.Abs(a - b)
	if abs > 180 {
		return 360 - abs
	}
	return abs
}

func maxAngle(points []Point, originIndex int) float64 {
	origin := points[originIndex]
	angles := make([]float64, 0, len(points)-1)
	for i, p := range points {
		if i == originIndex {
			continue
		}

		angles = append(angles, p.Sub(origin).Angle())
	}

	sort.Float64s(angles)

	max := 0.0
	for _, a := range angles {
		target := a + 180
		if target >= 360 {
			target -= 360
		}

		upperIndex := sort.SearchFloat64s(angles, target)
		if upperIndex > len(angles)-1 {
			upperIndex -= len(angles)
		}
		lowerIndex := upperIndex - 1
		if lowerIndex < 0 {
			lowerIndex += len(angles)
		}

		max = math.Max(max, Angle2(a, angles[upperIndex]))
		max = math.Max(max, Angle2(a, angles[lowerIndex]))
	}
	return max
}

func solve(N int64, X []int64, Y []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	points := make([]Point, N)
	for i := int64(0); i < N; i++ {
		points[i] = Point{
			X: X[i],
			Y: Y[i],
		}
	}

	max := 0.0
	for i := range points {
		max = math.Max(max, maxAngle(points, i))
	}

	fmt.Fprintln(w, max)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	X := make([]int64, N)
	Y := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		X[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		Y[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, X, Y)
}
