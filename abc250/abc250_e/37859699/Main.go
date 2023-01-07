package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	yes = "Yes"
	no  = "No"
)

const (
	// _     = iota
	debug = iota
)

type ZobristHash struct {
	fixedRandom uint64
}

func NewZobristHash() *ZobristHash {
	rand.Seed(time.Now().UnixNano())
	return &ZobristHash{
		fixedRandom: rand.Uint64(),
	}
}

func (h *ZobristHash) Values(s []uint64) []uint64 {
	m := make(map[uint64]struct{})
	hashes := make([]uint64, len(s))
	for i, v := range s {
		if _, ok := m[v]; ok {
			hashes[i] = hashes[i-1]
			continue
		}
		m[v] = struct{}{}
		if i == 0 {
			hashes[i] = h.Of(v)
		} else {
			hashes[i] = hashes[i-1] ^ h.Of(v)
		}
	}
	return hashes
}

func (h *ZobristHash) Of(x uint64) uint64 {
	return h.splitmix64(x + h.fixedRandom)
}

func (*ZobristHash) splitmix64(x uint64) uint64 {
	x += 0x9e3779b97f4a7c15
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	return x ^ (x >> 31)
}

func solve(o, lg Printer, N int, a []int, b []int, Q int, x []int, y []int) {
	h := NewZobristHash()
	a2, b2 := make([]uint64, N), make([]uint64, N)
	for i := range a {
		a2[i] = uint64(a[i])
		b2[i] = uint64(b[i])
	}
	lg.p(a2)
	lg.p(b2)
	ah, bh := h.Values(a2), h.Values(b2)
	lg.p(ah)
	lg.p(bh)
	for i := range x {
		lg.p(x[i]-1, y[i]-1, ah[x[i]-1], bh[y[i]-1], ah[x[i]-1] == bh[y[i]-1])
		if ah[x[i]-1] == bh[y[i]-1] {
			o.l(yes)
		} else {
			o.l(no)
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
	b := make([]int, N)
	for i := 0; i < N; i++ {
		b[i] = sc.Int()
	}
	Q := sc.Int()
	x := make([]int, Q)
	y := make([]int, Q)
	for i := 0; i < Q; i++ {
		x[i] = sc.Int()
		y[i] = sc.Int()
	}
	stdout := bufio.NewWriter(os.Stdout)
	out := NewPrinter(stdout)
	logger := NewLogger()
	solve(out, logger, N, a, b, Q, x, y)
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
