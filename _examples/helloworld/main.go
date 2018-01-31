package main

import (
	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/math"
	"github.com/influxdata/arrow/memory"
)

func main() {
	fb := array.NewFloat64Builder(memory.DefaultAllocator)

	fb.AppendValues([]float64{1, 3, 5, 7, 9, 11}, nil)

	vec := fb.Finish()
	math.Float64.Sum(vec)
}
