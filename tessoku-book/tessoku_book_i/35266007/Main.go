package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(o Printer, H int, W int, N int, A []int, B []int, C []int, D []int) {
	diffs := Make2D(H+2, W+2, 0)
	for i := range A {
		diffs[A[i]][B[i]]++
		diffs[A[i]][D[i]+1]--
		diffs[C[i]+1][B[i]]--
		diffs[C[i]+1][D[i]+1]++
	}

	cumulatives := Accumulated(diffs)
	for i := 1; i <= H; i++ {
		o.l(Join(cumulatives[i][1 : W+1]))
	}
}

func Make2D(len1, len2, value int) [][]int {
	result := make([][]int, len1)
	for i := range result {
		result[i] = make([]int, len2)
		for j := range result[i] {
			result[i][j] = value
		}
	}
	return result
}

func Accumulated(s [][]int) [][]int {
	result := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		result[i] = make([]int, len(s[i]))
		copy(result[i], s[i])
		for j := 1; j < len(s[i]); j++ {
			result[i][j] += result[i][j-1]
		}
		if i == 0 {
			continue
		}
		for j := 1; j < len(s[i]); j++ {
			result[i][j] += result[i-1][j]
		}
	}
	return result
}

func Join(s []int) string {
	if len(s) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprint(&b, s[0])
	for i := 1; i < len(s); i++ {
		fmt.Fprint(&b, " ", s[i])
	}
	return b.String()
}

func main() {
	sc := NewScanner()
	H := sc.Int()
	W := sc.Int()
	N := sc.Int()
	A := make([]int, N)
	B := make([]int, N)
	C := make([]int, N)
	D := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = sc.Int()
		B[i] = sc.Int()
		C[i] = sc.Int()
		D[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, H, W, N, A, B, C, D)
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
