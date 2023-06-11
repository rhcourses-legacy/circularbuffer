package circularbuffer_generic

// Circularbuffer is a struct that holds the buffer, size, and
// positions of first and last elements of the buffer.
type Circularbuffer[T any] struct {
	buffer []T
	size   int
	first  int
}

// New creates a new Circularbuffer with the given size.
func New[T any](size int) *Circularbuffer[T] {
	return &Circularbuffer[T]{
		buffer: make([]T, size),
		size:   0,
		first:  0,
	}
}

// Size returns the number of elements in the buffer.
func (cb *Circularbuffer[T]) Size() int {
	return cb.size
}

// Capacity returns the capacity of the buffer.
func (cb *Circularbuffer[T]) Capacity() int {
	return len(cb.buffer)
}

// IsEmpty returns true if the buffer is empty.
func (cb *Circularbuffer[T]) IsEmpty() bool {
	return cb.size == 0
}

// IsFull returns true if the buffer is full.
func (cb *Circularbuffer[T]) IsFull() bool {
	return cb.size == len(cb.buffer)
}

// Add adds an element to the buffer.
// If the buffer is full, the oldest element is overwritten.
func (cb *Circularbuffer[T]) Add(element T) {
	if cb.IsFull() {
		cb.RemoveOldest()
	}
	last := (cb.first + cb.size) % cb.Capacity()
	cb.buffer[last] = element
	cb.size++
}

// GetAndRemove returns the oldest element in the buffer and removes it.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer[T]) GetAndRemove() T {
	element := cb.GetOldest()
	cb.RemoveOldest()
	return element
}

// GetOldest returns the oldest element in the buffer.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer[T]) GetOldest() T {
	if cb.IsEmpty() {
		panic("Buffer is empty")
	}
	return cb.buffer[cb.first]
}

// RemoveOldest removes the oldest element in the buffer.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer[T]) RemoveOldest() {
	if cb.IsEmpty() {
		panic("Buffer is empty")
	}
	cb.first = (cb.first + 1) % cb.Capacity()
	cb.size--
}
