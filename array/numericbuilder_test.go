package array_test

import (
	"testing"

	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestNewFloat64Builder(t *testing.T) {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	defer mem.AssertSize(t, 0)

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

	// check state of builder before NewFloat64Array
	assert.Equal(t, 10, ab.Len(), "unexpected Len()")
	assert.Equal(t, 2, ab.NullN(), "unexpected NullN()")

	a := ab.NewFloat64Array()

	// check state of builder after NewFloat64Array
	assert.Zero(t, ab.Len(), "unexpected ArrayBuilder.Len(), NewFloat64Array did not reset state")
	assert.Zero(t, ab.Cap(), "unexpected ArrayBuilder.Cap(), NewFloat64Array did not reset state")
	assert.Zero(t, ab.NullN(), "unexpected ArrayBuilder.NullN(), NewFloat64Array did not reset state")

	// check state of array
	assert.Equal(t, 2, a.NullN(), "unexpected null count")
	assert.Equal(t, []float64{1, 2, 3, 0, 5, 6, 0, 8, 9, 10}, a.Float64Values(), "unexpected Float64Values")
	assert.Equal(t, []byte{0xb7}, a.NullBitmapBytes()[:1]) // 4 bytes due to minBuilderCapacity
	assert.Len(t, a.Float64Values(), 10, "unexpected length of Float64Values")

	a.Release()

	ab.Append(7)
	ab.Append(8)

	a = ab.NewFloat64Array()

	assert.Equal(t, 0, a.NullN())
	assert.Equal(t, []float64{7, 8}, a.Float64Values())
	assert.Len(t, a.Float64Values(), 2)

	a.Release()
}

func TestFloat32Builder_AppendValues(t *testing.T) {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	defer mem.AssertSize(t, 0)

	ab := array.NewFloat64Builder(mem)

	exp := []float64{1.0, 1.1, 1.2, 1.3}
	ab.AppendValues(exp, nil)
	a := ab.NewFloat64Array()
	assert.Equal(t, exp, a.Float64Values())

	a.Release()
	ab.Release()
}
