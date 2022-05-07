package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SegmentTree struct {
	size int64
	seg  []int64
	lazy []int64
}

func NewSegmentTree(n int64) *SegmentTree {
	size := int64(1)
	for size < n {
		size *= 2
	}
	return &SegmentTree{
		size: size,
		seg:  make([]int64, size*2),
		lazy: make([]int64, size*2),
	}
}

func (st *SegmentTree) push(k int64) {
	if k < st.size {
		st.lazy[k*2] = Max(st.lazy[k*2], st.lazy[k])
		st.lazy[k*2+1] = Max(st.lazy[k*2+1], st.lazy[k])
	}
	st.seg[k] = Max(st.seg[k], st.lazy[k])
	st.lazy[k] = 0
}

func (st *SegmentTree) update(a, b, x, k, l, r int64) {
	st.push(k)
	if r <= a || b <= l {
		return
	}
	if a <= l && r <= b {
		st.lazy[k] = x
		st.push(k)
		return
	}
	st.update(a, b, x, k*2, l, (l+r)>>1)
	st.update(a, b, x, k*2+1, (l+r)>>1, r)
	st.seg[k] = Max(st.seg[k*2], st.seg[k*2+1])
}

func (st *SegmentTree) rangeMax(a, b, k, l, r int64) int64 {
	st.push(k)
	if r <= a || b <= l {
		return 0
	}
	if a <= l && r <= b {
		return st.seg[k]
	}
	lc := st.rangeMax(a, b, k*2, l, (l+r)>>1)
	rc := st.rangeMax(a, b, k*2+1, (l+r)>>1, r)
	return Max(lc, rc)
}

func (st *SegmentTree) Update(l, r, x int64) {
	st.update(l, r, x, 1, 0, st.size)
}

func (st *SegmentTree) RangeMax(l, r int64) int64 {
	return st.rangeMax(l, r, 1, 0, st.size)
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

func solve(W int64, N int64, L []int64, R []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	st := NewSegmentTree(W)

	for i := int64(0); i < N; i++ {
		h := st.RangeMax(L[i]-1, R[i]) + 1
		st.Update(L[i]-1, R[i], h)
		fmt.Fprintln(w, h)
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
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(W, N, L, R)
}
