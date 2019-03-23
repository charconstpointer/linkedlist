package main

func main() {
	first := &item{1, nil}
	list := first
	second := &item{2, nil}
	list = addFirst(list, second)
	list = addNode(list, &item{3, nil})
	print(list)
}
