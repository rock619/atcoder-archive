package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N, K int, A, B []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	maxA := 5000
	grid := make([][]int, maxA+1)
	for i := range grid {
		grid[i] = make([]int, maxA+1)
	}

	for i := range A {
		a, b := A[i], B[i]
		grid[a][b]++
	}

	for i := range grid {
		for j := 1; j < len(grid); j++ {
			grid[i][j] += grid[i][j-1]
		}
	}

	for i := range grid {
		for j := 1; j < len(grid); j++ {
			grid[j][i] += grid[j-1][i]
		}
	}

	max := 0
	for i := K + 1; i < len(grid); i++ {
		for j := K + 1; j < len(grid); j++ {
			UpdateMax(&max, grid[i][j]-grid[i-K-1][j]-grid[i][j-K-1]+grid[i-K-1][j-K-1])
		}
	}
	fmt.Fprintln(w, max)
}

func UpdateMax(max *int, v int) {
	if v > *max {
		*max = v
	}
}

func main() {
	s := NewScanner()
	N := s.Int()
	K := s.Int()
	A, B := s.Ints2(N)

	solve(N, K, A, B)
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 4096), 1048576)
	scanner.Split(bufio.ScanWords)
	return &Scanner{
		Scanner: scanner,
	}
}

func (s *Scanner) Int() int {
	s.Scan()
	v, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) Ints(size int) []int {
	v := make([]int, size)
	for i := 0; i < size; i++ {
		v[i] = s.Int()
	}
	return v
}

func (s *Scanner) Ints2(size int) ([]int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
	}
	return v1, v2
}

func (s *Scanner) Ints3(size int) ([]int, []int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	v3 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
	}
	return v1, v2, v3
}

func (s *Scanner) Ints4(size int) ([]int, []int, []int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	v3 := make([]int, size)
	v4 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
		v4[i] = s.Int()
	}
	return v1, v2, v3, v4
}

func (s *Scanner) IntsN(h, w int) [][]int {
	v := make([][]int, h)
	for i := 0; i < h; i++ {
		v[i] = make([]int, w)
		for j := 0; j < w; j++ {
			v[i][j] = s.Int()
		}
	}
	return v
}

func (s *Scanner) Bytes() []byte {
	s.Scan()
	return s.Scanner.Bytes()
}

func (s *Scanner) BytesN(h int) [][]byte {
	v := make([][]byte, h)
	for i := 0; i < h; i++ {
		v[i] = s.Bytes()
	}
	return v
}

func (s *Scanner) Byte() byte {
	return s.Bytes()[0]
}

func (s *Scanner) ByteN(n int) []byte {
	v := make([]byte, n)
	for i := 0; i < n; i++ {
		v[i] = s.Byte()
	}
	return v
}
