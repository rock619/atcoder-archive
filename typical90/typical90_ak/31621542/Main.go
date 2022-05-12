package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

type S int64

const negInf S = -(1 << 62)

func (st *SegTree) op(l, r S) S {
	return S(Max(int64(l), int64(r)))
}

func (st *SegTree) e() S {
	return negInf
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

func (st *SegTree) copy() *SegTree {
	s := make([]S, st.n)
	for i := int64(0); i < st.n; i++ {
		s[i] = st.Get(i)
	}
	return NewSegTree(s)
}

func solve(W int64, N int64, L []int64, R []int64, V []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	s := make([]S, W+1)
	for j := int64(1); j <= W; j++ {
		s[j] = negInf
	}
	current := NewSegTree(s)

	for i := int64(1); i <= N; i++ {
		ss := make([]S, W+1)
		for j := int64(0); j <= W; j++ {
			ss[j] = current.Get(j)
		}
		prev := current
		current = current.copy()
		for j := int64(0); j <= W; j++ {
			l := Max(0, j-R[i-1])
			r := Max(0, j-L[i-1]+1)
			if l == r {
				continue
			}

			maxV := prev.Prod(l, r)
			if maxV == negInf {
				continue
			}

			current.Set(
				j,
				S(Max(int64(prev.Get(j)), V[i-1]+int64(maxV))),
			)
		}
	}

	if v := current.Get(W); v == negInf {
		fmt.Fprintln(w, -1)
	} else {
		fmt.Fprintln(w, v)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	W, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	L := make([]int64, N)
	R := make([]int64, N)
	V := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		L[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		R[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		V[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(W, N, L, R, V)
}
