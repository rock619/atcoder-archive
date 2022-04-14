package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1e9

type Matrix struct {
	c [2][2]int64
}

func (m Matrix) Multiply(m2 Matrix) Matrix {
	result := Matrix{}
	for i := 0; i < 2; i++ {
		for k := 0; k < 2; k++ {
			for j := 0; j < 2; j++ {
				result.c[i][j] += m.c[i][k] * m2.c[k][j]
				result.c[i][j] %= mod
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

func solve(N int64) {
	m := Matrix{
		c: [2][2]int64{
			{1, 1},
			{1, 0},
		},
	}

	result := m.Power(N - 1)
	fmt.Println((result.c[1][0] + result.c[1][1]) % mod)
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
	solve(N)
}
