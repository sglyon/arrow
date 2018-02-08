package memory

type Buffer struct {
	buf     []byte
	length  int
	mutable bool
	mem     Allocator
}

// NewBufferBytes creates a fixed-size buffer from the specified data.
func NewBufferBytes(data []byte) *Buffer {
	return &Buffer{buf: data, length: len(data)}
}

// NewBuffer creates a mutable, resizable buffer with an Allocator for managing memory.
func NewResizableBuffer(mem Allocator) *Buffer {
	return &Buffer{mutable: true, mem: mem}
}

// Buf returns the slice of memory allocated by the Buffer.
func (b *Buffer) Buf() []byte { return b.buf }

// Bytes returns a slice of size Len.
func (b *Buffer) Bytes() []byte { return b.buf[:b.length] }
func (b *Buffer) Mutable() bool { return b.mutable }
func (b *Buffer) Len() int      { return b.length }
func (b *Buffer) Cap() int      { return len(b.buf) }

func (b *Buffer) Reserve(capacity int) {
	if capacity > len(b.buf) {
		newCap := roundUpToMultipleOf64(capacity)
		if len(b.buf) == 0 {
			b.buf = b.mem.Allocate(newCap)
		} else {
			b.buf = b.mem.Reallocate(newCap, b.buf)
		}
	}
}

func (b *Buffer) Resize(newSize int) {
	b.resize(newSize, true)
}

func (b *Buffer) ResizeNoShrink(newSize int) {
	b.resize(newSize, false)
}

func (b *Buffer) resize(newSize int, shrink bool) {
	if !shrink || newSize > b.length {
		b.Reserve(newSize)
	} else {
		// Buffer is not growing, so shrink to the requested size without
		// excess space.
		newCap := roundUpToMultipleOf64(newSize)
		if len(b.buf) != newCap {
			if newSize == 0 {
				b.mem.Free(b.buf)
				b.buf = nil
			} else {
				b.buf = b.mem.Reallocate(newCap, b.buf)
			}
		}
	}
	b.length = newSize
}
