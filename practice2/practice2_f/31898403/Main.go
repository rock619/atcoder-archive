package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

const mod = 998244353

func solve(N int64, M int64, a []int64, b []int64) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	a2 := make([]*ModInt, N)
	for i, v := range a {
		a2[i] = NewModInt(v, mod)
	}
	b2 := make([]*ModInt, M)
	for i, v := range b {
		b2[i] = NewModInt(v, mod)
	}

	fft := NewFFT(mod)
	result := fft.Convolution(a2, b2)
	for i, r := range result {
		if i == 0 {
			fmt.Fprint(w, r.V)
		} else {
			fmt.Fprintf(w, " %d", r.V)
		}
	}
	fmt.Fprintln(w)
}

type ModInt struct {
	V   int64
	Mod int64
}

func NewModInt(v, mod int64) *ModInt {
	mi := &ModInt{
		V:   v % mod,
		Mod: mod,
	}
	if mi.V < 0 {
		mi.V += mi.Mod
	}
	return mi
}

func (mi ModInt) Clone() *ModInt {
	return NewModInt(mi.V, mi.Mod)
}

func (mi *ModInt) Add(x *ModInt) *ModInt {
	mi.V += x.V
	if mi.V >= mi.Mod {
		mi.V -= mi.Mod
	}
	return mi
}

func (mi *ModInt) AddInt(v int64) *ModInt {
	mi.V += v % mi.Mod
	if mi.V < 0 {
		mi.V += mi.Mod
	} else if mi.V >= mi.Mod {
		mi.V -= mi.Mod
	}
	return mi
}

func (mi *ModInt) Sub(x *ModInt) *ModInt {
	mi.V -= x.V
	if mi.V < 0 {
		mi.V += mi.Mod
	}
	return mi
}

func (mi *ModInt) SubInt(v int64) *ModInt {
	mi.V -= v % mi.Mod
	if mi.V < 0 {
		mi.V += mi.Mod
	} else if mi.V >= mi.Mod {
		mi.V -= mi.Mod
	}
	return mi
}

func (mi *ModInt) Mul(x *ModInt) *ModInt {
	mi.V *= x.V
	mi.V %= mi.Mod
	return mi
}

func (mi *ModInt) MulInt(v int64) *ModInt {
	mi.V *= v % mi.Mod
	mi.V %= mi.Mod
	if mi.V < 0 {
		mi.V += mi.Mod
	} else if mi.V >= mi.Mod {
		mi.V -= mi.Mod
	}
	return mi
}

func (mi *ModInt) Div(x *ModInt) *ModInt {
	return mi.Mul(x.Inv())
}

func (mi ModInt) Pow(n int64) *ModInt {
	x := mi.Clone()
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

func (mi ModInt) Inv() *ModInt {
	_, x, _ := ExtendedGCD(mi.V, mi.Mod)
	return NewModInt(x, mi.Mod)
}

// 拡張ユークリッドの互除法
// 参考: https://cp-algorithms.com/algebra/extended-euclid-algorithm.html
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

type FFT struct {
	mod           int64
	g             int64
	rank2         int64
	root, iRoot   []*ModInt
	rate2, iRate2 []*ModInt
	rate3, iRate3 []*ModInt
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
	f.root = make([]*ModInt, f.rank2+1)
	f.iRoot = make([]*ModInt, f.rank2+1)

	f.root[f.rank2] = NewModInt(f.g, f.mod).Pow((f.mod - 1) >> f.rank2)
	f.iRoot[f.rank2] = f.root[f.rank2].Inv()

	for i := f.rank2 - 1; i >= 0; i-- {
		f.root[i] = f.root[i+1].Clone().Mul(f.root[i+1])
		f.iRoot[i] = f.iRoot[i+1].Clone().Mul(f.iRoot[i+1])
	}
}

func (f *FFT) initRate2() {
	if l := f.rank2 - 2 + 1; l > 0 {
		f.rate2 = make([]*ModInt, l)
		f.iRate2 = make([]*ModInt, l)
	}

	prod := NewModInt(1, f.mod)
	iProd := NewModInt(1, f.mod)
	for i := int64(0); i <= f.rank2-2; i++ {
		f.rate2[i] = f.root[i+2].Clone().Mul(prod)
		f.iRate2[i] = f.iRoot[i+2].Clone().Mul(iProd)
		prod.Mul(f.iRoot[i+2])
		iProd.Mul(f.root[i+2])
	}
}

func (f *FFT) initRate3() {
	if l := f.rank2 - 3 + 1; l > 0 {
		f.rate3 = make([]*ModInt, l)
		f.iRate3 = make([]*ModInt, l)
	}

	prod := NewModInt(1, f.mod)
	iProd := NewModInt(1, f.mod)
	for i := int64(0); i <= f.rank2-3; i++ {
		f.rate3[i] = f.root[i+3].Clone().Mul(prod)
		f.iRate3[i] = f.iRoot[i+3].Clone().Mul(iProd)
		prod.Mul(f.iRoot[i+3])
		iProd.Mul(f.root[i+3])
	}
}

func (f *FFT) Convolution(a, b []*ModInt) []*ModInt {
	n := int64(len(a))
	m := int64(len(b))
	z := int64(1) << ceilPow2(n+m-1)
	a = append(a, make([]*ModInt, z-n)...)
	for i := n; i < z; i++ {
		a[i] = NewModInt(0, f.mod)
	}
	f.butterfly(a)
	b = append(b, make([]*ModInt, z-m)...)
	for i := m; i < z; i++ {
		b[i] = NewModInt(0, f.mod)
	}
	f.butterfly(b)
	for i := int64(0); i < z; i++ {
		a[i].Mul(b[i])
	}
	f.butterflyInv(a)
	a = a[:n+m-1]
	iz := NewModInt(z, f.mod).Inv()
	for i := int64(0); i < n+m-1; i++ {
		a[i].Mul(iz)
	}
	return a
}

func (f *FFT) butterfly(a []*ModInt) {
	n := int64(len(a))
	h := ceilPow2(n)

	// a[i, i+(n>>length), i+2*(n>>length), ..] is transformed
	for length := int64(0); length < h; {
		if h-length == 1 {
			p := int64(1) << (h - length - 1)
			rot := NewModInt(1, f.mod)
			for s := int64(0); s < (1 << length); s++ {
				offset := s << (h - length)
				for i := int64(0); i < p; i++ {
					l := a[i+offset].Clone()
					r := a[i+offset+p].Clone().Mul(rot)
					a[i+offset].Add(r)
					a[i+offset+p] = l.Sub(r)
				}
				if s+1 != (1 << length) {
					rot.Mul(f.rate2[bits.TrailingZeros64(uint64(^s))])
				}
			}
			length++
			continue
		}

		// 4-base
		p := int64(1) << (h - length - 2)
		rot := NewModInt(1, f.mod)
		imag := f.root[2]
		for s := int64(0); s < (1 << length); s++ {
			rot2 := rot.Clone().Mul(rot)
			rot3 := rot2.Clone().Mul(rot)
			offset := s << (h - length)
			for i := int64(0); i < p; i++ {
				mod2 := f.mod * f.mod
				a0 := a[i+offset].V
				a1 := a[i+offset+p].V * rot.V
				a2 := a[i+offset+2*p].V * rot2.V
				a3 := a[i+offset+3*p].V * rot3.V
				a1na3imag := NewModInt(a1+mod2-a3, f.mod).V * imag.V
				na2 := mod2 - a2
				a[i+offset] = NewModInt(a0+a2+a1+a3, f.mod)
				a[i+offset+1*p] = NewModInt(a0+a2+(2*mod2-(a1+a3)), f.mod)
				a[i+offset+2*p] = NewModInt(a0+na2+a1na3imag, f.mod)
				a[i+offset+3*p] = NewModInt(a0+na2+(mod2-a1na3imag), f.mod)
			}
			if s+1 != (1 << length) {
				rot.Mul(f.rate3[bits.TrailingZeros64(uint64(^s))])
			}
		}
		length += 2
	}
}

func (f *FFT) butterflyInv(a []*ModInt) {
	n := int64(len(a))
	h := ceilPow2(n)

	// a[i, i+(n>>length), i+2*(n>>length), ..] is transformed
	for length := h; length > 0; length-- {
		if length == 1 {
			p := int64(1) << (h - length)
			iRot := NewModInt(1, f.mod)
			for s := int64(0); s < (1 << (length - 1)); s++ {
				offset := s << (h - length + 1)
				for i := int64(0); i < p; i++ {
					l := a[i+offset].Clone()
					r := a[i+offset+p].Clone()
					a[i+offset].Add(r)
					a[i+offset+p] =
						newModIntFromUint64(uint64(f.mod+l.V-r.V)*
							uint64(iRot.V), f.mod)
				}
				if s+1 != (1 << (length - 1)) {
					iRot.Mul(f.iRate2[bits.TrailingZeros64(uint64(^s))])
				}
			}
			continue
		}

		// 4-base
		p := int64(1) << (h - length)
		iRot := NewModInt(1, f.mod)
		iImag := f.iRoot[2]
		for s := int64(0); s < (1 << (length - 2)); s++ {
			iRot2 := iRot.Clone().Mul(iRot)
			iRot3 := iRot2.Clone().Mul(iRot)
			offset := s << (h - length + 2)
			for i := int64(0); i < p; i++ {
				a0 := uint64(a[i+offset+0*p].V)
				a1 := uint64(a[i+offset+1*p].V)
				a2 := uint64(a[i+offset+2*p].V)
				a3 := uint64(a[i+offset+3*p].V)

				mod := uint64(f.mod)
				a2na3iImag :=
					uint64(newModIntFromUint64((mod+a2-a3)*uint64(iImag.V), f.mod).V)

				a[i+offset] = newModIntFromUint64(a0+a1+a2+a3, f.mod)
				a[i+offset+1*p] =
					newModIntFromUint64((a0+(mod-a1)+a2na3iImag)*uint64(iRot.V), f.mod)
				a[i+offset+2*p] =
					newModIntFromUint64((a0+a1+(mod-a2)+(mod-a3))*
						uint64(iRot2.V), f.mod)
				a[i+offset+3*p] =
					newModIntFromUint64((a0+(mod-a1)+(mod-a2na3iImag))*
						uint64(iRot3.V), f.mod)
			}
			if s+1 != (1 << (length - 2)) {
				iRot.Mul(f.iRate3[bits.TrailingZeros64(uint64(^s))])
			}

		}
		length--
	}
}

func newModIntFromUint64(v uint64, mod int64) *ModInt {
	v %= uint64(mod)
	return NewModInt(int64(v), mod)
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
	M, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	a := make([]int64, N)
	for i := int64(0); i < N; i++ {
		scanner.Scan()
		a[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	b := make([]int64, M)
	for i := int64(0); i < M; i++ {
		scanner.Scan()
		b[i], err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	solve(N, M, a, b)
}
