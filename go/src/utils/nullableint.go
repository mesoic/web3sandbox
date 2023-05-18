package utils

type NullableInt struct {
	value int
	null  bool
}

func (n *NullableInt) Value() interface{} {
	if n.null {
		return nil
	}
	return n.value
}

func NewInt(value int) NullableInt {
	return NullableInt{value, false}
}

func NewNull() NullableInt {
	return NullableInt{0, true}
}
