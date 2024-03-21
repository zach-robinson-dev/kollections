package _map

import "reflect"

func (m *Map[K, V]) Filter(predicate PredicateFunc[K, V]) Map[K, V] {
	result := make(Map[K, V], len(*m))

	for key, value := range *m {
		if predicate(key, value) {
			result[key] = value
		}
	}

	return result
}

func (m *Map[K, V]) Remove(key K, expectedValue V) bool {
	switch value, isPresent := (*m)[key]; {
	case isPresent && reflect.DeepEqual(expectedValue, value):
		delete(*m, key)
		return true
	default:
		return false
	}
}

func (m *Map[K, V]) All(predicate PredicateFunc[K, V]) bool {
	for key, value := range *m {
		if !predicate(key, value) {
			return false
		}
	}

	return true
}

func (m *Map[K, V]) Any(predicate PredicateFunc[K, V]) bool {
	for key, value := range *m {
		if predicate(key, value) {
			return true
		}
	}

	return false
}
