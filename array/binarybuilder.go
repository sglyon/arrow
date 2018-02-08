package array

import (
	"math"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/memory"
)

const (
	binaryArrayMaximumCapacity = math.MaxInt32
)

// A BinaryBuilder is used to build a Binary array using the Append methods.
type BinaryBuilder struct {
	builder

	typE    arrow.BinaryDataType
	offsets int32BufferBuilder
	values  byteBufferBuilder
}

func NewBinaryBuilder(mem memory.Allocator, typE arrow.BinaryDataType) *BinaryBuilder {
	b := &BinaryBuilder{typE: typE}
	b.builder.mem = mem
	b.offsets.mem = mem
	b.values.mem = mem
	return b
}

func (b *BinaryBuilder) Append(v []byte) {
	b.Reserve(1)
	b.appendNextOffset()
	b.values.Append(v)
	b.UnsafeAppendBoolToBitmap(true)
}

func (b *BinaryBuilder) AppendString(v string) {
	b.Append([]byte(v))
}

func (b *BinaryBuilder) AppendNull() {
	b.Reserve(1)
	b.appendNextOffset()
	b.UnsafeAppendBoolToBitmap(false)
}

// AppendValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *BinaryBuilder) AppendValues(v [][]byte, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	for _, vv := range v {
		b.appendNextOffset()
		b.values.Append(vv)
	}

	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

// AppendStringValues will append the values in the v slice. The valid slice determines which values
// in v are valid (not null). The valid slice must either be empty or be equal in length to v. If empty,
// all values in v are appended and considered valid.
func (b *BinaryBuilder) AppendStringValues(v []string, valid []bool) {
	if len(v) != len(valid) && len(valid) != 0 {
		panic("len(v) != len(valid) && len(valid) != 0")
	}

	b.Reserve(len(v))
	for _, vv := range v {
		b.appendNextOffset()
		b.values.Append([]byte(vv))
	}

	b.builder.unsafeAppendBoolsToBitmap(valid, len(v))
}

func (b *BinaryBuilder) Value(i int) []byte {
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

func (b *BinaryBuilder) init(capacity int) {
	b.builder.init(capacity)
	b.offsets.resize((capacity + 1) * arrow.Int32SizeBytes)
}

// Reserve ensures there is enough space for appending n elements
// by checking the capacity and calling Resize if necessary.
func (b *BinaryBuilder) Reserve(n int) {
	b.builder.reserve(n, b.Resize)
}

// Resize adjusts the space allocated by b to n elements. If n is greater than b.Cap(),
// additional memory will be allocated. If n is smaller, the allocated memory may reduced.
func (b *BinaryBuilder) Resize(n int) {
	b.offsets.resize((n + 1) * arrow.Int32SizeBytes)
	b.builder.resize(n, b.init)
}

// Finish completes the transfers ownership of the buffers used to build the arrow
func (b *BinaryBuilder) Finish() *Binary {
	data := b.finishInternal()
	return NewBinaryData(data)
}

func (b *BinaryBuilder) finishInternal() *Data {
	b.appendNextOffset()
	offsets, values := b.offsets.Finish(), b.values.Finish()
	res := NewData(b.typE, b.length, []*memory.Buffer{b.nullBitmap, offsets, values}, b.nullN)

	b.builder.reset()

	return res
}

func (b *BinaryBuilder) appendNextOffset() {
	numBytes := b.values.Len()
	// TODO(sgc): check binaryArrayMaximumCapacity?
	b.offsets.AppendValue(int32(numBytes))
}
