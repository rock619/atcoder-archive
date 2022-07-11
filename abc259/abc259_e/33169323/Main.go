package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, m []int, p, e [][]int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	pfMap := make(map[int]int)

	for i := range m {
		for j := range p[i] {
			if e[i][j] > pfMap[p[i][j]] {
				pfMap[p[i][j]] = e[i][j]
			}

		}
	}

	doubled := make(map[int]bool)
	for i := range m {
		for j := range p[i] {
			if e[i][j] == pfMap[p[i][j]] {
				if _, ok := doubled[p[i][j]]; ok {
					doubled[p[i][j]] = true
				} else {
					doubled[p[i][j]] = false
				}
			}
		}
	}

	count := 1
	for i := range m {
		for j := range p[i] {
			if e[i][j] == pfMap[p[i][j]] && !doubled[p[i][j]] {
				count++
				break
			}
		}
	}

	fmt.Fprintln(w, Min(N, count))
}

func Min(v ...int) int {
	switch len(v) {
	case 0:
		panic("Min: len(v) == 0")
	case 1:
		return v[0]
	case 2:
		if v[0] < v[1] {
			return v[0]
		}
		return v[1]
	default:
		m := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] < m {
				m = v[i]
			}
		}
		return m
	}
}

func main() {
	s := NewScanner()
	N := s.Int()
	m := make([]int, N)
	p := make([][]int, N)
	e := make([][]int, N)
	for i := 0; i < N; i++ {
		m[i] = s.Int()
		p[i] = make([]int, m[i])
		e[i] = make([]int, m[i])
		for j := 0; j < m[i]; j++ {
			p[i][j] = s.Int()
			e[i][j] = s.Int()
		}
	}
	solve(N, m, p, e)
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), 1048576)
	s.Split(bufio.ScanWords)
	return &Scanner{
		Scanner: s,
	}
}

func (s *Scanner) Int() int {
	if ok := s.Scan(); !ok {
		panic(s.Err())
	}
	v, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) IntN(n int) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = s.Int()
	}
	return v
}

func (s *Scanner) IntN2(n int) ([]int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	for i := 0; i < n; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
	}
	return v1, v2
}

func (s *Scanner) IntN3(n int) ([]int, []int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	v3 := make([]int, n)
	for i := 0; i < n; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
	}
	return v1, v2, v3
}

func (s *Scanner) IntN4(n int) ([]int, []int, []int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	v3 := make([]int, n)
	v4 := make([]int, n)
	for i := 0; i < n; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
		v4[i] = s.Int()
	}
	return v1, v2, v3, v4
}

func (s *Scanner) IntNN(h, w int) [][]int {
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
	if ok := s.Scan(); !ok {
		panic(s.Err())
	}

	b := s.Scanner.Bytes()
	v := make([]byte, len(b))
	copy(v, b)
	return v
}

func (s *Scanner) BytesN(n int) [][]byte {
	v := make([][]byte, n)
	for i := 0; i < n; i++ {
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

func (s Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}
