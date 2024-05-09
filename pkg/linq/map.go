package linq

import (
	"errors"
	"slices"
)

type Enumerable[T interface{}] struct {
	Internal []T
}

type FilterFunc[T interface{}] func(T) bool
type MapperFunc[T interface{}, R interface{}] func(T) R
type FlatMapperFunc[T interface{}, R interface{}] func(T) []R

func NewEnumerable[T interface{}](arr []T) Enumerable[T] {
	return Enumerable[T]{Internal: arr}
}

func Map[T interface{}, R interface{}](arr []T, mapperFunc MapperFunc[T, R]) []R {
	result := []R{}
	for _, element := range arr {
		result = append(result, mapperFunc(element))
	}
	return result
}

func FlatMap[T interface{}, R interface{}](arr []T, flatMapperFunc FlatMapperFunc[T, R]) []R {
	result := []R{}
	for _, element := range arr {
		result = append(result, flatMapperFunc(element)...)
	}
	return result
}

func Filter[T interface{}](arr []T, filterFunc FilterFunc[T]) []T {
	result := []T{}
	for _, element := range arr {
		if filterFunc(element) {
			result = append(result, element)
		}
	}
	return result
}

func AnyMatch[T interface{}](arr []T, matchFunc FilterFunc[T]) bool {
	for _, element := range arr {
		if matchFunc(element) {
			return true
		}
	}
	return false
}

func AllMatch[T interface{}](arr []T, matchFunc FilterFunc[T]) bool {
	for _, element := range arr {
		if !matchFunc(element) {
			return false
		}
	}
	return true
}

func ForEach[T interface{}](arr []T, matchFunc func(T)) {
	for _, element := range arr {
		matchFunc(element)
	}
}

func First[T interface{}](arr []T, matchFunc FilterFunc[T]) (T, error) {
	index := slices.IndexFunc(arr, matchFunc)
	if index == -1 {
		var t T
		return t, errors.New("squence does not contain matching element")
	}
	return arr[index], nil
}

func FirstOrDefault[T interface{}](arr []T, matchFunc FilterFunc[T], defaultVal T) T {
	result, err := First(arr, matchFunc)
	if err != nil {
		return defaultVal
	}
	return result
}

func Last[T interface{}](arr []T, matchFunc FilterFunc[T]) (T, error) {
	cloned := slices.Clone(arr)
	slices.Reverse(cloned)
	return First(cloned, matchFunc)
}

func LastOrDefault[T interface{}](arr []T, matchFunc FilterFunc[T], defaultVal T) T {
	cloned := slices.Clone(arr)
	slices.Reverse(cloned)
	return FirstOrDefault(cloned, matchFunc, defaultVal)
}
