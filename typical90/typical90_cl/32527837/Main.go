package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 998244353

func solve(N, K int) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	g := 3

	steps := make([]int, K+1)
	for n := 1; n <= K; n++ {
		steps[K/n]++
	}
	dp := make([]int, K+1)
	dp[0] = 1
	res := make([]int, K+1)
	for l := 1; l <= K; l++ {
		if steps[l] == 0 {
			continue
		}

		iMulti := PowMod(steps[l], mod-2, mod)
		for i := 0; i < Min(l, len(dp)); i++ {
			res[i+1] = dp[i] * steps[l] % mod
		}

		buf := powSumFPS(res, l+2, g)
		for i := 0; i <= l; i++ {
			dp[i] = buf[i+1] * iMulti % mod
		}
	}

	recurrentA := make([]int, len(dp))
	recurrentA[len(dp)-1] = 1
	result := kthTermOfLinearlyRecurrentSequence(recurrentA, dp, N+len(dp), g)
	fmt.Fprintln(w, result)
}

func main() {
	s := NewScanner()
	N := s.Int()
	K := s.Int()

	solve(N, K)
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), 1048576)
	s.Split(bufio.ScanWords)
	return &Scanner{
		Scanner: s,
	}
}

func (s *Scanner) Int() int {
	s.Scan()
	v, err := strconv.Atoi(s.Text())
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) IntN(size int) []int {
	v := make([]int, size)
	for i := 0; i < size; i++ {
		v[i] = s.Int()
	}
	return v
}

func (s *Scanner) IntN2(size int) ([]int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
	}
	return v1, v2
}

func (s *Scanner) IntN3(size int) ([]int, []int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	v3 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
	}
	return v1, v2, v3
}

func (s *Scanner) IntN4(size int) ([]int, []int, []int, []int) {
	v1 := make([]int, size)
	v2 := make([]int, size)
	v3 := make([]int, size)
	v4 := make([]int, size)
	for i := 0; i < size; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
		v4[i] = s.Int()
	}
	return v1, v2, v3, v4
}

func (s *Scanner) IntNN(h, w int) [][]int {
	v := make([][]int, h)
	for i := 0; i < h; i++ {
		v[i] = make([]int, w)
		for j := 0; j < w; j++ {
			v[i][j] = s.Int()
		}
	}
	return v
}

func (s *Scanner) Bytes() []byte {
	s.Scan()
	return s.Scanner.Bytes()
}

func (s *Scanner) BytesN(h int) [][]byte {
	v := make([][]byte, h)
	for i := 0; i < h; i++ {
		v[i] = s.Bytes()
	}
	return v
}

func (s *Scanner) Byte() byte {
	return s.Bytes()[0]
}

func (s *Scanner) ByteN(n int) []byte {
	v := make([]byte, n)
	for i := 0; i < n; i++ {
		v[i] = s.Byte()
	}
	return v
}

func (s Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}

func Min(v ...int) int {
	switch len(v) {
	case 0:
		panic("Min: len(v) == 0")
	case 1:
		return v[0]
	case 2:
		if v[0] < v[1] {
			return v[0]
		}
		return v[1]
	default:
		m := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] < m {
				m = v[i]
			}
		}
		return m
	}
}

func ceilPow2(n int) int {
	x := 0
	for (1 << x) < n {
		x++
	}
	return x
}

func SafeMod(x, m int) int {
	x %= m
	if x < 0 {
		return x + m
	}
	return x
}

func PowMod(x, n, m int) int {
	if m == 1 {
		return 0
	}

	r := 1
	for y := SafeMod(x, m); n > 0; n >>= 1 {
		if n&1 != 0 {
			r = (r * y) % m
		}
		y = (y * y) % m
	}
	return r
}

func NTT(a []int, g int) {
	n := len(a)

	for i, j := 0, 0; j < n; j++ {
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
		for k := n >> 1; ; k >>= 1 {
			i ^= k
			if k <= i {
				break
			}
		}
	}

	for i := 1; i < n; i <<= 1 {
		q := PowMod(g, (mod-1)/i/2, mod)
		for j, qj := 0, 1; j < i; j++ {
			for k := j; k < n; k += i * 2 {
				l, r := a[k], a[k+i]*qj%mod
				a[k] = l + r
				if a[k] >= mod {
					a[k] -= mod
				}
				a[k+i] = l + mod - r
				if a[k+i] >= mod {
					a[k+i] -= mod
				}
			}
			qj = qj * q % mod
		}
	}
}

func powSumFPS(a []int, n, g int) []int {
	if n <= 1 {
		return []int{1}[:n]
	}
	nn := 1 << ceilPow2(n)
	hn := nn / 2

	tgA := make([]int, nn)
	copy(tgA, a[:Min(nn, len(a))])
	NTT(tgA, g)

	iG := PowMod(g, mod-2, mod)
	hInv := powSumFPS(a, hn, g)
	htInv := make([]int, nn)
	copy(htInv, hInv[:hn])
	NTT(htInv, g)

	r := make([]int, nn)
	for i := 0; i < nn; i++ {
		r[i] = tgA[i] * htInv[i] % mod
	}
	NTT(r, iG)

	copy(r[:hn], r[hn:])
	for i := hn; i < nn; i++ {
		r[i] = 0
	}
	NTT(r, g)

	iNN := PowMod(nn*nn%mod, mod-2, mod)
	for i := 0; i < nn; i++ {
		r[i] = r[i] * htInv[i] % mod * iNN % mod
	}
	NTT(r, iG)

	res := make([]int, n)
	copy(res, hInv[:hn])
	copy(res[hn:], r[:hn])
	return res
}

func invFPS(a []int, n, g int) []int {
	xa := make([]int, Min(n, len(a)))
	ia0 := PowMod(a[0], mod-2, mod)
	for i := 0; i < len(xa); i++ {
		xa[i] = (mod - a[i]) * ia0 % mod
	}

	xa[0] = 0
	xa = powSumFPS(xa, n, g)
	for i := 0; i < len(xa); i++ {
		xa[i] = xa[i] * ia0 % mod
	}
	return xa
}

func bostanMori(a []int, g int) []int {
	n := len(a)
	aa := make([]int, n)
	copy(aa, a)
	iG := PowMod(g, mod-2, mod)
	NTT(aa, iG)

	invN := PowMod(n, mod-2, mod)
	for i := 0; i < n; i++ {
		aa[i] = aa[i] * invN % mod
	}

	b := make([]int, n)
	w := PowMod(g, (mod-1)/(2*n), mod)
	for i, wp := 0, 1; i < n; i++ {
		b[i] = aa[i] * wp % mod
		wp = wp * w % mod
	}
	NTT(b, g)

	res := make([]int, n*2)
	for i := 0; i < n; i++ {
		res[i*2] = a[i]
		res[i*2+1] = b[i]
	}
	return res
}

func kthTermOfLinearlyRecurrentSequence(a, c []int, k, g int) int {
	n := len(a)
	k2 := 1 << ceilPow2(2*n+1)

	p := make([]int, k2)
	copy(p, a)
	NTT(p, g)

	q := make([]int, k2)
	q[0] = 1
	for i := 0; i < n; i++ {
		q[i+1] = (mod - c[i]) % mod
	}
	NTT(q, g)

	inv2 := PowMod(2, mod-2, mod)
	invK := PowMod(k2, mod-2, mod)
	for i := 0; i < k2; i++ {
		p[i] = p[i] * q[i] % mod * invK % mod
	}

	iG := PowMod(g, mod-2, mod)
	NTT(p, iG)
	for i := n; i < k2; i++ {
		p[i] = 0
	}
	NTT(p, g)

	w := PowMod(g, (mod-1)/k2, mod)
	iW := PowMod(w, mod-2, mod)
	for h2k := k2 / 2; k >= n; k /= 2 {
		u := make([]int, k2)
		for i := 0; i < k2; i++ {
			u[i] = p[i] * q[i^h2k] % mod
		}

		if k%2 == 0 {
			uEven := make([]int, h2k)
			for i := 0; i < h2k; i++ {
				uEven[i] = (u[i] + u[i+h2k]) * inv2 % mod
			}
			p = bostanMori(uEven, g)
		} else {
			uOdd := make([]int, h2k)
			for i, wp := 0, inv2; i < h2k; i++ {
				uOdd[i] = (u[i] + mod - u[i+h2k]) * wp % mod
				wp = wp * iW % mod
			}
			p = bostanMori(uOdd, g)
		}

		qh := make([]int, h2k)
		for i := 0; i < h2k; i++ {
			qh[i] = q[i] * q[i^h2k] % mod
		}
		q = bostanMori(qh, g)
	}

	NTT(p, iG)
	NTT(q, iG)

	q = invFPS(q, k+1, g)

	res := 0
	for i := 0; i <= k; i++ {
		res = (res + p[i]*q[k-i]) % mod
	}
	return res
}
