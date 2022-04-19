package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	YES = "YES"
	NO  = "NO"
)

func solve(S string) {
	reversed := reverse(S)
	parts := []string{"dream", "dreamer", "erase", "eraser"}
	reversedParts := make([][]byte, len(parts))
	for i, p := range parts {
		reversedParts[i] = reverse(p)
	}

	for len(reversed) > 0 {
		l := len(reversed)
		for _, p := range reversedParts {
			reversed = bytes.TrimPrefix(reversed, p)
		}
		if len(reversed) == l {
			fmt.Println(NO)
			return
		}
	}
	fmt.Println(YES)
}

func reverse(s string) []byte {
	b := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b = append(b, s[i])
	}
	return b
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	S := scanner.Text()
	solve(S)
}
