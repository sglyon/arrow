package memory

type Buffer struct {
	buf     []byte
	bytes   []byte // slice of buf representing requested size
	mutable bool
}

func NewBuffer(data []byte) *Buffer {
	return &Buffer{bytes: data}
}

func (b *Buffer) Bytes() []byte { return b.bytes }
func (b *Buffer) Mutable() bool { return b.mutable }
func (b *Buffer) Len() int      { return len(b.bytes) }
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
		// retain existing buffer size
		b.bytes = b.buf[:len(b.buf):len(b.buf)]
	}
}

func (b *PoolBuffer) Resize(newSize int) {
	b.resize(newSize, true)
}

func (b *PoolBuffer) ResizeNoShrink(newSize int) {
	b.resize(newSize, false)
}

func (b *PoolBuffer) resize(newSize int, shrink bool) {
	if !shrink || newSize > len(b.buf) {
		b.Reserve(newSize)
	} else {
		newCap := roundUpToMultipleOf64(newSize)
		if newSize == 0 {
			b.pool.Free(b.buf)
			b.buf, b.bytes = nil, nil
		} else {
			b.buf = b.pool.Reallocate(newCap, b.buf)
		}
	}
	b.bytes = b.buf[:newSize:newSize]
}
