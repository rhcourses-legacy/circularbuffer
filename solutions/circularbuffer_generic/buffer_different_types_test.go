package circularbuffer_generic

import "fmt"

// ExampleCircularbuffer_string demonstrates the usage
// of the Circularbuffer with strings as elements.
func ExampleCircularbuffer_string() {
	cb := New[string](5)
	cb.Add("cow")
	cb.Add("sheep")
	cb.Add("horse")

	for !cb.IsEmpty() {
		el := cb.GetAndRemove()
		fmt.Printf("%v: %T\n", el, el)
	}

	// Output:
	// cow: string
	// sheep: string
	// horse: string
}

// ExampleCircularbuffer_int demonstrates the usage
// of the Circularbuffer with strings as elements.
func ExampleCircularbuffer_int() {
	cb := New[int](5)
	cb.Add(42)
	cb.Add(38)
	cb.Add(100)

	for !cb.IsEmpty() {
		el := cb.GetAndRemove()
		fmt.Printf("%v: %T\n", el, el)
	}

	// Output:
	// 42: int
	// 38: int
	// 100: int
}

// ExampleCircularbuffer_custom demonstrates the usage
// of the Circularbuffer with a custom struct as elements.
func ExampleCircularbuffer_custom() {
	type Address struct {
		Street string
		City   string
	}

	cb := New[Address](5)
	cb.Add(Address{"Haupstrasse", "Berlin"})
	cb.Add(Address{"Lindenstrasse", "München"})
	cb.Add(Address{"Kaiserstrasse", "Hamburg"})

	for !cb.IsEmpty() {
		el := cb.GetAndRemove()
		fmt.Printf("%v: %T\n", el, el)
	}

	// Output:
	// {Haupstrasse Berlin}: circularbuffer_generic.Address
	// {Lindenstrasse München}: circularbuffer_generic.Address
	// {Kaiserstrasse Hamburg}: circularbuffer_generic.Address
}
