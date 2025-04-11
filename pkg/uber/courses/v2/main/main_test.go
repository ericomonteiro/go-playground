package main

import "testing"

func TestGraph(t *testing.T) {
	tests := []struct {
		name     string
		courses  [][]int
		expected bool
	}{
		{
			name: "valid grade",
			courses: [][]int{
				{0, 1},
				{1, 2},
				{1, 4},
				{2, 3},
				{2, 4},
			},
			expected: true,
		},
		{
			name: "valid grade",
			courses: [][]int{
				{0, 1},
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			expected: true,
		},
		{
			name: "invalid grade",
			courses: [][]int{
				{0, 1},
				{1, 2},
				{1, 4},
				{2, 3},
				{3, 1},
			},
			expected: false,
		},
		{
			name: "invalid grade",
			courses: [][]int{
				{0, 1},
				{1, 0},
				{1, 4},
				{2, 3},
				{3, 1},
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := checkGradeIsValid(tc.courses)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
