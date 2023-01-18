package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	_     = iota
	debug = iota
)

func solve(o, lg Printer, Q int, t []int, x []int) {
	N := 1048576
	m := make(map[int]int)
	nexts := make([]int, 1048576)
	for i := 0; i < 1048576; i++ {
		nexts[i] = i + 1
	}
	nexts[len(nexts)-1] = 0
	for i := range t {
		ti, xi := t[i], x[i]
		if ti == 1 {
			h := xi % N
			for {
				_, ok := m[h]
				if !ok {
					m[h] = xi
					break
				}
				h, nexts[h] = nexts[h], nexts[nexts[h]]
			}
		} else {
			v, ok := m[xi%N]
			if !ok {
				o.l(-1)
			} else {
				o.l(v)
			}
		}
	}
}

func main() {
	sc := NewScanner()
	Q := sc.Int()
	t := make([]int, Q)
	x := make([]int, Q)
	for i := 0; i < Q; i++ {
		t[i] = sc.Int()
		x[i] = sc.Int()
	}
	stdout := bufio.NewWriter(os.Stdout)
	out := NewPrinter(stdout)
	logger := NewLogger()
	solve(out, logger, Q, t, x)
	if err := stdout.Flush(); err != nil {
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
	w io.Writer
}

func NewPrinter(w io.Writer) Printer {
	return &printer{w}
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
	if debug == 1 {
		l.Logger.Print(a...)
	}
}

func (l *logger) f(format string, a ...interface{}) {
	if debug == 1 {
		l.Logger.Printf(format, a...)
	}
}

func (l *logger) l(a ...interface{}) {
	if debug == 1 {
		l.Logger.Println(a...)
	}
}
