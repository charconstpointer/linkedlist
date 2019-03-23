package main

import (
	"fmt"
)

type item struct {
	value int64
	next  *item
}

func printList(list *item) {
	for item := list; item != nil; item = item.next {
		fmt.Println("->", item.value)
	}
}

func addLast(list, newItem *item) *item {
	if list == nil {
		return list
	}

	for item := list; item != nil; item = item.next {
		if item.next == nil {
			item.next = newItem
			return list
		}
	}
	return list
}

func addFirst(list, item *item) *item {
	item.next = list
	return item
}

func delete(list, toDelete *item) {
	i := 0
	for item := list; item != nil; item = item.next {
		i++
		if item.next == toDelete {
			if toDelete.next == nil {
				item.next = nil
			} else {
				item.next = toDelete.next
			}
		}
	}
	return
}

func get(list *item, index int) int64 {
	i := 0
	for item := list; item != nil; item = item.next {
		i++
		if i == index {
			return item.value
		}

	}
	return -1
}
