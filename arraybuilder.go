package arrow

import "github.com/influxdata/arrow/memory"

const (
	minBuilderCapacity = 1 << 5
)

type arrayBuilder struct {
	pool       memory.Allocator
	nullBitmap *memory.PoolBuffer
	nullN      int
	length     int
	capacity   int
}

func (b *arrayBuilder) init(capacity int) {
	toAlloc := ceilByte(capacity) / 8
	b.nullBitmap = memory.NewPoolBuffer(b.pool)
	b.nullBitmap.Resize(toAlloc)
	b.capacity = capacity
	memory.Set(b.nullBitmap.Bytes(), 0)
}

func (b *arrayBuilder) resize(newBits int, init func(int)) {
	if b.nullBitmap == nil {
		init(newBits)
		return
	}

	newBytesN := ceilByte(newBits) / 8
	oldBytesN := b.nullBitmap.Len()
	b.nullBitmap.Resize(newBytesN)
	b.capacity = newBits
	if oldBytesN < newBytesN {
		// TODO(sgc): necessary?
		memory.Set(b.nullBitmap.Bytes()[oldBytesN:], 0)
	}
}

func (b *arrayBuilder) reserve(elements int, resize func(int)) {
	if b.length+elements > b.capacity {
		newCap := nextPowerOf2(b.length + elements)
		resize(newCap)
	}
}

func (b *arrayBuilder) unsafeAppendBoolsToBitmap(valid []bool) {
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
			bitSet |= bitMask[bitOffset]
		} else {
			bitSet &= flippedBitMask[bitOffset]
			b.nullN++
		}
		bitOffset++
	}

	if bitOffset != 0 {
		nullBitmap[byteOffset] = bitSet
	}
	b.length += len(valid)
}

// unsafeSetValid sets the next length bits to valid (in the validity bitmap.
func (b *arrayBuilder) unsafeSetValid(length int) {
	padToByte := min(8-(b.length%8), length)
	if padToByte == 8 {
		padToByte = 0
	}
	bits := b.nullBitmap.Bytes()
	for i := b.length; i < b.length+padToByte; i++ {
		setBit(bits, i)
	}

	start := (length + padToByte) / 8
	fastLength := (length - padToByte) / 8
	memory.Set(bits[start:start+fastLength], 0xff)

	newLength := b.length + length
	// trailing bytes
	for i := b.length + padToByte + (fastLength * 8); i < newLength; i++ {
		setBit(bits, i)
	}

	b.length = newLength
}

func (b *arrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		setBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}
