package math

import (
	"github.com/influxdata/arrow/internal/cpu"
)

func init() {
	if cpu.X86.HasAVX2 {
		initAVX2()
	} else if cpu.X86.HasSSE42 {
		initSSE4()
	} else {
		initGo()
	}
}

func initAVX2() {
	Float64.sum = sum_float64_avx2
}

func initSSE4() {
	Float64.sum = sum_float64_sse4
}

func initGo() {
	Float64.sum = sum_go
}
