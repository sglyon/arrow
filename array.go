package arrow

type Array interface {
	DataType() DataType
	NullN() int
	NullBitmapBytes() []byte
	IsNull(i int) bool
	IsValid(i int) bool
	Data() *ArrayData
	Len() int
}
