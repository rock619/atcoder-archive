package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, K int64, S string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	indexes := make([][]int64, N+1)
	for i := range indexes {
		indexes[i] = make([]int64, 26)
		for j := range indexes[i] {
			indexes[i][j] = -1
		}
	}
	for i := N - 1; i >= 0; i-- {
		for j := int64(0); j < 26; j++ {
			if int64(S[i]-'a') == j {
				indexes[i][j] = i
			} else {
				indexes[i][j] = indexes[i+1][j]
			}
		}
	}

	result := make([]rune, 0, K)
	current := int64(0)
	for i := int64(0); i < K; i++ {
		maxIndex := N - K + i
		for j := int64(0); j < 26; j++ {
			if indexes[current][j] == -1 || indexes[current][j] > maxIndex {
				continue
			}

			current = indexes[current][j] + 1
			result = append(result, rune('a'+j))
			break
		}
	}
	fmt.Fprintln(w, string(result))
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
	scanner.Scan()
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	S := scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, K, S)
}
