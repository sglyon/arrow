package array

import (
	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/memory"
)

// A type which represents the memory and metadata for an Arrow array.
type Data struct {
	typE      arrow.DataType
	nullN     int
	length    int
	buffers   []*memory.Buffer // TODO(sgc): should this be an interface?
	childData []*Data          // TODO(sgc): managed by ListArray, StructArray and UnionArray types
}

func NewData(typE arrow.DataType, length int, buffers []*memory.Buffer, nullN int) *Data {
	return &Data{
		typE:    typE,
		nullN:   nullN,
		length:  length,
		buffers: buffers,
	}
}

func (a *Data) DataType() arrow.DataType { return a.typE }
func (a *Data) NullN() int               { return a.nullN }
func (a *Data) Len() int                 { return a.length }
