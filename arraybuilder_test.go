package arrow

import (
	"testing"

	"github.com/influxdata/arrow/internal/testing/tools"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestArrayBuilder_Init(t *testing.T) {
	type exp struct{ size int }
	tests := []struct {
		name string
		cap  int

		exp exp
	}{
		{"07 bits", 07, exp{size: 1}},
		{"19 bits", 19, exp{size: 3}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ab := &arrayBuilder{pool: memory.NewGoAllocator()}
			ab.init(test.cap)
			assert.Equal(t, test.cap, ab.Cap(), "invalid capacity")
			assert.Equal(t, test.exp.size, ab.nullBitmap.Len(), "invalid length")
		})
	}
}

func TestArrayBuilder_UnsafeSetValid(t *testing.T) {
	ab := &arrayBuilder{pool: memory.NewGoAllocator()}
	ab.init(32)
	ab.unsafeAppendBoolsToBitmap(tools.Bools(0, 0, 0, 0, 0), 5)
	assert.Equal(t, 5, ab.Len())
	assert.Equal(t, []byte{0, 0, 0, 0}, ab.nullBitmap.Bytes())

	ab.unsafeSetValid(17)
	assert.Equal(t, []byte{0xe0, 0xff, 0x3f, 0}, ab.nullBitmap.Bytes())
}
