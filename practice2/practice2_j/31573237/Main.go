package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

type S int64

func (st *SegTree) op(l, r S) S {
	return S(Max(int64(l), int64(r)))
}

func (st *SegTree) e() S {
	return -1
}

type SegTree struct {
	n    int64
	log  int64
	size int64
	data []S
}

func NewSegTree(v []S) *SegTree {
	st := &SegTree{}

	st.n = int64(len(v))
	st.log = int64(bits.Len(uint(st.n - 1)))
	st.size = int64(1) << st.log
	st.data = make([]S, 2*st.size)
	for i := range st.data {
		st.data[i] = st.e()
	}
	for i := int64(0); i < st.n; i++ {
		st.data[st.size+i] = v[i]
	}
	for i := st.size - 1; i >= 1; i-- {
		st.update(i)
	}
	return st
}

func (st *SegTree) Set(p int64, x S) {
	p += st.size
	st.data[p] = x
	for i := int64(1); i <= st.log; i++ {
		st.update(p >> i)
	}
}

func (st *SegTree) Get(p int64) S {
	return st.data[p+st.size]
}

func (st *SegTree) Prod(l, r int64) S {
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

func (st *SegTree) MaxRight(l int64, f func(x S) bool) int64 {
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

func (st *SegTree) MinLeft(r int64, f func(x S) bool) int64 {
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
				r = (2*r + 1)
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

func (st *SegTree) update(k int64) {
	st.data[k] = st.op(st.data[2*k], st.data[2*k+1])
}

func Max(ints ...int64) int64 {
	if len(ints) == 0 {
		panic("Max: len(ints) == 0")
	}
	m := ints[0]
	for i := 1; i < len(ints); i++ {
		if ints[i] > m {
			m = ints[i]
		}
	}
	return m
}

func solve(N, Q int64, a []int64, queries []Query) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	ss := make([]S, N)
	for i := int64(0); i < N; i++ {
		ss[i] = S(a[i])
	}

	st := NewSegTree(ss)
	for i := int64(0); i < Q; i++ {
		switch q := queries[i]; q.T {
		case 1:
			st.Set(q.X-1, S(q.V))
		case 2:
			fmt.Fprintln(w, st.Prod(q.L-1, q.R))
		default:
			fmt.Fprintln(w, st.MaxRight(q.X-1, func(x S) bool { return x < S(q.V) })+1)
		}
	}
}

type Query struct {
	T, X, V, L, R int64
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	a := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	queries := make([]Query, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		queries[i].T, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		if queries[i].T == 2 {
			scanner.Scan()
			queries[i].L, err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			scanner.Scan()
			queries[i].R, err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			continue
		}

		scanner.Scan()
		queries[i].X, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		queries[i].V, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, Q, a, queries)
}
