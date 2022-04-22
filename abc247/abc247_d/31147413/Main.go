package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Query struct {
	Type int64
	X    int64
	C    int64
}

func solve(Q int64, queries []Query) {
	cylinder := make([]Query, 0, Q)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	index := 0
	for _, q := range queries {
		if q.Type == 1 {
			cylinder = append(cylinder, q)
			continue
		}

		rest := q.C
		sum := int64(0)
		for {
			if cylinder[index].C >= rest {
				cylinder[index].C -= rest
				sum += cylinder[index].X * rest
				if cylinder[index].C == 0 {
					index++
				}
				break
			}

			rest -= cylinder[index].C
			sum += cylinder[index].X * cylinder[index].C
			index++
		}
		fmt.Fprintln(w, sum)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	Q, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	queries := make([]Query, Q)
	for i := int64(0); i < Q; i++ {
		scanner.Scan()
		queries[i].Type, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		if queries[i].Type == 2 {
			queries[i].C, _ = strconv.ParseInt(scanner.Text(), 10, 64)
			continue
		}
		queries[i].X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		scanner.Scan()
		queries[i].C, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	solve(Q, queries)
}
