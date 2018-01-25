package arrow

import "unsafe"

type BinaryArray struct {
	array
	valueOffsets []int32
	valueBytes   []byte
}

func NewBinaryArray(data *ArrayData) *BinaryArray {
	a := &BinaryArray{}
	a.setData(data)
	return a
}

// Value returns the slice at index i. This value should not be mutated.
func (a *BinaryArray) Value(i int) []byte {
	return a.valueBytes[a.valueOffsets[i]:a.valueOffsets[i+1]]
}

// ValueString returns the string at index i without performing additional allocations.
// The string is only valid for the lifetime of the BinaryArray.
func (a *BinaryArray) ValueString(i int) string {
	b := a.Value(i)
	return *(*string)(unsafe.Pointer(&b))
}

func (a *BinaryArray) ValueOffset(i int) int { return int(a.valueOffsets[i]) }
func (a *BinaryArray) ValueLen(i int) int    { return int(a.valueOffsets[i+1] - a.valueOffsets[i]) }
func (a *BinaryArray) ValueOffsets() []int32 { return a.valueOffsets }
func (a *BinaryArray) ValueBytes() []byte    { return a.valueBytes }

func (a *BinaryArray) setData(data *ArrayData) {
	if len(data.buffers) != 3 {
		panic("len(data.buffers) != 3")
	}

	a.array.setData(data)

	if valueData := data.buffers[2]; valueData != nil {
		a.valueBytes = valueData.Bytes()
	}

	if valueOffsets := data.buffers[1]; valueOffsets != nil {
		a.valueOffsets = Int32Traits{}.CastFromBytes(valueOffsets.Bytes())
	}
}
