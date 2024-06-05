package main

import "fmt"

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
	Value any
	Next  *ListItem
	Prev  *ListItem
}

type LinkedList struct {
	length    int
	frontItem *ListItem
	backItem  *ListItem
}

func NewList() *LinkedList {
	return new(LinkedList)
}

func (l *LinkedList) Len() int {
	return l.length
}

func (l *LinkedList) Front() *ListItem {
	return l.frontItem
}

func (l *LinkedList) Back() *ListItem {
	return l.backItem
}

func (l *LinkedList) PushFront(value any) *ListItem {
	newItem := &ListItem{Value: value}

	l.pushFront(newItem)

	l.length++
	return newItem
}

func (l *LinkedList) pushFront(newItem *ListItem) {
	if l.length == 0 {
		l.frontItem = newItem
		l.backItem = newItem
	} else {
		l.frontItem.Prev = newItem
		newItem.Next = l.frontItem
		l.frontItem = newItem
	}
}

func (l *LinkedList) PushBack(value any) *ListItem {
	newItem := &ListItem{Value: value}

	l.pushBack(newItem)

	l.length++
	return newItem
}

func (l *LinkedList) pushBack(newItem *ListItem) {
	if l.length == 0 {
		l.frontItem = newItem
		l.backItem = newItem
	} else {
		l.backItem.Next = newItem
		newItem.Prev = l.backItem
		l.backItem = newItem
	}
}

func (l *LinkedList) Remove(i *ListItem) {
	if l.length == 1 {
		l.Clear()
		return
	} else if l.frontItem == i {
		l.frontItem.Next.Prev = nil
		l.frontItem = l.frontItem.Next
	} else if l.backItem == i {
		l.backItem.Prev.Next = nil
		l.backItem = l.backItem.Prev
	} else {
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.length--
}

func (l *LinkedList) MoveToFront(i *ListItem) {
	l.Remove(i)
	i.Prev = nil
	l.length++
	l.pushFront(i)
}

func (l *LinkedList) Clear() {
	l.frontItem = nil
	l.backItem = nil
	l.length = 0
}

func (l *LinkedList) Print() {
	fmt.Print("Front ")
	for item := l.Back(); item != nil; item = item.Prev {
		fmt.Printf("%v ", item.Value)
	}
	fmt.Print(" Back\n")
}
