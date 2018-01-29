package arrow

import (
	"math"

	"github.com/influxdata/arrow/memory"
)

const (
	binaryArrayMaximumCapacity = math.MaxInt32
)

// A BinaryArrayBuilder is used to build a BinaryArray using the Append methods.
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

func (b *BinaryArrayBuilder) init(capacity int) {
	b.arrayBuilder.init(capacity)
	b.offsets.resize((capacity + 1) * Int32SizeBytes)
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *BinaryArrayBuilder) Reserve(n int) {
	b.arrayBuilder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *BinaryArrayBuilder) Resize(n int) {
	b.offsets.resize((n + 1) * Int32SizeBytes)
	b.arrayBuilder.resize(n, b.init)
}

// Finish completes the transfers ownership of the buffers used to build the arrow
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
