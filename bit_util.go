package arrow

import (
	"math/bits"
)

var (
	bitMask        = [8]byte{1, 2, 4, 8, 16, 32, 64, 128}
	flippedBitMask = [8]byte{254, 253, 251, 247, 239, 223, 191, 127}
)

func nextPowerOf2(x int) int { return 1 << uint(bits.Len(uint(x))) }

// ceilByte rounds size to the next multiple of 8
func ceilByte(size int) int { return (size + 7) &^ 7 }

func bitIsSet(bits []byte, i int) bool    { return (bits[uint(i)/8] & bitMask[byte(i)%8]) != 0 }
func bitIsNotSet(bits []byte, i int) bool { return (bits[uint(i)/8] & bitMask[byte(i)%8]) == 0 }
func setBit(bits []byte, i int)           { bits[uint(i)/8] |= bitMask[byte(i)%8] }
func clearBit(bits []byte, i int)         { bits[uint(i)/8] &= flippedBitMask[byte(i)%8] }
