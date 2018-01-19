package arrow_test

import (
	"testing"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/internal/testing/tools"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestBooleanArrayBuilder_AppendValues(t *testing.T) {
	pool := memory.NewGoAllocator()
	b := arrow.NewBooleanArrayBuilder(pool)

	exp := tools.Bools(1, 1, 0, 1, 1, 0, 1, 0)
	got := make([]bool, len(exp))

	b.AppendValues(exp, nil)
	a := b.Finish()
	for i := 0; i < a.Len(); i++ {
		got[i] = a.Value(i)
	}
	assert.Equal(t, exp, got)
}
