package circularbuffer

// Circularbuffer is a struct that holds the buffer, size, and
// positions of first and last elements of the buffer.
type Circularbuffer struct {
	buffer []interface{}
	size   int
	first  int
}

// New creates a new Circularbuffer with the given size.
func New(size int) *Circularbuffer {
	return &Circularbuffer{
		buffer: make([]interface{}, size),
		size:   0,
		first:  0,
	}
}

// Size returns the number of elements in the buffer.
func (cb *Circularbuffer) Size() int {
	return cb.size
}

// Capacity returns the capacity of the buffer.
func (cb *Circularbuffer) Capacity() int {
	return len(cb.buffer)
}

// IsEmpty returns true if the buffer is empty.
func (cb *Circularbuffer) IsEmpty() bool {
	return cb.size == 0
}

// IsFull returns true if the buffer is full.
func (cb *Circularbuffer) IsFull() bool {
	return cb.size == len(cb.buffer)
}

// Add adds an element to the buffer.
// If the buffer is full, the oldest element is overwritten.
func (cb *Circularbuffer) Add(element interface{}) {
	// Prüfen ob der Puffer voll ist.
	// Berechnen der Position des letzten Elements.
	// Hinzufügen des Elements an der berechneten Stelle.
	// Anpassen der Größe des Puffers.
}

// GetAndRemove returns the oldest element in the buffer and removes it.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer) GetAndRemove() interface{} {
	// Ältestes Element aus dem Puffer holen.
	// Ältestes Element aus dem Puffer entfernen.
	// Zuvor geholtes Element zurückgeben.
	return nil // (nil ist ein Platzhalter für den Rückgabewert)
}

// GetOldest returns the oldest element in the buffer.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer) GetOldest() interface{} {
	// Prüfen ob der Puffer leer ist.
	// Ältestes Element aus dem Puffer holen und zurückgeben.
	return nil // (nil ist ein Platzhalter für den Rückgabewert)
}

// RemoveOldest removes the oldest element in the buffer.
// If the buffer is empty, a panic occurs.
func (cb *Circularbuffer) RemoveOldest() {
	// Prüfen ob der Puffer leer ist.
	// Neue Position des ältesten Elements berechnen (Länge des Puffers beachten).
	// Größe des Puffers anpassen.
}
