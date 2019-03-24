package main

import (
	"fmt"
	"sync"
)

type list struct {
	start *item
	//end   *item TBI
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

func (l *list) addAtIndex(itemToAdd *item, index int) {
	prev, i := l.start, 0
	for item := l.start; item != nil; item = item.next {
		if i == index {
			prev.next = itemToAdd
			itemToAdd.next = item
		}
		prev = item
		i++
	}
}

func (l *list) delete(toDelete *item) {
	i := 0
	for item := l.start; item != nil; item = item.next {
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

func (l *list) deleteAtIndex(index int) {
	prev, i := l.start, 0
	for item := l.start; item != nil; item = item.next {
		if i == index {
			prev.next = item.next
		}
		prev = item
		i++
	}
}

func (l *list) get(index int) *item {
	i := 0
	for item := l.start; item != nil; item = item.next {
		if i == index {
			return item
		}
		i++
	}
	return nil
}
