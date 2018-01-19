package memory

func memory_memset_go(buf []byte, c byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = c
	}
}
