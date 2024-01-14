package lists

type node[T comparable] struct {
	Value T
	Next  *node[T]
}

type linkedList[T comparable] struct {
	head *node[T]
}

func (list *linkedList[T]) Add(value T) {
	if list.IsEmpty() {
		list.head = &node[T]{
			Value: value,
			Next:  nil,
		}
	} else {
		element := list.head
		for element.Next != nil {
			element = element.Next
		}

		element.Next = &node[T]{
			Value: value,
			Next:  nil,
		}
	}

}

func (list *linkedList[T]) IsEmpty() bool {
	return list.Size() == 0
}

func (list *linkedList[T]) Size() int {
	count := 0

	element := list.head
	for element != nil {
		element = element.Next
		count++
	}

	return count
}

func NewLinkedList[T comparable](values ...T) *linkedList[T] {
	list := linkedList[T]{}
	for _, value := range values {
		list.Add(value)
	}

	return &list
}
