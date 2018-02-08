package array_test

import (
	"testing"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestNewFloat64Data(t *testing.T) {
	exp := []float64{1.0, 2.0, 4.0, 8.0, 16.0}

	ad := array.NewData(arrow.PrimitiveTypes.Float64, len(exp), []*memory.Buffer{nil, memory.NewBufferBytes(arrow.Float64Traits.CastToBytes(exp))}, 0)
	fa := array.NewFloat64Data(ad)

	assert.Equal(t, len(exp), fa.Len(), "unexpected Len()")
	assert.Equal(t, exp, fa.Float64Values(), "unexpected Float64Values()")
}
