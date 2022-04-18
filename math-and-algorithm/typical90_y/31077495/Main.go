package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int64, B int64) {
	ints := DigitsIncreasingInts(100_000_000_000)
	fmSet := make(map[int64]struct{})
	for _, i := range ints {
		fmSet[DigitsProduct(i)] = struct{}{}
	}

	count := 0
	for fm := range fmSet {
		if m := fm + B; m-DigitsProduct(m) == B && m <= N {
			count++
		}
	}
	fmt.Println(count)
}

// DigitsIncreasingInts 0からmaxまでの整数のうち各位の数字が単調増加となるものを返す
// i.e. [0 1 2 ... 9 11 12 ... 19 22 23 ... 29 33 34 ...]
func DigitsIncreasingInts(max int64) []int64 {
	result := make([]int64, 1)
	result[0] = 0
	for {
		current := result[len(result)-1]
		for i := int64(1); i <= 10-current%10; i++ {
			if current+i > max {
				return result
			}
			result = append(result, current+i)
		}

		for d := int64(100_000_000_000); d >= 10; d /= 10 {
			if result[len(result)-1]%d != 0 {
				continue
			}
			result[len(result)-1] += ((result[len(result)-1] / d) % 10) * (d / 10)
		}

		if result[len(result)-1] > max {
			return result[:len(result)-1]
		}
	}
}

// DigitsProduct nの各位の数字の積を返す
func DigitsProduct(n int64) int64 {
	if n == 0 {
		return 0
	}
	result := int64(1)
	for n >= 1 {
		result *= n % 10
		n /= 10
	}
	return result
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
	var B int64
	scanner.Scan()
	B, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(N, B)
}
