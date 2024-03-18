package list

type List[T any] []T

type PredicateFunc[T any] func(item T) bool
