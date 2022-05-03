package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func radian(e, t int64) float64 {
	return math.Pi * (2*float64(e)/float64(t) - 0.5)
}

func solve(T int64, L int64, X int64, Y int64, Q int64, E []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	for _, e := range E {
		rad := radian(e, T)
		y := -float64(L) * math.Cos(rad) / 2.0
		yDiff := math.Abs(y - float64(Y))
		dist := float64(X*X) + yDiff*yDiff
		z := float64(L) * (math.Sin(rad) + float64(1)) / 2.0
		degree := math.Atan(math.Sqrt(z*z/dist)) / math.Pi * 180

		fmt.Fprintf(w, "%.15f\n", degree)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	T, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	L, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	X, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	Y, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	E := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		E[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(T, L, X, Y, Q, E)
}
