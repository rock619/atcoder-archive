package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const (
	yes = "Yes"
	no  = "No"
)

type ZobristHash struct {
	fixedRandom uint64
}

func NewZobristHash() *ZobristHash {
	return &ZobristHash{
		fixedRandom: rand.Uint64(),
	}
}

func (h *ZobristHash) Of(x uint64) uint64 {
	return h.splitmix64(x + h.fixedRandom)
}

func (_ *ZobristHash) splitmix64(x uint64) uint64 {
	x += 0x9e3779b97f4a7c15
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	return x ^ (x >> 31)
}

func hashes(h *ZobristHash, ints []int64) []uint64 {
	m := make(map[int64]struct{}, len(ints))
	result := make([]uint64, len(ints))
	result[0] = h.Of(uint64(ints[0]))
	m[ints[0]] = struct{}{}
	for i := 1; i < len(ints); i++ {
		if _, ok := m[ints[i]]; ok {
			result[i] = result[i-1]
			continue
		}
		result[i] = result[i-1] ^ h.Of(uint64(ints[i]))
		m[ints[i]] = struct{}{}
	}
	return result
}

func solve(N int64, a []int64, b []int64, Q int64, x []int64, y []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	h := NewZobristHash()
	hashesA, hashesB := hashes(h, a), hashes(h, b)

	for i := int64(0); i < Q; i++ {
		if hashesA[x[i]-1] == hashesB[y[i]-1] {
			fmt.Fprintln(w, yes)
		} else {
			fmt.Fprintln(w, no)
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
	a := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	b := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		b[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	scanner.Scan()
	Q, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	x := make([]int64, Q)
	y := make([]int64, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		x[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		y[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, a, b, Q, x, y)
}
