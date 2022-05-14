package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Submission struct {
	ID    int64
	Score int64
}

func solve(N int64, S []string, T []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	m := make(map[string]Submission, N)
	for i := int64(0); i < N; i++ {
		sub := Submission{
			ID:    i + 1,
			Score: T[i],
		}
		if _, ok := m[S[i]]; ok {
			continue
		}
		m[S[i]] = sub
	}

	s := make([]Submission, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i].Score == s[j].Score {
			return s[i].ID < s[j].ID
		}
		return s[i].Score > s[j].Score
	})

	fmt.Fprintln(w, s[0].ID)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	S := make([]string, N)
	T := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		S[i] = scanner.Text()
		scanner.Scan()
		T[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, S, T)
}
