package ds

type List[T comparable] interface {
	// Mutation
	Add(T)
	Insert(e T, index int)
	Remove(T)
	RemoveAt(int)
	Clear()

	// Checking
	Contains(T) bool
	IndexOf(T) int
	Count() int
	Get(int) T
}

type IntegerList struct {
	elements []int
}

// Iterable
func (l *IntegerList) GetIterator() Iterator[int] {
	return &GenericIterator[int]{
		currentIndex: 0,
		collection:   l.elements,
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

func (l *IntegerList) Get(index int) int {
	return l.elements[index]
}

func NewIntegerList(values []int) List[int] {
	// hardcoded fixed initial len
	//len := 100
	return &IntegerList{
		elements: values,
	}
}
