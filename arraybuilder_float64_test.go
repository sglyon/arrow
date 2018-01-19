package arrow_test

import (
	"testing"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestNewFloat64ArrayBuilder(t *testing.T) {
	pool := memory.NewGoAllocator()
	b := arrow.NewFloat64ArrayBuilder(pool)

	b.Append(1)
	b.Append(2)
	b.Append(3)
	b.AppendNull()
	b.Append(5)
	b.Append(6)
	b.AppendNull()
	b.Append(8)

	a := b.Finish()
	assert.Equal(t, 2, a.NullN())
	assert.Equal(t, []float64{1, 2, 3, 0, 5, 6, 0, 8}, a.Float64Values())
	assert.Equal(t, []byte{0xb7}, a.NullBitmapBytes()[:1]) // 4 bytes due to minBuilderCapacity

	b.Append(7)
	b.Append(8)

	a = b.Finish()
	assert.Equal(t, 0, a.NullN())
	assert.Equal(t, []float64{7, 8}, a.Float64Values())
}
