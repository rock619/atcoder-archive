package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(o Printer, H int, W int, X [][]int, Q int, A []int, B []int, C []int, D []int) {
	sums := make([][]int, H+1)
	for i := range sums {
		sums[i] = make([]int, W+1)
		if i == 0 {
			continue
		}
		for j := 1; j <= W; j++ {
			sums[i][j] = sums[i][j-1] + X[i-1][j-1]
		}
	}

	for i := 1; i <= W; i++ {
		for j := 1; j <= H; j++ {
			sums[j][i] += sums[j-1][i]
		}
	}

	for i := range A {
		o.l(sums[C[i]][D[i]] + sums[A[i]-1][B[i]-1] - sums[A[i]-1][D[i]] - sums[C[i]][B[i]-1])
	}
}

func main() {
	sc := NewScanner()
	H := sc.Int()
	W := sc.Int()
	X := make([][]int, H)
	for i := 0; i < H; i++ {
		X[i] = make([]int, W)
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			X[i][j] = sc.Int()
		}
	}
	Q := sc.Int()
	A := make([]int, Q)
	B := make([]int, Q)
	C := make([]int, Q)
	D := make([]int, Q)
	for i := 0; i < Q; i++ {
		A[i] = sc.Int()
		B[i] = sc.Int()
		C[i] = sc.Int()
		D[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, H, W, X, Q, A, B, C, D)
	if err := out.w.Flush(); err != nil {
		panic(err)
	}
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), math.MaxInt64)
	s.Split(bufio.ScanWords)
	return &Scanner{
		Scanner: s,
	}
}

func (s *Scanner) Scan() {
	if ok := s.Scanner.Scan(); !ok {
		panic(s.Err())
	}
}

func (s *Scanner) Int() int {
	s.Scan()
	v, err := strconv.Atoi(s.Scanner.Text())
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
	s.Scan()
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

func (s *Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) Text() string {
	s.Scan()
	return s.Scanner.Text()
}

type Printer interface {
	// p fmt.Print
	p(a ...interface{})
	// f fmt.Printf
	f(format string, a ...interface{})
	// l fmt.Println
	l(a ...interface{})
}

type printer struct {
	w *bufio.Writer
}

func NewPrinter() *printer {
	return &printer{bufio.NewWriter(os.Stdout)}
}

func (p *printer) p(a ...interface{}) {
	fmt.Fprint(p.w, a...)
}

func (p *printer) f(format string, a ...interface{}) {
	fmt.Fprintf(p.w, format, a...)
}

func (p *printer) l(a ...interface{}) {
	fmt.Fprintln(p.w, a...)
}
