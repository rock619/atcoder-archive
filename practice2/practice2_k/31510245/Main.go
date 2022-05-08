package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

const mod = 998244353

type S struct {
	A, Size int64
}

type F struct {
	A, B int64
}

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

func (lst *LazySegTree) op(l, r S) S { return S{A: (l.A + r.A) % mod, Size: l.Size + r.Size} }

func (lst *LazySegTree) e() S { return S{} }

func (lst *LazySegTree) mapping(l F, r S) S {
	return S{A: (r.A*l.A%mod + r.Size*l.B%mod) % mod, Size: r.Size}
}

func (lst *LazySegTree) composition(l, r F) F {
	return F{A: r.A * l.A % mod, B: (r.B*l.A%mod + l.B) % mod}
}

func (lst *LazySegTree) id() F { return F{A: 1, B: 0} }

type Query struct {
	Type, L, R, B, C int64
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
		ss[i] = S{
			A:    a[i],
			Size: 1,
		}
	}

	lst := NewLazySegTree(ss)

	for i := int64(0); i < Q; i++ {
		q := queries[i]
		l, r := q.L, q.R
		if q.Type == 0 {
			lst.Apply(l, r, F{A: q.B, B: q.C})
		} else {
			v := lst.Prod(l, r).A
			fmt.Fprintln(w, v)
		}
	}
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
		queries[i].Type, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
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
		if queries[i].Type == 1 {
			continue
		}

		scanner.Scan()
		queries[i].B, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		queries[i].C, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, Q, a, queries)
}
