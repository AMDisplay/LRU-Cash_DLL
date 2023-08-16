package main

type ListInterface interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	next, prev *ListItem
	data       any
	list       *List
}

type List struct {
	root ListItem // Абстрактный объект
	len  int
}

// длина списка
func (l *List) Len() int {
	return l.len
}

// первый элемент списка
func (l *List) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// последний элемент списка
func (l *List) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// добавить значение в начало
func (l *List) PushFront(v interface{}) *ListItem {
	newItems := &ListItem{
		next: nil,
		prev: nil,
		data: v,
		list: l,
	}
	if l.len == 0 {
		l.root.next = newItems
		l.root.prev = newItems
	} else {
		firstItem := l.Front()
		firstItem.prev = newItems
		newItems.next = firstItem
		l.root.next = newItems
	}
	l.len++
	return newItems

}

// добавить значение в конец
func (l *List) PushBack(v interface{}) *ListItem {
	newItems := &ListItem{
		next: nil,
		prev: nil,
		data: v,
		list: l,
	}
	if l.len == 0 {
		l.root.next = newItems
		l.root.prev = newItems
	} else {
		LastItem := l.Back()
		LastItem.next = newItems
		newItems.prev = LastItem
		l.root.prev = newItems
	}
	l.len++
	return newItems

}

// удалить элемент
func (l *List) Remove(i *ListItem) {
	nextItem := i.next
	prevItem := i.prev
	nextItem.prev = prevItem
	prevItem.next = nextItem
	i.next = nil
	i.prev = nil
}

// переместить элемент в начало
func (l *List) MoveToFront(i *ListItem) {
	nextItem := i.next
	prevItem := i.prev
	nextItem.prev = prevItem
	prevItem.next = nextItem
	i.prev = nil
	FirstElem := l.Front()
	FirstElem.prev = i
	i.next = FirstElem
	l.root.next = i
}
