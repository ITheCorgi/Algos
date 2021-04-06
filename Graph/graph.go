package graph

import (
	"fmt"
	"sync"
)

type Node struct {
	data interface{}
}

type GraphItem struct {
	nodes []*Node
	edges map[Node][]*Node
	sync.RWMutex
}

func (graph *GraphItem) AddNode(n *Node) {
	graph.Lock()
	defer graph.Unlock()
	graph.nodes = append(graph.nodes, n)
}

