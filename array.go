package arrow

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)

const (
	byteBitWidth   = 8
	bytePaddingLen = 8

	int32Width   = 4
	int64Width   = 8
	float64Width = 8

	startLenBytes        = 0
	startNullCountBytes  = int32Width
	startNullBitmapBytes = int32Width * 2
)

var order binary.ByteOrder = binary.LittleEndian

type Buffer []byte

// newBuffer allocates a new buffer for the given size, ensuring the buffer is padded to 8 bytes.
func newBuffer(size int) Buffer {
	n := size
	remainder := (size % bytePaddingLen)
	if remainder > 0 {
		n += bytePaddingLen - remainder
	}
	//TODO: Possibly allocate in larger 64 byte aligned []byte
	buf := make([]byte, n)
	return buf[:size]
}

type array struct {
	len       int32
	nullCount int32

	nullBitmap Buffer
	values     Buffer
}

func newEmptyArray(size int32, typeWidth int) (a array) {
	a.len = size
	a.nullCount = size
	a.nullBitmap = newBuffer(nullBitmapLen(size))
	a.values = newBuffer(int(size) * typeWidth)
	return
}

func (a *array) Len() int32 {
	return a.len
}

func (a *array) NullCount() int32 {
	return a.nullCount
}

func nullBitmapLen(l int32) int {
	n := int(l / byteBitWidth)
	if l%byteBitWidth != 0 {
		n++
	}
	return n
}

var bitmask = [8]byte{1, 2, 4, 8, 16, 32, 64, 128}

func (a *array) IsNull(i int) bool {
	b := i / byteBitWidth
	// Use a simple lookup instead of bit shifting.
	mask := bitmask[int(i%byteBitWidth)]
	return a.nullBitmap[b]&mask == 0
}

func (a *array) clearNullBit(i int32) {
	b := int(i / byteBitWidth)
	// Use a simple lookup instead of bit shifting.
	mask := bitmask[int(i%byteBitWidth)]
	if a.nullBitmap[b]&mask == 0 {
		a.nullCount--
		a.nullBitmap[b] = a.nullBitmap[b] | mask
	}
}

type Float64Array struct {
	array
}

func NewEmptyFloat64Array(size int32) *Float64Array {
	return &Float64Array{array: newEmptyArray(size, float64Width)}
}

func (a *Float64Array) Set(i int32, v float64) {
	start := i * float64Width
	stop := start + float64Width
	a.clearNullBit(i)
	order.PutUint64(a.values[start:stop], math.Float64bits(v))
}

func (a *Float64Array) At(i int32) float64 {
	start := i * float64Width
	stop := start + float64Width
	return math.Float64frombits(order.Uint64(a.values[start:stop]))
}

type NullChecker interface {
	IsNull(i int) bool
	NullCount() int32
}

func (a *Float64Array) Do(f func([]float64, NullChecker)) {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&a.values))
	header.Len /= float64Width
	header.Cap /= float64Width
	floatValues := *(*[]float64)(unsafe.Pointer(&header))
	f(floatValues, a)
}

type Int64Array struct {
	array
}

func NewEmptyInt64Array(size int32) *Int64Array {
	return &Int64Array{array: newEmptyArray(size, int64Width)}
}

func (a *Int64Array) Set(i int32, v int64) {
	start := i * int64Width
	stop := start + int64Width
	a.clearNullBit(i)
	order.PutUint64(a.values[start:stop], uint64(v))
}

func (a *Int64Array) At(i int32) int64 {
	start := i * int64Width
	stop := start + int64Width
	return int64(order.Uint64(a.values[start:stop]))
}

func (a *Int64Array) Do(f func([]int64, NullChecker)) {
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&a.values))
	header.Len /= int64Width
	header.Cap /= int64Width
	intValues := *(*[]int64)(unsafe.Pointer(&header))
	f(intValues, a)
}
