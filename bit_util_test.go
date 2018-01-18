package arrow

import "testing"

func BenchmarkSetBit(b *testing.B) {
	bits := make([]byte, 32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		setBit(bits, (i%32)&0x1a)
	}
}
