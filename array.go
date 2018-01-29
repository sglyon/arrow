package arrow

import (
	"github.com/influxdata/arrow/internal/bitutil"
)

// Array represents an immutable sequence of values.
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
func (a *array) DataType() DataType { return a.data.typE }

// NullN returns the number of null values in the array.
func (a *array) NullN() int { return a.data.nullN }

// NullBitmapBytes returns a byte slice of the validity bitmap.
func (a *array) NullBitmapBytes() []byte { return a.nullBitmapBytes }

func (a *array) Data() *ArrayData { return a.data }

// Len returns the number of elements in the array.
func (a *array) Len() int { return a.data.length }

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

type arrayConstructorFn func(*ArrayData) Array

var (
	makeArrayFn = [...]arrayConstructorFn{
		NULL:              unsupportedArrayType,
		BOOL:              func(data *ArrayData) Array { return NewBooleanArray(data) },
		UINT8:             unsupportedArrayType,
		INT8:              unsupportedArrayType,
		UINT16:            unsupportedArrayType,
		INT16:             unsupportedArrayType,
		UINT32:            unsupportedArrayType,
		INT32:             func(data *ArrayData) Array { return NewInt32Array(data) },
		UINT64:            func(data *ArrayData) Array { return NewUint64Array(data) },
		INT64:             func(data *ArrayData) Array { return NewInt64Array(data) },
		HALF_FLOAT:        unsupportedArrayType,
		FLOAT32:           unsupportedArrayType,
		FLOAT64:           func(data *ArrayData) Array { return NewFloat64Array(data) },
		STRING:            unsupportedArrayType,
		BINARY:            func(data *ArrayData) Array { return NewBinaryArray(data) },
		FIXED_SIZE_BINARY: unsupportedArrayType,
		DATE32:            unsupportedArrayType,
		DATE64:            unsupportedArrayType,
		TIMESTAMP:         func(data *ArrayData) Array { return NewTimestampArray(data) },
		TIME32:            unsupportedArrayType,
		TIME64:            unsupportedArrayType,
		INTERVAL:          unsupportedArrayType,
		DECIMAL:           unsupportedArrayType,
		LIST:              unsupportedArrayType,
		STRUCT:            unsupportedArrayType,
		UNION:             unsupportedArrayType,
		DICTIONARY:        unsupportedArrayType,
		MAP:               unsupportedArrayType,

		// invalid data types to fill out remaining
		28: invalidDataType,
		29: invalidDataType,
		30: invalidDataType,
		31: invalidDataType,
	}
)

func unsupportedArrayType(data *ArrayData) Array {
	panic("unsupported data type: " + data.typE.ID().String())
}

func invalidDataType(data *ArrayData) Array {
	panic("invalid data type: " + data.typE.ID().String())
}

// MakeArray constructs a strongly-typed Array instance from generic ArrayData.
func MakeArray(data *ArrayData) Array {
	return makeArrayFn[byte(data.typE.ID()&0x1f)](data)
}
