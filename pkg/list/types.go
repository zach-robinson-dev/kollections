package list

type List[T any] []T

type PredicateFunc[T any] func(item T) bool

type TransformFunc[T any, R any] func(item T) R

// Errors

type NoSuchElementError struct{}

func (e NoSuchElementError) Error() string {
	return "no element matching the provided predicate exists within the list"
}

type TooManyMatchingElementsError struct{}

func (e TooManyMatchingElementsError) Error() string {
	return "more than one element matching the provided predicate exists within the list"
}
