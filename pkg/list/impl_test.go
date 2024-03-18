package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
