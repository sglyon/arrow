package arrow

import "github.com/influxdata/arrow/memory"

type Pool interface {
	ArrayBuilderPool
	ArrayPool
}

type ArrayBuilderPool interface {
	NewFloat64ArrayBuilder(pool memory.Allocator) *Float64ArrayBuilder
}

type ArrayPool interface {
	NewFloat64Array(data *ArrayData) *Float64Array
}

var (
	DefaultPool Pool = &defaultPool{}
)

type defaultPool struct{}

func (*defaultPool) NewFloat64ArrayBuilder(pool memory.Allocator) *Float64ArrayBuilder {
	return NewFloat64ArrayBuilder(pool)
}

func (*defaultPool) NewFloat64Array(data *ArrayData) *Float64Array {
	return NewFloat64Array(data)
}
