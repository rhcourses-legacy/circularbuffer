package circularbuffer

/*
Diese Testdatei enthält Tests, die die Funktionen GetOldest, RemoveOldest und
GetAndRemove darauf testen, ob sie bei einem leeren Buffer paniken.
Die Tests sind in einer eigenen Datei, um die eigentlichen Tests für die
Kernfunktionalität übersichtlicher zu halten.
*/

import "testing"

// TestBuffer_GetOldest_Empty tests the GetOldest function.
// It creates a Circularbuffer with size 3 and calls GetOldest immediately.
// It expects a panic.
func TestBuffer_GetOldest_Empty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("cb.GetOldest did not panic")
		}
	}()

	cb := New(3)
	cb.GetOldest()
}

// TestBuffer_RemoveOldest_Empty tests the RemoveOldest function.
// It creates a Circularbuffer with size 3 and calls RemoveOldest immediately.
// It expects a panic.
func TestBuffer_RemoveOldest_Empty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("cb.RemoveOldest did not panic")
		}
	}()

	cb := New(3)
	cb.RemoveOldest()
}

// TestBuffer_GetAndRemove_Empty tests the GetAndRemove function.
// It creates a Circularbuffer with size 3 and calls GetAndRemove immediately.
// It expects a panic.
func TestBuffer_GetAndRemove_Empty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("cb.GetAndRemove did not panic")
		}
	}()

	cb := New(3)
	cb.GetAndRemove()
}
