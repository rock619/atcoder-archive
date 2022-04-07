package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, A []int64) {
	lenA := int64(len(A))
	count := 0
	for i1 := int64(0); i1 < lenA; i1++ {
		for i2 := i1 + 1; i2 < lenA; i2++ {
			for i3 := i2 + 1; i3 < lenA; i3++ {
				for i4 := i3 + 1; i4 < lenA; i4++ {
					for i5 := i4 + 1; i5 < lenA; i5++ {
						if A[i1]+A[i2]+A[i3]+A[i4]+A[i5] == 1000 {
							count++
						}
					}
				}
			}
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
	A := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(N, A)
}
