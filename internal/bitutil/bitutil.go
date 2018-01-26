package bitutil

import (
	"math/bits"
)

var (
	BitMask        = [8]byte{1, 2, 4, 8, 16, 32, 64, 128}
	FlippedBitMask = [8]byte{254, 253, 251, 247, 239, 223, 191, 127}
)

// NextPowerOf2 rounds x to the next power of two.
func NextPowerOf2(x int) int { return 1 << uint(bits.Len(uint(x))) }

// CeilByte rounds size to the next multiple of 8.
func CeilByte(size int) int { return (size + 7) &^ 7 }

// BitIsSet returns true if the bit at index i is set (1).
func BitIsSet(bits []byte, i int) bool { return (bits[uint(i)/8] & BitMask[byte(i)%8]) != 0 }

// BitIsNotSet returns true if the bit at index i is not set (0).
func BitIsNotSet(bits []byte, i int) bool { return (bits[uint(i)/8] & BitMask[byte(i)%8]) == 0 }

// SetBit sets the bit at index i to 1.
func SetBit(bits []byte, i int) { bits[uint(i)/8] |= BitMask[byte(i)%8] }

// ClearBit sets the bit at index i to 0.
func ClearBit(bits []byte, i int) { bits[uint(i)/8] &= FlippedBitMask[byte(i)%8] }

// SetBitTo sets the bit at index i to val.
func SetBitTo(bits []byte, i int, val bool) {
	if val {
		SetBit(bits, i)
	} else {
		ClearBit(bits, i)
	}
}
