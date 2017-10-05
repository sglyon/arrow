package arrow_test

import (
	"testing"

	"github.com/influxdata/arrow"
)

const (
	N = 1e6

	expSum   = 49999950000
	expCount = N
)

var data []float64
var floatArray arrow.Float64Array

func init() {
	data = make([]float64, N)
	floatArray = arrow.NewEmptyFloat64Array(int32(len(data)), nil)
	for i := 0; i < N; i++ {
		v := float64(i) / 10
		data[i] = v
		floatArray.Set(int32(i), v)
	}
}

type AggRangeFunc interface {
	Reset()
	Do(f []float64)
	Value() float64
}

type sumRangeFunc struct {
	sum float64
}

func (s *sumRangeFunc) Do(fs []float64) {
	for _, f := range fs {
		s.sum += f
	}
}
func (s *sumRangeFunc) Reset() {
	s.sum = 0
}
func (s *sumRangeFunc) Value() float64 {
	return s.sum
}

type countRangeFunc struct {
	count float64
}

func (c *countRangeFunc) Do(fs []float64) {
	c.count += float64(len(fs))
}
func (c *countRangeFunc) Reset() {
	c.count = 0
}
func (c *countRangeFunc) Value() float64 {
	return c.count
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
	for i, f := range fs {
		if nc.IsNull(i) {
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
		b.Fatalf("unexpected sum got %f exp %d", sum, expSum)
	}
}
func BenchmarkSumFuncArray(b *testing.B) {
	f := new(sumFunc)
	sum := doBenchAggFuncArray(b, f)
	if sum != expSum {
		b.Fatalf("unexpected sum got %f exp %d", sum, expSum)
	}
}

type nonNullChecker struct{}

func (nonNullChecker) IsNull(int) bool {
	return false
}
func (nonNullChecker) NullCount() int32 {
	return 0
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
