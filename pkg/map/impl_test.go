package _map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapFilter(t *testing.T) {
	cases := []struct {
		name      string
		input     Map[string, int]
		predicate PredicateFunc[string, int]
		expected  Map[string, int]
	}{
		{
			name: "FilterEvenValues",
			input: Map[string, int]{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			predicate: func(key string, value int) bool { return value%2 == 0 },
			expected: Map[string, int]{
				"two": 2,
			},
		},
		{
			name:      "FilterEmptyMap",
			input:     Map[string, int]{},
			predicate: func(key string, value int) bool { return value%2 == 0 },
			expected:  Map[string, int]{},
		},
		{
			name: "FilterAllValues",
			input: Map[string, int]{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			predicate: func(key string, value int) bool { return true },
			expected: Map[string, int]{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},
		{
			name: "FilterNoValues",
			input: Map[string, int]{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			predicate: func(key string, value int) bool { return false },
			expected:  Map[string, int]{},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Filter(tt.predicate)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemove(t *testing.T) {
	type testCase struct {
		mapIn          Map[int, int]
		removeKey      int
		removeValue    int
		expectedResult bool
		expectedMapOut Map[int, int]
	}

	tests := map[string]testCase{
		"ExistingPair": {
			mapIn: map[int]int{
				1: 1,
				2: 2,
				3: 3,
			},
			removeKey:      2,
			removeValue:    2,
			expectedResult: true,
			expectedMapOut: map[int]int{
				1: 1,
				3: 3,
			},
		},
		"NonexistentPair": {
			mapIn: map[int]int{
				1: 1,
				2: 2,
				3: 3,
			},
			removeKey:      4,
			removeValue:    4,
			expectedResult: false,
			expectedMapOut: map[int]int{
				1: 1,
				2: 2,
				3: 3,
			},
		},
		"ExistingKeyDifferentValue": {
			mapIn: map[int]int{
				1: 1,
				2: 2,
				3: 3,
			},
			removeKey:      2,
			removeValue:    3,
			expectedResult: false,
			expectedMapOut: map[int]int{
				1: 1,
				2: 2,
				3: 3,
			},
		},
		"EmptyMap": {
			mapIn:          map[int]int{},
			removeKey:      2,
			removeValue:    2,
			expectedResult: false,
			expectedMapOut: map[int]int{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assertions := assert.New(t)

			result := tc.mapIn.Remove(tc.removeKey, tc.removeValue)
			assertions.Equal(tc.expectedResult, result, "Incorrect result for %s", name)
			assertions.Equal(tc.expectedMapOut, tc.mapIn, "Incorrect map values for %s", name)
		})
	}
}

func TestAll(t *testing.T) {
	testCases := []struct {
		name      string
		mapInit   Map[string, int]
		predicate PredicateFunc[string, int]
		expected  bool
	}{
		{
			name: "All_Elements_Satisfy_Condition",
			mapInit: Map[string, int]{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			predicate: func(key string, value int) bool {
				return value > 0
			},
			expected: true,
		},
		{
			name: "Not_All_Elements_Satisfy_Condition",
			mapInit: Map[string, int]{
				"one":      1,
				"zero":     0,
				"negative": -1,
			},
			predicate: func(key string, value int) bool {
				return value >= 0
			},
			expected: false,
		},
		{
			name:    "Empty_Map",
			mapInit: Map[string, int]{},
			predicate: func(key string, value int) bool {
				return false
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.mapInit.All(tc.predicate)
			assert.Equal(t, tc.expected, result, tc.name)
		})
	}
}

func TestMap_Any(t *testing.T) {
	tests := []struct {
		name      string
		m         Map[int, int]
		predicate PredicateFunc[int, int]
		want      bool
	}{
		{
			name:      "Test_True_Condition",
			m:         Map[int, int]{1: 5, 2: 8, 3: 11},
			predicate: func(key int, value int) bool { return value > 10 },
			want:      true,
		},
		{
			name:      "Test_False_Condition",
			m:         Map[int, int]{1: 5, 2: 8, 3: 9},
			predicate: func(key int, value int) bool { return value > 10 },
			want:      false,
		},
		{
			name:      "Test_Empty_Map",
			m:         Map[int, int]{},
			predicate: func(key int, value int) bool { return value > 10 },
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.Any(tt.predicate)
			assert.Equal(t, tt.want, got)
		})
	}
}
