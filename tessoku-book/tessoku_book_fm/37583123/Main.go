package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Meeting struct {
	L, R int
}

func solve(o Printer, N int, K int, L []int, R []int) {
	meetings := make([]Meeting, N)
	for i := range L {
		meetings[i] = Meeting{L: L[i], R: R[i] + K}
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].R < meetings[j].R
	})
	dpL := make([]int, 86400+K+1)
	dpR := make([]int, 86400+K+1)
	for currentTime, i := 0, 0; i < N; i++ {
		if currentTime > meetings[i].L {
			continue
		}
		currentTime = meetings[i].R
		dpL[currentTime]++
	}

	for i := 1; i <= 86400+K; i++ {
		dpL[i] += dpL[i-1]
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].L > meetings[j].L
	})
	for currentTime, i := 86400+K, 0; i < N; i++ {
		if currentTime < meetings[i].R {
			continue
		}
		currentTime = meetings[i].L
		dpR[currentTime]++
	}
	for i := 86400 + K - 1; i >= 0; i-- {
		dpR[i] += dpR[i+1]
	}

	for i := range L {
		o.l(dpL[L[i]] + dpR[R[i]+K] + 1)
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	K := sc.Int()
	L := make([]int, N)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		L[i] = sc.Int()
		R[i] = sc.Int()
	}
	out := NewPrinter()
	solve(out, N, K, L, R)
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
