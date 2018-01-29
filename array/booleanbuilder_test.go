package array_test

import (
	"testing"

	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/internal/testing/tools"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestBooleanBuilder_AppendValues(t *testing.T) {
	mem := memory.NewGoAllocator()
	b := array.NewBooleanBuilder(mem)

	exp := tools.Bools(1, 1, 0, 1, 1, 0, 1, 0)
	got := make([]bool, len(exp))

	b.AppendValues(exp, nil)
	a := b.Finish()
	for i := 0; i < a.Len(); i++ {
		got[i] = a.Value(i)
	}
	assert.Equal(t, exp, got)
}
