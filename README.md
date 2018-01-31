Arrow
=====

[Apache Arrow][arrow] is a cross-language development platform for in-memory data. It specifies a 
standardized language-independent columnar memory format for flat and hierarchical data, 
organized for efficient analytic operations on modern hardware. It also provides computational 
libraries and zero-copy streaming messaging and inter-process communication.

Performance
-----------

The arrow package makes extensive use of [c2goasm][] to leverage LLVM's advanced optimizer and generate PLAN9 
assembly functions from C/C++ code. The arrow package can be compiled without these optimizations using the `noasm` 
build tag. Alternatively, by configuring an environment variable, it is possible to dynamically configure which 
architecture optimizations are used at runtime. 
See the `cpu` package [README](internal/cpu/README.md) for a description of this environment variable.

### Example Usage

The following benchmarks demonstrate summing an array of 8192 values using various optimizations. 

Disable no architecture optimizations (thus using AVX2):

```sh
$ INTEL_DISABLE_EXT=NONE go test -bench=8192 -run=. ./math
goos: darwin
goarch: amd64
pkg: github.com/influxdata/arrow/math
BenchmarkFloat64Funcs_Sum_8192-8   	 2000000	       687 ns/op	95375.41 MB/s
BenchmarkInt64Funcs_Sum_8192-8     	 2000000	       719 ns/op	91061.06 MB/s
BenchmarkUint64Funcs_Sum_8192-8    	 2000000	       691 ns/op	94797.29 MB/s
PASS
ok  	github.com/influxdata/arrow/math	6.444s
```

**NOTE:** `NONE` is simply ignored, thus enabling optimizations for AVX2 and SSE4

----

Disable AVX2 architecture optimizations:

```sh
$ INTEL_DISABLE_EXT=AVX2 go test -bench=8192 -run=. ./math
goos: darwin
goarch: amd64
pkg: github.com/influxdata/arrow/math
BenchmarkFloat64Funcs_Sum_8192-8   	 1000000	      1912 ns/op	34263.63 MB/s
BenchmarkInt64Funcs_Sum_8192-8     	 1000000	      1392 ns/op	47065.57 MB/s
BenchmarkUint64Funcs_Sum_8192-8    	 1000000	      1405 ns/op	46636.41 MB/s
PASS
ok  	github.com/influxdata/arrow/math	4.786s
```

----

Disable ALL architecture optimizations, thus using pure Go implementation:

```sh
$ INTEL_DISABLE_EXT=ALL go test -bench=8192 -run=. ./math
goos: darwin
goarch: amd64
pkg: github.com/influxdata/arrow/math
BenchmarkFloat64Funcs_Sum_8192-8   	  200000	     10285 ns/op	6371.41 MB/s
BenchmarkInt64Funcs_Sum_8192-8     	  500000	      3892 ns/op	16837.37 MB/s
BenchmarkUint64Funcs_Sum_8192-8    	  500000	      3929 ns/op	16680.00 MB/s
PASS
ok  	github.com/influxdata/arrow/math	6.179s
```

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
[c2goasm]:  https://github.com/minio/c2goasm