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

	// Finish building the int64 array and reset the builder.
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

// This example demonstrates creating an array, sourcing the values and
// null bitmaps directly from byte slices. The null count is set to
// UnknownNullCount, instructing the array to calculate the
// null count from the bitmap when NullN is called.
func Example_fromMemory() {
	// create LSB packed bits with the following pattern:
	// 01010011 11000101
	data := memory.NewBuffer([]byte{0xca, 0xa3})

	// create LSB packed validity (null) bitmap, where every 4th element is null:
	// 11101110 11101110
	nullBitmap := memory.NewBuffer([]byte{0x77, 0x77})

	// Create a boolean array and lazily determine NullN using UnknownNullCount
	bools := array.NewBoolean(16, data, nullBitmap, array.UnknownNullCount)

	// Show the null count
	fmt.Printf("NullN()  = %d\n", bools.NullN())

	// Enumerate the values.
	n := bools.Len()
	for i := 0; i < n; i++ {
		fmt.Printf("bools[%d] = ", i)
		if bools.IsNull(i) {
			fmt.Println("(null)")
		} else {
			fmt.Printf("%t\n", bools.Value(i))
		}
	}

	// Output:
	// NullN()  = 4
	// bools[0] = false
	// bools[1] = true
	// bools[2] = false
	// bools[3] = (null)
	// bools[4] = false
	// bools[5] = false
	// bools[6] = true
	// bools[7] = (null)
	// bools[8] = true
	// bools[9] = true
	// bools[10] = false
	// bools[11] = (null)
	// bools[12] = false
	// bools[13] = true
	// bools[14] = false
	// bools[15] = (null)
}
