// +build noasm

package memory

func Set(buf []byte, c byte) { memory_memset_go(buf, c) }
