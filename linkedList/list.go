package linkedList

type list struct {
	head *node
	tail *node
	size uint
}

type IList interface {
	Len() uint
	GetFirst() *int
	GetLast() *int
	Append(data int)
	Prepend(data int)
	Reverse()
	PopBack() *int
	PopFront() *int
}

func (l *list) Len() uint {
	return l.size
}

func (l *list) GetFirst() *int {
	if l.head == nil {
		return nil
	}

	return &l.head.data
}

func (l *list) GetLast() *int {
	if l.tail == nil {
		return nil
	}

	return &l.tail.data
}

func (l *list) Append(data int) {
	newNode := node{
		data: data,
	}

	l.size += 1

	if l.tail == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}

	newNode.prev = l.tail
	l.tail.next = &newNode
	l.tail = &newNode
}

func (l *list) Prepend(data int) {
	newNode := node{
		data: data,
	}

	l.size += 1

	if l.head == nil {
		l.head = &newNode
		l.tail = &newNode
		return
	}

	newNode.next = l.head
	l.head.prev = &newNode
	l.head = &newNode
}

func (l *list) Reverse() {
	if l.Len() <= 1 {
		return
	}

	l.tail = l.head

	var prev *node
	curr := l.head

	for curr.next != nil {
		next := curr.next

		curr.next = prev
		curr.prev = next

		prev = curr
		curr = next
	}

	curr.next = prev
	curr.prev = nil
	l.head = curr
}

func (l *list) PopBack() *int {
	if l.size == 0 {
		return nil
	}

	result := l.tail.data

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return &result
	}

	oldTail := l.tail
	newTail := l.tail.prev

	l.tail = newTail
	if l.tail != nil {
		l.tail.next = nil
	}

	oldTail.prev = nil

	l.size -= 1

	return &result
}

func (l *list) PopFront() *int {
	if l.size == 0 {
		return nil
	}

	result := l.head.data

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size = 0
		return &result
	}

	oldHead := l.head
	newHead := l.head.next

	l.head = newHead
	if l.head != nil {
		l.head.prev = nil
	}

	oldHead.next = nil

	l.size -= 1

	return &result
}

func New() IList {
	return &list{}
}
