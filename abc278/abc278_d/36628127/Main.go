package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
)

func solve(o Printer, N int, A []int, Q int, query []Query) {
	ss := make([]S, N)
	for i := range ss {
		ss[i] = S{
			Val:  A[i],
			Size: 1,
		}
	}
	lst := NewLazySegTree(ss)
	for _, q := range query {
		switch q.T {
		case 1:
			lst.Apply(0, N, q.X)
		case 2:
			lst.Set(q.I-1, S{
				Val:  lst.Get(q.I-1).Val + q.X,
				Size: 1,
			})
		case 3:
			o.l(lst.Prod(q.I-1, q.I).Val)
		}
	}
}

type Query struct {
	T, I, X int
}

func main() {
	sc := NewScanner()
	N := sc.Int()
	A := sc.IntN(N)
	Q := sc.Int()
	query := make([]Query, Q)
	for i := 0; i < Q; i++ {
		query[i].T = sc.Int()
		switch query[i].T {
		case 1:
			query[i].X = sc.Int()
		case 2:
			query[i].I = sc.Int()
			query[i].X = sc.Int()
		case 3:
			query[i].I = sc.Int()
		}
	}
	out := NewPrinter()
	solve(out, N, A, Q, query)
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

type S struct {
	Val, Size int
}

type F = int

const id F = 8e18

func (lst *LazySegTree) op(l, r S) S {
	return S{
		Val:  l.Val + r.Val,
		Size: l.Size + r.Size,
	}
}

func (lst *LazySegTree) e() S {
	return S{}
}

func (lst *LazySegTree) mapping(l F, r S) S {
	if l != id {
		r.Val = l * r.Size
	}
	return r
}

func (lst *LazySegTree) composition(l, r F) F {
	if l == id {
		return r
	}
	return l
}

func (lst *LazySegTree) id() F {
	return id
}

type LazySegTree struct {
	n    int
	log  int
	size int
	lz   []F
	data []S
}

func NewLazySegTree(
	v []S,
) *LazySegTree {
	lst := &LazySegTree{}

	lst.n = len(v)
	lst.log = bits.Len(uint(lst.n - 1))
	lst.size = 1 << lst.log
	lst.data = make([]S, 2*lst.size)
	lst.lz = make([]F, lst.size)
	for i := range lst.data {
		lst.data[i] = lst.e()
	}
	for i := range lst.lz {
		lst.lz[i] = lst.id()
	}
	for i := 0; i < lst.n; i++ {
		lst.data[lst.size+i] = v[i]
	}
	for i := lst.size - 1; i >= 1; i-- {
		lst.update(i)
	}
	return lst
}

func (lst *LazySegTree) Set(p int, x S) {
	p += lst.size
	lst.data[p] = x
	for i := lst.log; i >= 1; i-- {
		lst.push(p >> i)
	}
	lst.data[p] = x
	for i := 1; i <= lst.log; i++ {
		lst.update(p >> i)
	}
}

func (lst *LazySegTree) Get(p int) S {
	p += lst.size
	for i := lst.log; i >= 1; i-- {
		lst.push(p >> i)
	}
	return lst.data[p]
}

func (lst *LazySegTree) Prod(l, r int) S {
	if l == r {
		return lst.e()
	}
	l += lst.size
	r += lst.size
	for i := lst.log; i >= 1; i-- {
		if ((l >> i) << i) != l {
			lst.push(l >> i)
		}
		if ((r >> i) << i) != r {
			lst.push((r - 1) >> i)
		}
	}
	sml, smr := lst.e(), lst.e()
	for l < r {
		if l&1 != 0 {
			sml = lst.op(sml, lst.data[l])
			l++
		}
		if r&1 != 0 {
			r--
			smr = lst.op(lst.data[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return lst.op(sml, smr)
}

func (lst *LazySegTree) AllProd() S {
	return lst.data[1]
}

func (lst *LazySegTree) Apply(l, r int, f F) {
	if l == r {
		return
	}
	l += lst.size
	r += lst.size
	for i := lst.log; i >= 1; i-- {
		if ((l >> i) << i) != l {
			lst.push(l >> i)
		}
		if ((r >> i) << i) != r {
			lst.push((r - 1) >> i)
		}
	}

	l2, r2 := l, r
	for l < r {
		if l&1 != 0 {
			lst.allApply(l, f)
			l++
		}
		if r&1 != 0 {
			r--
			lst.allApply(r, f)
		}
		l >>= 1
		r >>= 1
	}
	l, r = l2, r2

	for i := 1; i <= lst.log; i++ {
		if ((l >> i) << i) != l {
			lst.update(l >> i)
		}
		if ((r >> i) << i) != r {
			lst.update((r - 1) >> i)
		}
	}
}

func (lst *LazySegTree) MaxRight(l int, g func(x S) bool) int {
	if l == lst.n {
		return lst.n
	}
	l += lst.size
	for i := lst.log; i >= 1; i-- {
		lst.push(l >> i)
	}
	sm := lst.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !g(lst.op(sm, lst.data[l])) {
			for l < lst.size {
				lst.push(l)
				l *= 2
				if g(lst.op(sm, lst.data[l])) {
					sm = lst.op(sm, lst.data[l])
					l++
				}
			}
			return l - lst.size
		}
		sm = lst.op(sm, lst.data[l])
		l++
		if l&-l == l {
			break
		}
	}
	return lst.n
}

func (lst *LazySegTree) MinLeft(r int, g func(x S) bool) int {
	if r == 0 {
		return 0
	}
	r += lst.size
	for i := lst.log; i >= 1; i-- {
		lst.push((r - 1) >> i)
	}
	sm := lst.e()
	for {
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !g(lst.op(lst.data[r], sm)) {
			for r < lst.size {
				lst.push(r)
				r *= 2
				r++
				if g(lst.op(lst.data[r], sm)) {
					sm = lst.op(lst.data[r], sm)
					r--
				}
			}
			return r + 1 - lst.size
		}
		sm = lst.op(lst.data[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}

func (lst *LazySegTree) update(k int) {
	lst.data[k] = lst.op(lst.data[2*k], lst.data[2*k+1])
}

func (lst *LazySegTree) push(k int) {
	lst.allApply(2*k, lst.lz[k])
	lst.allApply(2*k+1, lst.lz[k])
	lst.lz[k] = lst.id()
}

func (lst *LazySegTree) allApply(k int, f F) {
	lst.data[k] = lst.mapping(f, lst.data[k])
	if k < lst.size {
		lst.lz[k] = lst.composition(f, lst.lz[k])
	}
}
