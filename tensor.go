package arrow

// Tensor is an n-dimensional array
// Arrow implementations in general are not required to implement this type
type Tensor struct {
	typ     DataType
	data    Buffr
	shape   int64
	strides int64

	/// These names are optional
	dimNames []string
}
