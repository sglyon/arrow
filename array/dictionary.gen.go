// Code generated by array/dictionary.gen.go.tmpl.
// DO NOT EDIT.

package array

import (
	"github.com/influxdata/arrow"
)

// A type which represents an immutable sequence of int64 values, encoded using DictEncoding
type Int64Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []int64
}

// NewInt64DictData construct a new dict encoded array using reference and Pool data
func NewInt64DictData(data, poolData *Data) *Int64Dict {
	d := &Int64Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Int64Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Int64Dict) Value(i int) int64 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Int64Dict) Int64Values() []int64 {
	out := make([]int64, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Int64Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Int64Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type int64
func (d *Int64Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Int64Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of uint64 values, encoded using DictEncoding
type Uint64Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []uint64
}

// NewUint64DictData construct a new dict encoded array using reference and Pool data
func NewUint64DictData(data, poolData *Data) *Uint64Dict {
	d := &Uint64Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Uint64Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Uint64Dict) Value(i int) uint64 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Uint64Dict) Uint64Values() []uint64 {
	out := make([]uint64, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Uint64Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Uint64Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type uint64
func (d *Uint64Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Uint64Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of float64 values, encoded using DictEncoding
type Float64Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []float64
}

// NewFloat64DictData construct a new dict encoded array using reference and Pool data
func NewFloat64DictData(data, poolData *Data) *Float64Dict {
	d := &Float64Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Float64Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Float64Dict) Value(i int) float64 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Float64Dict) Float64Values() []float64 {
	out := make([]float64, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Float64Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Float64Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type float64
func (d *Float64Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Float64Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of int32 values, encoded using DictEncoding
type Int32Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []int32
}

// NewInt32DictData construct a new dict encoded array using reference and Pool data
func NewInt32DictData(data, poolData *Data) *Int32Dict {
	d := &Int32Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Int32Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Int32Dict) Value(i int) int32 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Int32Dict) Int32Values() []int32 {
	out := make([]int32, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Int32Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Int32Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type int32
func (d *Int32Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of uint32 values, encoded using DictEncoding
type Uint32Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []uint32
}

// NewUint32DictData construct a new dict encoded array using reference and Pool data
func NewUint32DictData(data, poolData *Data) *Uint32Dict {
	d := &Uint32Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Uint32Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Uint32Dict) Value(i int) uint32 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Uint32Dict) Uint32Values() []uint32 {
	out := make([]uint32, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Uint32Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Uint32Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type uint32
func (d *Uint32Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Uint32Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of float32 values, encoded using DictEncoding
type Float32Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []float32
}

// NewFloat32DictData construct a new dict encoded array using reference and Pool data
func NewFloat32DictData(data, poolData *Data) *Float32Dict {
	d := &Float32Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Float32Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Float32Dict) Value(i int) float32 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Float32Dict) Float32Values() []float32 {
	out := make([]float32, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Float32Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Float32Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type float32
func (d *Float32Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Float32Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of int16 values, encoded using DictEncoding
type Int16Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []int16
}

// NewInt16DictData construct a new dict encoded array using reference and Pool data
func NewInt16DictData(data, poolData *Data) *Int16Dict {
	d := &Int16Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Int16Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Int16Dict) Value(i int) int16 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Int16Dict) Int16Values() []int16 {
	out := make([]int16, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Int16Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Int16Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type int16
func (d *Int16Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Int16Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of uint16 values, encoded using DictEncoding
type Uint16Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []uint16
}

// NewUint16DictData construct a new dict encoded array using reference and Pool data
func NewUint16DictData(data, poolData *Data) *Uint16Dict {
	d := &Uint16Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Uint16Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Uint16Dict) Value(i int) uint16 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Uint16Dict) Uint16Values() []uint16 {
	out := make([]uint16, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Uint16Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Uint16Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type uint16
func (d *Uint16Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Uint16Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of int8 values, encoded using DictEncoding
type Int8Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []int8
}

// NewInt8DictData construct a new dict encoded array using reference and Pool data
func NewInt8DictData(data, poolData *Data) *Int8Dict {
	d := &Int8Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Int8Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Int8Dict) Value(i int) int8 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Int8Dict) Int8Values() []int8 {
	out := make([]int8, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Int8Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Int8Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type int8
func (d *Int8Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Int8Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of uint8 values, encoded using DictEncoding
type Uint8Dict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []uint8
}

// NewUint8DictData construct a new dict encoded array using reference and Pool data
func NewUint8DictData(data, poolData *Data) *Uint8Dict {
	d := &Uint8Dict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *Uint8Dict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *Uint8Dict) Value(i int) uint8 { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *Uint8Dict) Uint8Values() []uint8 {
	out := make([]uint8, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *Uint8Dict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on Uint8Dict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type uint8
func (d *Uint8Dict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.Uint8Traits.CastFromBytes(vals.Bytes())
	}
}

// A type which represents an immutable sequence of arrow.Timestamp values, encoded using DictEncoding
type TimestampDict struct {
	array
	poolData *Data
	Refs     []int32
	Pool     []arrow.Timestamp
}

// NewTimestampDictData construct a new dict encoded array using reference and Pool data
func NewTimestampDictData(data, poolData *Data) *TimestampDict {
	d := &TimestampDict{}
	d.refCount = 1
	d.setData(data)
	d.setPoolData(poolData)
	return d
}

func (d *TimestampDict) PoolData() *Data { return d.poolData }

// Value returns the value of the DictEncoded variable at index i by first consulting the reference and then extracting the appropriate element from the Pool
func (d *TimestampDict) Value(i int) arrow.Timestamp { return d.Pool[d.Refs[i]] }

// Values returns the all values in the DictEncoded column (see Value for more information)
func (d *TimestampDict) TimestampValues() []arrow.Timestamp {
	out := make([]arrow.Timestamp, len(d.Refs))
	for ix := 0; ix < len(d.Refs); ix++ {
		out[ix] = d.Pool[d.Refs[ix]]
	}
	return out
}

// setData updates the references. The data must have type int32
func (d *TimestampDict) setData(data *Data) {
	if data.typE.ID() != arrow.INT32 {
		panic("Can only call setData on TimestampDict with int32 data")
	}
	d.array.setData(data)
	vals := data.buffers[1]
	if vals != nil {
		d.Refs = arrow.Int32Traits.CastFromBytes(vals.Bytes())
	}
}

// setData updates the Pool. The data must be of type arrow.Timestamp
func (d *TimestampDict) setPoolData(poolData *Data) {
	if d.poolData != nil {
		d.poolData.Release()
	}

	d.poolData = poolData
	poolData.Retain()
	vals := poolData.buffers[1]
	if vals != nil {
		d.Pool = arrow.TimestampTraits.CastFromBytes(vals.Bytes())
	}
}