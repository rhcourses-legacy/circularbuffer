package circularbuffer

import "fmt"

// ExampleCircularbuffer_string_int demonstrates the usage of the Circularbuffer
// with different types of elements.
func ExampleCircularbuffer_string_int() {
	cb := New(5)
	cb.Add(42)
	cb.Add(38)
	cb.Add("horse")

	for !cb.IsEmpty() {
		el := cb.GetAndRemove()
		fmt.Printf("%v: %T\n", el, el)
	}

	// Output:
	// 42: int
	// 38: int
	// horse: string
}
