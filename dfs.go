package datastructure

import (
	"fmt"
)

var (
	count = 0
)

func DFSUndirectedGraph(g *UndirectedGraph) []int {
	visited := make(map[*UndirectedGraphNode]bool)
	done := make(map[*UndirectedGraphNode]bool)
	tmp := make([]int, len(g.nodes))

	for _, node := range g.nodes {
		if _, exist := done[node]; !exist {
			if _, exist := visited[node]; !exist {
				tmp[count] = node.value
				count++
				visited[node] = true
			}
			dfsUndirected(node, g, tmp, visited, done)
		}
	}

	count = 0
	return tmp
}

func dfsUndirected(n *UndirectedGraphNode, g *UndirectedGraph, tmp []int, visited map[*UndirectedGraphNode]bool, done map[*UndirectedGraphNode]bool) {
	fmt.Printf("running on %v\n", n.value)
	for _, node := range g.edges[n] {
		if _, isVisited := visited[node]; !isVisited {
			fmt.Printf("node %v is not visited\n", node.value)
			visited[node] = true
			tmp[count] = node.value
			count++
			fmt.Println(tmp)
			dfsUndirected(node, g, tmp, visited, done)
		} else {
			fmt.Printf("node %v is visited\n", node.value)
			continue
		}
	}
	fmt.Printf("node %v is done\n", n.value)
	done[n] = true
}

func DFSDirectedGraph(g *DirectedGraph) []int {
	visited := make(map[*DirectedGraphNode]bool)
	done := make(map[*DirectedGraphNode]bool)
	tmp := make([]int, len(g.nodes))

	for _, node := range g.nodes {
		if _, exist := done[node]; !exist {
			if _, exist := visited[node]; !exist {
				tmp[count] = node.value
				count++
				visited[node] = true
			}
			dfsDirected(node, g, tmp, visited, done)
		}
	}

	count = 0
	return tmp
}

func dfsDirected(n *DirectedGraphNode, g *DirectedGraph, tmp []int, visited map[*DirectedGraphNode]bool, done map[*DirectedGraphNode]bool) {
	fmt.Printf("running on %v\n", n.value)
	for _, node := range g.edges[n] {
		if _, isVisited := visited[node]; !isVisited {
			fmt.Printf("node %v is not visited\n", node.value)
			visited[node] = true
			tmp[count] = node.value
			count++
			fmt.Println(tmp)
			dfsDirected(node, g, tmp, visited, done)
		} else {
			fmt.Printf("node %v is visited\n", node.value)
			continue
		}
	}
	fmt.Printf("node %v is done\n", n.value)
	done[n] = true
}

func DFSDirectedGraphByStack(g *DirectedGraph) []int {
	visited := make(map[*DirectedGraphNode]bool)
	// done := make(map[*DirectedGraphNode]bool)
	tmp := make([]int, len(g.nodes))
	s := NewStack()

	for _, node := range g.nodes {
		// if _, exist := done[node]; !exist {
		if _, isVisited := visited[node]; !isVisited {
			tmp[count] = node.value
			count++
			visited[node] = true
			s.Push(node)
			// }
			dfsDirectedByStack(s, g, tmp, visited)
		}
	}

	count = 0
	return tmp
}

func dfsDirectedByStack(s *Stack, g *DirectedGraph, tmp []int, visited map[*DirectedGraphNode]bool) {

	fmt.Printf("running on %v\n", (s.GetTop().(*DirectedGraphNode)).value)
	for _, node := range g.edges[s.GetTop().(*DirectedGraphNode)] {
		if _, isVisited := visited[node]; !isVisited {
			fmt.Printf("node %v is not visited\n", node.value)
			visited[node] = true
			tmp[count] = node.value
			count++
			s.Push(node)
			fmt.Println(tmp)
			dfsDirectedByStack(s, g, tmp, visited)
		} else {
			fmt.Printf("node %v is visited\n", node.value)
			continue
		}
	}

	fmt.Printf("node %v is done\n", (s.Pop().(*DirectedGraphNode)).value)
}
