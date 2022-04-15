package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, A []int64) {
	visited := make(map[int64]int64, N)
	current := int64(1)
	visited[1] = 0
	for i := int64(1); i <= N; i++ {
		current = A[current-1]
		if i == K {
			fmt.Println(current)
			return
		}

		if prev, ok := visited[current]; ok {
			n := prev + (K-prev)%(i-prev)
			for k, v := range visited {
				if v == n {
					fmt.Println(k)
					return
				}
			}
		}

		visited[current] = i
	}

	fmt.Println("fail")
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
	var K int64
	scanner.Scan()
	K, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, K, A)
}
