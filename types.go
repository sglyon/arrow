package arrow

// Type is a *logical* type. They can be expressed as
// either a primitive physical type (bytes or bits of some fixed size), a
// nested type consisting of other data types, or another data type (e.g. a
// timestamp encoded as an int64)
type Type int

const (
	// NULL type having no physical storage
	NULL Type = iota

	// BOOL is a 1 bit, LSB bit-packed ordering
	BOOL

	// UINT8 is an Unsigned 8-bit little-endian integer
	UINT8

	// INT8 is a Signed 8-bit little-endian integer
	INT8

	// UINT16 is an Unsigned 16-bit little-endian integer
	UINT16

	// INT16 is a Signed 16-bit little-endian integer
	INT16

	// UINT32 is an Unsigned 32-bit little-endian integer
	UINT32

	// INT32 is a Signed 32-bit little-endian integer
	INT32

	// UINT64 is an Unsigned 64-bit little-endian integer
	UINT64

	// INT64 is a Signed 64-bit little-endian integer
	INT64

	// HALF_FLOAT is a 2-byte floating point value
	HALF_FLOAT

	// FLOAT is a 4-byte floating point value
	FLOAT

	// DOUBLE is an 8-byte floating point value
	DOUBLE

	// STRING is a UTF8 variable-length string
	STRING

	// BINARY is a Variable-length byte type (no guarantee of UTF8-ness)
	BINARY

	// FIXED_SIZE_BINARY is a binary where each value occupies the same number of bytes
	FIXED_SIZE_BINARY

	// DATE32 is int32 days since the UNIX epoch
	DATE32

	// DATE64 is int64 milliseconds since the UNIX epoch
	DATE64

	// TIMESTAMP is an exact timestamp encoded with int64 since UNIX epoch
	// Default unit millisecond
	TIMESTAMP

	// TIME32 is a signed 32-bit integer, representing either seconds or
	// milliseconds since midnight
	TIME32

	// TIME64 is a signed 64-bit integer, representing either microseconds or
	// nanoseconds since midnight
	TIME64

	// INTERVAL is YEAR_MONTH or DAY_TIME interval in SQL style
	INTERVAL

	// DECIMAL is a precision- and scale-based decimal type. Storage type depends on the
	// parameters.
	DECIMAL

	// LIST is a list of some logical data type
	LIST

	// STRUCT of logical types
	STRUCT

	// UNION of logical types
	UNION

	// DICTIONARY aka Category type
	DICTIONARY

	// MAP is a repeated struct logical type
	MAP
)

type ArrayData struct {
	typ        DataType
	length     int64
	null_count int64
	offset     int64
	buffers    []Buffr
	childData  []ArrayData
}

type Array struct {
	data           ArrayData
	nullBitmapData []byte
}

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

type Field struct {
	name     string           // Field name
	typ      DataType         // The field's data type
	nullable bool             // Fields can be nullable
	metadata KeyValueMetadata // The field's metadata, if any
}

type DataType struct {
	id       Type
	children []Field
}

type KeyValueMetadata struct {
	keys   []string
	values []string
}

type Column struct {
	field Field
	data  []Array
}
