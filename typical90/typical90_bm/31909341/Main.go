package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

const mod = 998244353

type ModCombination struct {
	fact    []int64
	inv     []int64
	factInv []int64
	mod     int64
}

func NewModCombination(maxN, mod int64) *ModCombination {
	fact := make([]int64, maxN+1)
	fact[0], fact[1] = 1, 1

	inv := make([]int64, maxN+1)
	inv[1] = 1

	factInv := make([]int64, maxN+1)
	factInv[0], factInv[1] = 1, 1

	for i := int64(2); i <= maxN; i++ {
		fact[i] = fact[i-1] * i % mod
		inv[i] = mod - inv[mod%i]*(mod/i)%mod
		factInv[i] = factInv[i-1] * inv[i] % mod
	}

	return &ModCombination{
		fact:    fact,
		inv:     inv,
		factInv: factInv,
		mod:     mod,
	}
}

func (mc *ModCombination) Do(n, k int64) int64 {
	return mc.fact[n] * (mc.factInv[k] * mc.factInv[n-k] % mc.mod) % mc.mod
}

func solve(R int64, G int64, B int64, K int64, X int64, Y int64, Z int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	mc := NewModCombination(Max(R, G, B), mod)

	rMax := Min(R, K)
	a := make([]int64, rMax+1)
	for i := K - Y; i <= rMax; i++ {
		a[i] = mc.Do(R, i)
	}

	gMax := Min(G, K)
	b := make([]int64, gMax+1)
	for i := K - Z; i <= gMax; i++ {
		b[i] = mc.Do(G, i)
	}

	fft := NewFFT(mod)
	d := fft.Convolution(a, b)

	sum := int64(0)
	bMax := Min(B, K)
	for i := K - X; i <= bMax; i++ {
		sum += d[K-i] * mc.Do(B, i) % mod
		sum %= mod
	}
	fmt.Fprintln(w, sum)
}

func Max(v ...int64) int64 {
	switch len(v) {
	case 0:
		panic("Max: len(v) == 0")
	case 1:
		return v[0]
	case 2:
		if v[0] > v[1] {
			return v[0]
		}
		return v[1]
	default:
		m := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] > m {
				m = v[i]
			}
		}
		return m
	}
}

func Min(v ...int64) int64 {
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

type FFT struct {
	mod           int64
	g             int64
	rank2         int64
	root, iRoot   []int64
	rate2, iRate2 []int64
	rate3, iRate3 []int64
}

func NewFFT(mod int64) *FFT {
	f := &FFT{
		mod:   mod,
		g:     primitiveRoot(mod),
		rank2: int64(bits.TrailingZeros64(uint64(mod) - 1)),
	}
	f.initRoot()
	f.initRate2()
	f.initRate3()
	return f
}

func (f *FFT) initRoot() {
	f.root = make([]int64, f.rank2+1)
	f.iRoot = make([]int64, f.rank2+1)

	f.root[f.rank2] = PowMod(f.g, (f.mod-1)>>f.rank2, f.mod)
	f.iRoot[f.rank2] = InvMod(f.root[f.rank2], f.mod)

	for i := f.rank2 - 1; i >= 0; i-- {
		f.root[i] = f.root[i+1] * f.root[i+1] % f.mod
		f.iRoot[i] = f.iRoot[i+1] * f.iRoot[i+1] % f.mod
	}
}

func (f *FFT) initRate2() {
	if l := f.rank2 - 2 + 1; l > 0 {
		f.rate2 = make([]int64, l)
		f.iRate2 = make([]int64, l)
	}

	prod := int64(1)
	iProd := int64(1)
	for i := int64(0); i <= f.rank2-2; i++ {
		f.rate2[i] = f.root[i+2] * prod % f.mod
		f.iRate2[i] = f.iRoot[i+2] * iProd % f.mod
		prod = prod * f.iRoot[i+2] % f.mod
		iProd = iProd * f.root[i+2] % f.mod
	}
}

func (f *FFT) initRate3() {
	if l := f.rank2 - 3 + 1; l > 0 {
		f.rate3 = make([]int64, l)
		f.iRate3 = make([]int64, l)
	}

	prod := int64(1)
	iProd := int64(1)
	for i := int64(0); i <= f.rank2-3; i++ {
		f.rate3[i] = f.root[i+3] * prod % f.mod
		f.iRate3[i] = f.iRoot[i+3] * iProd % f.mod
		prod = prod * f.iRoot[i+3] % f.mod
		iProd = iProd * f.root[i+3] % f.mod
	}
}

func (f *FFT) Convolution(a, b []int64) []int64 {
	n := int64(len(a))
	m := int64(len(b))
	z := int64(1) << ceilPow2(n+m-1)
	a = append(a, make([]int64, z-n)...)
	b = append(b, make([]int64, z-m)...)

	f.butterfly(a)
	f.butterfly(b)
	for i := int64(0); i < z; i++ {
		a[i] = a[i] * b[i] % f.mod
	}
	f.butterflyInv(a)

	a = a[:n+m-1]
	iz := InvMod(z, f.mod)
	for i := int64(0); i < n+m-1; i++ {
		a[i] = a[i] * iz % f.mod
	}
	return a
}

func (f *FFT) butterfly(a []int64) {
	n := int64(len(a))
	h := ceilPow2(n)

	for length := int64(0); length < h; length++ {
		if h-length == 1 {
			p := int64(1) << (h - length - 1)
			rot := int64(1)
			for s := int64(0); s < (1 << length); s++ {
				offset := s << (h - length)
				for i := int64(0); i < p; i++ {
					l := a[i+offset]
					r := a[i+offset+p] * rot % f.mod
					a[i+offset] = (l + r) % f.mod
					a[i+offset+p] = (l + f.mod - r) % f.mod
				}
				if s+1 != (1 << length) {
					rot = rot * f.rate2[bits.TrailingZeros64(uint64(^s))] % f.mod
				}
			}
			continue
		}

		p := int64(1) << (h - length - 2)
		rot := int64(1)
		imag := f.root[2]
		for s := int64(0); s < (1 << length); s++ {
			rot2 := rot * rot % f.mod
			rot3 := rot2 * rot % f.mod
			offset := s << (h - length)
			for i := int64(0); i < p; i++ {
				mod2 := f.mod * f.mod
				a0 := a[i+offset]
				a1 := a[i+offset+p] * rot % f.mod
				a2 := a[i+offset+2*p] * rot2 % f.mod
				a3 := a[i+offset+3*p] * rot3 % f.mod

				a1na3imag := (a1 + mod2 - a3) % f.mod * imag % f.mod
				na2 := mod2 - a2

				a[i+offset] = (a0 + a2 + a1 + a3) % f.mod
				a[i+offset+1*p] = (a0 + a2 + (2*mod2 - (a1 + a3))) % f.mod
				a[i+offset+2*p] = (a0 + na2 + a1na3imag) % f.mod
				a[i+offset+3*p] = (a0 + na2 + (mod2 - a1na3imag)) % f.mod
			}
			if s+1 != (1 << length) {
				rot = rot * f.rate3[bits.TrailingZeros64(uint64(^s))] % f.mod
			}
		}
		length++
	}
}

func (f *FFT) butterflyInv(a []int64) {
	n := int64(len(a))
	h := ceilPow2(n)

	for length := h; length > 0; length-- {
		if length == 1 {
			p := int64(1) << (h - length)
			iRot := int64(1)
			for s := int64(0); s < (1 << (length - 1)); s++ {
				offset := s << (h - length + 1)
				for i := int64(0); i < p; i++ {
					l := a[i+offset]
					r := a[i+offset+p]
					a[i+offset] = (l + r) % f.mod
					a[i+offset+p] = (f.mod + l - r) % f.mod * iRot % f.mod
				}
				if s+1 != (1 << (length - 1)) {
					iRot = iRot * f.iRate2[bits.TrailingZeros64(uint64(^s))] % f.mod
				}
			}
			continue
		}

		p := int64(1) << (h - length)
		iRot := int64(1)
		iImag := f.iRoot[2]
		for s := int64(0); s < (1 << (length - 2)); s++ {
			iRot2 := iRot * iRot % f.mod
			iRot3 := iRot2 * iRot % f.mod
			offset := s << (h - length + 2)
			for i := int64(0); i < p; i++ {
				a0 := a[i+offset+0*p]
				a1 := a[i+offset+1*p]
				a2 := a[i+offset+2*p]
				a3 := a[i+offset+3*p]

				a2na3iImag := (f.mod + a2 - a3) % f.mod * iImag % f.mod

				a[i+offset] = (a0 + a1 + a2 + a3) % f.mod
				a[i+offset+1*p] = (a0 + (f.mod - a1) + a2na3iImag) % f.mod * iRot % f.mod
				a[i+offset+2*p] = (a0 + a1 + (f.mod - a2) + (f.mod - a3)) % f.mod * iRot2 % f.mod
				a[i+offset+3*p] = (a0 + (f.mod - a1) + (f.mod - a2na3iImag)) % f.mod * iRot3 % f.mod
			}
			if s+1 != (1 << (length - 2)) {
				iRot = iRot * f.iRate3[bits.TrailingZeros64(uint64(^s))] % f.mod
			}

		}
		length--
	}
}

// ceilPow2 returns minimum non-negative x s.t. n <= 2**x
func ceilPow2(n int64) int64 {
	x := int64(0)
	for (1 << x) < n {
		x++
	}
	return x
}

func primitiveRoot(m int64) int64 {
	switch m {
	case 2:
		return 1
	case 167772161, 469762049, 998244353:
		return 3
	case 754974721:
		return 11
	}

	divs := make([]int64, 20)
	divs[0] = 2
	cnt := 1
	x := (m - 1) / 2
	for x%2 == 0 {
		x /= 2
	}
	for i := int64(3); i*i <= x; i += 2 {
		if x%i == 0 {
			divs[cnt] = i
			cnt++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		divs[cnt] = x
		cnt++
	}
	for g := int64(2); ; g++ {
		ok := true
		for i := 0; i < cnt; i++ {
			if PowMod(g, (m-1)/divs[i], m) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
	}
}

func SafeMod(x, m int64) int64 {
	x %= m
	if x < 0 {
		return x + m
	}
	return x
}

// PowMod (x ** n) % m
func PowMod(x, n, m int64) int64 {
	if m == 1 {
		return 0
	}

	r := int64(1)
	for y := SafeMod(x, m); n > 0; n >>= 1 {
		if n&1 != 0 {
			r = (r * y) % m
		}
		y = (y * y) % m
	}
	return r
}

func InvMod(x, mod int64) int64 {
	y := mod
	p := int64(1)
	q := int64(0)
	for y > 0 {
		c := x / y
		d := x
		x = y
		y = d % y
		d = p
		p = q
		q = d - c*q
	}
	if p < 0 {
		return p + mod
	}
	return p
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1048576
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)

	var err error
	scanner.Scan()
	R, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	G, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	B, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	K, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	X, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	Y, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	Z, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(R, G, B, K, X, Y, Z)
}
