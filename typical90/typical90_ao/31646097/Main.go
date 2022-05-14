package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	X, Y int64
}

func (p Point) Sub(p2 Point) Point {
	return Point{
		X: p.X - p2.X,
		Y: p.Y - p2.Y,
	}
}

func Cross(a, b Point) int64 {
	return a.X*b.Y - a.Y*b.X
}

func Reverse(s []Point) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func GCD(ints ...int64) int64 {
	result := int64(0)
	for _, i := range ints {
		result = gcd(result, i)
	}
	return result
}

func gcd(a, b int64) int64 {
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

func solve(N int64, X []int64, Y []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	g := make([]Point, N)
	for i := int64(0); i < N; i++ {
		g[i] = Point{X: X[i], Y: Y[i]}
	}
	sort.Slice(g, func(i, j int) bool {
		return g[i].X == g[j].X && g[i].Y < g[j].Y || g[i].X < g[j].X
	})

	g1, g2 := []Point{g[0], g[1]}, []Point{g[0], g[1]}
	for i := int64(2); i < N; i++ {
		for len(g1) >= 2 {
			if Cross(g1[len(g1)-1].Sub(g1[len(g1)-2]), g[i].Sub(g1[len(g1)-1])) > 0 {
				break
			}
			g1 = g1[:len(g1)-1]
		}
		g1 = append(g1, g[i])

		for len(g2) >= 2 {
			if Cross(g2[len(g2)-1].Sub(g2[len(g2)-2]), g[i].Sub(g2[len(g2)-1])) < 0 {
				break
			}
			g2 = g2[:len(g2)-1]
		}
		g2 = append(g2, g[i])
	}

	convexHull := make([]Point, 0, N)
	convexHull = append(convexHull, g1...)
	Reverse(g2)
	convexHull = append(convexHull, g2[1:len(g2)-1]...)

	area := int64(0)
	edgePoint := int64(len(convexHull))
	for i := range convexHull {
		a, b := convexHull[i], convexHull[(i+1)%len(convexHull)]
		area += (b.X - a.X) * (b.Y + a.Y)
		vx, vy := Abs(b.X-a.X), Abs(b.Y-a.Y)
		edgePoint += GCD(vx, vy) - 1
	}

	fmt.Fprintln(w, (Abs(area)+edgePoint+2)/2-N)
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
