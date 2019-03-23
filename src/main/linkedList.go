package main

import (
	"fmt"
	"sync"
)

type list struct {
	start *item
}

type item struct {
	value int64
	next  *item
}

func printList(list *list) {
	i := 0
	for item := list.start; item != nil; item = item.next {
		fmt.Printf("[%v] -> %v\n", i, item.value)
		i++
	}
}

func (l *list) addLast(item *item) {
	if l.start == nil {
		l.start = item
		return
	}

	for i := l.start; i != nil; i = i.next {
		if i.next == nil {
			i.next = item
			return
		}
	}
}

func (l *list) addFirst(item *item) {
	if l.start == nil {
		l.start = item
	} else {
		item.next = l.start
		l.start = item
	}

}

func delete(list, toDelete *item) {
	i := 0
	for item := list; item != nil; item = item.next {
		i++
		if item.next == toDelete {
			mutex := sync.Mutex{}
			mutex.Lock()
			if toDelete.next == nil {
				item.next = nil
			} else {
				item.next = toDelete.next
			}
			mutex.Unlock()
		}
	}
	return
}

func deleteAtIndex(list *item, index int) {
	prev, i := list, 0
	for list.next != nil {
		if i == index {
			if list.next == nil {
				prev.next = nil
			} else {
				prev.next = list.next
			}
			return
		}
		prev = list
		list = list.next
		i++
	}
	return
}

func get(list *item, index int) int64 {
	i := 0
	for item := list; item != nil; item = item.next {
		if i == index {
			return item.value
		}
		i++
	}
	return -1
}
