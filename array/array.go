package array

import (
	"sync/atomic"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/internal/bitutil"
	"github.com/influxdata/arrow/internal/debug"
)

// A type which satisfies array.Interface represents an immutable sequence of values.
type Interface interface {
	// DataType returns the type metadata for this instance.
	DataType() arrow.DataType

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

	Data() *Data

	// Len returns the number of elements in the array.
	Len() int

	// Retain increases the reference count by 1.
	// Retain may be called simultaneously from multiple goroutines.
	Retain()

	// Release decreases the reference count by 1.
	// Release may be called simultaneously from multiple goroutines.
	// When the reference count goes to zero, the memory is freed.
	Release()
}

const (
	// UnknownNullCount specifies the NullN should be calculated from the null bitmap buffer.
	UnknownNullCount = -1
)

type array struct {
	refCount        int64
	data            *Data
	nullBitmapBytes []byte
}

// Retain increases the reference count by 1.
// Retain may be called simultaneously from multiple goroutines.
func (a *array) Retain() {
	atomic.AddInt64(&a.refCount, 1)
}

// Release decreases the reference count by 1.
// Release may be called simultaneously from multiple goroutines.
// When the reference count goes to zero, the memory is freed.
func (a *array) Release() {
	debug.Assert(atomic.LoadInt64(&a.refCount) > 0, "too many releases")

	if atomic.AddInt64(&a.refCount, -1) == 0 {
		a.data.Release()
		a.data, a.nullBitmapBytes = nil, nil
	}
}

// DataType returns the type metadata for this instance.
func (a *array) DataType() arrow.DataType { return a.data.typE }

// NullN returns the number of null values in the array.
func (a *array) NullN() int {
	if a.data.nullN < 0 {
		a.data.nullN = a.data.length - bitutil.CountSetBits(a.nullBitmapBytes, a.data.length)
	}
	return a.data.nullN
}

// NullBitmapBytes returns a byte slice of the validity bitmap.
func (a *array) NullBitmapBytes() []byte { return a.nullBitmapBytes }

func (a *array) Data() *Data { return a.data }

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

func (a *array) setData(data *Data) {
	if a.data != nil {
		a.data.Release()
	}

	data.Retain()
	if len(data.buffers) > 0 && data.buffers[0] != nil {
		a.nullBitmapBytes = data.buffers[0].Bytes()
	}
	a.data = data
}

type arrayConstructorFn func(*Data) Interface

var (
	makeArrayFn = [...]arrayConstructorFn{
		arrow.NULL:              unsupportedArrayType,
		arrow.BOOL:              func(data *Data) Interface { return NewBooleanData(data) },
		arrow.UINT8:             func(data *Data) Interface { return NewUint8Data(data) },
		arrow.INT8:              func(data *Data) Interface { return NewInt8Data(data) },
		arrow.UINT16:            func(data *Data) Interface { return NewUint16Data(data) },
		arrow.INT16:             func(data *Data) Interface { return NewInt16Data(data) },
		arrow.UINT32:            func(data *Data) Interface { return NewUint32Data(data) },
		arrow.INT32:             func(data *Data) Interface { return NewInt32Data(data) },
		arrow.UINT64:            func(data *Data) Interface { return NewUint64Data(data) },
		arrow.INT64:             func(data *Data) Interface { return NewInt64Data(data) },
		arrow.HALF_FLOAT:        unsupportedArrayType,
		arrow.FLOAT32:           func(data *Data) Interface { return NewFloat32Data(data) },
		arrow.FLOAT64:           func(data *Data) Interface { return NewFloat64Data(data) },
		arrow.STRING:            unsupportedArrayType,
		arrow.BINARY:            func(data *Data) Interface { return NewBinaryData(data) },
		arrow.FIXED_SIZE_BINARY: unsupportedArrayType,
		arrow.DATE32:            unsupportedArrayType,
		arrow.DATE64:            unsupportedArrayType,
		arrow.TIMESTAMP:         func(data *Data) Interface { return NewTimestampData(data) },
		arrow.TIME32:            unsupportedArrayType,
		arrow.TIME64:            unsupportedArrayType,
		arrow.INTERVAL:          unsupportedArrayType,
		arrow.DECIMAL:           unsupportedArrayType,
		arrow.LIST:              unsupportedArrayType,
		arrow.STRUCT:            unsupportedArrayType,
		arrow.UNION:             unsupportedArrayType,
		arrow.DICTIONARY:        unsupportedArrayType,
		arrow.MAP:               unsupportedArrayType,

		// invalid data types to fill out array size 2⁵-1
		28: invalidDataType,
		29: invalidDataType,
		30: invalidDataType,
		31: invalidDataType,
	}
)

func unsupportedArrayType(data *Data) Interface {
	panic("unsupported data type: " + data.typE.ID().String())
}

func invalidDataType(data *Data) Interface {
	panic("invalid data type: " + data.typE.ID().String())
}

// MakeFromData constructs a strongly-typed array instance from generic Data.
func MakeFromData(data *Data) Interface {
	return makeArrayFn[byte(data.typE.ID()&0x1f)](data)
}

type dictArrayConstructorFn func(*Data, *Data) Interface

var (
	makeDictArrayFn = [...]dictArrayConstructorFn{
		arrow.NULL: unsupportedDictArrayType,
		// arrow.BOOL:              func(data, poolData *Data) Interface { return NewBooleanDictData(data, poolData) },
		arrow.BOOL:       unsupportedDictArrayType,
		arrow.UINT8:      func(data, poolData *Data) Interface { return NewUint8DictData(data, poolData) },
		arrow.INT8:       func(data, poolData *Data) Interface { return NewInt8DictData(data, poolData) },
		arrow.UINT16:     func(data, poolData *Data) Interface { return NewUint16DictData(data, poolData) },
		arrow.INT16:      func(data, poolData *Data) Interface { return NewInt16DictData(data, poolData) },
		arrow.UINT32:     func(data, poolData *Data) Interface { return NewUint32DictData(data, poolData) },
		arrow.INT32:      func(data, poolData *Data) Interface { return NewInt32DictData(data, poolData) },
		arrow.UINT64:     func(data, poolData *Data) Interface { return NewUint64DictData(data, poolData) },
		arrow.INT64:      func(data, poolData *Data) Interface { return NewInt64DictData(data, poolData) },
		arrow.HALF_FLOAT: unsupportedDictArrayType,
		arrow.FLOAT32:    func(data, poolData *Data) Interface { return NewFloat32DictData(data, poolData) },
		arrow.FLOAT64:    func(data, poolData *Data) Interface { return NewFloat64DictData(data, poolData) },
		arrow.STRING:     unsupportedDictArrayType,
		// arrow.BINARY:            func(data, poolData *Data) Interface { return NewBinaryDictData(data, poolData) },
		arrow.BINARY:            unsupportedDictArrayType,
		arrow.FIXED_SIZE_BINARY: unsupportedDictArrayType,
		arrow.DATE32:            unsupportedDictArrayType,
		arrow.DATE64:            unsupportedDictArrayType,
		arrow.TIMESTAMP:         func(data, poolData *Data) Interface { return NewTimestampDictData(data, poolData) },
		arrow.TIME32:            unsupportedDictArrayType,
		arrow.TIME64:            unsupportedDictArrayType,
		arrow.INTERVAL:          unsupportedDictArrayType,
		arrow.DECIMAL:           unsupportedDictArrayType,
		arrow.LIST:              unsupportedDictArrayType,
		arrow.STRUCT:            unsupportedDictArrayType,
		arrow.UNION:             unsupportedDictArrayType,
		arrow.DICTIONARY:        unsupportedDictArrayType,
		arrow.MAP:               unsupportedDictArrayType,

		// invalid data types to fill out array size 2⁵-1
		28: invalidDictDataType,
		29: invalidDictDataType,
		30: invalidDictDataType,
		31: invalidDictDataType,
	}
)

func unsupportedDictArrayType(data, poolData *Data) Interface {
	panic("unsupported data type: " + poolData.typE.ID().String())
}

func invalidDictDataType(data, poolData *Data) Interface {
	panic("invalid data type: " + poolData.typE.ID().String())
}

// MakeFromData constructs a strongly-typed array instance from generic Data.
func MakeDictFromData(data, poolData *Data) Interface {
	return makeDictArrayFn[byte(poolData.typE.ID()&0x1f)](data, poolData)
}
