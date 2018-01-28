package arrow

import (
	"github.com/influxdata/arrow/internal/bitutil"
)

type BooleanTraits struct{}

// BytesRequired returns the number of bytes required to store the requested number of elements.
func (t BooleanTraits) BytesRequired(elements int) int { return bitutil.CeilByte(elements) / 8 }
func (t BooleanTraits) CastFromBytes(b []byte) []byte  { return b }
func (t BooleanTraits) CastToBytes(b []byte) []byte    { return b }
func (t BooleanTraits) Copy(dst, src []byte)           { copy(dst, src) }
