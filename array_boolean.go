package arrow

import "github.com/influxdata/arrow/internal/bitutil"

// BooleanArray represents an immutable sequence of boolean values.
type BooleanArray struct {
	array
	values []byte
}

func NewBooleanArray(data *ArrayData) *BooleanArray {
	a := &BooleanArray{}
	a.setData(data)
	return a
}

func (a *BooleanArray) Value(i int) bool { return bitutil.BitIsSet(a.values, i) }

func (a *BooleanArray) setData(data *ArrayData) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = BooleanTraits{}.CastFromBytes(vals.Bytes())
	}
}
