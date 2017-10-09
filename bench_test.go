package arrow_test

import (
	"testing"

	"github.com/influxdata/arrow"
)

const (
	N = int(1e6)

	nullSpacing = 37

	expSum            = float64(49999950000)
	expCount          = float64(N)
	expSumWithNulls   = 48648551351.400002
	expCountWithNulls = float64(N - int(N/nullSpacing) - 1)
)

var data []float64
var floatArray *arrow.Float64Array
var floatArrayWithNulls *arrow.Float64Array

func init() {
	data = make([]float64, N)
	floatArray = arrow.NewEmptyFloat64Array(int32(len(data)))
	floatArrayWithNulls = arrow.NewEmptyFloat64Array(int32(len(data)))
	for i := 0; i < N; i++ {
		v := float64(i) / 10
		data[i] = v
		floatArray.Set(int32(i), v)
		if i%nullSpacing == 0 {
			//Every nullSpacing'th value is null
			continue
		}
		floatArrayWithNulls.Set(int32(i), v)
	}
}

type AggFunc interface {
	Reset()
	Do(f []float64, nc arrow.NullChecker)
	Value() float64
}

type sumFunc struct {
	sum float64
}

func (s *sumFunc) Do(fs []float64, nc arrow.NullChecker) {
	if nc.NullCount() == 0 {
		for _, f := range fs {
			s.sum += f
		}
		return
	}
	anc := nc.(*arrow.Float64Array)
	for i, f := range fs {
		if anc.IsNull(i) {
			continue
		}
		s.sum += f
	}
}

func (s *sumFunc) Reset() {
	s.sum = 0
}
func (s *sumFunc) Value() float64 {
	return s.sum
}

type countFunc struct {
	count int
}

func (c *countFunc) Do(fs []float64, nc arrow.NullChecker) {
	c.count += len(fs) - int(nc.NullCount())
}
func (c *countFunc) Reset() {
	c.count = 0
}
func (c *countFunc) Value() float64 {
	return float64(c.count)
}

type nonNullChecker struct{}

func (nonNullChecker) IsNull(int) bool {
	return false
}
func (nonNullChecker) NullCount() int32 {
	return 0
}

func BenchmarkCountFuncData(b *testing.B) {
	f := new(countFunc)
	count := doBenchAggFuncData(b, f)
	if count != expCount {
		b.Fatalf("unexpected count got %f exp %f", count, expCount)
	}
}
func BenchmarkCountFuncArray(b *testing.B) {
	f := new(countFunc)
	count := doBenchAggFuncArray(b, f)
	if count != expCount {
		b.Fatalf("unexpected count got %f exp %f", count, expCount)
	}
}
func BenchmarkSumFuncData(b *testing.B) {
	f := new(sumFunc)
	sum := doBenchAggFuncData(b, f)
	if sum != expSum {
		b.Fatalf("unexpected sum got %f exp %f", sum, expSum)
	}
}
func BenchmarkSumFuncArray(b *testing.B) {
	f := new(sumFunc)
	sum := doBenchAggFuncArray(b, f)
	if sum != expSum {
		b.Fatalf("unexpected sum got %f exp %f", sum, expSum)
	}
}

func BenchmarkCountFuncArrayWithNulls(b *testing.B) {
	f := new(countFunc)
	count := doBenchAggFuncArrayWithNulls(b, f)
	if count != expCountWithNulls {
		b.Fatalf("unexpected count got %f exp %f", count, expCountWithNulls)
	}
}
func BenchmarkSumFuncArrayWithNulls(b *testing.B) {
	f := new(sumFunc)
	sum := doBenchAggFuncArrayWithNulls(b, f)
	if sum != expSumWithNulls {
		b.Fatalf("unexpected sum got %f exp %f", sum, expSumWithNulls)
	}
}

func doBenchAggFuncData(b *testing.B, f AggFunc) float64 {
	for i := 0; i < b.N; i++ {
		f.Reset()
		f.Do(data, nonNullChecker{})
	}
	return f.Value()
}

func doBenchAggFuncArray(b *testing.B, f AggFunc) float64 {
	for i := 0; i < b.N; i++ {
		f.Reset()
		floatArray.Do(f.Do)
	}
	return f.Value()
}

func doBenchAggFuncArrayWithNulls(b *testing.B, f AggFunc) float64 {
	for i := 0; i < b.N; i++ {
		f.Reset()
		floatArrayWithNulls.Do(f.Do)
	}
	return f.Value()
}
