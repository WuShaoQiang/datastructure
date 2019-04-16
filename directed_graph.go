package datastructure

import (
	"fmt"
	"log"
)

type DirectedGraphNode struct {
	value int
}

type DirectedGraph struct {
	nodes  []*DirectedGraphNode                        // 节点集
	edges  map[*DirectedGraphNode][]*DirectedGraphNode // 邻接表表示的无向图
	weight map[*DirectedGraphNode]map[*DirectedGraphNode]int
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		nodes:  make([]*DirectedGraphNode, 0),
		edges:  make(map[*DirectedGraphNode][]*DirectedGraphNode),
		weight: make(map[*DirectedGraphNode]map[*DirectedGraphNode]int),
	}
}

func NewDirectedGraphNode(i int) *DirectedGraphNode {
	return &DirectedGraphNode{
		value: i,
	}
}

func (g *DirectedGraph) AddNode(n *DirectedGraphNode) {
	g.nodes = append(g.nodes, n)
}

func (g *DirectedGraph) AddEdge(u, v *DirectedGraphNode, weight int) {
	g.edges[u] = append(g.edges[u], v)
	if g.weight[u] == nil {
		g.weight[u] = make(map[*DirectedGraphNode]int)
	}
	g.weight[u][v] = weight
}

func (g *DirectedGraph) DeleteEdge(u, v *DirectedGraphNode) {
	pos := isDirectedGraphNodeExist(g.edges[u], v)
	if pos == -1 {
		log.Println("u didn't point to v, nothing to delete")
	} else {
		if pos == len(g.edges[u])-1 {
			g.edges[u] = append(g.edges[u][:pos])
		} else {
			g.edges[u] = append(g.edges[u][:pos], g.edges[u][pos+1:]...)
		}
	}

	g.weight[u][v] = -1
}

func isDirectedGraphNodeExist(a []*DirectedGraphNode, b *DirectedGraphNode) int {
	for index, single := range a {
		if single.value == b.value {
			return index
		}
	}
	return -1
}

func (g *DirectedGraph) String() {
	for _, node := range g.nodes {
		for _, point := range g.edges[node] {
			fmt.Printf("%v --%v--> %v\n", node.value, g.weight[node][point], point.value)
		}
	}
}
