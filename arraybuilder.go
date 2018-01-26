package arrow

import (
	"github.com/influxdata/arrow/internal/bitutil"
	"github.com/influxdata/arrow/memory"
)

const (
	minBuilderCapacity = 1 << 5
)

// arrayBuilder provides common functionality to array builders for managing the validity bitmap (nulls).
type arrayBuilder struct {
	pool       memory.Allocator
	nullBitmap *memory.PoolBuffer
	nullN      int
	length     int
	capacity   int

	// TODO(sgc): store DataType per C++ implementation?
}

// Len returns the length of the array.
func (b *arrayBuilder) Len() int { return b.length }

// Cap returns the total number of elements that can be stored without allocating additional memory.
func (b *arrayBuilder) Cap() int { return b.capacity }

// NullN returns the number of null values in the array.
func (b *arrayBuilder) NullN() int { return b.nullN }

func (b *arrayBuilder) init(capacity int) {
	toAlloc := bitutil.CeilByte(capacity) / 8
	b.nullBitmap = memory.NewPoolBuffer(b.pool)
	b.nullBitmap.Resize(toAlloc)
	b.capacity = capacity
	memory.Set(b.nullBitmap.Buf(), 0)
}

func (b *arrayBuilder) reset() {
	b.nullBitmap = nil
	b.nullN = 0
	b.length = 0
	b.capacity = 0
}

func (b *arrayBuilder) resize(newBits int, init func(int)) {
	if b.nullBitmap == nil {
		init(newBits)
		return
	}

	newBytesN := bitutil.CeilByte(newBits) / 8
	oldBytesN := b.nullBitmap.Len()
	b.nullBitmap.Resize(newBytesN)
	b.capacity = newBits
	if oldBytesN < newBytesN {
		// TODO(sgc): necessary?
		memory.Set(b.nullBitmap.Buf()[oldBytesN:], 0)
	}
}

func (b *arrayBuilder) reserve(elements int, resize func(int)) {
	if b.length+elements > b.capacity {
		newCap := bitutil.NextPowerOf2(b.length + elements)
		resize(newCap)
	}
}

// unsafeAppendBoolsToBitmap appends the contents of valid to the validity bitmap.
// As an optimization, if the valid slice is empty, the next length bits will be set to valid (not null).
func (b *arrayBuilder) unsafeAppendBoolsToBitmap(valid []bool, length int) {
	if len(valid) == 0 {
		b.unsafeSetValid(length)
		return
	}

	byteOffset := b.length / 8
	bitOffset := byte(b.length % 8)
	nullBitmap := b.nullBitmap.Bytes()
	bitSet := nullBitmap[byteOffset]

	for _, v := range valid {
		if bitOffset == 8 {
			bitOffset = 0
			nullBitmap[byteOffset] = bitSet
			byteOffset++
			bitSet = nullBitmap[byteOffset]
		}

		if v {
			bitSet |= bitutil.BitMask[bitOffset]
		} else {
			bitSet &= bitutil.FlippedBitMask[bitOffset]
			b.nullN++
		}
		bitOffset++
	}

	if bitOffset != 0 {
		nullBitmap[byteOffset] = bitSet
	}
	b.length += len(valid)
}

// unsafeSetValid sets the next length bits to valid in the validity bitmap.
func (b *arrayBuilder) unsafeSetValid(length int) {
	padToByte := min(8-(b.length%8), length)
	if padToByte == 8 {
		padToByte = 0
	}
	bits := b.nullBitmap.Bytes()
	for i := b.length; i < b.length+padToByte; i++ {
		bitutil.SetBit(bits, i)
	}

	start := (b.length + padToByte) / 8
	fastLength := (length - padToByte) / 8
	memory.Set(bits[start:start+fastLength], 0xff)

	newLength := b.length + length
	// trailing bytes
	for i := b.length + padToByte + (fastLength * 8); i < newLength; i++ {
		bitutil.SetBit(bits, i)
	}

	b.length = newLength
}

func (b *arrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}
