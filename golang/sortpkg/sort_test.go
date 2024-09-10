package sortpkg_test

import (
	"testing"

	"github.com/redbonzai/sortpkg/sortpkg"
)

func TestSort(t *testing.T) {
	tests := []struct {
		width, height, length, mass int
		expected                    string
	}{
		// Standard packages (neither bulky nor heavy)
		{90, 90, 90, 10, "STANDARD"},
		{50, 50, 50, 5, "STANDARD"},
		{100, 100, 100, 15, "STANDARD"},

		// Special packages (bulky but not heavy)
		{200, 100, 100, 10, "SPECIAL"},
		{150, 90, 90, 10, "SPECIAL"},
		{90, 150, 90, 15, "SPECIAL"},
		{90, 90, 150, 15, "SPECIAL"},
		{100, 100, 1000, 15, "SPECIAL"}, // Volume = 1,000,000 cm³

		// Special packages (heavy but not bulky)
		{90, 90, 90, 20, "SPECIAL"},  // Mass is exactly 20 kg
		{95, 95, 95, 25, "SPECIAL"},  // Mass is greater than 20 kg, but not bulky

		// Rejected packages (both bulky and heavy)
		{200, 200, 200, 25, "REJECTED"},
		{150, 150, 150, 21, "REJECTED"}, // Dimensions exactly 150 cm, mass over 20 kg

		// Edge cases (Invalid inputs)
		{0, 90, 90, 10, "Invalid input: dimensions must be positive"},
		{90, 90, 90, -5, "Invalid input: mass must be positive"},
		{90, 0, 90, 10, "Invalid input: dimensions must be positive"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := sortpkg.Sort(test.width, test.height, test.length, test.mass)
			if result != test.expected {
				t.Errorf("expected %s but got %s", test.expected, result)
			}
		})
	}
}
