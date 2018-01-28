package memory

var (
	memset func(b []byte, c byte)
)

// Set assigns the value c to every element of the slice buf.
func Set(buf []byte, c byte) {
	memset(buf, c)
}
