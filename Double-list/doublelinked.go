package doublelinked

import (
	"fmt"
	"errors"
)

type Node struct {
	data interface{}
	prev, next *Node
}

type DoubleLinkedList struct {
	head, tail *Node
	syze int
	sync.RWMutex
}

//Head returns first node of double linked list
//Time complexity O(1)
func (dList *DoubleLinkedList) Head() *Node {
	dList.RLock()
	defer dList.RUnlock()
	return dList.head
}

//Tail returns last node of double linked list
//Time complexity O(1)
func (dList *DoubleLinkedList) Tail() *Node {
	dList.RLock()
	defer dList.RUnlock()
	return dList.tail
}

//Append is a concorrency safe inserting data at the end of a double linked list
//Time complexity O(n)
func (dList *DoubleLinkedList) Append(data interface{}) {
	dList.Lock()
	defer dList.Unlock()
	
	appendNode := &Node{
		data: data,
		prev: nil,
		next: nil,
	}
	if dList.head == nil {
		dList.head = appendNode
		dList.tail = appendNode
	} else {
		lastNode := dList.head
		for dList.tail != nil {
			lastNode = lastNode.next
		}
		lastNode.next = appendNode
		appendNode.prev = lastNode
	}
	list.size++
}

//DeleteAt is a concorrency safe deleting node at the defined position
//Time complexity O(n)
func (dList *DoubleLinkedList) DeleteAt(indexAt int) error {
	dList.Lock()
	defer dList.Unlock()

	if indexAt < 0 {
		return errors.New("Index is out of limits")
	}
	
	currNode := dList.head
	for i := 1; currNode != nil && i < indexAt; i++ {
		currNode = currNode.next
	}
	currNode.prev.next = currNode.next
	currNode.next.prev = currNode.prev
	currNode.next = nil
	currNode.prev = nil

	list.size--
}