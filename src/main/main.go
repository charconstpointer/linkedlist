package main

func main() {
	linked := LinkedList()
	linked.addFirst(&item{"first item", nil})
	linked.addLast(&item{"second item", nil})
	linked.addLast(&item{"third item", nil})
	linked.addAtIndex(&item{"at index", nil}, 1)
	printList(&linked)
}
