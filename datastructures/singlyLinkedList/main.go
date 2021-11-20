package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (l *linkedList) prepend(n *node) {
	n.next = l.head
	l.head = n

}

func (l linkedList) printElements() {
	n := l.head

	for n.next != nil {
		fmt.Printf("%d\n", n.data)
		n = n.next
	}

	fmt.Printf("%d\n", n.data)

}

func (l *linkedList) append(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		current := l.head

		for current.next != nil {
			current = current.next
		}

		current.next = n

	}

}

func (l *linkedList) deleteNode(position int32) {
	// Write your code here
	head := l.head

	if head == nil {
		fmt.Println("deletion not possible")
		return
	}

	length := l.size()
	if position > int32(length) {
		fmt.Println("deletion not possible.")
		return
	}

	if position == 0 {
		n := head
		head.next = nil
		head = n
	} else {
		current := head
		next := head.next
		prev := head

		for i := 0; i < int(position); i++ {
			prev = current
			current = current.next
			next = current.next
		}
		current.next = nil
		prev.next = next
	}

}

func (l *linkedList) size() int {
	current := l.head
	length := 0
	for current.next != nil {
		current = current.next
		length++
	}

	return length + 1

}

func (l *linkedList) insertNode(position int, n *node) {
	length := l.size()
	if position > length {
		fmt.Println("Insertion not possible")
		return
	}

	if position == 0 {
		n.next = l.head
		l.head = n
	} else {
		current := l.head
		for i := 1; i < position; i++ {
			current = current.next
		}
		temp := current.next
		current.next = n
		n.next = temp

	}

}

func main() {
	mylist := linkedList{}

	node1 := &node{data: 11}
	node2 := &node{data: 12}
	node3 := &node{data: 8}
	node4 := &node{data: 18}
	node5 := &node{data: 16}
	node6 := &node{data: 5}
	node7 := &node{data: 18}
	node8 := &node{data: 0}
	mylist.append(node1)
	mylist.append(node2)
	mylist.append(node3)
	mylist.append(node4)
	mylist.append(node5)
	mylist.append(node6)
	mylist.append(node7)
	mylist.append(node8)
	mylist.printElements()
	mylist.deleteNode(7)
	mylist.printElements()

}
