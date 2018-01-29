package array

import (
	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/internal/bitutil"
)

// A type which represents an immutable sequence of boolean values.
type Boolean struct {
	array
	values []byte
}

func NewBooleanData(data *Data) *Boolean {
	a := &Boolean{}
	a.setData(data)
	return a
}

func (a *Boolean) Value(i int) bool { return bitutil.BitIsSet(a.values, i) }

func (a *Boolean) setData(data *Data) {
	a.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = arrow.BooleanTraits{}.CastFromBytes(vals.Bytes())
	}
}
