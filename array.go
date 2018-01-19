package arrow

import "github.com/influxdata/arrow/memory"

type Array interface {
	DataType() DataType
	NullN() int
	NullBitmapBytes() []byte
	IsNull(i int) bool
	IsValid(i int) bool
	Data() *ArrayData
	Len() int
	Values() *memory.Buffer
}
