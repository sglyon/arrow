package array_test

import (
	"testing"

	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestNewFloat64Builder(t *testing.T) {
	mem := memory.NewGoAllocator()
	ab := array.NewFloat64Builder(mem)

	ab.Append(1)
	ab.Append(2)
	ab.Append(3)
	ab.AppendNull()
	ab.Append(5)
	ab.Append(6)
	ab.AppendNull()
	ab.Append(8)
	ab.Append(9)
	ab.Append(10)

	// check state of builder before Finish
	assert.Equal(t, 10, ab.Len(), "unexpected Len()")
	assert.Equal(t, 2, ab.NullN(), "unexpected NullN()")

	a := ab.Finish()

	// check state of builder after Finish
	assert.Zero(t, ab.Len(), "unexpected ArrayBuilder.Len(), Finish did not reset state")
	assert.Zero(t, ab.Cap(), "unexpected ArrayBuilder.Cap(), Finish did not reset state")
	assert.Zero(t, ab.NullN(), "unexpected ArrayBuilder.NullN(), Finish did not reset state")

	// check state of array
	assert.Equal(t, 2, a.NullN(), "unexpected null count")
	assert.Equal(t, []float64{1, 2, 3, 0, 5, 6, 0, 8, 9, 10}, a.Float64Values(), "unexpected Float64Values")
	assert.Equal(t, []byte{0xb7}, a.NullBitmapBytes()[:1]) // 4 bytes due to minBuilderCapacity
	assert.Len(t, a.Float64Values(), 10, "unexpected length of Float64Values")

	ab.Append(7)
	ab.Append(8)

	a = ab.Finish()

	assert.Equal(t, 0, a.NullN())
	assert.Equal(t, []float64{7, 8}, a.Float64Values())
	assert.Len(t, a.Float64Values(), 2)
}
