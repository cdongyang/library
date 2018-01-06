package rbtree

type IntKey int

func (k IntKey) Compare(key Keyer) int {
	return int(k - key.(IntKey))
}
