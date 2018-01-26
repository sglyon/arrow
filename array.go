package arrow

import (
	"github.com/influxdata/arrow/internal/bitutil"
	"github.com/influxdata/arrow/memory"
)

// Array is the most basic data structure
type Array interface {
	DataType() DataType

	// NullN returns the number of null values in the array.
	NullN() int

	// TODO
	NullBitmapBytes() []byte

	// IsNull returns true if value at index is null.
	// IsNull will panic if NullBitmapBytes is not empty and 0 > i ≥ Len.
	IsNull(i int) bool

	// IsValid returns true if value at index is not null.
	// IsValid will panic if NullBitmapBytes is not empty and 0 > i ≥ Len.\
	IsValid(i int) bool

	Data() *ArrayData

	// Len returns the number of elements in the array.
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

// IsNull returns true if value at index is null.
// IsNull will panic if NullBitmapBytes is not empty and 0 > i ≥ Len.
func (a *array) IsNull(i int) bool {
	return len(a.nullBitmapBytes) != 0 && bitutil.BitIsNotSet(a.nullBitmapBytes, i)
}

// IsValid returns true if value at index is not null.
// IsValid will panic if NullBitmapBytes is not empty and 0 > i ≥ Len.
func (a *array) IsValid(i int) bool {
	return len(a.nullBitmapBytes) == 0 || bitutil.BitIsSet(a.nullBitmapBytes, i)
}

func (a *array) setData(data *ArrayData) {
	if len(data.buffers) > 0 && data.buffers[0] != nil {
		a.nullBitmapBytes = data.buffers[0].Bytes()
	}
	a.data = data
}
