package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	T := MustScanInt(scanner)
	for i := int64(0); i < T; i++ {
		N := MustScanInt(scanner)
		if N == -1 {
			return
		}
		m := make(map[int64]int64)
		if N == 1 {
			ans := getOrQuery(scanner, m, 0)
			fmt.Printf("! %d\n", ans)
			continue
		}

		left, right := int64(0), N-1
		for right-left >= 3 {
			centerLeft := left + (right-left)/3
			centerRight := right - (right-left)/3
			clv := getOrQuery(scanner, m, centerLeft)
			crv := getOrQuery(scanner, m, centerRight)
			if clv > crv {
				right = centerRight
			} else {
				left = centerLeft
			}
		}

		for j := left; j <= right; j++ {
			current := getOrQuery(scanner, m, j)
			if j == right {
				fmt.Printf("! %d\n", current)
				break
			}
			next := getOrQuery(scanner, m, j+1)
			if current > next {
				fmt.Printf("! %d\n", current)
				break
			}
		}
	}
}

func getOrQuery(scanner *bufio.Scanner, m map[int64]int64, i int64) int64 {
	v, ok := m[i]
	if ok {
		return v
	}
	fmt.Printf("? %d\n", i+1)
	v = MustScanInt(scanner)
	m[i] = v
	return v
}

func MustScanInt(scanner *bufio.Scanner) int64 {
	scanner.Scan()
	i, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
