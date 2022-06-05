package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	fmt.Fprintln(w, 1)
	mc := NewModCombination(30, 1000000007)
	for i := 1; i < N; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				fmt.Fprint(w, mc.Do(int64(i), int64(j)))
			} else {
				fmt.Fprintf(w, " %d", mc.Do(int64(i), int64(j)))
			}
		}
		fmt.Fprintln(w)
	}
}

type ModCombination struct {
	fact    []int64
	inv     []int64
	factInv []int64
	mod     int64
}

func NewModCombination(maxN, mod int64) *ModCombination {
	fact := make([]int64, maxN+1)
	fact[0], fact[1] = 1, 1

	inv := make([]int64, maxN+1)
	inv[1] = 1

	factInv := make([]int64, maxN+1)
	factInv[0], factInv[1] = 1, 1

	for i := int64(2); i <= maxN; i++ {
		fact[i] = fact[i-1] * i % mod
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
		factInv[i] = factInv[i-1] * inv[i] % mod
	}

	return &ModCombination{
		fact:    fact,
		inv:     inv,
		factInv: factInv,
		mod:     mod,
	}
}

func (mc *ModCombination) Do(n, k int64) int64 {
	return mc.fact[n] * (mc.factInv[k] * mc.factInv[n-k] % mc.mod) % mc.mod
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 4096), 1048576)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N)
}
