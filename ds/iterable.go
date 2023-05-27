package ds

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type Iterable[T any] interface {
	GetIterator() Iterator[T]
}

type GenericIterator[T any] struct {
	currentIndex int
	collection   []T
}

func (iter *GenericIterator[T]) HasNext() bool {
	return iter.currentIndex+1 < len(iter.collection)
}

func (iter *GenericIterator[T]) Next() T {
	if iter.HasNext() {
		iter.currentIndex++
	}
	return iter.collection[iter.currentIndex]
}

// (value, index)
type MapPredicate[T any] func(T, ...int) T

// (accumulator, value)
type ReducePredicate[T any] func(T, T) T

// (value)
type BoolPredicate[T any] func(T) bool

// Query
func Map[T any](col Iterable[T], p MapPredicate[T]) []T {
	resultMap := []T{}
	iter := col.GetIterator()
	index := 0

	for iter.HasNext() {
		resultMap = append(resultMap, p(iter.Next(), index))
		index++
	}

	return resultMap
}

func Filter[T any](col Iterable[T], p BoolPredicate[T]) []T {
	resultFilter := []T{}
	iter := col.GetIterator()

	for iter.HasNext() {
		next := iter.Next()
		if p(next) {
			resultFilter = append(resultFilter, next)
		}
	}

	return resultFilter
}

func Reduce[T any](col Iterable[T], p ReducePredicate[T]) T {
	iter := col.GetIterator()
	result := iter.Next()

	for iter.HasNext() {
		result = p(result, iter.Next())
	}

	return result
}

func Some[T any](col Iterable[T], p BoolPredicate[T]) bool {
	iter := col.GetIterator()

	for iter.HasNext() {
		if p(iter.Next()) {
			return true
		}
	}

	return false
}

func Every[T any](col Iterable[T], p BoolPredicate[T]) bool {
	iter := col.GetIterator()

	for iter.HasNext() {
		if !p(iter.Next()) {
			return false
		}
	}

	return true
}
