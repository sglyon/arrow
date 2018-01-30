package math

import "github.com/influxdata/arrow/array"

type float64Funcs struct {
	sum func(a *array.Float64) float64
}

var (
	Float64 float64Funcs
)

//
func (f float64Funcs) Sum(a *array.Float64) float64 {
	return f.sum(a)
}

func sum_go(a *array.Float64) float64 {
	acc := float64(0)
	for _, v := range a.Float64Values() {
		acc += v
	}
	return acc
}
