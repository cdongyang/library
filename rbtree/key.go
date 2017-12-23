package rbtree

type IntKey int

func (k IntKey) Compare(key Keyer) int {
	return int(k - key.(IntKey))
}

func (k IntKey) GetKey() Keyer {
	return k
}

type IntPoiterKey int

func (k *IntPoiterKey) Compare(key Keyer) int {
	return int(*k - *(key.(*IntPoiterKey)))
}

func (k *IntPoiterKey) GetKey() Keyer {
	return k
}
