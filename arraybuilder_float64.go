package arrow

import (
	"github.com/influxdata/arrow/memory"
)

type Float64ArrayBuilder struct {
	arrayBuilder

	data    *memory.PoolBuffer
	rawData []float64
}

func NewFloat64ArrayBuilder(pool memory.Allocator) *Float64ArrayBuilder {
	return &Float64ArrayBuilder{arrayBuilder: arrayBuilder{pool: pool}}
}

//region: append

func (b *Float64ArrayBuilder) Append(v float64) {
	b.Reserve(1)
	b.UnsafeAppend(v)
}

func (b *Float64ArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.UnsafeAppendBoolToBitmap(false)
}

func (b *Float64ArrayBuilder) UnsafeAppend(v float64) {
	setBit(b.nullBitmap.Bytes(), b.length)
	b.rawData[b.length] = v
	b.length++
}

func (b *Float64ArrayBuilder) UnsafeAppendBoolToBitmap(isValid bool) {
	if isValid {
		setBit(b.nullBitmap.Bytes(), b.length)
	} else {
		b.nullN++
	}
	b.length++
}

func (b *Float64ArrayBuilder) AppendValues(v []float64, valid []bool) {
	b.Reserve(len(v))
	if len(v) != len(valid) {
		panic("len(v) != len(valid)")
	}

	if len(v) > 0 {
		Float64Traits{}.Copy(b.rawData[b.length:], v)
	}
	b.arrayBuilder.unsafeAppendBoolsToBitmap(valid)
}

//endregion

func (b *Float64ArrayBuilder) Init(capacity int) {
	b.arrayBuilder.init(capacity)

	b.data = memory.NewPoolBuffer(b.pool)
	bytesN := Float64Traits{}.BytesRequired(capacity)
	b.data.Resize(bytesN)
	b.rawData = Float64Traits{}.CastFromBytes(b.data.Bytes())
}

// Reserve ensures there is enough space for adding the specified number of elements
// by checking the capacity and calling Resize if necessary.
func (b *Float64ArrayBuilder) Reserve(elements int) {
	b.arrayBuilder.reserve(elements, b.Resize)
}

func (b *Float64ArrayBuilder) Resize(capacity int) {
	if capacity < minBuilderCapacity {
		capacity = minBuilderCapacity
	}

	if b.capacity == 0 {
		b.Init(capacity)
	} else {
		b.arrayBuilder.resize(capacity, b.Init)
		b.data.Resize(Float64Traits{}.BytesRequired(capacity))
		b.rawData = Float64Traits{}.CastFromBytes(b.data.Bytes())
	}
}

func (b *Float64ArrayBuilder) Finish() *Float64Array {
	data := b.finishInternal()
	return NewFloat64Array(data)
}

func (b *Float64ArrayBuilder) finishInternal() *ArrayData {
	bytesRequired := Float64Traits{}.BytesRequired(b.length)
	if bytesRequired > 0 && bytesRequired < b.data.Len() {
		// trim buffers
		b.data.Resize(bytesRequired)
	}
	res := NewArrayData(PrimitiveTypes.Float64, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, &b.data.Buffer}, b.nullN)

	*b = Float64ArrayBuilder{arrayBuilder: arrayBuilder{pool: b.pool}} // clear

	return res
}
