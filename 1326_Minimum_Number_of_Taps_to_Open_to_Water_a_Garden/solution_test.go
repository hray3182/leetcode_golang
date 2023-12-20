package leetcode

import "testing"

func TestMinTaps(t *testing.T) {
	var result int
	result = minTaps(5, []int{3, 0, 1, 1, 0, 0})
	if result != 1 {
		t.Errorf("MinTaps failed, got %d, expected 1", result)
	}
	result = minTaps(3, []int{0, 0, 0, 0})
	if result != -1 {
		t.Errorf("MinTaps failed, got %d, expected -1", result)
	}
	result = minTaps(7, []int{1, 2, 1, 0, 2, 1, 0, 1})
	if result != 3 {
		t.Errorf("MinTaps failed, got %d, expected 3", result)
	}
	result = minTaps(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4})
	if result != 2 {
		t.Errorf("MinTaps failed, got %d, expected 2", result)
	}
	result = minTaps(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4})
	if result != 1 {
		t.Errorf("MinTaps failed, got %d, expected 1", result)
	}
}

func TestMinTaps2(t *testing.T) {
	var result int
	result = minTaps2(5, []int{3, 0, 1, 1, 0, 0})
	// [[0,3],[1,1],[1,3],[2,4],[4,4],[5,5]]

	if result != -1 {
		t.Errorf("MinTaps failed, got %d, expected 1", result)
	}
	result = minTaps2(3, []int{0, 0, 0, 0})
	if result != -1 {
		t.Errorf("MinTaps failed, got %d, expected -1", result)
	}
	result = minTaps2(7, []int{1, 2, 1, 0, 2, 1, 0, 1})
	if result != 3 {
		t.Errorf("MinTaps failed, got %d, expected 3", result)
	}
	result = minTaps2(8, []int{4, 0, 0, 0, 0, 0, 0, 0, 4})
	if result != 2 {
		t.Errorf("MinTaps failed, got %d, expected 2", result)
	}
	result = minTaps2(8, []int{4, 0, 0, 0, 4, 0, 0, 0, 4})
	if result != 1 {
		t.Errorf("MinTaps failed, got %d, expected 1", result)
	}
}
