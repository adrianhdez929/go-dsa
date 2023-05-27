package ds

type Iterator interface {
	HasNext() bool
	Next() int
}

type Iterable interface {
	GetIterator() Iterator
}

// (value, index)
type MapPredicate func(int, ...int) int

// (accumulator, value)
type ReducePredicate func(int, int) int

// (value)
type BoolPredicate func(int) bool

func Map(col Iterable, p MapPredicate) []int {
	resultMap := []int{}
	iter := col.GetIterator()
	index := 0

	for iter.HasNext() {
		resultMap = append(resultMap, p(iter.Next(), index))
		index++
	}

	return resultMap
}

func Filter(col Iterable, p BoolPredicate) []int {
	resultFilter := []int{}
	iter := col.GetIterator()

	for iter.HasNext() {
		next := iter.Next()
		if p(next) {
			resultFilter = append(resultFilter, next)
		}
	}

	return resultFilter
}

func Reduce(col Iterable, p ReducePredicate, initial int) int {
	result := initial
	iter := col.GetIterator()

	for iter.HasNext() {
		result = p(result, iter.Next())
	}

	return result
}

func Some(col Iterable, p BoolPredicate) bool {
	iter := col.GetIterator()

	for iter.HasNext() {
		if p(iter.Next()) {
			return true
		}
	}

	return false
}

func Every(col Iterable, p BoolPredicate) bool {
	iter := col.GetIterator()

	for iter.HasNext() {
		if !p(iter.Next()) {
			return false
		}
	}

	return true
}
