package arrow_test

import (
	"reflect"
	"testing"

	"github.com/influxdata/arrow"
)

func TestArrowNulls(t *testing.T) {
	size := int32(10)
	a := arrow.NewEmptyFloat64Array(size)

	n := 0
	a.Do(func(fs []float64, nc arrow.NullChecker) {
		n = len(fs) - int(nc.NullCount())
	})
	if n != 0 {
		t.Fatal("array should be empty", n)
	}

	for i := int32(0); i < size; i++ {
		if i%2 == 0 {
			a.Set(i, float64(i*i))
		}
	}

	exp := []float64{0, 4, 16, 36, 64}
	var got []float64
	a.Do(func(fs []float64, nc arrow.NullChecker) {
		for i, f := range fs {
			if nc.IsNull(i) {
				continue
			}
			got = append(got, f)
		}
	})

	if !reflect.DeepEqual(exp, got) {
		t.Errorf("unexpected set of values got %v exp %v", got, exp)
	}
}
