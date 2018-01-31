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


### Array support

- [x] Null bitmap support
- [x] Signed and unsigned 8, 16, 32 and 64 bit integers
- [x] Packed LSB booleans
- [x] Variable-length binary
- [x] 64-bit timestamps
- Parametric types
    - [ ] List
    - [ ] Struct
    - [ ] Union


### Array builder support

- [x] Signed and unsigned 8, 16, 32 and 64 bit integers
- [x] Packed LSB booleans
- [x] Variable-length binary
- [x] 64-bit timestamps
- Parametric types
    - [ ] List
    - [ ] Struct
    - [ ] Union


### Type metadata

- [x] Data types for existing arrays
- [ ] Field
- [ ] Schema
  

### I/O 

Serialization is planned for a future iteration.

- [ ] Flatbuffers for serializing metadata
- [ ] Record Batch
- [ ] Table



[arrow]:    https://arrow.apache.org
[ifql]:     https://github.com/influxdata/ifql
[InfluxDB]: https://github.com/influxdata/influxdb