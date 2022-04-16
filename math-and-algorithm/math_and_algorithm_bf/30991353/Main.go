package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(N int64, a []int64, b []int64, c []int64) {
	max := 0.0
	for i := int64(0); i < N; i++ {
		for j := i + 1; j < N; j++ {
			if a[i]*b[j] == a[j]*b[i] {
				continue
			}
			x := float64(b[j]*c[i]-b[i]*c[j]) / float64(a[i]*b[j]-a[j]*b[i])
			y := float64(a[i]*c[j]-a[j]*c[i]) / float64(a[i]*b[j]-a[j]*b[i])
			if Satisfy(x, y, a, b, c) {
				max = math.Max(max, x+y)
			}
		}
	}
	fmt.Println(max)
}

func Satisfy(x, y float64, a, b, c []int64) bool {
	for i := range a {
		if float64(a[i])*x+float64(b[i])*y > float64(c[i]) {
			return false
		}
	}
	return true
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
	a := make([]int64, N)
	b := make([]int64, N)
	c := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		b[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		c[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, a, b, c)
}
