package arrow_test

import (
	"fmt"

	"github.com/influxdata/arrow/array"
	"github.com/influxdata/arrow/memory"
)

// This example demonstrates how to build an array of int64 values using a builder and Append.
// Whilst convenient for small arrays,
func Example_minimal() {
	// Create an allocator.
	pool := memory.NewGoAllocator()

	// Create an int64 array builder.
	builder := array.NewInt64Builder(pool)

	builder.Append(1)
	builder.Append(2)
	builder.Append(3)
	builder.AppendNull()
	builder.Append(5)
	builder.Append(6)
	builder.Append(7)
	builder.Append(8)

	// Finish populating the array and reset the builder.
	ints := builder.Finish()

	// Enumerate the values.
	for i, v := range ints.Int64Values() {
		fmt.Printf("ints[%d] = ", i)
		if ints.IsNull(i) {
			fmt.Println("(null)")
		} else {
			fmt.Println(v)
		}
	}

	// Output:
	// ints[0] = 1
	// ints[1] = 2
	// ints[2] = 3
	// ints[3] = (null)
	// ints[4] = 5
	// ints[5] = 6
	// ints[6] = 7
	// ints[7] = 8
}
