package util

import "reflect"

func IsZero(value any) bool {
	reflectValue := reflect.ValueOf(value)
	return !reflectValue.IsValid() || reflectValue.IsZero()
}
