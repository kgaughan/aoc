package helpers

import "testing"

func TestLcm(t *testing.T) {
	if lcm([]int{1, 2, 8, 3}) != 24 {
		t.Errorf("Expected LCM=24")
	}
	if lcm([]int{1, 5, 6, 8}) != 120 {
		t.Errorf("Expected LCM=120")
	}
	if lcm([]int{2, 7, 3, 9, 4}) != 252 {
		t.Errorf("Expected LCM=252")
	}
}
