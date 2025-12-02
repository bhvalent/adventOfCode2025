package day1_test

import (
	"testing"

	"github.com/bhvalent/adventOfCode2025/src/internal/app/day1"
)

func TestHowManyTimesAtZero_ShouldReturnCorrectAnswer_WithTestData(t *testing.T) {
	result := day1.HowManyTimesAtZero("input_test_data.txt")
	expected := 3
	if result != expected {
		t.Errorf("HowManyTimesAtZero() = %d; Expected: %d", result, expected)
	}
}

func TestHowManyTimesAtZero_ShouldReturnCorrectAnswer_WithRealData(t *testing.T) {
	result := day1.HowManyTimesAtZero("input_real_data.txt")
	expected := 1154
	if result != expected {
		t.Errorf("HowManyTimesAtZero() = %d; Expected: %d", result, expected)
	}
}

