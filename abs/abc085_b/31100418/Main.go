package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, d []int64) {
	m := make(map[int64]struct{}, N)
	for _, v := range d {
		m[v] = struct{}{}
	}
	fmt.Println(len(m))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	d := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		d[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, d)
}
