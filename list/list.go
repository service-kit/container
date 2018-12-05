package list

import "github.com/service-kit/container/iterator"

type ListIterator struct {
	data     interface{}
	nextIter iterator.Iterator
	preIter  iterator.Iterator
}

func (li *ListIterator) Pre() iterator.Iterator {
	return li.preIter
}

func (li *ListIterator) Next() iterator.Iterator {
	return li.nextIter
}

func (li *ListIterator) SetValue(val interface{}) {
	li.data = val
}

func (li *ListIterator) SetPre(pre iterator.Iterator) {
	li.preIter = pre
}

func (li *ListIterator) SetNext(next iterator.Iterator) {
	li.nextIter = next
}

func (li *ListIterator) Value() interface{} {
	return li.data
}

type ListContainer interface {
	Init()
	First() iterator.Iterator
	End() iterator.Iterator
	Len() int
	Insert(interface{}, iterator.Iterator)
	Delete(iterator.Iterator)
	PushBack(interface{})
	PushFront(interface{})
	InsertBefore(interface{}, iterator.Iterator)
	InsertAfter(interface{}, iterator.Iterator)
	InsertList(ListContainer, iterator.Iterator)
	MoveToFront(iterator.Iterator)
	MoveToBack(iterator.Iterator)
	MoveBefore(iterator.Iterator)
	MoveAfter(iterator.Iterator)
}

type list struct {
	len  int
	head iterator.Iterator
	tail iterator.Iterator
}

func (l *list) Init() {
	l.len = 0
	l.head = new(ListIterator)
	l.tail = new(ListIterator)
	l.head.SetNext(l.tail)
	l.tail.SetPre(l.head)
}

func (l *list) First() iterator.Iterator {
	return l.head.Next()
}

func (l *list) End() iterator.Iterator {
	return l.tail.Pre()
}

func (l *list) Len() int {
	return l.len
}

func (l *list) insert(cur, pos iterator.Iterator) {
	if nil == cur || l.head == pos {
		return
	}
	pos.Pre().SetNext(cur)
	cur.SetPre(pos.Pre())
	cur.SetNext(pos)
	pos.SetPre(cur)
	l.len++
}

func (l *list) delete(pos iterator.Iterator) {
	if l.head == pos || l.tail == pos {
		return
	}
	pos.Pre().SetNext(pos.Next())
	pos.Next().SetPre(pos.Pre())
	pos.SetPre(nil)
	pos.SetNext(nil)
	l.len--
}

func (l *list) Insert(data interface{}, iter iterator.Iterator) {
	if nil == iter || l.head == iter {
		return
	}
	newIter := new(ListIterator)
	newIter.SetValue(data)
	l.insert(newIter, iter)
}

func (l *list) Delete(cur iterator.Iterator) {
	l.delete(cur)
}

func (l *list) PushBack(data interface{}) {
	l.Insert(data, l.tail)
}

func (l *list) PushFront(data interface{}) {
	if nil == l.head {
		return
	}
	l.Insert(data, l.head.Next())
}

func (l *list) InsertBefore(data interface{}, pos iterator.Iterator) {
	l.Insert(data, pos)
}

func (l *list) InsertAfter(data interface{}, pos iterator.Iterator) {
	if nil == pos {
		return
	}
	l.Insert(data, pos.Next())
}

func (l *list) InsertList(lis ListContainer, pos iterator.Iterator) {
	if nil == lis || l.head == pos || 0 == lis.Len() {
		return
	}
	pos.Pre().SetNext(lis.First())
	lis.First().SetPre(pos.Pre())
	lis.End().SetNext(pos)
	pos.SetPre(lis.End())
	l.len += lis.Len()
}

func (l *list) MoveToFront(pos iterator.Iterator) {
	if 0 == l.len || l.head.Next() == pos {
		return
	}
	l.delete(pos)
	l.insert(pos, l.First())
}

func (l *list) MoveToBack(pos iterator.Iterator) {
	if 0 == l.len || l.tail.Pre() == pos {
		return
	}
	l.delete(pos)
	l.insert(pos, l.tail)
}

func (l *list) MoveBefore(pos iterator.Iterator) {
	if 0 == l.len || l.head.Next() == pos {
		return
	}
	pre := pos.Pre()
	l.delete(pos)
	l.insert(pos, pre)
}

func (l *list) MoveAfter(pos iterator.Iterator) {
	if 0 == l.len || l.tail.Pre() == pos {
		return
	}
	next := pos.Next().Next()
	l.delete(pos)
	l.insert(pos, next)
}

func NewList() ListContainer {
	l := new(list)
	l.Init()
	return l
}
