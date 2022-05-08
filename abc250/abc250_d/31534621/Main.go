package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	count := int64(0)
	primes := primesUnder(1000000)
	for i := 0; i < len(primes); i++ {
		if primes[i]*primes[i]*primes[i]*primes[i] >= N {
			break
		}
		for j := i + 1; j < len(primes); j++ {
			if primes[i]*primes[j]*primes[j]*primes[j] > N {
				break
			}
			count++
		}
	}
	fmt.Fprintln(w, count)
}

func primesUnder(n int64) []int64 {
	flags := make([]bool, n+1)
	flags[0] = true
	flags[1] = true
	for i := int64(2); i*i <= n; i++ {
		if flags[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			flags[j] = true
		}
	}

	results := make([]int64, 0, n)
	for i, f := range flags {
		if !f {
			results = append(results, int64(i))
		}
	}
	return results
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
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N)
}
