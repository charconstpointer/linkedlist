package main

func main() {
	a := &item{2, nil}
	linkedList := &list{nil}
	linkedList.addFirst(a)
	linkedList.addFirst(&item{44, nil})
	linkedList.addLast(&item{55, nil})
	//linkedList.delete(a)
	//linkedList.deleteAtIndex(2)
	linkedList.addAtIndex(&item{66, nil}, 2)
	printList(linkedList)
}
