package main

import (
	"fmt"
	"time"
)

type item struct {
	value int64
	next  *item
}

func printList(list *item) {
	for item := list; item != nil; item = item.next {
		fmt.Println("->", item.value)
		time.Sleep(time.Second)
	}
}

func addLast(list, newItem *item) *item {
	if list == nil {
		return list
	}

	for item := list; item != nil; item = item.next {
		if item.next == nil {
			fmt.Println(item.value)
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

//func get(list *item, index int) int64 {
//	i := 0
//	for item := list; item != nil; item = item.next {
//		i++
//		if i == index {
//			return item.value
//		}
//
//	}
//	return -1
//}
