package arrow

// ListArray is a nested type in which each array slot contains a variable-size
// sequence of values all having the same relative type
type ListArray struct {
	data            ArrayData
	nullBitmapData  []byte
	rawValueOffsets []int32
	values          Array
}

type BinaryArray struct {
	data            ArrayData
	nullBitmapData  []byte
	rawValueOffsets []int32
	rawData         []byte
}

type FixedSizeBinaryArray struct {
	data           ArrayData
	nullBitmapData []byte
	rawValues      []byte
	byteWidth      int32
}

type StructArray struct {
	data           ArrayData
	nullBitmapData []byte
	boxedFields    []Array
}

type UnionArray struct {
	data            ArrayData
	nullBitmapData  []byte
	rawTypeIds      []byte
	rawValueOffsets []int32
	boxedFields     []Array // For caching boxed child data
}

type DictionaryArray struct {
	data           ArrayData
	nullBitmapData []byte
	dictType       DictionaryType
	indices        Array
}

type DictionaryType struct {
	indexType  DataType
	dictionary Array
	ordered    bool
}
