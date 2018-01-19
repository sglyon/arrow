// +build !noasm

package memory

import (
	"github.com/influxdata/arrow/internal/cpu"
)

var (
	memset func(b []byte, c byte)
)

func init() {
	if cpu.X86.HasAVX2 {
		memset = memory_memset_avx2
	} else if cpu.X86.HasSSE3 {
		memset = memory_memset_sse3
	} else {
		memset = memory_memset_go
	}
}

func Set(buf []byte, c byte) {
	memset(buf, c)
}
