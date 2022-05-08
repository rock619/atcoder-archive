package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

type S struct {
	Zero, One, Inversion int64
}

type F bool

func (lst *LazySegTree) op(l, r S) S {
	return S{
		Zero:      l.Zero + r.Zero,
		One:       l.One + r.One,
		Inversion: l.Inversion + r.Inversion + l.One*r.Zero,
	}
}

func (lst *LazySegTree) e() S { return S{} }

func (lst *LazySegTree) mapping(l F, r S) S {
	if !l {
		return r
	}
	return S{
		Zero:      r.One,
		One:       r.Zero,
		Inversion: r.One*r.Zero - r.Inversion,
	}
}

func (lst *LazySegTree) composition(l, r F) F {
	return (l && !r) || (!l && r)
}

func (lst *LazySegTree) id() F { return false }

type LazySegTree struct {
	size int64
	log  int64
	lz   []F
	data []S
}

func NewLazySegTree(
	v []S,
) *LazySegTree {
	lst := &LazySegTree{}

	n := int64(len(v))
	lst.log = int64(bits.Len(uint(n - 1)))
	lst.size = int64(1) << lst.log
	lst.data = make([]S, 2*lst.size)
	lst.lz = make([]F, lst.size)
	for i := range lst.data {
		lst.data[i] = lst.e()
	}
	for i := range lst.lz {
		lst.lz[i] = lst.id()
	}
	for i := int64(0); i < n; i++ {
		lst.data[lst.size+i] = v[i]
	}
	for i := lst.size - 1; i >= 1; i-- {
		lst.update(i)
	}
	return lst
}

func (lst *LazySegTree) Set(p int64, x S) {
	p += lst.size
	lst.data[p] = x
	for i := lst.log; i >= 1; i-- {
		lst.push(p >> i)
	}
	lst.data[p] = x
	for i := int64(1); i <= lst.log; i++ {
		lst.update(p >> i)
	}
}

func (lst *LazySegTree) Get(p int64) S {
	p += lst.size
	for i := lst.log; i >= 1; i-- {
		lst.push(p >> i)
	}
	return lst.data[p]
}

func (lst *LazySegTree) Prod(l, r int64) S {
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

func (lst *LazySegTree) Apply(l, r int64, f F) {
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
	l = l2
	r = r2
	for i := int64(1); i <= lst.log; i++ {
		if ((l >> i) << i) != l {
			lst.update(l >> i)
		}
		if ((r >> i) << i) != r {
			lst.update((r - 1) >> i)
		}
	}
}

func (lst *LazySegTree) update(k int64) {
	lst.data[k] = lst.op(lst.data[2*k], lst.data[2*k+1])
}

func (lst *LazySegTree) push(k int64) {
	lst.allApply(2*k, lst.lz[k])
	lst.allApply(2*k+1, lst.lz[k])
	lst.lz[k] = lst.id()
}

func (lst *LazySegTree) allApply(k int64, f F) {
	lst.data[k] = lst.mapping(f, lst.data[k])
	if k < lst.size {
		lst.lz[k] = lst.composition(f, lst.lz[k])
	}
}

func solve(N int64, Q int64, A []int64, T []int64, L []int64, R []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	a := make([]S, N)
	for i := int64(0); i < N; i++ {
		if A[i] == 0 {
			a[i] = S{Zero: 1}
		} else {
			a[i] = S{One: 1}
		}
	}

	lst := NewLazySegTree(a)
	for i := int64(0); i < Q; i++ {
		if T[i] == 1 {
			lst.Apply(L[i]-1, R[i], true)
		} else {
			fmt.Fprintln(w, lst.Prod(L[i]-1, R[i]).Inversion)
		}
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
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	T := make([]int64, Q)
	L := make([]int64, Q)
	R := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		T[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
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
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, Q, A, T, L, R)
}
