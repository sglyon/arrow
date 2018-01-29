package array

import "github.com/influxdata/arrow/memory"

type ByteBufferBuilder struct {
	bufferBuilder
}

func NewByteBufferBuilder(mem memory.Allocator) *ByteBufferBuilder {
	return &ByteBufferBuilder{bufferBuilder: bufferBuilder{mem: mem}}
}

func (b *ByteBufferBuilder) Values() []byte   { return b.Bytes() }
func (b *ByteBufferBuilder) Value(i int) byte { return b.bytes[i] }
