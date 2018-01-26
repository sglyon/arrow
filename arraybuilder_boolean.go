package arrow

import (
	"github.com/influxdata/arrow/internal/bitutil"
	"github.com/influxdata/arrow/memory"
)

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
	bitutil.SetBit(b.nullBitmap.Bytes(), b.length)
	if v {
		bitutil.SetBit(b.rawData, b.length)
	} else {
		bitutil.ClearBit(b.rawData, b.length)
	}
	b.length++
}

func (b *BooleanArrayBuilder) AppendValues(v []bool, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	for i, vv := range v {
		bitutil.SetBitTo(b.rawData, b.length+i, vv)
	}
	b.arrayBuilder.unsafeAppendBoolsToBitmap(valid, len(v))
}

//endregion

func (b *BooleanArrayBuilder) Init(capacity int) {
	b.arrayBuilder.init(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := BooleanTraits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = BooleanTraits{}.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for adding the specified number of elements
// by checking the capacity and calling Resize if necessary.
func (b *BooleanArrayBuilder) Reserve(elements int) {
	b.arrayBuilder.reserve(elements, b.Resize)
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

func (b *BooleanArrayBuilder) Finish() *BooleanArray {
	data := b.finishInternal()
	return NewBooleanArray(data)
}

func (b *BooleanArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := BooleanTraits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(FixedWidthTypes.Boolean, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)
	b.reset()

	return res
}
