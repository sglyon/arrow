Arrow
=====

[Apache Arrow][arrow] is a cross-language development platform for in-memory data. It specifies a 
standardized language-independent columnar memory format for flat and hierarchical data, 
organized for efficient analytic operations on modern hardware. It also provides computational 
libraries and zero-copy streaming messaging and inter-process communication.

Status
------

The first milestone was to implement the necessary Array types in order to use
them internally in the [ifql][] execution engine and storage layers of [InfluxDB][].


### Memory Management

- [x] Allocations are 64-byte aligned and padded to 8-bytes


### Array and builder support

**Primitive types**

- [x] Signed and unsigned 8, 16, 32 and 64 bit integers
- [x] 32 and 64 bit floats
- [x] Packed LSB booleans
- [x] Variable-length binary
- [ ] String (valid UTF-8)
- [ ] Half-float (16-bit)
- [ ] Null (no physical storage)

**Parametric types**

- [x] Timestamp
- [ ] Interval (year/month or day/time)
- [ ] Date32 (days since UNIX epoch)
- [ ] Date64 (milliseconds since UNIX epoch)
- [ ] Time32 (seconds or milliseconds since midnight)
- [ ] Time64 (microseconds or nanoseconds since midnight)
- [ ] Decimal (128-bit)
- [ ] Fixed-sized binary
- [ ] List
- [ ] Struct
- [ ] Union
    - [ ] Dense
    - [ ] Sparse
- [ ] Dictionary 
    - [ ] Dictionary encoding

### Type metadata

- [x] Data types (implemented arrays)
- [ ] Field
- [ ] Schema
  

### I/O 

Serialization is planned for a future iteration.

- [ ] Flat buffers for serializing metadata
- [ ] Record Batch
- [ ] Table



[arrow]:    https://arrow.apache.org
[ifql]:     https://github.com/influxdata/ifql
[InfluxDB]: https://github.com/influxdata/influxdb