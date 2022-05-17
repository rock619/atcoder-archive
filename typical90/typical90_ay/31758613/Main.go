package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, K int64, P int64, A []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	g1, g2 := A[:N/2], A[N/2:]
	g1Sums := sums(g1, K, P)
	for i := range g1Sums {
		sort.Slice(g1Sums[i], func(i2, j int) bool {
			return g1Sums[i][i2] < g1Sums[i][j]
		})
	}
	g2l := int64(len(g2))
	result := int64(0)
	for i := int64(0); i < (1 << g2l); i++ {
		count := int64(bits.OnesCount64(uint64(i)))
		if count > K {
			continue
		}
		sum := int64(0)
		for j := int64(0); j < g2l; j++ {
			if i&(1<<j) != 0 {
				sum += g2[j]
			}
		}
		if sum > P {
			continue
		}

		index := sort.Search(len(g1Sums[K-count]), func(j int) bool {
			return g1Sums[K-count][j]+sum > P
		})
		// fmt.Fprintln(w, i, sum, index)
		result += int64(index)
	}
	// fmt.Fprintln(w, g1)
	// fmt.Fprintln(w, g2)
	// fmt.Fprintln(w, g1Sums)
	fmt.Fprintln(w, result)
}

func sums(a []int64, K, P int64) [][]int64 {
	result := make([][]int64, K+1)
	al := int64(len(a))
	for i := int64(0); i < (1 << al); i++ {
		count := int64(bits.OnesCount64(uint64(i)))
		if count > K {
			continue
		}
		sum := int64(0)
		for j := int64(0); j < al; j++ {
			if i&(1<<j) != 0 {
				sum += a[j]
			}
		}
		if sum > P {
			continue
		}
		result[count] = append(result[count], sum)
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
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	P, err := strconv.ParseInt(scanner.Text(), 10, 64)
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
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K, P, A)
}
