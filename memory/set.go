package memory

func Set(b []byte, c byte) {
	// TODO: optimize
	for i := 0; i < len(b); i++ {
		b[i] = c
	}
}
