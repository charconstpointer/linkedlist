package main

func main() {

	linkedList := &list{nil}
	linkedList.addFirst(&item{2, nil})
	linkedList.addFirst(&item{44, nil})
	linkedList.addLast(&item{55, nil})
	printList(linkedList)
}
