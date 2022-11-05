package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Pair struct {
	A, B int
}

func solve(o Printer, N int, a []int) {
	gcd := a[0]
	for i := 1; i < N; i++ {
		gcd = GCD(gcd, a[i])
	}
	for i := range a {
		a[i] /= gcd
	}

	m := make(map[int]Pair)
	for i := 0; i <= 1000_000_000; i++ {
		vi := 1 << i
		if vi > 1000_000_000 {
			break
		}
		for j, vj := 0, 1; j <= 1000_000_000; j++ {
			v := vi * vj
			if v > 1000_000_000 {
				break
			}
			m[v] = Pair{A: i, B: j}
			vj *= 3
		}
	}

	// log.Print(m)

	min := Pair{A: 1 << 62, B: 1 << 62}
	sum := 0
	for _, aa := range a {
		v, ok := m[aa]
		if !ok {
			o.l(-1)
			return
		}
		sum += v.A + v.B
		UpdateMin(&min.A, v.A)
		UpdateMin(&min.B, v.B)
	}
	o.l(sum - N*(min.A+min.B))
}

func GCD(a, b int) int {
	for a > 0 && b > 0 {
		if a < b {
			b = b % a
		} else {
			a = a % b
		}
	}

	if a > 0 {
		return a
	}
	return b
}

func UpdateMin(min *int, v ...int) {
	for i := 0; i < len(v); i++ {
		if v[i] < *min {
			*min = v[i]
		}
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, N, a)
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
