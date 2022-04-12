package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Deque struct {
	s []int64
}

func NewDeque(cap int64) *Deque {
	return &Deque{
		s: make([]int64, 0, cap),
	}
}

func (d Deque) Size() int64 {
	return int64(len(d.s))
}

func (d *Deque) PopFront() int64 {
	if d.Size() == 0 {
		panic("Deque.PopFront(): deque is empty")
	}
	e := d.s[0]
	d.s = d.s[1:]
	return e
}

func (d *Deque) PushBack(e int64) {
	d.s = append(d.s, e)
}

func (d *Deque) PushFront(e int64) {
	d.s = append([]int64{e}, d.s...)
}

func solve(K int64) {
	dists := make([]int64, K)
	for i := 0; i < len(dists); i++ {
		dists[i] = 1 << 62
	}
	dists[1] = 1

	deque := NewDeque(10)
	deque.PushFront(1)

	for deque.Size() > 0 {
		v := deque.PopFront()

		if v2 := (v * 10) % K; dists[v2] > dists[v] {
			dists[v2] = dists[v]
			deque.PushFront(v2)
		}

		if v2 := (v + 1) % K; dists[v2] > dists[v]+1 {
			dists[v2] = dists[v] + 1
			deque.PushBack(v2)
		}
	}

	fmt.Println(dists[0])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(K)
}
