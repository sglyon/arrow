package arrow

type BooleanArray struct {
	array
	values []byte
}

func NewBooleanArray(data *ArrayData) *BooleanArray {
	a := &BooleanArray{}
	a.setData(data)
	return a
}

func (a *BooleanArray) Value(i int) bool { return bitIsSet(a.values, i) }

func (a *BooleanArray) setData(data *ArrayData) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = BooleanTraits{}.CastFromBytes(vals.Bytes())
	}
}
