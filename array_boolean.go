package arrow

import "github.com/influxdata/arrow/memory"

type BooleanArray struct {
	data            *ArrayData
	nullBitmapBytes []byte
	values          []byte
}

func NewBooleanArray(data *ArrayData) *BooleanArray {
	a := &BooleanArray{}
	a.setData(data)
	return a
}

func (a *BooleanArray) DataType() DataType      { return a.data.typE }
func (a *BooleanArray) NullN() int              { return a.data.nullN }
func (a *BooleanArray) NullBitmapBytes() []byte { return a.nullBitmapBytes }
func (a *BooleanArray) Data() *ArrayData        { return a.data }
func (a *BooleanArray) Len() int                { return a.data.length }
func (a *BooleanArray) Values() *memory.Buffer  { return a.data.buffers[1] }
func (a *BooleanArray) Value(i int) bool        { return bitIsSet(a.values, i) }

// IsNull returns true if value at index is null. Does not check bounds.
func (a *BooleanArray) IsNull(i int) bool {
	return len(a.nullBitmapBytes) != 0 && bitIsNotSet(a.nullBitmapBytes, i)
}

// IsValid returns true if value at index is not null. Does not check bounds.
func (a *BooleanArray) IsValid(i int) bool {
	return len(a.nullBitmapBytes) == 0 || bitIsSet(a.nullBitmapBytes, i)
}

func (a *BooleanArray) arraySetData(data *ArrayData) {
	if len(data.buffers) > 0 && data.buffers[0] != nil {
		a.nullBitmapBytes = data.buffers[0].Bytes()
	}
	a.data = data
}

func (a *BooleanArray) setData(data *ArrayData) {
	a.arraySetData(data)
	vals := data.buffers[1]
	if vals != nil {
		a.values = BooleanTraits{}.CastFromBytes(vals.Bytes())
	}
}
