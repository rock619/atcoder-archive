package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

type ModInt64 struct {
	i   int64
	mod int64
}

func NewModInt64(i, mod int64) *ModInt64 {
	zero := &ModInt64{
		i:   0,
		mod: mod,
	}
	return zero.Add(i)
}

func (mi *ModInt64) Int64() int64 {
	return mi.i
}

func (mi *ModInt64) Add(n int64) *ModInt64 {
	mi.i += n % mi.mod
	switch {
	case mi.i >= mi.mod:
		mi.i -= mi.mod
	case mi.i < 0:
		mi.i += mi.mod
	}
	return mi
}

func (mi *ModInt64) Sub(n int64) *ModInt64 {
	return mi.Add(-n)
}

type ModCombi struct {
	factorials []int64
	mod        int64
}

func NewModCombi(mod, maxN int64) *ModCombi {
	f := make([]int64, 1, maxN)
	f[0] = 1
	return &ModCombi{
		factorials: f,
		mod:        mod,
	}
}

func (c *ModCombi) Do(n, k int64) int64 {
	return ModDiv(c.factorial(n), c.factorial(k)*c.factorial(n-k)%c.mod, c.mod)
}

func (c *ModCombi) factorial(n int64) int64 {
	for i := int64(len(c.factorials)); i <= n; i++ {
		c.factorials = append(c.factorials, c.factorials[i-1]*i%c.mod)
	}

	return c.factorials[n]
}

// ModDiv a/b mod m
func ModDiv(a, b, m int64) int64 {
	return a * ModPow(b, m-2, m) % m
}

// ModPow a**b mod m
func ModPow(a, b, m int64) int64 {
	p := a
	result := int64(1)
	for i := 0; i < 63; i++ {
		if b&(1<<i) != 0 {
			result *= p
			result %= m
		}
		p *= p
		p %= m
	}

	return result
}

func solve(N int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	modCombi := NewModCombi(mod, N)
	for k := int64(1); k <= N; k++ {
		sum := NewModInt64(0, mod)
		for i := int64(1); i*k <= N+k-1; i++ {
			sum.Add(modCombi.Do(N-(k-1)*(i-1), i))
		}
		fmt.Fprintln(w, sum.Int64())
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
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N)
}
