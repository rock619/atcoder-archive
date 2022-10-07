package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func solve(logger *Logger, o Printer, N int, A []int, D int, L []int, R []int) {
	accL := make([]int, N+1)
	for i := 0; i < N; i++ {
		accL[i+1] = Max(accL[i], A[i])
	}
	accR := make([]int, N+1)
	for i := N; i > 0; i-- {
		accR[i-1] = Max(accR[i], A[i-1])
	}
	logger.p(accL)
	logger.p(accR)
	for i := range L {
		o.l(Max(accL[L[i]-1], accR[R[i]]))
	}
	logger.Disable()
}

func Max(v ...int) int {
	switch len(v) {
	case 0:
		panic("Max: len(v) == 0")
	case 1:
		return v[0]
	case 2:
		if v[0] > v[1] {
			return v[0]
		}
		return v[1]
	default:
		m := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] > m {
				m = v[i]
			}
		}
		return m
	}
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = sc.Int()
	}
	D := sc.Int()
	L := make([]int, D)
	R := make([]int, D)
	for i := 0; i < D; i++ {
		L[i] = sc.Int()
		R[i] = sc.Int()
	}
	out := NewPrinter()
	logger := NewLogger()
	defer logger.Flush()
	solve(logger, out, N, A, D, L, R)
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

func (p *printer) Flush() {
	if err := p.w.Flush(); err != nil {
		panic(err)
	}
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

type Logger struct {
	logger  *log.Logger
	Enabled bool
}

func NewLogger() *Logger {
	return &Logger{
		logger:  log.New(bufio.NewWriter(os.Stderr), "", log.Ltime|log.Lmicroseconds|log.Lshortfile),
		Enabled: true,
	}
}

func (l *Logger) Flush() {
	if l.Enabled {
		l.logger.Writer().(*bufio.Writer).Flush()
	}
}

func (l *Logger) Disable() {
	l.Enabled = false
}

func (l *Logger) p(v ...interface{}) {
	l.logger.Print(v...)
}

func (l *Logger) f(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) l(v ...interface{}) {
	l.logger.Println(v...)
}
