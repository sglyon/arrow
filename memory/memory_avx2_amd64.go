// +build !noasm

package memory

import "unsafe"

//go:noescape
func _memset_avx2(buf, len, c unsafe.Pointer)

func memory_memset_avx2(buf []byte, c byte) {
	_memset_avx2(unsafe.Pointer(&buf[0]), unsafe.Pointer(uintptr(len(buf))), unsafe.Pointer(uintptr(c)))
}
