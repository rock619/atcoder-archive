package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 1000000007

type ModInt struct {
	V   int64
	Mod int64
}

func NewModInt(v, mod int64) ModInt {
	x := v % mod
	if x < 0 {
		x += mod
	}
	return ModInt{
		V:   x,
		Mod: mod,
	}
}

func (mi ModInt) Add(x ModInt) ModInt {
	return NewModInt(mi.V+x.V, mi.Mod)
}

func (mi ModInt) AddInt(v int64) ModInt {
	return mi.Add(NewModInt(v, mi.Mod))
}

func (mi ModInt) Sub(x ModInt) ModInt {
	return NewModInt(mi.V-x.V, mi.Mod)
}

func (mi ModInt) SubInt(v int64) ModInt {
	return mi.Sub(NewModInt(v, mi.Mod))
}

func (mi ModInt) Mul(x ModInt) ModInt {
	return NewModInt(mi.V*x.V, mi.Mod)
}

func (mi ModInt) MulInt(v int64) ModInt {
	return mi.Mul(NewModInt(v, mi.Mod))
}

func (mi ModInt) Div(x ModInt) ModInt {
	return mi.Mul(x.Inv())
}

func (mi ModInt) Pow(n int64) ModInt {
	x := mi
	r := NewModInt(1, mi.Mod)
	for n > 0 {
		if (n & 1) != 0 {
			r = r.Mul(x)
		}
		x = x.Mul(x)
		n >>= 1
	}
	return r
}

func (mi ModInt) Inv() ModInt {
	_, x, _ := ExtendedGCD(mi.V, mi.Mod)
	return NewModInt(x, mi.Mod)
}

func ExtendedGCD(a, b int64) (gcd, x, y int64) {
	x, y = 1, 0
	x1, y1, a1, b1 := y, x, a, b
	for b1 != 0 {
		q := a1 / b1
		x, x1 = x1, x-q*x1
		y, y1 = y1, y-q*y1
		a1, b1 = b1, a1-q*b1
	}
	return a1, x, y
}

func solve(N int64) {
	ints := make([]ModInt, N)
	ints[0] = NewModInt(1, MOD)
	ints[1] = NewModInt(1, MOD)
	for i := int64(2); i < N; i++ {
		ints[i] = ints[i-1].Add(ints[i-2])
	}

	fmt.Println(ints[N-1].V)
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
	solve(N)
}
