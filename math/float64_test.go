package math_test

import (
	"testing"

	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/math"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestFloat64Funcs_Sum(t *testing.T) {
	vec := makeArrayFloat64(10000)
	res := math.Float64.Sum(vec)
	assert.Equal(t, res, float64(49995000.0))
}

func makeArrayFloat64(l int) *array.Float64 {
	fb := array.NewFloat64Builder(memory.NewGoAllocator())
	fb.Reserve(l)
	for i := 0; i < l; i++ {
		fb.Append(float64(i))
	}
	return fb.Finish()
}

func benchmarkFloat64Funcs_Sum(b *testing.B, n int) {
	vec := makeArrayFloat64(n)
	b.SetBytes(int64(vec.Len() * 8))
	for i := 0; i < b.N; i++ {
		math.Float64.Sum(vec)
	}
}

func BenchmarkFloat64Funcs_Sum_256(b *testing.B) {
	benchmarkFloat64Funcs_Sum(b, 256)
}

func BenchmarkFloat64Funcs_Sum_1024(b *testing.B) {
	benchmarkFloat64Funcs_Sum(b, 1024)
}
