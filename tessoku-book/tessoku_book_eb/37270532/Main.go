package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(o Printer, Q int, queries []Query) {
	rbt := NewRedBlackTree()
	for _, q := range queries {
		switch q.T {
		case 1:
			rbt.Add(q.X, q.X)
		case 2:
			gteNode, foundGTE := rbt.GetGTE(q.X)
			lteNode, foundLTE := rbt.GetLTE(q.X)
			switch {
			case foundGTE && foundLTE:
				o.l(Min(gteNode.Key-q.X, q.X-lteNode.Key))
			case foundGTE:
				o.l(gteNode.Key - q.X)
			case foundLTE:
				o.l(q.X - lteNode.Key)
			default:
				o.l(-1)
			}
		}
	}
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

type RedBlackTreeKey = int
type RedBlackTreeValue = int

type RedBlackTree struct {
	root *RedBlackTreeNode
	nil  *RedBlackTreeNode
}

func NewRedBlackTree() *RedBlackTree {
	n := &RedBlackTreeNode{}
	n.left = n
	n.right = n
	n.parent = n
	return &RedBlackTree{
		root: n,
		nil:  n,
	}
}

func (t *RedBlackTree) newNode(key RedBlackTreeKey, value RedBlackTreeValue) *RedBlackTreeNode {
	return &RedBlackTreeNode{
		Key:   key,
		Value: value,
		left:  t.nil,
		right: t.nil,
		red:   true,
	}
}

func (t *RedBlackTree) Nil(x *RedBlackTreeNode) bool {
	return x == t.nil
}

func (t *RedBlackTree) Get(key RedBlackTreeKey) (node *RedBlackTreeNode, found bool) {
	x := t.root
	for !t.Nil(x) {
		switch {
		case key == x.Key:
			return x, true
		case key < x.Key:
			x = x.left
		default:
			x = x.right
		}
	}
	return x, false
}

func (t *RedBlackTree) GetGT(key RedBlackTreeKey) (node *RedBlackTreeNode, found bool) {
	x := t.root
	y := t.nil
	for !t.Nil(x) {
		if key < x.Key {
			y = x
			x = x.left
		} else {
			x = x.right
		}
	}
	return y, !t.Nil(y)
}

func (t *RedBlackTree) GetGTE(key RedBlackTreeKey) (node *RedBlackTreeNode, found bool) {
	x := t.root
	y := t.nil
	for !t.Nil(x) {
		switch {
		case key == x.Key:
			return x, true
		case key < x.Key:
			y = x
			x = x.left
		default:
			x = x.right
		}
	}
	return y, !t.Nil(y)
}

func (t *RedBlackTree) GetLT(key RedBlackTreeKey) (node *RedBlackTreeNode, found bool) {
	x := t.root
	y := t.nil
	for !t.Nil(x) {
		if key > x.Key {
			y = x
			x = x.right
		} else {
			x = x.left
		}
	}
	return y, !t.Nil(y)
}

func (t *RedBlackTree) GetLTE(key RedBlackTreeKey) (node *RedBlackTreeNode, found bool) {
	x := t.root
	y := t.nil
	for !t.Nil(x) {
		switch {
		case key == x.Key:
			return x, true
		case key < x.Key:
			x = x.left
		default:
			y = x
			x = x.right
		}
	}
	return y, !t.Nil(y)
}

func (t *RedBlackTree) Next(x *RedBlackTreeNode) (node *RedBlackTreeNode, found bool) {
	if !t.Nil(x.right) {
		return t.min(x.right), true
	}
	y := x.parent
	for !t.Nil(y) && x == y.right {
		x = y
		y = y.parent
	}
	return y, !t.Nil(y)
}

func (t *RedBlackTree) Prev(x *RedBlackTreeNode) (node *RedBlackTreeNode, found bool) {
	if !t.Nil(x.left) {
		return t.max(x.left), true
	}
	y := x.parent
	for !t.Nil(y) && x == y.left {
		x = y
		y = y.parent
	}
	return y, !t.Nil(y)
}

func (t *RedBlackTree) Max() *RedBlackTreeNode {
	return t.max(t.root)
}

func (t *RedBlackTree) max(x *RedBlackTreeNode) *RedBlackTreeNode {
	for !t.Nil(x.right) {
		x = x.right
	}
	return x
}

func (t *RedBlackTree) Min() *RedBlackTreeNode {
	return t.min(t.root)
}

func (t *RedBlackTree) min(x *RedBlackTreeNode) *RedBlackTreeNode {
	for !t.Nil(x.left) {
		x = x.left
	}
	return x
}

func (t *RedBlackTree) Set(key RedBlackTreeKey, value RedBlackTreeValue) (added bool) {
	y := t.nil
	x := t.root
	for !t.Nil(x) {
		y = x
		switch {
		case key == x.Key:
			x.Value = value
			return false
		case key < x.Key:
			x = x.left
		default:
			x = x.right
		}
	}

	z := t.newNode(key, value)
	z.parent = y
	switch {
	case t.Nil(y):
		t.root = z
	case z.Key < y.Key:
		y.left = z
	default:
		y.right = z
	}
	t.addFixup(z)
	return true
}

func (t *RedBlackTree) Add(key RedBlackTreeKey, value RedBlackTreeValue) (added bool) {
	y := t.nil
	x := t.root
	for !t.Nil(x) {
		y = x
		switch {
		case key == x.Key:
			return false
		case key < x.Key:
			x = x.left
		default:
			x = x.right
		}
	}

	z := t.newNode(key, value)
	z.parent = y
	switch {
	case t.Nil(y):
		t.root = z
	case z.Key < y.Key:
		y.left = z
	default:
		y.right = z
	}
	t.addFixup(z)
	return true
}

func (t *RedBlackTree) addFixup(z *RedBlackTreeNode) {
	for z.parent.Red() {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.Red() {
				z.parent.setBlack()
				y.setBlack()
				z.parent.parent.setRed()
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.rotateLeft(z)
				}
				z.parent.setBlack()
				z.parent.parent.setRed()
				t.rotateRight(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.Red() {
				z.parent.setBlack()
				y.setBlack()
				z.parent.parent.setRed()
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rotateRight(z)
				}
				z.parent.setBlack()
				z.parent.parent.setRed()
				t.rotateLeft(z.parent.parent)
			}
		}
	}
	t.root.setBlack()
}

func (t *RedBlackTree) Delete(key RedBlackTreeKey) (deleted bool) {
	node, found := t.Get(key)
	if !found {
		return false
	}
	t.delete(node)
	return true
}

func (t *RedBlackTree) delete(z *RedBlackTreeNode) {
	y := z
	yOriginBlack := y.Black()
	var x *RedBlackTreeNode
	switch {
	case t.Nil(z.left):
		x = z.right
		t.transplant(z, z.right)
	case t.Nil(z.right):
		x = z.left
		t.transplant(z, z.left)
	default:
		y := t.min(z.right)
		yOriginBlack = y.Black()
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.setColorOf(z)
	}
	if yOriginBlack {
		t.deleteFixup(x)
	}
}

func (t *RedBlackTree) deleteFixup(x *RedBlackTreeNode) {
	for x != t.root && x.Black() {
		if x == x.parent.left {
			w := x.parent.right
			if w.Red() {
				w.setBlack()
				x.parent.setRed()
				t.rotateLeft(x.parent)
				w = x.parent.right
			}
			if w.left.Black() && w.right.Black() {
				w.setRed()
				x = x.parent
			} else {
				if w.right.Black() {
					w.left.setBlack()
					w.setRed()
					t.rotateRight(w)
					w = x.parent.right
				}
				w.setColorOf(x.parent)
				x.parent.setBlack()
				w.right.setBlack()
				t.rotateLeft(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.Red() {
				w.setBlack()
				x.parent.setRed()
				t.rotateRight(x.parent)
				w = x.parent.left
			}
			if w.left.Black() && w.right.Black() {
				w.setRed()
				x = x.parent
			} else {
				if w.left.Black() {
					w.right.setBlack()
					w.setRed()
					t.rotateLeft(w)
					w = x.parent.left
				}
				w.setColorOf(x.parent)
				x.parent.setBlack()
				w.left.setBlack()
				t.rotateRight(x.parent)
				x = t.root
			}
		}
	}
	x.setBlack()
}

func (t *RedBlackTree) transplant(u, v *RedBlackTreeNode) {
	switch {
	case t.Nil(u.parent):
		t.root = v
	case u == u.parent.left:
		u.parent.left = v
	default:
		u.parent.right = v
	}
	v.parent = u.parent
}

func (t *RedBlackTree) rotateLeft(x *RedBlackTreeNode) {
	y := x.right
	x.right = y.left
	if !t.Nil(y.left) {
		y.left.parent = x
	}
	y.parent = x.parent
	switch {
	case t.Nil(x.parent):
		t.root = y
	case x == x.parent.left:
		x.parent.left = y
	default:
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RedBlackTree) rotateRight(x *RedBlackTreeNode) {
	y := x.left
	x.left = y.right
	if !t.Nil(y.right) {
		y.right.parent = x
	}
	y.parent = x.parent
	switch {
	case t.Nil(x.parent):
		t.root = y
	case x == x.parent.right:
		x.parent.right = y
	default:
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

type RedBlackTreeNode struct {
	Key                 RedBlackTreeKey
	Value               RedBlackTreeValue
	left, right, parent *RedBlackTreeNode
	red                 bool
}

func (n *RedBlackTreeNode) Black() bool {
	return !n.red
}

func (n *RedBlackTreeNode) Red() bool {
	return n.red
}

func (n *RedBlackTreeNode) setBlack() {
	n.red = false
}

func (n *RedBlackTreeNode) setRed() {
	n.red = true
}

func (n *RedBlackTreeNode) setColorOf(x *RedBlackTreeNode) {
	n.red = x.red
}

func (t *RedBlackTree) String() string {
	if t.Nil(t.root) {
		return "[Nil]"
	}

	depth := func(t *RedBlackTree, x *RedBlackTreeNode) int {
		result := 0
		for !t.Nil(x) {
			result++
			x = x.parent
		}
		return result
	}

	lines := make([]string, 0)
	s := make([]*RedBlackTreeNode, 1)
	s[0] = t.root
	for len(s) > 0 {
		n := s[len(s)-1]
		s = s[:len(s)-1]
		lines = append(lines, strings.Repeat("-", depth(t, n))+fmt.Sprint(n))
		if !t.Nil(n.left) {
			s = append(s, n.left)
		}
		if !t.Nil(n.right) {
			s = append(s, n.right)
		}
	}
	return strings.Join(lines, "\n")
}

func (n *RedBlackTreeNode) String() string {
	prefix := "(🔲)"
	if n.Red() {
		prefix = "(🔴)"
	}
	s := fmt.Sprintf("%s key:%v val:%v", prefix, n.Key, n.Value)
	return s
}

func (t *RedBlackTree) InorderWalk(x *RedBlackTreeNode, fn func(node *RedBlackTreeNode)) {
	if t.Nil(x) {
		return
	}
	t.InorderWalk(x.left, fn)
	fn(x)
	t.InorderWalk(x.right, fn)
}

func (t *RedBlackTree) PreorderWalk(x *RedBlackTreeNode, fn func(node *RedBlackTreeNode)) {
	if t.Nil(x) {
		return
	}
	fn(x)
	t.PreorderWalk(x.left, fn)
	t.PreorderWalk(x.right, fn)
}

func (t *RedBlackTree) PostorderWalk(x *RedBlackTreeNode, fn func(node *RedBlackTreeNode)) {
	if t.Nil(x) {
		return
	}
	t.PostorderWalk(x.left, fn)
	t.PostorderWalk(x.right, fn)
	fn(x)
}

type Query struct {
	T int
	X int
}

func main() {
	sc := NewScanner()
	Q := sc.Int()
	queries := make([]Query, Q)
	for i := range queries {
		queries[i].T = sc.Int()
		queries[i].X = sc.Int()
	}
	out := NewPrinter()
	solve(out, Q, queries)
	if err := out.w.Flush(); err != nil {
		panic(err)
	}
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner() *Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Buffer(make([]byte, 4096), math.MaxInt64)
	s.Split(bufio.ScanWords)
	return &Scanner{
		Scanner: s,
	}
}

func (s *Scanner) Scan() {
	if ok := s.Scanner.Scan(); !ok {
		panic(s.Err())
	}
}

func (s *Scanner) Int() int {
	s.Scan()
	v, err := strconv.Atoi(s.Scanner.Text())
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) IntN(n int) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = s.Int()
	}
	return v
}

func (s *Scanner) IntN2(n int) ([]int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	for i := 0; i < n; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
	}
	return v1, v2
}

func (s *Scanner) IntN3(n int) ([]int, []int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	v3 := make([]int, n)
	for i := 0; i < n; i++ {
		v1[i] = s.Int()
		v2[i] = s.Int()
		v3[i] = s.Int()
	}
	return v1, v2, v3
}

func (s *Scanner) IntN4(n int) ([]int, []int, []int, []int) {
	v1 := make([]int, n)
	v2 := make([]int, n)
	v3 := make([]int, n)
	v4 := make([]int, n)
	for i := 0; i < n; i++ {
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
	b := s.Scanner.Bytes()
	v := make([]byte, len(b))
	copy(v, b)
	return v
}

func (s *Scanner) BytesN(n int) [][]byte {
	v := make([][]byte, n)
	for i := 0; i < n; i++ {
		v[i] = s.Bytes()
	}
	return v
}

func (s *Scanner) Float() float64 {
	s.Scan()
	v, err := strconv.ParseFloat(s.Text(), 64)
	if err != nil {
		panic(err)
	}
	return v
}

func (s *Scanner) Text() string {
	s.Scan()
	return s.Scanner.Text()
}

type Printer interface {
	// p fmt.Print
	p(a ...interface{})
	// f fmt.Printf
	f(format string, a ...interface{})
	// l fmt.Println
	l(a ...interface{})
}

type printer struct {
	w *bufio.Writer
}

func NewPrinter() *printer {
	return &printer{bufio.NewWriter(os.Stdout)}
}

func (p *printer) p(a ...interface{}) {
	fmt.Fprint(p.w, a...)
}

func (p *printer) f(format string, a ...interface{}) {
	fmt.Fprintf(p.w, format, a...)
}

func (p *printer) l(a ...interface{}) {
	fmt.Fprintln(p.w, a...)
}
