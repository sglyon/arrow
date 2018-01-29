package array_test

import (
	"testing"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/memory"
	"github.com/stretchr/testify/assert"
)

type testDataType struct {
	id arrow.Type
}

func (d *testDataType) ID() arrow.Type { return d.id }
func (d *testDataType) Name() string   { panic("implement me") }

func TestMakeFromData(t *testing.T) {
	tests := []struct {
		name     string
		d        arrow.DataType
		expPanic bool
		expError string
	}{
		// unsupported types
		{name: "null", d: &testDataType{arrow.NULL}, expPanic: true, expError: "unsupported data type: NULL"},
		{name: "map", d: &testDataType{arrow.MAP}, expPanic: true, expError: "unsupported data type: MAP"},

		// supported types
		{name: "bool", d: &testDataType{arrow.BOOL}},

		// invalid types
		{name: "invalid(-1)", d: &testDataType{arrow.Type(-1)}, expPanic: true, expError: "invalid data type: Type(-1)"},
		{name: "invalid(28)", d: &testDataType{arrow.Type(28)}, expPanic: true, expError: "invalid data type: Type(28)"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var b [4]*memory.Buffer
			data := array.NewData(test.d, 0, b[:], 0)

			if test.expPanic {
				assert.PanicsWithValue(t, test.expError, func() {
					array.MakeFromData(data)
				})
			} else {
				assert.NotNil(t, array.MakeFromData(data))
			}
		})
	}
}
