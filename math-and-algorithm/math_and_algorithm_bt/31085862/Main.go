package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(L int64, R int64) {
	nums := make([]bool, R-L+1)
	if L == 1 {
		nums[0] = true
	}
	for i := int64(2); i*i <= R; i++ {
		if index := i - L; index >= 0 && nums[index] {
			continue
		}
		for j := i * Max(L/i, 2); j <= R; j += i {
			if index := j - L; index >= 0 {
				nums[index] = true
			}
		}
	}

	count := 0
	for _, v := range nums {
		if !v {
			count++
		}
	}
	fmt.Println(count)
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var L int64
	scanner.Scan()
	L, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var R int64
	scanner.Scan()
	R, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(L, R)
}
