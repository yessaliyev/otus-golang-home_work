package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	Clear()
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Key   Key
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(data interface{}) *ListItem {
	item := &ListItem{Value: data}

	if l.head == nil {
		l.head = item
		l.tail = item
	} else {
		item.Next = l.head
		l.head.Prev = item
		l.head = item
	}

	l.len++

	return item
}

func (l *list) PushBack(data interface{}) *ListItem {
	item := &ListItem{
		Value: data,
	}

	if l.tail == nil {
		l.head = item
		l.tail = item
	} else {
		item.Prev = l.tail
		l.tail.Next = item
		l.tail = item
	}

	l.len++

	return item
}

func (l *list) Remove(i *ListItem) {
	l.len--

	if i.Next == nil {
		i.Prev.Next = nil
		return
	}

	if i.Prev == nil {
		i.Next.Prev = nil
		return
	}

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	// обрываем связи
	i.Prev.Next = i.Next

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Next = l.head
	l.head.Prev = i
	l.head = i
}

func (l *list) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}
