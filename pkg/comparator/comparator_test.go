package comparator

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zach-robinson-dev/kollections/pkg/util"
)

func TestAscendingOrder(t *testing.T) {
	type testCase[C interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
	}] struct {
		name string
		a    int
		b    int
		want int
	}
	tests := []testCase[int]{
		{
			name: "less_than",
			a:    1,
			b:    2,
			want: -1,
		},
		{
			name: "equal_to",
			a:    1,
			b:    1,
			want: 0,
		},
		{
			name: "greater_than",
			a:    2,
			b:    1,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AscendingOrder[int]()(tt.a, tt.b))
		})
	}
}

func TestCaseInsensitiveOrder(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		{
			name: "less_than",
			a:    "a",
			b:    "B",
			want: -1,
		},
		{
			name: "equal_to",
			a:    "a",
			b:    "A",
			want: 0,
		},
		{
			name: "greater_than",
			a:    "B",
			b:    "a",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CaseInsensitiveOrder()(tt.a, tt.b))
		})
	}
}

func TestComparator_Reversed(t *testing.T) {
	type testCase[C interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
	}] struct {
		name string
		a    int
		b    int
		want int
	}
	tests := []testCase[int]{
		{
			name: "less_than",
			b:    2,
			a:    1,
			want: 1,
		},
		{
			name: "equal_to",
			b:    1,
			a:    1,
			want: 0,
		},
		{
			name: "greater_than",
			b:    1,
			a:    2,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, AscendingOrder[int]().Reversed()(tt.a, tt.b))
		})
	}
}

func TestComparator_Then(t *testing.T) {
	type someStruct struct {
		first  int
		second int
	}
	ascendingOrderFirstElement := AscendingOrderBy(func(it someStruct) int { return it.first })
	descendingOrderSecondElement := DescendingOrderBy(func(it someStruct) int { return it.second })
	ascendingOrderFirstElementThenSecondElement := ascendingOrderFirstElement.Then(descendingOrderSecondElement)
	type testCase struct {
		name string
		a    someStruct
		b    someStruct
		want int
	}
	tests := []testCase{
		{
			name: "first_element_less_than",
			a:    someStruct{first: 1},
			b:    someStruct{first: 2},
			want: -1,
		},
		{
			name: "second_element_less_than",
			a:    someStruct{first: 1, second: 1},
			b:    someStruct{first: 1, second: 2},
			want: 1,
		},
		{
			name: "second_element_equal_to",
			a:    someStruct{first: 1, second: 1},
			b:    someStruct{first: 1, second: 1},
			want: 0,
		},
		{
			name: "second_element_greater_than",
			a:    someStruct{first: 1, second: 2},
			b:    someStruct{first: 1, second: 1},
			want: -1,
		},
		{
			name: "first_element_greater_than",
			a:    someStruct{first: 2},
			b:    someStruct{first: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ascendingOrderFirstElementThenSecondElement(tt.a, tt.b))
		})
	}
}

func TestComparator_ThenDescending(t *testing.T) {
	type someStruct struct {
		first  int
		second int
	}
	ascendingOrderFirstElement := AscendingOrderBy(func(it someStruct) int { return it.first })
	descendingOrderSecondElement := DescendingOrderBy(func(it someStruct) int { return it.second })
	ascendingOrderFirstElementThenSecondElementDescending := ascendingOrderFirstElement.ThenDescending(descendingOrderSecondElement)
	type testCase struct {
		name string
		a    someStruct
		b    someStruct
		want int
	}
	tests := []testCase{
		{
			name: "first_element_less_than",
			a:    someStruct{first: 1},
			b:    someStruct{first: 2},
			want: -1,
		},
		{
			name: "second_element_less_than",
			a:    someStruct{first: 1, second: 1},
			b:    someStruct{first: 1, second: 2},
			want: -1,
		},
		{
			name: "second_element_equal_to",
			a:    someStruct{first: 1, second: 1},
			b:    someStruct{first: 1, second: 1},
			want: 0,
		},
		{
			name: "second_element_greater_than",
			a:    someStruct{first: 1, second: 2},
			b:    someStruct{first: 1, second: 1},
			want: 1,
		},
		{
			name: "first_element_greater_than",
			a:    someStruct{first: 2},
			b:    someStruct{first: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ascendingOrderFirstElementThenSecondElementDescending(tt.a, tt.b))
		})
	}
}

func TestDescendingOrder(t *testing.T) {
	type testCase[C interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
	}] struct {
		name string
		a    int
		b    int
		want int
	}
	tests := []testCase[int]{
		{
			name: "less_than",
			b:    2,
			a:    1,
			want: 1,
		},
		{
			name: "equal_to",
			b:    1,
			a:    1,
			want: 0,
		},
		{
			name: "greater_than",
			b:    1,
			a:    2,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DescendingOrder[int]()(tt.a, tt.b))
		})
	}
}

func TestZeroValuesFirst(t *testing.T) {
	type testCase struct {
		name string
		a    *bool
		b    *bool
		want int
	}
	tests := []testCase{
		{
			name: "both_zero",
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			name: "a_is_zero",
			a:    nil,
			b:    util.PointerTo(false),
			want: -1,
		},
		{
			name: "b_is_zero",
			a:    util.PointerTo(true),
			b:    nil,
			want: 1,
		},
		{
			name: "both_non_zero",
			a:    util.PointerTo(false),
			b:    util.PointerTo(true),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ZeroValuesFirst[*bool]()(tt.a, tt.b))
		})
	}
}

func TestZeroValuesLast(t *testing.T) {
	type testCase struct {
		name string
		a    *bool
		b    *bool
		want int
	}
	tests := []testCase{
		{
			name: "both_zero",
			a:    nil,
			b:    nil,
			want: 0,
		},
		{
			name: "a_is_zero",
			a:    nil,
			b:    util.PointerTo(false),
			want: 1,
		},
		{
			name: "b_is_zero",
			a:    util.PointerTo(true),
			b:    nil,
			want: -1,
		},
		{
			name: "both_non_zero",
			a:    util.PointerTo(false),
			b:    util.PointerTo(true),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ZeroValuesLast[*bool]()(tt.a, tt.b))
		})
	}
}

func TestBySelector(t *testing.T) {
	type someStruct struct {
		value int
	}
	selector := func(it someStruct) int { return it.value }
	comparator := BySelector[someStruct, int](AscendingOrder[int](), selector)
	type testCase struct {
		name string
		a    someStruct
		b    someStruct
		want int
	}
	tests := []testCase{
		{
			name: "less_than",
			a:    someStruct{value: 1},
			b:    someStruct{value: 2},
			want: -1,
		},
		{
			name: "equal_to",
			a:    someStruct{value: 1},
			b:    someStruct{value: 1},
			want: 0,
		},
		{
			name: "greater_than",
			a:    someStruct{value: 2},
			b:    someStruct{value: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, comparator(tt.a, tt.b))
		})
	}
}

func TestZeroValuesFirstBy(t *testing.T) {
	type someStruct struct {
		value *bool
	}
	selector := func(it someStruct) *bool { return it.value }
	comparator := ZeroValuesFirstBy[someStruct, *bool](selector)
	type testCase struct {
		name string
		a    someStruct
		b    someStruct
		want int
	}
	tests := []testCase{
		{
			name: "both_zero",
			a:    someStruct{value: nil},
			b:    someStruct{value: nil},
			want: 0,
		},
		{
			name: "a_is_zero",
			a:    someStruct{value: nil},
			b:    someStruct{value: util.PointerTo(false)},
			want: -1,
		},
		{
			name: "b_is_zero",
			a:    someStruct{value: util.PointerTo(true)},
			b:    someStruct{value: nil},
			want: 1,
		},
		{
			name: "both_non_zero",
			a:    someStruct{value: util.PointerTo(false)},
			b:    someStruct{value: util.PointerTo(true)},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, comparator(tt.a, tt.b))
		})
	}
}

func TestCaseInsensitiveOrderBy(t *testing.T) {
	type someStruct struct {
		value string
	}
	selector := func(it someStruct) string { return it.value }
	tests := []struct {
		name string
		a    someStruct
		b    someStruct
		want int
	}{
		{
			name: "less_than",
			a:    someStruct{value: "a"},
			b:    someStruct{value: "B"},
			want: -1,
		},
		{
			name: "equal_to",
			a:    someStruct{value: "a"},
			b:    someStruct{value: "A"},
			want: 0,
		},
		{
			name: "greater_than",
			a:    someStruct{value: "B"},
			b:    someStruct{value: "a"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, CaseInsensitiveOrderBy(selector)(tt.a, tt.b))
		})
	}
}

func TestZeroValuesLastBy(t *testing.T) {
	type someStruct struct {
		value *bool
	}
	selector := func(it someStruct) *bool { return it.value }
	tests := []struct {
		name string
		a    someStruct
		b    someStruct
		want int
	}{
		{
			name: "both_zero",
			a:    someStruct{value: nil},
			b:    someStruct{value: nil},
			want: 0,
		},
		{
			name: "a_is_zero",
			a:    someStruct{value: nil},
			b:    someStruct{value: util.PointerTo(false)},
			want: 1,
		},
		{
			name: "b_is_zero",
			a:    someStruct{value: util.PointerTo(true)},
			b:    someStruct{value: nil},
			want: -1,
		},
		{
			name: "both_non_zero",
			a:    someStruct{value: util.PointerTo(false)},
			b:    someStruct{value: util.PointerTo(true)},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ZeroValuesLastBy(selector)(tt.a, tt.b))
		})
	}
}
