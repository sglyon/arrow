package memory

type Buffer struct {
	buf     []byte
	length  int
	mutable bool
}

func NewBuffer(data []byte) *Buffer {
	return &Buffer{buf: data, length: len(data)}
}

// Buf returns the slice of memory allocated by the Buffer.
func (b *Buffer) Buf() []byte { return b.buf }

// Bytes returns a slice of size Len.
func (b *Buffer) Bytes() []byte { return b.buf[:b.length] }
func (b *Buffer) Mutable() bool { return b.mutable }
func (b *Buffer) Len() int      { return b.length }
func (b *Buffer) Cap() int      { return len(b.buf) }

type PoolBuffer struct {
	Buffer
	pool Allocator
}

func NewPoolBuffer(pool Allocator) *PoolBuffer {
	return &PoolBuffer{pool: pool, Buffer: Buffer{mutable: true}}
}

func (b *PoolBuffer) Reserve(capacity int) {
	if capacity > len(b.buf) {
		newCap := roundUpToMultipleOf64(capacity)
		if len(b.buf) == 0 {
			b.buf = b.pool.Allocate(newCap)
		} else {
			b.buf = b.pool.Reallocate(newCap, b.buf)
		}
	}
}

func (b *PoolBuffer) Resize(newSize int) {
	b.resize(newSize, true)
}

func (b *PoolBuffer) ResizeNoShrink(newSize int) {
	b.resize(newSize, false)
}

func (b *PoolBuffer) resize(newSize int, shrink bool) {
	if !shrink || newSize > b.length {
		b.Reserve(newSize)
	} else {
		// Buffer is not growing, so shrink to the requested size without
		// excess space.
		newCap := roundUpToMultipleOf64(newSize)
		if len(b.buf) != newCap {
			if newSize == 0 {
				b.pool.Free(b.buf)
				b.buf = nil
			} else {
				b.buf = b.pool.Reallocate(newCap, b.buf)
			}
		}
	}
	b.length = newSize
}
