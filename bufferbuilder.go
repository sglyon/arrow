package arrow

import "github.com/influxdata/arrow/memory"

// bufferBuilder provides common functionality for a buffer builder for populating memory with type-specific values.
// Specialized versions provide type-safe APIs.
type bufferBuilder struct {
	pool     memory.Allocator
	buffer   *memory.PoolBuffer
	length   int
	capacity int

	bytes []byte
}

// Len returns the length of the memory buffer.
func (b *bufferBuilder) Len() int { return b.length }

// Cap returns the total number of elements that can be stored without allocating additional memory.
func (b *bufferBuilder) Cap() int { return b.capacity }

// Bytes returns a slice of length b.Len().
func (b *bufferBuilder) Bytes() []byte { return b.bytes[:b.length] }

func (b *bufferBuilder) resize(elements int) {
	if b.buffer == nil {
		b.buffer = memory.NewPoolBuffer(b.pool)
	}

	b.buffer.Resize(elements)
	oldCapacity := b.capacity
	b.capacity = b.buffer.Cap()
	b.bytes = b.buffer.Buf()

	if b.capacity > oldCapacity {
		memory.Set(b.bytes[oldCapacity:], 0)
	}
}

// Advance increases the buffer by length and initializes the skipped bytes to zero.
func (b *bufferBuilder) Advance(length int) {
	if b.capacity < b.length+length {
		newCapacity := nextPowerOf2(b.length + length)
		b.resize(newCapacity)
	}
	b.length += length
}

func (b *bufferBuilder) Append(data []byte) {
	if b.capacity < b.length+len(data) {
		newCapacity := nextPowerOf2(b.length + len(data))
		b.resize(newCapacity)
	}
	b.UnsafeAppend(data)
}

func (b *bufferBuilder) UnsafeAppend(data []byte) {
	copy(b.bytes[b.length:], data)
	b.length += len(data)
}

func (b *bufferBuilder) Reset() {
	b.buffer, b.bytes = nil, nil
	b.capacity, b.length = 0, 0
}

func (b *bufferBuilder) Finish() *memory.Buffer {
	if b.length > 0 {
		b.buffer.ResizeNoShrink(b.length)
	}
	res := &b.buffer.Buffer
	b.Reset()
	return res
}
