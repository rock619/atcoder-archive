package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Matrix struct {
	c [3][3]float64
}

func (m Matrix) Multiply(m2 Matrix) Matrix {
	result := Matrix{}
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			for j := 0; j < 3; j++ {
				result.c[i][j] += m.c[i][k] * m2.c[k][j]
			}
		}
	}
	return result
}

func (m Matrix) Power(n int64) Matrix {
	p := m
	q := p
	flag := false
	for i := 0; i < 60; i++ {
		if (n & (1 << i)) != 0 {
			if !flag {
				q = p
				flag = true
			} else {
				q = q.Multiply(p)
			}
		}

		p = p.Multiply(p)
	}

	return q
}

func solve(Q int64, X []float64, Y []float64, Z []float64, T []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := int64(0); i < Q; i++ {
		m := Matrix{c: [3][3]float64{
			{1 - X[i], Y[i], 0},
			{0, 1 - Y[i], Z[i]},
			{X[i], 0, 1 - Z[i]},
		}}
		m = m.Power(T[i])
		fmt.Fprintf(w, "%.15f %.15f %.15f\n", Sum(m.c[0]), Sum(m.c[1]), Sum(m.c[2]))
	}
}

func Sum(ns [3]float64) float64 {
	sum := 0.0
	for _, n := range ns {
		sum += n
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var Q int64
	scanner.Scan()
	Q, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	X := make([]float64, Q)
	Y := make([]float64, Q)
	Z := make([]float64, Q)
	T := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		X[i], _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		Y[i], _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		Z[i], _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		T[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(Q, X, Y, Z, T)
}
