package arrow

import "github.com/influxdata/arrow/memory"

type ByteBufferBuilder struct {
	bufferBuilder
}

func NewByteBufferBuilder(pool memory.Allocator) *ByteBufferBuilder {
	return &ByteBufferBuilder{bufferBuilder: bufferBuilder{pool: pool}}
}

func (b *ByteBufferBuilder) Values() []byte   { return b.Bytes() }
func (b *ByteBufferBuilder) Value(i int) byte { return b.bytes[i] }
