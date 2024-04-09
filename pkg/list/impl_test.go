package list

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zach-robinson-dev/kollections/pkg/comparator"
	"github.com/zach-robinson-dev/kollections/pkg/util"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		list      List[int]
		predicate PredicateFunc[int]
		want      List[int]
	}{
		{
			name:      "Empty list",
			list:      List[int]{},
			predicate: func(x int) bool { return true },
			want:      List[int]{},
		},
		{
			name:      "Filter evens",
			list:      List[int]{1, 2, 3, 4, 5},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      List[int]{2, 4},
		},
		{
			name:      "Filter odds",
			list:      List[int]{1, 2, 3, 4, 5},
			predicate: func(x int) bool { return x%2 != 0 },
			want:      List[int]{1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Filter(tt.predicate)
			assert.Equal(t, tt.want, got, "Filter() should return expected result")
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name    string
		list    List[int]
		element int
		want    bool
	}{
		{"Contains in empty list", List[int]{}, 1, false},
		{"Contains in list without element", List[int]{2, 3, 4}, 1, false},
		{"Contains in list with element", List[int]{1, 2, 3, 4}, 1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.Contains(tt.element)
			assert.Equal(t, tt.want, got, "Contains() should return expected result")
		})
	}
}

func TestRemoveAll(t *testing.T) {
	tests := []struct {
		name       string
		list       List[int]
		toBeRemove []int
		want       List[int]
		modified   bool
	}{
		{"Remove from empty list", List[int]{}, []int{1}, List[int]{}, false},
		{"Remove nonexistent elements", List[int]{1, 2}, []int{3, 4}, List[int]{1, 2}, false},
		{"Remove existing elements", List[int]{1, 2, 3, 4}, []int{1, 3}, List[int]{2, 4}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.RemoveAll(tt.toBeRemove...)
			assert.Equal(t, tt.modified, got, "RemoveAll() should return expected modification flag")
			assert.Equal(t, tt.want, tt.list, "RemoveAll() should modify list as expected")
		})
	}
}

func TestList_All(t *testing.T) {
	type testCase struct {
		name     string
		list     List[int]
		pred     PredicateFunc[int]
		expected bool
	}

	testCases := []testCase{
		{
			name:     "all elements satisfy predicate",
			list:     List[int]{2, 4, 8, 10},
			pred:     func(item int) bool { return item%2 == 0 },
			expected: true,
		},
		{
			name:     "not all elements satisfy predicate",
			list:     List[int]{1, 2, 3, 4},
			pred:     func(item int) bool { return item > 2 },
			expected: false,
		},
		{
			name:     "no elements in list",
			list:     List[int]{},
			pred:     func(item int) bool { return item > 2 },
			expected: true,
		},
		{
			name:     "predicate always returns true",
			list:     List[int]{1, 2, 3, 4},
			pred:     func(item int) bool { return true },
			expected: true,
		},
		{
			name:     "predicate always returns false",
			list:     List[int]{1, 2, 3, 4},
			pred:     func(item int) bool { return false },
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := test.list.All(test.pred)
			assert.Equal(t, test.expected, result, test.name)
		})
	}
}

func TestAny(t *testing.T) {
	var lessThanTen PredicateFunc[int] = func(item int) bool {
		return item < 10
	}

	var isEven PredicateFunc[int] = func(item int) bool {
		return item%2 == 0
	}

	tt := []struct {
		name      string
		list      List[int]
		predicate PredicateFunc[int]
		expect    bool
	}{
		{name: "EmptyList", list: List[int]{}, predicate: lessThanTen, expect: false},
		{name: "SingleElementList", list: List[int]{8}, predicate: lessThanTen, expect: true},
		{name: "MultipleElementsAllLessThanTen", list: List[int]{1, 2, 3, 4, 5}, predicate: lessThanTen, expect: true},
		{name: "MultipleElementsNoneLessThanTen", list: List[int]{10, 11, 12, 13, 14}, predicate: lessThanTen, expect: false},
		{name: "MultipleElementsSomeLessThanTen", list: List[int]{10, 2, 12, 4, 14}, predicate: lessThanTen, expect: true},
		{name: "SingleElementListIsEven", list: List[int]{8}, predicate: isEven, expect: true},
		{name: "SingleElementListIsNotEven", list: List[int]{7}, predicate: isEven, expect: false},
		{name: "MultipleElementsAllAreEven", list: List[int]{2, 4, 6, 8}, predicate: isEven, expect: true},
		{name: "MultipleElementsNoneAreEven", list: List[int]{1, 3, 5, 7}, predicate: isEven, expect: false},
		{name: "MultipleElementsSomeAreEven", list: List[int]{2, 3, 6, 7, 8}, predicate: isEven, expect: true},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			result := test.list.Any(test.predicate)
			assert.Equal(t, test.expect, result)
		})
	}
}

func TestMinWithOrNil(t *testing.T) {
	var compareInts comparator.Comparator[int] = func(a, b int) int {
		return a - b
	}

	tests := []struct {
		name       string
		list       List[int]
		comparator comparator.Comparator[int]
		want       *int
	}{
		{
			name:       "Empty list",
			list:       List[int]{},
			comparator: compareInts,
			want:       nil,
		},
		{
			name:       "Single element list",
			list:       List[int]{5},
			comparator: compareInts,
			want:       util.PointerTo(5),
		},
		{
			name:       "Multiple elements in ascending order",
			list:       List[int]{1, 2, 3, 4, 5},
			comparator: compareInts,
			want:       util.PointerTo(1),
		},
		{
			name:       "Multiple elements in descending order",
			list:       List[int]{5, 4, 3, 2, 1},
			comparator: compareInts,
			want:       util.PointerTo(1),
		},
		{
			name:       "Multiple elements in random order",
			list:       List[int]{3, 1, 5, 2, 4},
			comparator: compareInts,
			want:       util.PointerTo(1),
		},
		{
			name:       "Multiple elements with duplicates",
			list:       List[int]{3, 1, 5, 2, 1, 4},
			comparator: compareInts,
			want:       util.PointerTo(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.MinWithOrNil(tt.comparator)
			assert.Equal(t, tt.want, got, "MinWithOrNil() should return expected result")
		})
	}
}

func TestMaxWithOrNil(t *testing.T) {
	var compareInts comparator.Comparator[int] = func(a, b int) int {
		return a - b
	}

	tests := []struct {
		name       string
		list       List[int]
		comparator comparator.Comparator[int]
		want       *int
	}{
		{
			name:       "Empty list",
			list:       List[int]{},
			comparator: compareInts,
			want:       nil,
		},
		{
			name:       "Single element list",
			list:       List[int]{1},
			comparator: compareInts,
			want:       util.PointerTo(1),
		},
		{
			name:       "Multiple elements in ascending order",
			list:       List[int]{1, 2, 3, 4, 5},
			comparator: compareInts,
			want:       util.PointerTo(5),
		},
		{
			name:       "Multiple elements in descending order",
			list:       List[int]{5, 4, 3, 2, 1},
			comparator: compareInts,
			want:       util.PointerTo(5),
		},
		{
			name:       "Multiple elements in random order",
			list:       List[int]{3, 1, 5, 2, 4},
			comparator: compareInts,
			want:       util.PointerTo(5),
		},
		{
			name:       "Multiple elements with duplicates",
			list:       List[int]{3, 1, 5, 2, 1, 4},
			comparator: compareInts,
			want:       util.PointerTo(5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.MaxWithOrNil(tt.comparator)
			assert.Equal(t, tt.want, got, "MaxWithOrNil() should return expected result")
		})
	}
}
