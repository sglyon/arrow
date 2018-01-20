package arrow

import (
	"math"

	"github.com/influxdata/arrow/memory"
)

const (
	binaryArrayMaximumCapacity = math.MaxInt32
)

type BinaryArrayBuilder struct {
	arrayBuilder

	offsets Int32BufferBuilder
	values  ByteBufferBuilder
}

func NewBinaryArrayBuilder(pool memory.Allocator) *BinaryArrayBuilder {
	b := &BinaryArrayBuilder{}
	b.arrayBuilder.pool = pool
	b.offsets.pool = pool
	b.values.pool = pool
	return b
}

//region: append

func (b *BinaryArrayBuilder) Append(v []byte) {
	b.Reserve(1)
	b.appendNextOffset()
	b.values.Append(v)
	b.UnsafeAppendBoolToBitmap(true)
}

func (b *BinaryArrayBuilder) AppendNull() {
	b.Reserve(1)
	b.appendNextOffset()
	b.UnsafeAppendBoolToBitmap(false)
}

//endregion

func (b *BinaryArrayBuilder) Value(i int) []byte {
	offsets := b.offsets.Values()
	start := int(offsets[i])
	var end int
	if i == (b.length - 1) {
		end = b.values.Len()
	} else {
		end = int(offsets[i+1])
	}
	return b.values.Bytes()[start:end]
}

func (b *BinaryArrayBuilder) Init(capacity int) {
	b.arrayBuilder.init(capacity)
	b.offsets.resize((capacity + 1) * int32SizeBytes)
}

// Reserve ensures there is enough space for adding the specified number of elements
// by checking the capacity and calling Resize if necessary.
func (b *BinaryArrayBuilder) Reserve(elements int) {
	b.arrayBuilder.reserve(elements, b.Resize)
}

func (b *BinaryArrayBuilder) Resize(capacity int) {
	b.offsets.resize((capacity + 1) * int32SizeBytes)
	b.arrayBuilder.resize(capacity, b.Init)
}

func (b *BinaryArrayBuilder) Finish() *BinaryArray {
	data := b.finishInternal()
	return NewBinaryArray(data)
}

func (b *BinaryArrayBuilder) finishInternal() *ArrayData {
	b.appendNextOffset()
	offsets, values := b.offsets.Finish(), b.values.Finish()
	res := NewArrayData(BinaryTypes.Binary, b.length, []*memory.Buffer{&b.nullBitmap.Buffer, offsets, values}, b.nullN)

	b.arrayBuilder.reset()

	return res
}

func (b *BinaryArrayBuilder) appendNextOffset() {
	numBytes := b.values.Len()
	// TODO(sgc): check binaryArrayMaximumCapacity?
	b.offsets.AppendValue(int32(numBytes))
}
