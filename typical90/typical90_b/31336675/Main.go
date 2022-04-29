package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	for i := int64(0); i < (1 << N); i++ {
		runes := make([]rune, N)
		for j := int64(0); j < N; j++ {
			if i&(1<<j) != 0 {
				runes[j] = ')'
			} else {
				runes[j] = '('
			}
		}

		Reverse(runes)
		if valid(runes) {
			fmt.Fprintln(w, string(runes))
		}
	}
}

func Reverse(s []rune) {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

func valid(s []rune) bool {
	count := 0
	for i, r := range s {
		if r == '(' {
			count++
			if count > len(s)-1-i {
				return false
			}
			continue
		}

		count--
		if count < 0 {
			return false
		}
	}

	return count == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N)
}
