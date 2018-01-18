package arrow

import "github.com/influxdata/arrow/memory"

type BooleanArrayBuilder struct {
	pool       memory.Allocator
	nullBitmap *memory.PoolBuffer
	nullN      int
	length     int
	capacity   int

	data    *memory.PoolBuffer
	rawData []float64
}

func NewBooleanArrayBuilder(pool memory.Allocator) *BooleanArrayBuilder {
	return &BooleanArrayBuilder{pool: pool}
}

//region: append

func (b *BooleanArrayBuilder) arrayBuilderUnsafeAppendBoolsToBitmap(valid []bool) {
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

func (b *BooleanArrayBuilder) Append(v float64) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *BooleanArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *BooleanArrayBuilder) UnsafeAppend(v float64) {
	setBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *BooleanArrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		setBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

func (b *BooleanArrayBuilder) AppendValues(v []float64, valid []bool) {
	b.Reserve(len(v))
	if len(v) != len(valid) {
		panic("len(v) != len(valid)")
	}

	if len(v) > 0 {
		Float64Traits{}.Copy(b.rawData[b.length:], v)
	}
	b.arrayBuilderUnsafeAppendBoolsToBitmap(valid)
}

//endregion

func (b *BooleanArrayBuilder) arrayBuilderInit(capacity int) {
	toAlloc := ceilByte(capacity) / 8
	b.nullBitmap = memory.NewPoolBuffer(b.pool)
	b.nullBitmap.Resize(toAlloc)
	b.capacity = capacity
	memory.Set(b.nullBitmap.Bytes(), 0)
}

func (b *BooleanArrayBuilder) Init(capacity int) {
	b.arrayBuilderInit(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := Float64Traits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = Float64Traits{}.CastFromBytes(b.data.Bytes())
}

func (b *BooleanArrayBuilder) arrayBuilderReserve(elements int) {
	if b.length+elements > b.capacity {
		newCap := nextPowerOf2(b.length + elements)
		b.Resize(newCap)
	}
}

// Reserve ensures there is enough space for adding the specified number of elements
// by checking the capacity and calling Resize if necessary.
func (b *BooleanArrayBuilder) Reserve(elements int) {
	b.arrayBuilderReserve(elements)
}

func (b *BooleanArrayBuilder) arrayBuilderResize(newBits int) {
	if b.nullBitmap == nil {
		b.Init(newBits)
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

func (b *BooleanArrayBuilder) Resize(capacity int) {
	if capacity < minBuilderCapacity {
		capacity = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.Init(capacity)
	} else {
		b.arrayBuilderResize(capacity)
		b.data.Resize(Float64Traits{}.BytesRequired(capacity))
		b.rawData = Float64Traits{}.CastFromBytes(b.data.Bytes())
	}
}

func (b *BooleanArrayBuilder) Finish() *Float64Array {
	data := b.finishInternal()
	return NewFloat64Array(data)
}

func (b *BooleanArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := Float64Traits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(PrimitiveTypes.Float64, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)

	*b = BooleanArrayBuilder{pool: b.pool} // clear

	return res
}
