package day2_test

import (
	"testing"

	"github.com/bhvalent/adventOfCode2025/src/internal/app/day2"
)

func TestGetInvalidRangeSum_ShouldReturnCorrectAnswer_WithTestData(t *testing.T) {
	result := day2.GetInvalidRangeSum("input_test_data.txt")
	expected := 1227775554
	if result != expected {
		t.Errorf("GetInvalidRangeSum() = %d; Expected: %d", result, expected)
	}
}

func TestGetInvalidRangeSum_ShouldReturnCorrectAnswer_WithRealData(t *testing.T) {
	result := day2.GetInvalidRangeSum("input_real_data.txt")
	expected := 40398804950
	if result != expected {
		t.Errorf("GetInvalidRangeSum() = %d; Expected: %d", result, expected)
	}
}

