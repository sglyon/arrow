package arrow_test

import (
	"fmt"

	"github.com/influxdata/arrow"
	"github.com/influxdata/arrow/memory"
)

// This example demonstrates how to build an array of int64 values using an array builder and Append.
// Whilst convenient for small arrays,
func Example_minimal() {
	// Create an allocator.
	pool := memory.NewGoAllocator()

	// Create an int64 buffer builder.
	builder := arrow.NewInt64ArrayBuilder(pool)

	builder.Append(1)
	builder.Append(2)
	builder.Append(3)
	builder.AppendNull()
	builder.Append(5)
	builder.Append(6)
	builder.Append(7)
	builder.Append(8)

	// Finish populating the array and reset the builder.
	array := builder.Finish()

	// Enumerate the values.
	for i, v := range array.Int64Values() {
		fmt.Printf("array[%d] = ", i)
		if array.IsNull(i) {
			fmt.Println("(null)")
		} else {
			fmt.Println(v)
		}
	}

	// Output:
	// array[0] = 1
	// array[1] = 2
	// array[2] = 3
	// array[3] = (null)
	// array[4] = 5
	// array[5] = 6
	// array[6] = 7
	// array[7] = 8
}
