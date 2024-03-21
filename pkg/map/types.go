package _map

type Map[K comparable, V any] map[K]V

type PredicateFunc[K comparable, V any] func(key K, value V) bool
