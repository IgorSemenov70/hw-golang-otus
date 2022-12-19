package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size  int
	front *ListItem
	back  *ListItem
}

func NewListItem(value interface{}, next *ListItem, prev *ListItem) *ListItem {
	return &ListItem{
		value,
		next,
		prev,
	}
}

func NewList() List {
	return new(list)
}

func (list *list) Len() int {
	// Возвращает длину списка
	return list.size
}

func (list *list) Front() *ListItem {
	// Возвращает первый элемент списка
	return list.front
}

func (list *list) Back() *ListItem {
	// Возвращает последний элемент списка
	return list.back
}

func (list *list) PushFront(value interface{}) *ListItem {
	// Добавляет элемент в начало списка
	listItem := NewListItem(value, list.front, nil)

	if list.front != nil {
		list.front.Prev = listItem
	}
	if list.back == nil {
		list.back = listItem
	}
	list.front = listItem
	list.size++

	return listItem
}

func (list *list) PushBack(value interface{}) *ListItem {
	// Добавляет элемент в конец списка
	listItem := NewListItem(value, nil, list.back)

	if list.back != nil {
		list.back.Next = listItem
	}
	if list.front == nil {
		list.front = listItem
	}
	list.back = listItem
	list.size++

	return listItem
}

func (list *list) Remove(i *ListItem) {
	// Удаляет элемент по индексу
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Prev == nil {
		list.front = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Next == nil {
		list.back = i.Prev
	}
	if list.size != 0 {
		list.size--
	}
}

func (list *list) MoveToFront(i *ListItem) {
	// Перемещает элемент в начало списка
	if i.Prev != nil {
		value := i.Value
		list.Remove(i)
		list.PushFront(value)
	}
}
