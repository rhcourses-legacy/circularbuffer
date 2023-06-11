package circularbuffer

import "testing"

// TestBuffer_New tests the New function.
// It creates Circularbuffer with different sizes and checks the results
// of the Size, Empty, and IsFull methods.
func TestBuffer_New_Size(t *testing.T) {
	cb5 := New(5)
	CheckSizeEmptyFull(cb5, 0, true, false, t)

	cb10 := New(10)
	CheckSizeEmptyFull(cb10, 0, true, false, t)
}

// TestBuffer_Add_One tests the Add function.
// It creates a Circularbuffer with size 3 and adds one element.
// Then it checks the results of the Size, Empty, and IsFull methods.
// Furthermore, it checks if GetOldest returns the added element.
func TestBuffer_Add_One(t *testing.T) {
	cb := New(3)
	cb.Add(11)

	CheckSizeEmptyFull(cb, 1, false, false, t)

	if cb.GetOldest() != 11 {
		t.Errorf("cb.GetOldest returned %d, expected 11", cb.GetOldest())
	}
}

// TestBuffer_Add_Two tests the Add function.
// It creates a Circularbuffer with size 3 and adds 2 elements.
// Then it checks the results of the Size, Empty, and IsFull methods.
// Furthermore, it checks if GetOldest returns the first added element.
func TestBuffer_Add_Two(t *testing.T) {
	cb := New(3)
	cb.Add(11)
	cb.Add(22)

	CheckSizeEmptyFull(cb, 2, false, false, t)

	if cb.GetOldest() != 11 {
		t.Errorf("cb.GetOldest returned %d, expected 11", cb.GetOldest())
	}
}

// TestBuffer_Add_Three tests the Add function.
// It creates a Circularbuffer with size 3 and adds 3 elements.
// Then it checks the results of the Size, Empty, and IsFull methods.
// Furthermore, it checks if GetOldest returns the first added element.
func TestBuffer_Add_Three(t *testing.T) {
	cb := New(3)
	cb.Add(11)
	cb.Add(22)
	cb.Add(33)

	CheckSizeEmptyFull(cb, 3, false, true, t)

	if cb.GetOldest() != 11 {
		t.Errorf("cb.GetOldest returned %d, expected 11", cb.GetOldest())
	}
}

// TestBuffer_Add_Full_Overwrite tests the Add function.
// It creates a Circularbuffer with size 3 and adds 4 elements.
// Then it checks the results of the Size, Empty, and IsFull methods.
// Furthermore, it checks if GetOldest returns the second added element.
func TestBuffer_Add_Full_Overwrite(t *testing.T) {
	cb := New(3)
	cb.Add(11)
	cb.Add(22)
	cb.Add(33)
	cb.Add(44)

	CheckSizeEmptyFull(cb, 3, false, true, t)

	if cb.GetOldest() != 22 {
		t.Errorf("cb.GetOldest returned %d, expected 22", cb.GetOldest())
	}
}

// TestBuffer_Add2_Remove1 tests the GetAndRemove function.
// It creates a Circularbuffer with size 3 and adds 2 elements.
//
// Then it checks the the following things:
// * if GetAndRemove returns the first added element.
// * results of the Size, Empty, and IsFull methods after the removal.
// * if GetOldest returns the second added element after the removal.
func TestBuffer_Add2_Remove1(t *testing.T) {
	cb := New(3)
	cb.Add(11)
	cb.Add(22)

	if cb.GetAndRemove() != 11 {
		t.Errorf("cb.GetAndRemove returned %d, expected 11", cb.GetAndRemove())
	}

	CheckSizeEmptyFull(cb, 1, false, false, t)

	if cb.GetOldest() != 22 {
		t.Errorf("cb.GetOldest returned %d, expected 22", cb.GetOldest())
	}
}

// TestBuffer_Add2_Remove2 tests the GetAndRemove function.
// It creates a Circularbuffer with size 3 and adds 2 elements.
// It calls GetAndRemove twice and checks the expected results.
func TestBuffer_Add2_Remove2(t *testing.T) {
	cb := New(3)
	cb.Add(11)
	cb.Add(22)

	if cb.GetAndRemove() != 11 {
		t.Errorf("cb.GetAndRemove returned %d, expected 11", cb.GetAndRemove())
	}

	if cb.GetAndRemove() != 22 {
		t.Errorf("cb.GetAndRemove returned %d, expected 22", cb.GetAndRemove())
	}

	CheckSizeEmptyFull(cb, 0, true, false, t)
}
