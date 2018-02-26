package rbtree

type SetNode struct {
	n _node
}

func (n SetNode) GetData() interface{} {
	return n.n.tree.getKey(n.n.node)
}

func (n SetNode) Next() SetNode {
	return SetNode{n.n.tree.Next(n.n)}
}

func (n SetNode) Last() SetNode {
	return SetNode{n.n.tree.Last(n.n)}
}

type baseSet struct {
	tree
}

func (s *baseSet) pack(n _node) SetNode {
	return SetNode{n: n}
}

func (s *baseSet) Init(unique bool, data interface{}, compare func(a, b interface{}) int) {
	s.tree.Init(unique, data, nil, compare)
}

func (s *baseSet) Begin() SetNode {
	return s.pack(s.tree.Begin())
}

func (s *baseSet) End() SetNode {
	return s.pack(s.tree.End())
}

func (s *baseSet) EqualRange(data interface{}) (beg, end SetNode) {
	a, b := s.tree.EqualRange(data)
	return s.pack(a), s.pack(b)
}

func (s *baseSet) EraseNode(n SetNode) {
	s.tree.EraseNode(n.n)
}

func (s *baseSet) EraseNodeRange(beg, end SetNode) (count int) {
	return s.tree.EraseNodeRange(beg.n, end.n)
}

func (s *baseSet) Find(data interface{}) SetNode {
	return s.pack(s.tree.Find(data))
}

func (s *baseSet) Insert(data interface{}) (SetNode, bool) {
	n, ok := s.tree.Insert(data, nil)
	return s.pack(n), ok
}

func (s *baseSet) LowerBound(data interface{}) SetNode {
	return s.pack(s.tree.LowerBound(data))
}

func (s *baseSet) Last(n SetNode) SetNode {
	return s.pack(s.tree.Last(n.n))
}

func (s *baseSet) Next(n SetNode) SetNode {
	return s.pack(s.tree.Next(n.n))
}

func (s *baseSet) UpperBound(data interface{}) SetNode {
	return s.pack(s.tree.UpperBound(data))
}

func (s *baseSet) Count(data interface{}) (count int) {
	return s.tree.Count(data)
}

func (s *baseSet) Erase(data interface{}) (count int) {
	return s.tree.Erase(data)
}

type BaseSet struct {
	baseSet
}

type Set struct {
	baseSet
}

func (s *Set) Init(unique bool, data interface{}, compare func(a, b interface{}) int) {
	s.ifacedata = true
	s.baseSet.Init(unique, data, compare)
}

func NewSet(data interface{}, compare func(a, b interface{}) int) *Set {
	var s = &Set{}
	s.Init(true, data, compare)
	return s
}

func NewMultiSet(data interface{}, compare func(a, b interface{}) int) *Set {
	var s = &Set{}
	s.Init(false, data, compare)
	return s
}

type IntSetNode struct {
	n SetNode
}

func (n IntSetNode) GetData() int {
	return n.n.GetData().(int)
}

func (n IntSetNode) Next() IntSetNode {
	return IntSetNode{n.n.Next()}
}

func (n IntSetNode) Last() IntSetNode {
	return IntSetNode{n.n.Last()}
}

type IntSet struct {
	baseSet
}

func NewIntSet(compare func(a, b int) int) *IntSet {
	var s = &IntSet{}
	s.Init(true, compare)
	return s
}

func NewIntMultiSet(compare func(a, b int) int) *IntSet {
	var s = &IntSet{}
	s.Init(false, compare)
	return s
}

var tmpInt interface{} = 1

func (s *IntSet) pack(n SetNode) IntSetNode {
	return IntSetNode{n: n}
}

func (s *IntSet) Init(unique bool, compare func(a, b int) int) {
	s.baseSet.Init(unique, tmpInt, func(a, b interface{}) int {
		return compare(a.(int), b.(int))
	})
}

func (s *IntSet) Begin() IntSetNode {
	return s.pack(s.baseSet.Begin())
}

func (s *IntSet) End() IntSetNode {
	return s.pack(s.baseSet.End())
}

func (s *IntSet) EqualRange(data int) (beg, end IntSetNode) {
	a, b := s.baseSet.EqualRange(data)
	return s.pack(a), s.pack(b)
}

func (s *IntSet) EraseNode(n IntSetNode) {
	s.baseSet.EraseNode(n.n)
}

func (s *IntSet) Erase(data int) (count int) {
	return s.baseSet.Erase(NoescapeInterface(data))
}

func (s *IntSet) EraseNodeRange(beg, end IntSetNode) (count int) {
	return s.baseSet.EraseNodeRange(beg.n, end.n)
}

func (s *IntSet) Find(data int) IntSetNode {
	return s.pack(s.baseSet.Find(NoescapeInterface(data)))
}

func (s *IntSet) Insert(data int) (IntSetNode, bool) {
	n, ok := s.baseSet.Insert(NoescapeInterface(data))
	return s.pack(n), ok
}

func (s *IntSet) LowerBound(data int) IntSetNode {
	return s.pack(s.baseSet.LowerBound(NoescapeInterface(data)))
}

func (s *IntSet) Last(n IntSetNode) IntSetNode {
	return s.pack(s.baseSet.Last(n.n))
}

func (s *IntSet) Next(n IntSetNode) IntSetNode {
	return s.pack(s.baseSet.Next(n.n))
}

func (s *IntSet) UpperBound(data int) IntSetNode {
	return s.pack(s.baseSet.UpperBound(NoescapeInterface(data)))
}
