package list

type List[T any] []T

type PredicateFunc[T any] func(item T) bool

type TransformFunc[T any, R any] func(item T) R
