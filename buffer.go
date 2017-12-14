package arrow

// Buffr is an immutable chunk of bytes
type Buffr struct {
	mutable     bool
	data        []byte
	mutableDAta []byte
	size        int64
	capacity    int64
	parent      *Buffr
}
