package main

func main() {
	a := &item{1, nil}
	c := &item{3, a}
	b := &item{2, c}

	list := b
	list = addLast(list, &item{4, nil})
	printList(list)
}
