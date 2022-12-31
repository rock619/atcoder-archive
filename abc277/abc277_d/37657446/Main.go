package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type Interval struct {
	L, R, Sum int
}

func UpdateMin(min *int, v ...int) {
	for i := 0; i < len(v); i++ {
		if v[i] < *min {
			*min = v[i]
		}
	}
}

func Sum(v ...int) int {
	switch len(v) {
	case 0:
		return 0
	case 1:
		return v[0]
	case 2:
		return v[0] + v[1]
	default:
		s := v[0]
		for i := 1; i < len(v); i++ {
			s += v[i]
		}
		return s
	}
}

func solve(o, lg Printer, N int, M int, A []int) {
	if N == 1 {
		o.l(0)
		return
	}
	a := make([]int, N)
	copy(a, A)
	sort.Ints(a)
	sum := Sum(a...)
	// lg.p(a, sum)
	a = append(a, a...)
	subs := make([]Interval, 0)

	for i, v := range a {
		if len(subs) == 0 {
			subs = append(subs, Interval{L: i, R: i, Sum: v})
			continue
		}
		if i > 0 && (a[i] == a[i-1] || a[i] == a[i-1]+1 || (a[i] == 0 && a[i-1] == M-1)) {
			subs[len(subs)-1].R = i
			subs[len(subs)-1].Sum += v
		} else {
			subs = append(subs, Interval{L: i, R: i, Sum: v})
		}
	}

	// lg.f("%+v", subs)

	min := sum
	for _, v := range subs {
		UpdateMin(&min, sum-v.Sum)
	}
	if min < 0 {
		o.l(0)
	} else {
		o.l(min)
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	M := sc.Int()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = sc.Int()
	}
	out := NewPrinter()
	logger := NewLogger()
	solve(out, logger, N, M, A)
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

type logger struct {
	*log.Logger
}

func NewLogger() Printer {
	return &logger{
		log.New(os.Stderr, "", log.Lmicroseconds|log.Lshortfile),
	}
}

func (l *logger) p(a ...interface{}) {
	l.Logger.Print(a...)
}

func (l *logger) f(format string, a ...interface{}) {
	l.Logger.Printf(format, a...)
}

func (l *logger) l(a ...interface{}) {
	l.Logger.Println(a...)
}
