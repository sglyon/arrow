/*
Package arrow provides an implementation of Apache Arrow.

Apache Arrow is a cross-language development platform for in-memory data. It specifies a standardized
language-independent columnar memory format for flat and hierarchical data, organized for efficient analytic
operations on modern hardware. It also provides computational libraries and zero-copy streaming
messaging and inter-process communication.

Fundamentals

The most basic data structure in Arrow is an Array. It holds a sequence of values of the same type. An array
consists of memory holding the data and an additional validity bitmap that indicates if the corresponding entry in the
array is valid (not null). If the array has no null entries, it is possible to omit this bitmap.

*/
package arrow

//go:generate tmpl -data=@array_numeric.gen.go.tmpldata array_numeric.gen.go.tmpl arraybuilder_numeric.gen.go.tmpl type_traits_numeric.gen.go.tmpl bufferbuilder_numeric.gen.go.tmpl
//go:generate tmpl -data=@datatype_numeric.gen.go.tmpldata datatype_numeric.gen.go.tmpl
