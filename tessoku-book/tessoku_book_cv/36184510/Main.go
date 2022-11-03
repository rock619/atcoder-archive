package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(o Printer, N int, X []int, Y []int) {
	inf := float64(1 << 62)
	min := inf
	for l := 0; l < N; l++ {
		dp := Make2D(1<<N, N, inf)
		dp[1<<l][l] = 0
		for j := 0; j < (1 << N); j++ {
			for i := 0; i < N; i++ {
				for k := 0; k < N; k++ {
					if k != i && j&(1<<k) == 0 {
						UpdateMin(&dp[j+(1<<k)][k], dp[j][i]+Distance(X[k], X[i], Y[k], Y[i]))
					}
				}
			}
		}
		result := dp[(1<<N)-1]
		for i := 0; i < N; i++ {
			if i == l {
				result[i] = 1 << 62
				continue
			}
			result[i] += Distance(X[l], X[i], Y[l], Y[i])
		}
		UpdateMin(&min, Min(result...))
	}
	o.l(min)
}

type T = float64

func Make2D(len1, len2 int, value T) [][]T {
	result := make([][]T, len1)
	for i := range result {
		result[i] = make([]T, len2)
		for j := range result[i] {
			result[i][j] = value
		}
	}
	return result
}

func UpdateMin(min *T, v T) {
	if v < *min {
		*min = v
	}
}

func Distance(xa, xb, ya, yb int) float64 {
	return math.Sqrt(float64((xa-xb)*(xa-xb) + (ya-yb)*(ya-yb)))
}

func Min(v ...T) T {
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
	sc := NewScanner()
	N := sc.Int()
	X := make([]int, N)
	Y := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = sc.Int()
		Y[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, N, X, Y)
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
