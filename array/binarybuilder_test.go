package array_test

import (
	"testing"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestBinaryBuilder(t *testing.T) {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	defer mem.AssertSize(t, 0)

	ab := array.NewBinaryBuilder(mem, arrow.BinaryTypes.Binary)

	exp := [][]byte{[]byte("foo"), []byte("bar"), nil, []byte("sydney"), []byte("cameron")}
	for _, v := range exp {
		if v == nil {
			ab.AppendNull()
		} else {
			ab.Append(v)
		}
	}

	assert.Equal(t, len(exp), ab.Len(), "unexpected Len()")
	assert.Equal(t, 1, ab.NullN(), "unexpected NullN()")

	for i, v := range exp {
		if v == nil {
			v = []byte{}
		}
		assert.Equal(t, v, ab.Value(i), "unexpected BinaryArrayBuilder.Value(%d)", i)
	}

	ar := ab.NewArray()
	ab.Release()
	ar.Release()

	// check state of builder after finish
	assert.Zero(t, ab.Len(), "unexpected ArrayBuilder.Len(), NewArray did not reset state")
	assert.Zero(t, ab.Cap(), "unexpected ArrayBuilder.Cap(), NewArray did not reset state")
	assert.Zero(t, ab.NullN(), "unexpected ArrayBuilder.NullN(), NewArray did not reset state")
}
