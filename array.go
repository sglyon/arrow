package arrow

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)

const (
	byteBitWidth = 8

	int32Width   = 4
	int64Width   = 8
	float64Width = 8

	startLenBytes        = 0
	startNullCountBytes  = int32Width
	startNullBitmapBytes = int32Width * 2
)

var order binary.ByteOrder = binary.LittleEndian

type Buffer []byte

type array Buffer

func newEmptyArray(size int32, typeWidth int, buf Buffer) (a array) {
	n := int32Width*2 + nullBitmapLen(size) + int(size)*typeWidth
	if buf != nil {
		if cap(buf) < n {
			panic("provided buffer not large enough")
		}
		a = array(buf[:n])
		// Zero null counter and null bitmap
		order.PutUint32(a[startNullCountBytes:startNullBitmapBytes], 0)
		nbm := a.nullBitmap()
		for i := range nbm {
			nbm[i] = 0
		}
	} else {
		a = make(array, n)
	}
	order.PutUint32(a[startLenBytes:startNullCountBytes], uint32(size))
	a.setNullCount(size)
	return
}

func (a array) Len() int32 {
	return int32(order.Uint32(a[startLenBytes:startNullCountBytes]))
}

func (a array) NullCount() int32 {
	return int32(order.Uint32(a[startNullCountBytes:startNullBitmapBytes]))
}
func (a array) setNullCount(c int32) {
	order.PutUint32(a[startNullCountBytes:startNullBitmapBytes], uint32(c))
}

func nullBitmapLen(l int32) int {
	n := int(l / byteBitWidth)
	if l%byteBitWidth != 0 {
		n++
	}
	return n
}

func (a array) nullBitmap() Buffer {
	n := nullBitmapLen(a.Len())
	return Buffer(a[startNullBitmapBytes : startNullBitmapBytes+n])
}

func (a array) values() Buffer {
	n := nullBitmapLen(a.Len())
	return Buffer(a[startNullBitmapBytes+n:])
}

func (a array) IsNull(i int) bool {
	bm := a.nullBitmap()
	b := i / byteBitWidth
	mask := byte(1 << uint8(i%byteBitWidth))
	return bm[b]&mask == 0
}

func (a array) clearNull(i int32) {
	bm := a.nullBitmap()
	b := int(i / byteBitWidth)
	mask := byte(1 << uint8(i%byteBitWidth))
	bm[b] = bm[b] | mask
	a.setNullCount(a.NullCount() - 1)
}

type Float64Array array

func NewEmptyFloat64Array(size int32, buf Buffer) Float64Array {
	return Float64Array(newEmptyArray(size, float64Width, buf))
}

func (a Float64Array) Set(i int32, v float64) {
	start := i * int64Width
	stop := start + int64Width
	array(a).clearNull(i)
	order.PutUint64(array(a).values()[start:stop], math.Float64bits(v))
}

func (a Float64Array) At(i int32) float64 {
	start := i * float64Width
	stop := start + float64Width
	return math.Float64frombits(order.Uint64(array(a).values()[start:stop]))
}

type NullChecker interface {
	IsNull(i int) bool
	NullCount() int32
}

func (a Float64Array) Do(f func([]float64, NullChecker)) {
	values := array(a).values()
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&values))
	header.Len /= float64Width
	header.Cap /= float64Width
	floatValues := *(*[]float64)(unsafe.Pointer(&header))
	f(floatValues, array(a))
}

type Int64Array array

func NewEmptyInt64Array(size int32, buf Buffer) Int64Array {
	return Int64Array(newEmptyArray(size, int64Width, buf))
}

func (a Int64Array) Set(i int32, v int64) {
	start := i * int64Width
	stop := start + int64Width
	array(a).clearNull(i)
	order.PutUint64(array(a).values()[start:stop], uint64(v))
}

func (a Int64Array) At(i int32) int64 {
	start := i * int64Width
	stop := start + int64Width
	return int64(order.Uint64(array(a).values()[start:stop]))
}

func (a Int64Array) Do(f func([]int64, NullChecker)) {
	values := array(a).values()
	header := *(*reflect.SliceHeader)(unsafe.Pointer(&values))
	header.Len /= int64Width
	header.Cap /= int64Width
	intValues := *(*[]int64)(unsafe.Pointer(&header))
	f(intValues, array(a))
}
