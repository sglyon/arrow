package memory

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundToPowerOf2(t *testing.T) {
	tests := []struct {
		v, round int
		exp      int
	}{
		{60, 64, 64},
		{122, 64, 128},
		{16, 64, 64},
		{64, 64, 64},
		{13, 8, 16},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("v%d_r%d", test.v, test.round), func(t *testing.T) {
			a := roundToPowerOf2(test.v, test.round)
			assert.Equal(t, test.exp, a)
		})
	}
}
