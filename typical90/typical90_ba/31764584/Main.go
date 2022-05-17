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

		max := int64(0)
		left, right := int64(0), N-1
		for j := 0; j < 11; j++ {
			center := (left + right) / 2
			d1 := getOrQuery(scanner, m, center)
			if d1 > max {
				max = d1
			}
			d2 := getOrQuery(scanner, m, center+1)
			if d2 > max {
				max = d2
			}

			if d1 > d2 {
				right = center
			} else {
				left = center
			}
		}
		fmt.Printf("! %d\n", max)
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
