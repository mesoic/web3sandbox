package utils

type NilableInt struct {
	value int
	null  bool
}

func (n *NilableInt) Value() interface{} {
	if n.null {
		return nil
	}
	return n.value
}

func NewInt(value int) NilableInt {
	return NilableInt{value, false}
}

func NewNil() NilableInt {
	return NilableInt{0, true}
}
