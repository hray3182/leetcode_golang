package leetcode

import "testing"

func TestMinTaps(t *testing.T) {
	var result int
	result = MinTaps(5, []int{3, 4, 1, 1, 0, 0})
	if result != 1 {
		t.Errorf("MinTaps failed, got %d, expected 1", result)
	}
	result = MinTaps(3, []int{0, 0, 0, 0})
	if result != -1 {
		t.Errorf("MinTaps failed, got %d, expected -1", result)
	}
	result = MinTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1})
	if result != 3 {
		t.Errorf("MinTaps failed, got %d, expected 3", result)
	}
	result = MinTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4})
	if result != 2 {
		t.Errorf("MinTaps failed, got %d, expected 2", result)
	}
	result = MinTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4})
	if result != 1 {
		t.Errorf("MinTaps failed, got %d, expected 1", result)
	}
}
