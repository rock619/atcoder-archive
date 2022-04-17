package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Movie struct {
	L, R int64
}

func solve(N int64, L []int64, R []int64) {
	movies := make([]Movie, N)
	for i := int64(0); i < N; i++ {
		movies[i] = Movie{L: L[i], R: R[i]}
	}
	sort.Slice(movies, func(i, j int) bool {
		return movies[i].R < movies[j].R
	})

	count := 0
	current := int64(0)
	for _, m := range movies {
		if m.L >= current {
			current = m.R
			count++
		}
	}
	fmt.Println(count)
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
	L := make([]int64, N)
	R := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		L[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		R[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, L, R)
}
