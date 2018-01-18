package arrow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitIsSet(t *testing.T) {
	bits := make([]byte, 2)
	bits[0] = 0xa1
	bits[1] = 0xc2
	exp := []bool{true, false, false, false, false, true, false, true, false, true, false, false, false, false, true, true}
	var got []bool
	for i := 0; i < 0x10; i++ {
		got = append(got, bitIsSet(bits, i))
	}
	assert.Equal(t, exp, got)
}

func TestBitIsNotSet(t *testing.T) {
	bits := make([]byte, 2)
	bits[0] = 0xa1
	bits[1] = 0xc2
	exp := []bool{false, true, true, true, true, false, true, false, true, false, true, true, true, true, false, false}
	var got []bool
	for i := 0; i < 0x10; i++ {
		got = append(got, bitIsNotSet(bits, i))
	}
	assert.Equal(t, exp, got)
}

func TestSetBit(t *testing.T) {
	bits := make([]byte, 2)
	bits[0] = 0xff
	bits[1] = 0xff
	for i, v := range []bool{false, true, true, true, true, false, true, false, true, false, true, true, true, true, false, false} {
		if v {
			clearBit(bits, i)
		}
	}
	assert.Equal(t, []byte{0xa1, 0xc2}, bits)
}

func TestClearBit(t *testing.T) {
	bits := make([]byte, 2)
	for i, v := range []bool{true, false, false, false, false, true, false, true, false, true, false, false, false, false, true, true} {
		if v {
			setBit(bits, i)
		}
	}
	assert.Equal(t, []byte{0xa1, 0xc2}, bits)
}

func BenchmarkBitIsSet(b *testing.B) {
	bits := make([]byte, 32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bitIsSet(bits, (i%32)&0x1a)
	}
}

func BenchmarkSetBit(b *testing.B) {
	bits := make([]byte, 32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setBit(bits, (i%32)&0x1a)
	}
}
