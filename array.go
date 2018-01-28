package arrow

import (
	"github.com/influxdata/arrow/internal/bitutil"
)

// Array is the most basic data structure and represents an immutable sequence of values.
type Array interface {
	// DataType returns the type metadata for this instance.
	DataType() DataType

	// NullN returns the number of null values in the array.
	NullN() int

	// NullBitmapBytes returns a byte slice of the validity bitmap.
	NullBitmapBytes() []byte

	// IsNull returns true if value at index is null.
	// NOTE: IsNull will panic if NullBitmapBytes is not empty and 0 > i ≥ Len.
	IsNull(i int) bool

	// IsValid returns true if value at index is not null.
	// NOTE: IsValid will panic if NullBitmapBytes is not empty and 0 > i ≥ Len.
	IsValid(i int) bool

	Data() *ArrayData

	// Len returns the number of elements in the array.
	Len() int
}

type array struct {
	data            *ArrayData
	nullBitmapBytes []byte
}

// DataType returns the type metadata for this instance.
func (a *array) DataType() DataType      { return a.data.typE }
func (a *array) NullN() int              { return a.data.nullN }
func (a *array) NullBitmapBytes() []byte { return a.nullBitmapBytes }
func (a *array) Data() *ArrayData        { return a.data }
func (a *array) Len() int                { return a.data.length }

// IsNull returns true if value at index is null.
// NOTE: IsNull will panic if NullBitmapBytes is not empty and 0 > i ≥ Len.
func (a *array) IsNull(i int) bool {
	return len(a.nullBitmapBytes) != 0 && bitutil.BitIsNotSet(a.nullBitmapBytes, i)
}

// IsValid returns true if value at index is not null.
// NOTE: IsValid will panic if NullBitmapBytes is not empty and 0 > i ≥ Len.
func (a *array) IsValid(i int) bool {
	return len(a.nullBitmapBytes) == 0 || bitutil.BitIsSet(a.nullBitmapBytes, i)
}

func (a *array) setData(data *ArrayData) {
	if len(data.buffers) > 0 && data.buffers[0] != nil {
		a.nullBitmapBytes = data.buffers[0].Bytes()
	}
	a.data = data
}
