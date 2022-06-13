package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func solve(N, Q int, x, y, z, ww []int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	w2 := make([]bool, Q)
	result := 1
	for i := 0; i < 60; i++ {
		for j := range ww {
			w2[j] = (ww[j]/(1<<i))%2 == 1
		}
		result *= bitwiseSearch(N, Q, x, y, z, w2)
		result %= mod
	}

	fmt.Fprintln(w, result)
}

func bitwiseSearch(N, Q int, x, y, z []int, w []bool) int {
	count := 0
	for i := 0; i < (1 << N); i++ {
		bit := make([]bool, 15)
		for j := 0; j < N; j++ {
			bit[j+1] = (i/(1<<j))%2 == 1
		}

		ok := true
		for j := range x {
			if (bit[x[j]] || bit[y[j]] || bit[z[j]]) != w[j] {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}
	return count
}

func main() {
	s := NewScanner()
	N := s.Int()
	Q := s.Int()
	x, y, z, w := s.Ints4(Q)

	solve(N, Q, x, y, z, w)
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
