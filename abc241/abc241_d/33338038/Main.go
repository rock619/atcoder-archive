package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func solve(Q int, query []Query) {
	w := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := w.Flush(); err != nil {
			panic(err)
		}
	}()

	rbt := NewRedBlackTree()
	for _, q := range query {
		switch key := q.X; q.C {
		case 1:
			if v, ok := rbt.Find(key); ok {
				rbt.Set(key, v+1)
			} else {
				rbt.Add(key, 1)
			}
		case 2:
			for node, count := rbt.FindLTENode(key), 0; ; node = node.Prev() {
				if node.isNil {
					fmt.Fprintln(w, -1)
					break
				}
				count += node.Value
				if count >= q.K {
					fmt.Fprintln(w, node.Key)
					break
				}
			}
		default:
			for node, count := rbt.FindGTENode(key), 0; ; node = node.Next() {
				if node.isNil {
					fmt.Fprintln(w, -1)
					break
				}
				count += node.Value
				if count >= q.K {
					fmt.Fprintln(w, node.Key)
					break
				}
			}
		}
	}
}

func Compress(s []int) []int {
	s2 := make([]int, len(s))
	copy(s2, s)
	sort.Ints(s2)
	j := 0
	for i := 1; i < len(s2); i++ {
		if s2[j] == s2[i] {
			continue
		}
		j++
		s2[j] = s2[i]
	}
	s2 = s2[:j+1]

	result := make([]int, len(s))
	for i, v := range s {
		result[i] = sort.SearchInts(s2, v)
	}
	return result
}

type Query struct {
	C, X, K int
}

func main() {
	s := NewScanner()
	Q := s.Int()
	query := make([]Query, Q)
	for i := range query {
		query[i].C = s.Int()
		query[i].X = s.Int()
		if query[i].C == 2 || query[i].C == 3 {
			query[i].K = s.Int()
		}
	}

	solve(Q, query)
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
	if ok := s.Scan(); !ok {
		panic(s.Err())
	}
	v, err := strconv.Atoi(s.Text())
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
	if ok := s.Scan(); !ok {
		panic(s.Err())
	}

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

type RedBlackTreeKey = int
type RedBlackTreeValue = int

type RedBlackTree struct {
	r       *RedBlackTreeNode
	n       int
	nilNode *RedBlackTreeNode
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{
		r:       NewNilRedBlackTreeNode(),
		nilNode: NewNilRedBlackTreeNode(),
	}
}

func (t *RedBlackTree) NewNode(key RedBlackTreeKey, val RedBlackTreeValue) *RedBlackTreeNode {
	return &RedBlackTreeNode{
		Key:    key,
		Value:  val,
		left:   t.nilNode,
		right:  t.nilNode,
		parent: t.nilNode,
	}
}

func (t *RedBlackTree) Find(key RedBlackTreeKey) (v RedBlackTreeValue, ok bool) {
	u := t.r
	for !u.isNil {
		switch {
		case key < u.Key:
			u = u.left
		case key > u.Key:
			u = u.right
		default:
			return u.Value, true
		}
	}
	return v, false
}

func (t *RedBlackTree) FindGTE(key RedBlackTreeKey) (v RedBlackTreeValue, ok bool) {
	if w := t.FindGTENode(key); !w.isNil {
		return w.Value, true
	}
	return v, false
}

func (t *RedBlackTree) FindGTENode(key RedBlackTreeKey) *RedBlackTreeNode {
	w, z := t.r, t.nilNode
	for !w.isNil {
		switch {
		case key < w.Key:
			z = w
			w = w.left
		case key > w.Key:
			w = w.right
		default:
			return w
		}
	}
	return z
}

func (t *RedBlackTree) FindLTE(key RedBlackTreeKey) (v RedBlackTreeValue, ok bool) {
	if w := t.FindLTENode(key); !w.isNil {
		return w.Value, true
	}
	return v, false
}

func (t *RedBlackTree) FindLTENode(key RedBlackTreeKey) *RedBlackTreeNode {
	u, z := t.r, t.nilNode
	for !u.isNil {
		switch {
		case key < u.Key:
			u = u.left
		case key > u.Key:
			z = u
			u = u.right
		default:
			return u
		}
	}
	return z
}

func (t *RedBlackTree) Set(key RedBlackTreeKey, val RedBlackTreeValue) {
	w := t.r
	prev := t.nilNode
	for !w.isNil {
		prev = w
		switch {
		case key < w.Key:
			w = w.left
		case key > w.Key:
			w = w.right
		default:
			w.Value = val
			return
		}
	}

	u := t.NewNode(key, val)
	u.setRed()
	t.addChild(prev, u)
	t.addFixup(u)
}

func (t *RedBlackTree) Add(key RedBlackTreeKey, val RedBlackTreeValue) bool {
	u := t.NewNode(key, val)
	u.setRed()
	added := t.add(u)
	if added {
		t.addFixup(u)
	}
	return added
}

func (t *RedBlackTree) add(u *RedBlackTreeNode) bool {
	p := t.findLast(u.Key)
	return t.addChild(p, u)
}

func (t *RedBlackTree) findLast(key RedBlackTreeKey) *RedBlackTreeNode {
	w := t.r
	prev := t.nilNode
	for !w.isNil {
		prev = w
		if key < w.Key {
			w = w.left
		} else if key > w.Key {
			w = w.right
		} else {
			return w
		}
	}
	return prev
}

func (t *RedBlackTree) addChild(p, u *RedBlackTreeNode) bool {
	if p.isNil {
		t.r = u
	} else {
		if u.Key < p.Key {
			p.left = u
		} else if u.Key > p.Key {
			p.right = u
		} else {
			return false
		}
		u.parent = p
	}
	t.n++
	return true
}

func (t *RedBlackTree) addFixup(u *RedBlackTreeNode) {
	for u.Red() {
		if u == t.r {
			u.setBlack()
			return
		}
		w := u.parent
		if w.left.Black() { // ensure left-leaning
			t.flipLeft(w)
			u = w
			w = u.parent
		}
		if w.Black() {
			return // no red-red edge = done
		}

		g := w.parent // grandparent of u
		if g.right.Black() {
			t.flipRight(g)
			return
		} else {
			t.pushBlack(g)
			u = g
		}
	}
}

func (t *RedBlackTree) Remove(key RedBlackTreeKey) bool {
	u := t.findLast(key)
	if u.isNil || u.Key != key {
		return false
	}

	w := u.right
	if w.isNil {
		w = u
		u = w.left
	} else {
		for !w.left.isNil {
			w = w.left
		}
		u.Key = w.Key // TODO: check
		u = w.right
	}
	t.splice(w)
	if u.Red() && w.Black() {
		u.setBlack()
	}
	u.parent = w.parent
	t.removeFixup(u)
	return true
}

func (t *RedBlackTree) splice(u *RedBlackTreeNode) {
	var s, p *RedBlackTreeNode
	if !u.left.isNil {
		s = u.left
	} else {
		s = u.right
	}
	if u == t.r {
		t.r = s
		p = t.nilNode
	} else {
		p = u.parent
		if p.left == u {
			p.left = s
		} else {
			p.right = s
		}
	}
	if !s.isNil {
		s.parent = p
	}
	t.n--
}

func (t *RedBlackTree) removeFixup(u *RedBlackTreeNode) {
	for u.Black() {
		switch {
		case u == t.r:
			u.setBlack()
		case u.parent.left.Red():
			u = t.removeFixupCase1(u)
		case u == u.parent.left:
			u = t.removeFixupCase2(u)
		default:
			u = t.removeFixupCase3(u)
		}
	}
	if u != t.r {
		if w := u.parent; w.right.Red() && w.left.Black() {
			t.flipLeft(w)
		}
	}
}

func (t *RedBlackTree) removeFixupCase1(u *RedBlackTreeNode) *RedBlackTreeNode {
	t.flipRight(u.parent)
	return u
}

func (t *RedBlackTree) removeFixupCase2(u *RedBlackTreeNode) *RedBlackTreeNode {
	w := u.parent
	v := w.right
	t.pullBlack(w)
	t.flipLeft(w)
	q := w.right
	if q.Red() {
		t.rotateLeft(w)
		t.flipRight(v)
		t.pushBlack(q)
		if v.right.Red() {
			t.flipLeft(v)
		}
		return q
	} else {
		return v
	}
}

func (t *RedBlackTree) removeFixupCase3(u *RedBlackTreeNode) *RedBlackTreeNode {
	w := u.parent
	v := w.left
	t.pullBlack(w)
	t.flipRight(w)
	q := w.left
	if q.Red() {
		t.rotateRight(w)
		t.flipLeft(v)
		t.pushBlack(q)
		return q
	} else {
		if v.left.Red() {
			t.pushBlack(v)
			return v
		} else {
			t.flipLeft(v)
			return w
		}
	}
}

func (t *RedBlackTree) pushBlack(u *RedBlackTreeNode) {
	u.setRed()
	u.left.setBlack()
	u.right.setBlack()
}

func (t *RedBlackTree) pullBlack(u *RedBlackTreeNode) {
	u.setBlack()
	u.left.setRed()
	u.right.setRed()
}

func (t *RedBlackTree) flipLeft(u *RedBlackTreeNode) {
	t.swapColors(u, u.right)
	t.rotateLeft(u)
}

func (t *RedBlackTree) rotateLeft(u *RedBlackTreeNode) {
	w := u.right
	w.parent = u.parent
	if !w.parent.isNil {
		if w.parent.left.Key == u.Key {
			w.parent.left = w
		} else {
			w.parent.right = w
		}
	}
	u.right = w.left
	if !u.right.isNil {
		u.right.parent = u
	}
	u.parent = w
	w.left = u
	if u.Key == t.r.Key {
		t.r = w
		t.r.parent = t.nilNode
	}
}

func (t *RedBlackTree) flipRight(u *RedBlackTreeNode) {
	t.swapColors(u, u.left)
	t.rotateRight(u)
}

func (t *RedBlackTree) rotateRight(u *RedBlackTreeNode) {
	w := u.left
	w.parent = u.parent
	if !w.parent.isNil {
		if w.parent.left.Key == u.Key {
			w.parent.left = w
		} else {
			w.parent.right = w
		}
	}
	u.left = w.right
	if !u.left.isNil {
		u.left.parent = u
	}
	u.parent = w
	w.right = u
	if u.Key == t.r.Key {
		t.r = w
		t.r.parent = t.nilNode
	}
}

func (t *RedBlackTree) swapColors(u, w *RedBlackTreeNode) {
	u.color, w.color = w.color, u.color
}

type RedBlackTreeNode struct {
	Key                 RedBlackTreeKey
	Value               RedBlackTreeValue
	color               RedBlackTreeColor
	left, right, parent *RedBlackTreeNode
	isNil               bool
}

func NewNilRedBlackTreeNode() *RedBlackTreeNode {
	n := &RedBlackTreeNode{
		color: RedBlackTreeColorBlack,
		isNil: true,
	}
	n.left, n.right, n.parent = n, n, n
	return n
}

func (n *RedBlackTreeNode) Equal(other *RedBlackTreeNode) bool {
	if n.isNil && other.isNil {
		return true
	}
	return n.isNil == other.isNil && n.Key == other.Key
}

func (n *RedBlackTreeNode) Less(other *RedBlackTreeNode) bool {
	if other.isNil {
		return false
	}
	return n.isNil || n.Key < other.Key
}

func (n *RedBlackTreeNode) Next() *RedBlackTreeNode {
	w := n
	if w.right.isNil {
		for !w.parent.isNil && !w.parent.left.Equal(w) {
			w = w.parent
		}
		w = w.parent
	} else {
		w = w.right
		for !w.left.isNil {
			w = w.left
		}
	}
	return w
}

func (n *RedBlackTreeNode) Prev() *RedBlackTreeNode {
	w := n
	if w.left.isNil {
		for !w.parent.isNil && !w.parent.right.Equal(w) {
			w = w.parent
		}
		w = w.parent
	} else {
		w = w.left
		for !w.right.isNil {
			w = w.right
		}
	}
	return w
}

func (n RedBlackTreeNode) String() string {
	if n.isNil {
		return "Node{nil}"
	}

	l := "Left:nil"
	if !n.left.isNil {
		l = fmt.Sprintf("Left:{Key:%v Value:%v color:%v}", n.left.Key, n.left.Value, n.left.color)
	}
	r := "Right:nil"
	if !n.right.isNil {
		r = fmt.Sprintf("Right:{Key:%v Value:%v color:%v}", n.right.Key, n.right.Value, n.right.color)
	}
	p := "Parent:nil"
	if !n.parent.isNil {
		p = fmt.Sprintf("Parent:{Key:%v Value:%v color:%v}", n.parent.Key, n.parent.Value, n.parent.color)
	}
	return fmt.Sprintf("Node{Key:%v Value:%v color:%v %s %s %s}", n.Key, n.Value, n.color, l, r, p)
}

func (n *RedBlackTreeNode) Black() bool {
	return bool(n.color)
}

func (n *RedBlackTreeNode) Red() bool {
	return !n.Black()
}

func (n *RedBlackTreeNode) setBlack() {
	n.color = RedBlackTreeColorBlack
}

func (n *RedBlackTreeNode) setRed() {
	n.color = RedBlackTreeColorRed
}

type RedBlackTreeColor bool

const (
	RedBlackTreeColorBlack RedBlackTreeColor = true
	RedBlackTreeColorRed   RedBlackTreeColor = false
)

func (c RedBlackTreeColor) String() string {
	if bool(c) {
		return "black"
	}
	return "red"
}

type RedBlackTreeIterator struct {
	w, prev *RedBlackTreeNode
}

func NewRedBlackTreeIterator(n *RedBlackTreeNode) *RedBlackTreeIterator {
	return &RedBlackTreeIterator{
		w: n,
	}
}

func (i *RedBlackTreeIterator) HasNext() bool {
	return !i.w.isNil
}
