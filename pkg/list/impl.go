package list

import (
	"reflect"

	"github.com/zach-robinson-dev/kollections/pkg/comparator"
)

func (l *List[T]) Filter(predicate PredicateFunc[T]) List[T] {
	result := make(List[T], 0, len(*l))

	for _, element := range *l {
		if predicate(element) {
			result = append(result, element)
		}
	}

	return result
}

func (l *List[T]) Contains(t T) bool {
	for _, element := range *l {
		if reflect.DeepEqual(t, element) {
			return true
		}
	}

	return false
}

func (l *List[T]) RemoveAll(elements ...T) bool {
	elementsToRemove := List[T](elements)

	result := make(List[T], 0, len(*l))

	wasModified := false

	for _, element := range *l {
		switch elementsToRemove.Contains(element) {
		case true:
			wasModified = true
		default:
			result = append(result, element)
		}
	}

	*l = result

	return wasModified
}

func (l *List[T]) All(predicate PredicateFunc[T]) bool {
	for _, element := range *l {
		if !predicate(element) {
			return false
		}
	}

	return true
}

func (l *List[T]) Any(predicate PredicateFunc[T]) bool {
	for _, element := range *l {
		if predicate(element) {
			return true
		}
	}

	return false
}

func (l *List[T]) MinWithOrNil(comparator comparator.Comparator[T]) *T {
	if len(*l) == 0 {
		return nil
	}

	minEntry := (*l)[0]

	for _, element := range *l {
		if comparator(element, minEntry) < 0 {
			minEntry = element
		}
	}

	return &minEntry
}

func (l *List[T]) MinWith(comparator comparator.Comparator[T]) T {
	switch result := l.MinWithOrNil(comparator); result {
	case nil:
		var zero T
		return zero
	default:
		return *result
	}
}
