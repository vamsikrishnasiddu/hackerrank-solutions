package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

/*
 * Complete the 'insertNodeAtPosition' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_SINGLY_LINKED_LIST llist
 *  2. INTEGER data
 *  3. INTEGER position
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	// Write your code here
	head := llist
	node := &SinglyLinkedListNode{data: data}

	if position == 0 {
		node.next = head
		head = node
	} else {
		current := head
		for i := 1; i < int(position); i++ {
			current = current.next
		}
		temp := current.next
		current.next = node
		node.next = temp
	}

	return head

}

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	// Write your code here
	head := llist

	if position == 0 {
		n := head.next
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

	return head

}

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here
	head := llist

	if head == nil {
		return nil
	}

	var dummy *SinglyLinkedListNode
	var next *SinglyLinkedListNode

	for head != nil {
		next = head.next
		head.next = dummy
		dummy = head
		head = next
	}

	return dummy

}

func reversePrint(llist *SinglyLinkedListNode) {
	// Write your code here
	if llist != nil {
		reversePrint(llist.next)
		fmt.Println(llist.data)
	}
}

func calcSize(head *SinglyLinkedListNode) int {
	current := head
	length := 0
	for current != nil {
		length++
		current = current.next
	}
	return length
}

func compare_lists(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) bool {
	size1 := calcSize(head1)
	size2 := calcSize(head2)

	if size1 != size2 {
		return false
	}

	current1 := head1
	current2 := head2
	count := 0
	for current1 != nil && current2 != nil {
		if current1.data == current2.data {
			count++
		}
		current1 = current1.next
		current2 = current2.next
	}

	if count == size1 {
		return true
	}

	return false

}

func mergeTwoSortedLists(l1 *SinglyLinkedListNode, l2 *SinglyLinkedListNode) *SinglyLinkedListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.data > l2.data {
		l1, l2 = l2, l1
	}

	result := l1
	var temp *SinglyLinkedListNode

	for l1 != nil && l2 != nil {
		temp = nil
		for (l1 != nil) && (l1.data <= l2.data) {
			temp = l1
			l1 = l1.next
		}

		temp.next = l2
		l1, l2 = l2, l1
	}

	return result

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	llistCount1, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	llist1 := SinglyLinkedList{}
	for i := 0; i < int(llistCount1); i++ {
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		llistItem := int32(llistItemTemp)
		llist1.insertNodeIntoSinglyLinkedList(llistItem)
	}

	llistCount2, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	llist2 := SinglyLinkedList{}
	for i := 0; i < int(llistCount2); i++ {
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		llistItem := int32(llistItemTemp)
		llist2.insertNodeIntoSinglyLinkedList(llistItem)
	}
	llist3 := mergeTwoSortedLists(llist1.head, llist2.head)

	// dataTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)
	// data := int32(dataTemp)

	// positionTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)
	// position := int32(positionTemp)

	// llist_head := insertNodeAtPosition(llist.head, data, position)

	printSinglyLinkedList(llist3, " ", writer)
	// fmt.Fprintf(writer, "\n")
	// printSinglyLinkedList(llist2.head, " ", writer)
	// fmt.Fprintf(writer, "\n")

	// result := compare_lists(llist1.head, llist2.head)
	// fmt.Println("result", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
