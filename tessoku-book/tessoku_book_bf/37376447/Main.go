package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
)

func solve(o Printer, N, Q int, queries []Query) {
	a := make([]int, N)
	segTree := NewSegTree(a)
	for _, q := range queries {
		if q.T == 1 {
			segTree.Set(q.Pos-1, q.X)
		} else {
			o.l(segTree.Prod(q.L-1, q.R-1))
		}
	}
}

type S = int

func (st *SegTree) op(l, r S) S {
	if l > r {
		return l
	}
	return r
}

func (st *SegTree) e() S {
	return -1
}

type SegTree struct {
	n    int
	log  int
	size int
	data []S
}

func NewSegTree(v []S) *SegTree {
	st := &SegTree{}

	st.n = len(v)
	st.log = bits.Len(uint(st.n - 1))
	st.size = 1 << st.log
	st.data = make([]S, 2*st.size)
	for i := range st.data {
		st.data[i] = st.e()
	}
	for i := 0; i < st.n; i++ {
		st.data[st.size+i] = v[i]
	}
	for i := st.size - 1; i >= 1; i-- {
		st.update(i)
	}
	return st
}

func (st *SegTree) Set(p int, x S) {
	p += st.size
	st.data[p] = x
	for i := 1; i <= st.log; i++ {
		st.update(p >> i)
	}
}

func (st *SegTree) Get(p int) S {
	return st.data[p+st.size]
}

func (st *SegTree) Prod(l, r int) S {
	sml, smr := st.e(), st.e()
	l += st.size
	r += st.size

	for l < r {
		if l&1 != 0 {
			sml = st.op(sml, st.data[l])
			l++
		}
		if r&1 != 0 {
			r--
			smr = st.op(st.data[r], smr)
		}
		l >>= 1
		r >>= 1
	}

	return st.op(sml, smr)
}

func (st *SegTree) AllProd() S {
	return st.data[1]
}

func (st *SegTree) MaxRight(l int, f func(x S) bool) int {
	if l == st.n {
		return st.n
	}

	l += st.size
	sm := st.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !f(st.op(sm, st.data[l])) {
			for l < st.size {
				l *= 2
				if f(st.op(sm, st.data[l])) {
					sm = st.op(sm, st.data[l])
					l++
				}
			}
			return l - st.size
		}
		sm = st.op(sm, st.data[l])
		l++
		if l&-l == l {
			break
		}
	}
	return st.n
}

func (st *SegTree) MinLeft(r int, f func(x S) bool) int {
	if r == 0 {
		return 0
	}

	r += st.size
	sm := st.e()
	for {
		r--
		for r > 0 && r%2 != 0 {
			r >>= 1
		}
		if !f(st.op(st.data[r], sm)) {
			for r < st.size {
				r *= 2
				r++
				if f(st.op(st.data[r], sm)) {
					sm = st.op(st.data[r], sm)
					r--
				}
			}
			return r + 1 - st.size
		}
		sm = st.op(st.data[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}

func (st *SegTree) update(k int) {
	st.data[k] = st.op(st.data[2*k], st.data[2*k+1])
}

type Query struct {
	T, Pos, X, L, R int
}

func main() {
	sc := NewScanner()
	N, Q := sc.Int(), sc.Int()
	queries := make([]Query, Q)
	for i := range queries {

		if queries[i].T = sc.Int(); queries[i].T == 1 {
			queries[i].Pos, queries[i].X = sc.Int(), sc.Int()
		} else {
			queries[i].L, queries[i].R = sc.Int(), sc.Int()
		}
	}
	out := NewPrinter()
	solve(out, N, Q, queries)
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
