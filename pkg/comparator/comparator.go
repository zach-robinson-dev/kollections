package comparator

import (
	"strings"

	"github.com/zach-robinson-dev/kollections/pkg/util"
)

type Comparator[T any] func(a T, b T) int

type Selector[X any, Y any] func(it X) Y

func (c Comparator[T]) Reversed() Comparator[T] {
	return func(a T, b T) int {
		return c(a, b) * -1
	}
}

func (c Comparator[T]) Then(comparator Comparator[T]) Comparator[T] {
	return func(a T, b T) int {
		switch res := c(a, b); res {
		case 0:
			return comparator(a, b)
		default:
			return res
		}
	}
}

func (c Comparator[T]) ThenDescending(comparator Comparator[T]) Comparator[T] {
	return c.Then(comparator.Reversed())
}

func BySelector[X any, Y any](comparator Comparator[Y], selector Selector[X, Y]) Comparator[X] {
	return func(a X, b X) int {
		return comparator(selector(a), selector(b))
	}
}

func AscendingOrder[C int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string]() Comparator[C] {
	return func(a C, b C) int {
		switch {
		case a == b:
			return 0
		case a < b:
			return -1
		default:
			return 1
		}
	}
}

func AscendingOrderBy[T any, C int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string](selector Selector[T, C]) Comparator[T] {
	return BySelector[T, C](AscendingOrder[C](), selector)
}

func DescendingOrder[C int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string]() Comparator[C] {
	return AscendingOrder[C]().Reversed()
}

func DescendingOrderBy[T any, C int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string](selector Selector[T, C]) Comparator[T] {
	return BySelector[T, C](DescendingOrder[C](), selector)
}

func CaseInsensitiveOrder() Comparator[string] {
	return func(a string, b string) int {
		return AscendingOrder[string]()(strings.ToLower(a), strings.ToLower(b))
	}
}

func CaseInsensitiveOrderBy[T any](selector Selector[T, string]) Comparator[T] {
	return BySelector[T, string](CaseInsensitiveOrder(), selector)
}

func ZeroValuesFirst[T any]() Comparator[T] {
	return func(a T, b T) int {
		switch aIsZero, bIsZero := func() (bool, bool) { return util.IsZero(a), util.IsZero(b) }(); {
		case aIsZero == bIsZero:
			return 0
		case aIsZero:
			return -1
		default:
			return 1
		}
	}
}

func ZeroValuesFirstBy[X any, Y any](selector Selector[X, Y]) Comparator[X] {
	return BySelector[X, Y](ZeroValuesFirst[Y](), selector)
}

func ZeroValuesLast[T any]() Comparator[T] {
	return ZeroValuesFirst[T]().Reversed()
}

func ZeroValuesLastBy[X any, Y any](selector Selector[X, Y]) Comparator[X] {
	return BySelector[X, Y](ZeroValuesLast[Y](), selector)
}
