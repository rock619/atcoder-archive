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
	// _     = iota
	debug = iota
)

func Rotate(n int) (int, bool) {
	if n < 10 || n%10 == 0 {
		return n, false
	}
	s := strconv.Itoa(n)
	i, _ := strconv.Atoi(string(s[len(s)-1]) + s[:len(s)-1])
	return i, true
}

type QueueElement = int

type Queue struct {
	Size int
	head int
	s    []QueueElement
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		s: make([]QueueElement, 0, capacity),
	}
}

func (q *Queue) Empty() bool {
	return q.Size == 0
}

func (q *Queue) Clear() {
	q.Size = 0
}

func (q *Queue) Push(x QueueElement) {
	q.Size++
	if q.Size > len(q.s) {
		q.reorder()
		q.s = append(q.s, x)
		return
	}
	tail := q.head + q.Size - 1
	if tail >= len(q.s) {
		tail -= len(q.s)
	}
	q.s[tail] = x
}

func (q *Queue) Pop() (x QueueElement, ok bool) {
	if q.Empty() {
		return x, false
	}
	x = q.s[q.head]
	if q.head++; q.head >= len(q.s) {
		q.head -= len(q.s)
	}
	q.Size--
	return x, true
}

func (q *Queue) Peek() (x QueueElement, ok bool) {
	if q.Empty() {
		return x, false
	}
	return q.s[q.head], true
}

func (q *Queue) reorder() {
	q.s = append(q.s[q.head:], q.s[:q.head]...)
	q.head = 0
}

func solve(o, lg Printer, a int, N int) {
	s := make([]int, 1000000)
	q := NewQueue(1)
	q.Push(1)
	for !q.Empty() {
		x, _ := q.Pop()
		i := s[x]
		lg.p(i, x)
		if x == N {
			o.l(i)
			return
		}
		xa := x * a
		if xa < 1000000 {
			if s[xa] == 0 {
				s[xa] = i + 1
				q.Push(xa)
			}
		}
		for {
			rotated, ok := Rotate(x)
			if !ok {
				break
			}
			if s[rotated] != 0 {
				break
			}
			s[rotated] = i + 1
			q.Push(rotated)
		}
	}
	o.l(-1)
}

func main() {
	sc := NewScanner()
	a := sc.Int()
	N := sc.Int()
	stdout := bufio.NewWriter(os.Stdout)
	out := NewPrinter(stdout)
	logger := NewLogger()
	solve(out, logger, a, N)
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
