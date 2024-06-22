package arrayandslices

import "testing"

func TestAddArray(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	result := 15
	slice_result := AddArray(slice)
	if result != slice_result {
		t.Errorf("The desired value is %v and the value got is %v", result, slice_result)
	}
}
