package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	yes = "Yes"
	no  = "No"
)

const (
	_     = iota
	debug = iota
)

type StackElement = byte

type Stack struct {
	Size int
	s    []StackElement
}

func NewStack(capacity int) *Stack {
	return &Stack{
		s: make([]StackElement, 0, capacity),
	}
}

func (s *Stack) Empty() bool {
	return s.Size == 0
}

func (s *Stack) Clear() {
	s.Size = 0
}

func (s *Stack) Push(x StackElement) {
	if s.Size >= len(s.s) {
		s.s = append(s.s, x)
	} else {
		s.s[s.Size] = x
	}
	s.Size++
}

func (s *Stack) Pop() (x StackElement, ok bool) {
	if s.Empty() {
		return x, false
	}
	s.Size--
	return s.s[s.Size], true
}

func (s *Stack) Peek() (x StackElement, ok bool) {
	if s.Empty() {
		return x, false
	}
	return s.s[s.Size-1], true
}

func (s *Stack) String() string {
	if s.Empty() {
		return "Stack(0)[]"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "Stack(%d)[%v", s.Size, s.s[0])
	for i := 1; i < s.Size; i++ {
		fmt.Fprintf(&b, ", %v", s.s[i])
	}
	fmt.Fprint(&b, "]<-top")
	return b.String()
}

func solve(o, lg Printer, S []byte) {
	m := make(map[byte]bool)
	stack := NewStack(1)
	for _, b := range S {
		switch b {
		case '(':
			stack.Push('(')
		case ')':
			for {
				x, ok := stack.Pop()
				if !ok {
					o.l(yes)
					return
				}
				if x == '(' {
					break
				}
				m[x] = false
			}
		default:
			if m[b] {
				o.l(no)
				return
			}
			m[b] = true
			stack.Push(b)
		}
	}
	o.l(yes)
}

func main() {
	sc := NewScanner()
	S := sc.Bytes()
	stdout := bufio.NewWriter(os.Stdout)
	out := NewPrinter(stdout)
	logger := NewLogger()
	solve(out, logger, S)
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
