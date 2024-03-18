package list

import "reflect"

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
