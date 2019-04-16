package datastructure

import (
	"fmt"
	"log"
	"sync"
)

type UndirectedGraphNode struct {
	value int
}

type UndirectedGraph struct {
	nodes []*UndirectedGraphNode                          // 节点集
	edges map[*UndirectedGraphNode][]*UndirectedGraphNode // 邻接表表示的无向图
	lock  sync.RWMutex                                    // 保证线程安全
}

func NewUndirectedGraph() *UndirectedGraph {
	return &UndirectedGraph{
		nodes: make([]*UndirectedGraphNode, 0),
		edges: make(map[*UndirectedGraphNode][]*UndirectedGraphNode),
	}
}

func NewUndirectedGraphNode(i int) *UndirectedGraphNode {
	return &UndirectedGraphNode{
		value: i,
	}
}

func (g *UndirectedGraph) AddNode(n *UndirectedGraphNode) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

func (g *UndirectedGraph) AddEdge(u, v *UndirectedGraphNode) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if g.edges == nil {
		g.edges = make(map[*UndirectedGraphNode][]*UndirectedGraphNode)
	}
	g.edges[u] = append(g.edges[u], v)
	g.edges[v] = append(g.edges[v], u)
}

func (g *UndirectedGraph) DeleteEdge(u, v *UndirectedGraphNode) {
	g.lock.Lock()
	defer g.lock.Unlock()

	pos := isUndirectedGraphExist(g.edges[u], v)

	if pos == -1 {
		log.Println("no connection between u and v")
	} else {
		if pos == len(g.edges[u])-1 {
			g.edges[u] = append(g.edges[u][:pos])
		} else {

			g.edges[u] = append(g.edges[u][:pos], g.edges[u][pos+1:]...)
		}
		pos = isUndirectedGraphExist(g.edges[v], u)
		if pos == len(g.edges[v])-1 {
			g.edges[v] = append(g.edges[v][:pos])
		} else {

			g.edges[v] = append(g.edges[v][:pos], g.edges[v][pos+1:]...)
		}
	}
}

func isUndirectedGraphExist(a []*UndirectedGraphNode, b *UndirectedGraphNode) int {
	for index, single := range a {
		if single.value == b.value {
			return index
		}
	}
	return -1
}

func (g *UndirectedGraph) String() {
	g.lock.RLock()
	defer g.lock.RUnlock()

	for _, node := range g.nodes {
		fmt.Print(node.value)
		for _, other := range g.edges[node] {
			fmt.Print("-->", other.value)
		}
		fmt.Println("")
	}
}
