package main

import "fmt"

func main() {
	a := &item{1, nil}
	c := &item{3, a}
	b := &item{2, c}

	list := b
	list = addLast(list, &item{4, nil})
	list = addFirst(list, &item{55, nil})
	fmt.Println(get(list, 2))
	deleteAtIndex(list, 2)
	printList(list)
}
