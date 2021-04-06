package main

import (
  "fmt"
  "sync"
  "strings"
  "errors"
)

type Node struct {
  data interface{}
  next *Node
}

type List struct {
  head *Node
  size int
  sync.RWMutex
}

//New is a Linked List constructor
func New() *List {
  return new(List)
}

//Append adds *Node into the end of a Linked List
//Concurrency safe
//Time complexity is O(n)
func (l *List) Append(n interface{}) {
  l.Lock()
  defer l.Unlock()

  if l.head == nil {
    l.head = &Node{n, nil}
  } else {
    curnode := l.head
    for curnode.next != nil {
      curnode = curnode.next
    }
    curnode.next = getNode(n)
  }
  l.size++
}

//InsertAt is a concorrency safe inserting data by a given index into a linked list
//Time complexity O(n)
func (l *List) InsertAt(n interface{}, idx int) error {
  l.Lock()
  defer l.Unlock()
  
  if idx > l.size || idx < 0 {
    return errors.New("Index is out of range")
  }
  insnode := getNode(n)
  if idx == 0 {
    insnode.next = l.head
    l.head = insnode
  } else {
    node := l.head
  	for i := 0; i < idx-2; i++ {
    	node = node.next
  	}
    insnode.next = node.next
    node.next = insnode
  }
  l.size++
  return nil
}

//Reverse function
//Concurrency safe
//Time complexity is O(n)
func (l *List) Reverse() {
  l.Lock()
  defer l.Unlock()
  
  curnode := l.head
  var prev *Node
  var next *Node
  
  if curnode == nil {
    return
  }
  for curnode != nil {
    next, curnode.next = curnode.next, prev
    prev, curnode = curnode, next
  }
  l.head = prev
}

//IsEmpty checks for linked list
//Concurrency safe
//Time complexity O(1)
func (l *List) IsEmpty() bool {
	l.RLock()
	defer l.RUnlock()

	if l.head == nil { 
		return true
	}
	return false
}

//ToString uses to add *Node elements into a string
//Concurrency safe
//Time Complexity is O(n)
func (l *List) ToString() string {
  l.Lock()
  defer l.Unlock()
  
  result := ""
  values := []string{}
  for el := l.head; el != nil; el = el.next {
    values = append(values, fmt.Sprintf("%v", el.data))
  }
  result += strings.Join(values, ", ")
  return result
}

//getNode returns a *Node element
func getNode(n interface{}) *Node {
	switch n.(type) {
    case *Node: return n.(*Node)
    default: {
      node := &Node{n, nil}
      return node
    }
  }  
}

func main() {
  list := New()
  for i:=0; i<11; i++ {
    list.Append(i)
  }
  fmt.Printf("Original Linked List is: %s\n", list.ToString())
  list.Reverse()
  fmt.Printf("Reversed Linked List is: %s\n", list.ToString())
  list.InsertAt(100, 8)
	fmt.Printf("InsertAt works!: %s\n", list.ToString())
}