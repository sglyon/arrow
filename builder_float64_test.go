package arrow_test

import (
	"testing"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/memory"
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
	b.Append(7)
	b.Append(8)

	a := b.Finish()
	t.Log(a.Float64Values())
}
