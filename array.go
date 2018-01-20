package arrow

import "github.com/influxdata/arrow/memory"

type Array interface {
	DataType() DataType
	NullN() int
	NullBitmapBytes() []byte
	IsNull(i int) bool
	IsValid(i int) bool
	Data() *ArrayData
	Len() int
	Values() *memory.Buffer
}

type array struct {
	data            *ArrayData
	nullBitmapBytes []byte
}

func (a *array) DataType() DataType      { return a.data.typE }
func (a *array) NullN() int              { return a.data.nullN }
func (a *array) NullBitmapBytes() []byte { return a.nullBitmapBytes }
func (a *array) Data() *ArrayData        { return a.data }
func (a *array) Len() int                { return a.data.length }
func (a *array) Values() *memory.Buffer  { return a.data.buffers[1] }

// IsNull returns true if value at index is null. Does not check bounds.
func (a *array) IsNull(i int) bool {
	return len(a.nullBitmapBytes) != 0 && bitIsNotSet(a.nullBitmapBytes, i)
}

// IsValid returns true if value at index is not null. Does not check bounds.
func (a *array) IsValid(i int) bool {
	return len(a.nullBitmapBytes) == 0 || bitIsSet(a.nullBitmapBytes, i)
}

func (a *array) setData(data *ArrayData) {
	if len(data.buffers) > 0 && data.buffers[0] != nil {
		a.nullBitmapBytes = data.buffers[0].Bytes()
	}
	a.data = data
}
