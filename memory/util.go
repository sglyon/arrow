package memory

import "unsafe"

func roundToPowerOf2(v, round int) int {
	forceCarry := round - 1
	truncateMask := ^forceCarry
	return (v + forceCarry) & truncateMask
}

func roundUpToMultipleOf64(v int) int {
	return roundToPowerOf2(v, 64)
}

func addressOf(b []byte) uintptr {
	return uintptr(unsafe.Pointer(&b[0]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
