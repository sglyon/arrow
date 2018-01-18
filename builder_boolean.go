package arrow

import "github.com/influxdata/arrow/memory"

type BooleanArrayBuilder struct {
	arrayBuilder

	data    *memory.PoolBuffer
	rawData []byte
}

func NewBooleanArrayBuilder(pool memory.Allocator) *BooleanArrayBuilder {
	return &BooleanArrayBuilder{arrayBuilder: arrayBuilder{pool: pool}}
}

//region: append

func (b *BooleanArrayBuilder) Append(v bool) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *BooleanArrayBuilder) AppendByte(v byte) {
	b.Reserve(1)
	b.UnsafeAppend(v != 0)
}

func (b *BooleanArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *BooleanArrayBuilder) UnsafeAppend(v bool) {
	setBit(b.nullBitmap.Bytes(), b.length)
	if v {
		setBit(b.rawData, b.length)
	} else {
		clearBit(b.rawData, b.length)
	}
	b.length++
}

func (b *BooleanArrayBuilder) AppendValues(v []bool, valid []bool) {
	b.Reserve(len(v))
	if len(v) != len(valid) {
		panic("len(v) != len(valid)")
	}

	if len(v) > 0 {
		panic("not implemented")
		//BooleanTraits{}.Copy(b.rawData[b.length:], v)
	}
	b.arrayBuilder.unsafeAppendBoolsToBitmap(valid)
}

//endregion

func (b *BooleanArrayBuilder) Init(capacity int) {
	b.arrayBuilder.init(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := BooleanTraits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = BooleanTraits{}.CastFromBytes(b.data.Bytes())
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
	b.reserve(elements, b.Resize)
}

func (b *BooleanArrayBuilder) Resize(capacity int) {
	if capacity < minBuilderCapacity {
		capacity = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.Init(capacity)
	} else {
		b.arrayBuilder.resize(capacity, b.Init)
		b.data.Resize(BooleanTraits{}.BytesRequired(capacity))
		b.rawData = BooleanTraits{}.CastFromBytes(b.data.Bytes())
	}
}

func (b *BooleanArrayBuilder) Finish() *Float64Array {
	data := b.finishInternal()
	return NewFloat64Array(data)
}

func (b *BooleanArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := BooleanTraits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(PrimitiveTypes.Float64, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)

	*b = BooleanArrayBuilder{arrayBuilder: arrayBuilder{pool: b.pool}} // clear

	return res
}
