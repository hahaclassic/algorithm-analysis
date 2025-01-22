package bruteforce

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input: []int{1, 2},
			expected: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for _, test := range tests {
		var results [][]int
		permute(test.input, func(p []int) {
			permutation := make([]int, len(p))
			copy(permutation, p)
			results = append(results, permutation)
		})

		sort.Slice(results, func(i, j int) bool {
			return fmt.Sprint(results[i]) < fmt.Sprint(results[j])
		})
		sort.Slice(test.expected, func(i, j int) bool {
			return fmt.Sprint(test.expected[i]) < fmt.Sprint(test.expected[j])
		})

		if !reflect.DeepEqual(results, test.expected) {
			t.Errorf("permute(%v) = %v; want %v", test.input, results, test.expected)
		}
	}
}
