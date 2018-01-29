package array

type Buffer interface {
	Buf() []byte
	Bytes() []byte
	Mutable() bool
	Len() int
	Cap() int
}

type ResizableBuffer interface {
	Buffer
	Reserve(capacity int)
	Resize(newSize int)
	ResizeNoShrink(newSize int)
}
