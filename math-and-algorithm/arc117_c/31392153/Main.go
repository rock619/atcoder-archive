package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Combi struct {
	combi [][]int64
	mod   int64
}

func NewCombi(mod int64) *Combi {
	combi := make([][]int64, mod)
	for i := range combi {
		combi[i] = make([]int64, mod)
	}
	combi[0][0] = 1
	for i := int64(1); i < mod; i++ {
		combi[i][0] = 1
		for j := i; j > 0; j-- {
			combi[i][j] = (combi[i-1][j-1] + combi[i-1][j]) % mod
		}
	}
	return &Combi{
		combi,
		mod,
	}
}

func (c *Combi) Do(n, k int64) int64 {
	result := int64(1)
	for n > 0 {
		result *= c.combi[n%c.mod][k%c.mod]
		result %= c.mod
		n /= c.mod
		k /= c.mod
	}
	return result
}

func colorToInt64(color byte) int64 {
	switch color {
	case 'B':
		return 0
	case 'W':
		return 1
	case 'R':
		return 2
	default:
		panic(fmt.Sprintf("colorToInt64: %q", color))
	}
}

func int64ToColor(i int64) byte {
	switch i {
	case 0:
		return 'B'
	case 1:
		return 'W'
	case 2:
		return 'R'
	default:
		panic(fmt.Sprintf("int64ToColor: %d", i))
	}
}

func solve(N int64, c []byte) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	combi := NewCombi(3)
	result := int64(0)
	for i := int64(0); i < N; i++ {
		n := combi.Do(N-1, i) * colorToInt64(c[i])
		result += n % 3
		result %= 3
	}
	if N%2 == 0 && result > 0 {
		result = 3 - result
	}
	fmt.Fprintln(w, string(int64ToColor(result)))
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
	scanner.Scan()
	c := scanner.Bytes()
	solve(N, c)
}
