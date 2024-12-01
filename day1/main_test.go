package main

import (
	"testing"
)

func TestGetDistanceSum(t *testing.T) {
	tests := map[string]struct {
		a        []int
		b        []int
		expected int
	}{
		"advent test case": {
			a:        []int{3, 4, 2, 1, 3, 3},
			b:        []int{4, 3, 5, 3, 9, 3},
			expected: 11,
		},
		"minimal test":                {a: []int{1}, b: []int{2}, expected: 1},
		"mininal test 2":              {a: []int{5}, b: []int{2}, expected: 3},
		"multiple input in both list": {a: []int{2, 5}, b: []int{9, 3}, expected: 5},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := getDistanceSum(test.a, test.b)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.expected {
				t.Fatalf("getDistanceSum(%+v, %+v) returned %d; expected %d", test.a, test.b, got, test.expected)
			}
		})
	}
}
