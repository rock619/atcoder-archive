package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

type ModCombination struct {
	factorials []int64
}

func NewModCombination(n int64) *ModCombination {
	f := make([]int64, 1, n)
	f[0] = 1
	return &ModCombination{
		factorials: f,
	}
}

func (c *ModCombination) Do(n, r, m int64) int64 {
	return ModDiv(c.Factorial(n, m), c.Factorial(r, m)*c.Factorial(n-r, m)%m, m)
}

func (c *ModCombination) Factorial(n, m int64) int64 {
	for i := int64(len(c.factorials)); i <= n; i++ {
		c.factorials = append(c.factorials, c.factorials[i-1]*i%m)
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
	for i := 0; i < 62; i++ {
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
	defer w.Flush()

	modCombi := NewModCombination(N)
	for k := int64(1); k <= N; k++ {
		sum := int64(0)
		for i := int64(1); i*k <= N+k-1; i++ {
			sum += modCombi.Do(N-(k-1)*(i-1), i, MOD)
			sum %= MOD
		}
		fmt.Fprintln(w, sum)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N)
}
