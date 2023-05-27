package ds

type List interface {
	// Mutation
	Add(int)
	Insert(e int, index int)
	Remove(int)
	RemoveAt(int)
	Clear()

	// Checking
	Contains(int) bool
	IndexOf(int) int
	Count() int
	//Get(int) int
}

type IntegerList struct {
	elements []int
}

type IntegerListIterator struct {
	currentIndex int
	elements     []int
}

// Iterator
func (iter *IntegerListIterator) HasNext() bool {
	return iter.currentIndex+1 < len(iter.elements)
}

func (iter *IntegerListIterator) Next() int {
	if iter.HasNext() {
		iter.currentIndex++
		return iter.elements[iter.currentIndex]
	}
	// hardcoded, but should raise error
	return -1
}

// Iterable
func (l *IntegerList) GetIterator() Iterator {
	return &IntegerListIterator{
		currentIndex: 0,
		elements:     l.elements,
	}
}

// Mutation
func (l *IntegerList) Add(el int) {
	l.elements = append(l.elements, el)
}

func (l *IntegerList) Insert(e int, index int) {
	if index >= 0 && index < l.Count() {
		l.elements = append(l.elements[:index], append([]int{e}, l.elements[index:]...)...)
	}
}

func (l *IntegerList) Remove(el int) {
	i := l.IndexOf(el)

	if i >= 0 {
		l.elements = append(l.elements[:i], l.elements[i+1:]...)
	}
}

func (l *IntegerList) RemoveAt(index int) {
	if index >= 0 && index < l.Count() {
		l.elements = append(l.elements[:index], l.elements[index+1:]...)
	}
}

func (l *IntegerList) Clear() {
	l.elements = []int{}
}

// Checking
func (l *IntegerList) Contains(el int) bool {
	for _, v := range l.elements {
		if el == v {
			return true
		}
	}

	return false
}

func (l *IntegerList) IndexOf(el int) int {
	for i, v := range l.elements {
		if v == el {
			return i
		}
	}

	return -1
}

func (l *IntegerList) Count() int {
	return len(l.elements)
}

func NewList(values []int) List {
	// hardcoded fixed initial len
	//len := 100
	return &IntegerList{
		elements: values,
	}
}
