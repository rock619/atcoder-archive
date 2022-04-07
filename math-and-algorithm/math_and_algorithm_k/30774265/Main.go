package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(N int64) {
	primes := make([]int64, 0)
	for n := int64(2); n <= N; n++ {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}
	print(primes)
}

func isPrime(N int64) bool {
	for i := int64(2); i*i <= N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}

func print(primes []int64) {
	parts := make([]string, len(primes))
	for i, v := range primes {
		parts[i] = strconv.FormatInt(v, 10)
	}
	fmt.Println(strings.Join(parts, " "))
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
