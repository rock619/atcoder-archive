package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 2

func solve(N int64, A []int64) {
	min := int64(1 << 62)
	for _, a := range A {
		count := int64(0)
		for ; a%MOD == 0; a /= MOD {
			count++
		}
		if count < min {
			min = count
		}
	}
	fmt.Println(min)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
