package list

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type List[T any] struct {
	head *Node[T]
	size int
}

func New[T any]() *List[T] {
	return &List[T]{
		head: nil,
		size: 0,
	}
}

func (l *List[T]) InsertAtEnd(data T) {
	newNode := &Node[T]{Data: data, Next: nil}
	
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.size++
}

func (l *List[T]) InsertAtBeginning(data T) {
	newNode := &Node[T]{Data: data, Next: l.head}
	l.head = newNode
	l.size++
}

func (l *List[T]) RemoveFirst() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}
	
	data := l.head.Data
	l.head = l.head.Next
	l.size--
	return data, true
}

func (l *List[T]) Display(separator string) string {
	if l.head == nil {
		return ""
	}
	
	result := ""
	current := l.head
	for current != nil {
		if result != "" {
			result += separator
		}
		result += fmt.Sprintf("%v", current.Data)
		current = current.Next
	}
	return result
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}
