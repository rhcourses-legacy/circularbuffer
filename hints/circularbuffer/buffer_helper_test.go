package circularbuffer

/*
Diese Datei enthält Hilfsfunktionen für die Tests.
*/

import "testing"

// CheckSizeEmptyFull checks the results of the Size, Empty, and IsFull methods.
// It is a helper function used by the tests.
// It expects a buffer, the expected result of Size, Empty, and IsFull,
// and the testing object.
func CheckSizeEmptyFull(cb *Circularbuffer, size int, empty bool, full bool, t *testing.T) {
	if cb.Size() != size {
		t.Errorf("cb has size %d, expected %d", cb.Size(), size)
	}
	if cb.IsEmpty() != empty {
		t.Errorf("cb.IsEmpty returned %t, expected %t", cb.IsEmpty(), empty)
	}
	if cb.IsFull() != full {
		t.Errorf("cb.IsFull returned %t, expected %t", cb.IsFull(), full)
	}
}
