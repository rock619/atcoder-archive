package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(N int64, a []int64) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	diff := int64(0)
	for i, n := range a {
		if i%2 == 0 {
			diff += n
		} else {
			diff -= n
		}
	}
	fmt.Println(diff)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	a := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, a)
}
