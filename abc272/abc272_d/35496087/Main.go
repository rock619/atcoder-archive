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

func solve(o Printer, N int, M int) {
	pairs := make([]Pair, 0)
	for i := 0; i < 400; i++ {
		for j := i; j < 400; j++ {
			if i*i+j*j == M {
				pairs = append(pairs, Pair{A: i, B: j}, Pair{A: -i, B: j}, Pair{A: i, B: -j}, Pair{A: -i, B: -j})
				if i != j {
					pairs = append(pairs, Pair{A: j, B: i}, Pair{A: -j, B: i}, Pair{A: j, B: -i}, Pair{A: -j, B: -i})
				}
			}
		}
	}

	results := Make2D(N, N, -1)
	results[0][0] = 0
	points := make([]Pair, 1)
	points[0] = Pair{0, 0}
	for i := 1; i < 160000; i++ {
		nexts := make([]Pair, 0)
		for _, point := range points {
			x, y := point.A, point.B
			for _, diff := range pairs {
				nextX, nextY := x-diff.A, y-diff.B
				if nextX < 0 || nextX >= N || nextY < 0 || nextY >= N {
					continue
				}
				if results[nextY][nextX] != -1 {
					continue
				}
				results[nextY][nextX] = i
				nexts = append(nexts, Pair{nextX, nextY})
			}
		}
		if len(nexts) == 0 {
			break
		}
		points = nexts
	}
	for i := range results {
		for j := range results[i] {
			if j == 0 {
				o.p(results[i][j])
			} else {
				o.p(" ", results[i][j])
			}
		}
		o.l()
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

func main() {
	sc := NewScanner()
	N := sc.Int()
	M := sc.Int()
	out := NewPrinter()
	solve(out, N, M)
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
