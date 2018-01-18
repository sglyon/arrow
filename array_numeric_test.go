package arrow_test

import (
	"testing"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/memory"
)

func TestFloat64Array_SetData(t *testing.T) {
	data := []float64{1.0, 2.0, 4.0, 8.0, 16.0}

	ad := arrow.NewArrayData(arrow.PrimitiveTypes.Float64, len(data), []*memory.Buffer{nil, memory.NewBuffer(arrow.Float64Traits{}.CastToBytes(data))}, 0)
	fa := arrow.NewFloat64Array(ad)
	t.Logf("name: %v", fa.Float64Values())
}
