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

func TestHowManyTimesPassingZero_ShouldReturnCorrectAnswer_WithTestData(t *testing.T) {
	result := day1.HowManyTimesPassingZero("input_test_data.txt")
	expected := 6
	if result != expected {
		t.Errorf("HowManyTimesPassingZero() = %d; Expected: %d", result, expected)
	}
}

func TestHowManyTimesPassingZero_ShouldReturnCorrectAnswer_WithAdditionalTestData(t *testing.T) {
	result := day1.HowManyTimesPassingZero("input_additional_test_data.txt")
	expected := 31
	if result != expected {
		t.Errorf("HowManyTimesPassingZero() = %d; Expected: %d", result, expected)
	}
}

func TestHowManyTimesPassingZero_ShouldReturnCorrectAnswer_WithRealData(t *testing.T) {
	result := day1.HowManyTimesPassingZero("input_real_data.txt")
	expected := 6819
	if result != expected {
		t.Errorf("HowManyTimesPassingZero() = %d; Expected: %d", result, expected)
	}
}

