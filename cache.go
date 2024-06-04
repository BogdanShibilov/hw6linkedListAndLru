package main

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	sync.RWMutex
}

func (l *lruCache) Set(key Key, value any) bool {
	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()
	if item, ok := l.items[key]; ok {
		l.replaceItemValue(item, value)
		return true
	} else {
		l.addNewItem(key, value)
		return false
	}
}

func (l *lruCache) replaceItemValue(replacedItem *ListItem, newValue any) {
	replacedItem.Value = newValue
	l.queue.MoveToFront(replacedItem)
}

func (l *lruCache) addNewItem(key Key, value any) {
	if l.queue.Len() >= l.capacity {
		l.removeOldestItem()
	}
	newItem := l.queue.PushFront(value)
	l.items[key] = newItem
}

func (l *lruCache) removeOldestItem() {
	oldestItem := l.queue.Back()
	oldestItemKey := l.findKeyOfItem(oldestItem)
	l.queue.Remove(oldestItem)
	delete(l.items, oldestItemKey)
}

func (l *lruCache) findKeyOfItem(item *ListItem) Key {
	for key, mapItem := range l.items {
		if item == mapItem {
			return key
		}
	}
	return ""
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.RWMutex.RLock()
	defer l.RWMutex.RUnlock()
	if item, ok := l.items[key]; ok {
		l.queue.MoveToFront(item)
		return item.Value, true
	} else {
		return nil, false
	}
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem)
	l.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
