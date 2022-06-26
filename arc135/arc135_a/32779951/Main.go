package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 998244353

func solve(X int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	maxes := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}
	if X <= 4 {
		fmt.Fprintln(w, maxes[X])
		return
	}

	c := 0
	m := make(map[int]int)
	m[X] = 1
	for done := false; !done; {
		c++
		m2 := make(map[int]int)
		for k, v := range m {
			f, c := FloorCeil(k)
			m2[f] += v
			m2[c] += v
		}

		done = true
		for k := range m2 {
			if k > 4 {
				done = false
			}
		}
		m = m2
	}

	result := 1
	for k, v := range m {
		result *= PowMod(maxes[k], v, mod)
		result %= mod
	}
	fmt.Fprintln(w, result)
}

func SafeMod(x, m int) int {
	x %= m
	if x < 0 {
		return x + m
	}
	return x
}

// PowMod (x ** n) % m
func PowMod(x, n, m int) int {
	if m == 1 {
		return 0
	}

	r := 1
	for y := SafeMod(x, m); n > 0; n >>= 1 {
		if n&1 != 0 {
			r = (r * y) % m
		}
		y = (y * y) % m
	}
	return r
}

func FloorCeil(n int) (floor, ceil int) {
	if f := n / 2; n%2 == 0 {
		return f, f
	} else {
		return f, f + 1
	}
}

func main() {
	s := NewScanner()

	X := s.Int()

	solve(X)
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
