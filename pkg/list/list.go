package list

type List[T any] []T

func (l *List[T]) Filter(predicate PredicateFunc[T]) List[T] {
	result := make(List[T], 0, len(*l))

	for _, item := range *l {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
