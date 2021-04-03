package linkedlist

import (
	"fmt"
	"sync"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

//InsertAt is a concorrency safe inserting data by a given index into a linked list
//Time complexity O(n)
func (list *LinkedList) InsertAt(data, indexAt int) error {
	list.lock.Lock()
	defer list.lock.Unlock()
	
	if indexAt > list.size || indexAt < 0 {
		return errors.New("Index is out of limits")
	}
	insertNode := Node{data, nil}
	if indexAt == 0 {
		insertNode.next = list.head
		list.head = &insertNode
		return nil
	}
	//getting a previous node for insertNode
	node := list.head
	for i := 0; i < indexAt-2; i++ {
		node = node.next
	}
	insertNode.next = node.next
	node.next = &insertNode
	list.size++
	return nil
}

//Append is a concorrency safe inserting data at the end of a linked list
//Time complexity O(n)
func (list *LinkedList) Append(data int) {
	list.lock.Lock()
	defer list.lock.Unlock()

	appendNode := Node{data,nil}
	if list.head == nil {
		list.head = &node
	} else {
		currList := list.head
		for currList.head != nil {
			currList = currList.next 
		}
		currList.next = &appendNode
	}
	list.size++
}

//DeleteAt is a concorrency safe removing node by a given index
//Time complexity O(n)
func (list *LinkedList) DeleteAt(indexAt int) error {
	list.lock.Lock()
	defer list.lock.Lock()

	if indexAt > list.size || indexAt < 0 {
		return errors.New("Index is out of limits")
	}

	currNode := list.head
	for i := 0; i < indexAt - 1; i++ {
		currNode = currNode.next
	}
	deleteNode = currNode.next
	currNode.next = deleteNode.next
	list.size--
}

//IsEmpty is a concorrency safe check for linked list
//Time complexity O(1)
func (list *LinkedList) IsEmpty() bool {
	list.lock.RLock()
	defer list.lock.RUnlock()

	if list.head == nil { 
		return true
	}
	return false
}
