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

	fib := []int64{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597}
	T := MustScanInt(scanner)
	for i := int64(0); i < T; i++ {
		N := MustScanInt(scanner)
		if N == -1 {
			return
		}
		m := make(map[int64]int64)
		if N == 1 {
			ans := getOrQuery(scanner, m, N, 0)
			fmt.Printf("! %d\n", ans)
			continue
		}

		max := int64(0)
		left := int64(0)
		right := fib[len(fib)-1]
		for j := len(fib) - 3; j >= 0; j-- {
			leftCenter := left + fib[j] - 1
			rightCenter := right - fib[j] - 1

			leftV := getOrQuery(scanner, m, N, leftCenter)
			if leftV > max {
				max = leftV
			}
			rightV := getOrQuery(scanner, m, N, rightCenter)
			if rightV > max {
				max = rightV
			}

			if leftV > rightV {
				right = rightCenter + 1
			} else {
				left = leftCenter + 1
			}
		}
		fmt.Printf("! %d\n", max)
	}
}

func getOrQuery(scanner *bufio.Scanner, m map[int64]int64, N, i int64) int64 {
	if i >= N {
		return -i
	}

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
