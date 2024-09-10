package sortpkg

import (
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		width, height, length, mass int
		expected                    string
	}{
		{90, 90, 90, 10, "STANDARD"},
		{95, 95, 95, 20, "SPECIAL"},
		{95, 95, 100, 20, "SPECIAL"},
		{200, 200, 200, 25, "REJECTED"},
		{160, 50, 50, 10, "SPECIAL"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := Sort(test.width, test.height, test.length, test.mass)
			if result != test.expected {
				t.Errorf("expected %s but got %s", test.expected, result)
			}
		})
	}
}
