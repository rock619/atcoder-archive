package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 699999953

func solve(N int64, S string, T string) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	seq1 := make([]int64, N)
	seq3 := make([]int64, N)
	for i := int64(0); i < N; i++ {
		switch S[i] {
		case 'R':
			seq1[i] = 0
		case 'G':
			seq1[i] = 1
		default:
			seq1[i] = 2
		}

		switch T[i] {
		case 'R':
			seq3[i] = 0
		case 'G':
			seq3[i] = 1
		default:
			seq3[i] = 2
		}
	}

	k := 0
	for i := int64(0); i < 3; i++ {
		seq2 := make([]int64, N)
		for j := int64(0); j < N; j++ {
			seq2[j] = (i - seq3[j] + 3) % 3
		}
		power3 := int64(1)
		hash1 := int64(0)
		hash2 := int64(0)
		for j := int64(0); j < N; j++ {
			hash1 = (hash1*3 + seq1[j]) % mod
			hash2 = (hash2 + power3*seq2[N-j-1]) % mod
			if hash1 == hash2 {
				k++
			}
			power3 = power3 * 3 % mod
		}
		power3 = 1
		hash1 = 0
		hash2 = 0
		for j := int64(0); j < N-1; j++ {
			hash1 = (hash1 + power3*seq1[N-j-1]) % mod
			hash2 = (hash2*3 + seq2[j]) % mod
			if hash1 == hash2 {
				k++
			}
			power3 = power3 * 3 % mod
		}
	}
	fmt.Fprintln(w, k)
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
	S := scanner.Text()
	scanner.Scan()
	T := scanner.Text()
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, S, T)
}
