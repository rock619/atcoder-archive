package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 998244353

type ModCombi struct {
	fact    []int64
	inv     []int64
	factInv []int64
	mod     int64
}

func NewModCombi(maxN, mod int64) *ModCombi {
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
	return &ModCombi{
		fact:    fact,
		inv:     inv,
		factInv: factInv,
		mod:     mod,
	}
}

func (mc *ModCombi) Do(n, k int64) int64 {
	return mc.fact[n] * (mc.factInv[k] * mc.factInv[n-k] % mc.mod) % mc.mod
}

func solve(N, M int64, T []int64, A [][]int64, S []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	s := make([]bool, M)
	for i := int64(0); i < M; i++ {
		s[i] = S[i] == 1
	}

	buttons := make([][]bool, N)
	for i := int64(0); i < N; i++ {
		buttons[i] = make([]bool, M)
		for _, v := range A[i] {
			buttons[i][v-1] = true
		}
	}

	pos := int64(0)
	for i := int64(0); i < M; i++ {
		found := false
		for j := pos; j < N; j++ {
			if buttons[j][i] {
				if j != pos {
					buttons[j], buttons[pos] = buttons[pos], buttons[j]
				}
				found = true
				break
			}
		}

		if found {
			for j := int64(0); j < N; j++ {
				if j != pos && buttons[j][i] {
					buttons[j] = xor(buttons[pos], buttons[j])
				}
			}
			if s[i] {
				s = xor(s, buttons[pos])
			}

			pos++
		}
	}

	for _, v := range s {
		if v {
			fmt.Fprintln(w, 0)
			return
		}
	}

	result := int64(1)
	for i := pos; i < N; i++ {
		result *= 2
		result %= mod
	}

	fmt.Fprintln(w, result)
}

func xor(a, b []bool) []bool {
	result := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] != b[i]
	}
	return result
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
	M, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	T := make([]int64, N)
	A := make([][]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		T[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		A[i] = make([]int64, T[i])
		for j := int64(0); j < T[i]; j++ {
			scanner.Scan()
			A[i][j], err = strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
		}
	}
	S := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		S[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, M, T, A, S)
}
