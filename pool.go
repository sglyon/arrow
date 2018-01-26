package arrow

type Pool interface {
	ArrayPool
}

type ArrayPool interface {
	NewFloat64Array(data *ArrayData) *Float64Array
}

var (
	DefaultPool Pool = &defaultPool{}
)

type defaultPool struct{}

func (*defaultPool) NewFloat64Array(data *ArrayData) *Float64Array {
	return NewFloat64Array(data)
}
