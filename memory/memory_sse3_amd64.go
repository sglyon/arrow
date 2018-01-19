// +build !noasm

package memory

import "unsafe"

//go:noescape
func _memset_sse3(buf, len, c unsafe.Pointer)

func memory_memset_sse3(buf []byte, c byte) {
	_memset_sse3(unsafe.Pointer(&buf[0]), unsafe.Pointer(uintptr(len(buf))), unsafe.Pointer(uintptr(c)))
}
