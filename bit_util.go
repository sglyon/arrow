package arrow

import (
	"math/bits"
)

var (
	bitMask = []byte{1, 2, 4, 8, 16, 32, 64, 128}
	// the ~i byte version of bitMask
	flippedBitMask = []byte{254, 253, 251, 247, 239, 223, 191, 127}
)

func nextPowerOf2(x int) int { return 1 << uint(bits.Len(uint(x))) }

// ceilByte rounds size to the next multiple of 8
func ceilByte(size int) int { return (size + 7) &^ 7 }

func bitIsSet(bits []byte, i int) bool    { return (bits[i/8] & bitMask[i%8]) != 0 }
func bitIsNotSet(bits []byte, i int) bool { return (bits[i/8] & bitMask[i%8]) == 0 }
func setBit(bits []byte, i int)           { bits[i/8] |= bitMask[i%8] }
func clearBit(bits []byte, i int)         { bits[i/8] &= flippedBitMask[i%8] }
