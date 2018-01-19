// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: datatype_numeric.gen.go.tmpl

package arrow

type Int8Type struct{}

func (t *Int8Type) ID() Type     { return INT8 }
func (t *Int8Type) Name() string { return "int8" }

type Int16Type struct{}

func (t *Int16Type) ID() Type     { return INT16 }
func (t *Int16Type) Name() string { return "int16" }

type Int32Type struct{}

func (t *Int32Type) ID() Type     { return INT32 }
func (t *Int32Type) Name() string { return "int32" }

type Int64Type struct{}

func (t *Int64Type) ID() Type     { return INT64 }
func (t *Int64Type) Name() string { return "int64" }

type Uint8Type struct{}

func (t *Uint8Type) ID() Type     { return UINT8 }
func (t *Uint8Type) Name() string { return "uint8" }

type Uint16Type struct{}

func (t *Uint16Type) ID() Type     { return UINT16 }
func (t *Uint16Type) Name() string { return "uint16" }

type Uint32Type struct{}

func (t *Uint32Type) ID() Type     { return UINT32 }
func (t *Uint32Type) Name() string { return "uint32" }

type Uint64Type struct{}

func (t *Uint64Type) ID() Type     { return UINT64 }
func (t *Uint64Type) Name() string { return "uint64" }

type Float32Type struct{}

func (t *Float32Type) ID() Type     { return FLOAT32 }
func (t *Float32Type) Name() string { return "float32" }

type Float64Type struct{}

func (t *Float64Type) ID() Type     { return FLOAT64 }
func (t *Float64Type) Name() string { return "float64" }

var (
	PrimitiveTypes = struct {
		Int8    DataType
		Int16   DataType
		Int32   DataType
		Int64   DataType
		Uint8   DataType
		Uint16  DataType
		Uint32  DataType
		Uint64  DataType
		Float32 DataType
		Float64 DataType
	}{

		Int8:    &Int8Type{},
		Int16:   &Int16Type{},
		Int32:   &Int32Type{},
		Int64:   &Int64Type{},
		Uint8:   &Uint8Type{},
		Uint16:  &Uint16Type{},
		Uint32:  &Uint32Type{},
		Uint64:  &Uint64Type{},
		Float32: &Float32Type{},
		Float64: &Float64Type{},
	}
)
