package datastructure

var (
	bfscount = 0
)

func BFSUndirectedGraphByQueue(g *UndirectedGraph) []int {
	visited := make(map[*UndirectedGraphNode]bool)
	tmp := make([]int, len(g.nodes))

	q := NewQueue()

	for _, node := range g.nodes {
		if _, isVisited := visited[node]; !isVisited {
			tmp[bfscount] = node.value
			bfscount++
			visited[node] = true
			q.EnQueue(node)
			bfsUndirectedGraphByQueue(q, g, tmp, visited)
		}
	}
	bfscount = 0
	return tmp
}

func bfsUndirectedGraphByQueue(q *Queue, g *UndirectedGraph, tmp []int, visited map[*UndirectedGraphNode]bool) {
	var next = false
	for i := 0; i < q.Len(); i++ {
		for _, node := range g.edges[q.DeQueue().(*UndirectedGraphNode)] {
			next = true
			if _, isVisited := visited[node]; !isVisited {
				tmp[bfscount] = node.value
				bfscount++
				visited[node] = true
				q.EnQueue(node)
			}
		}
	}

	if next {
		bfsUndirectedGraphByQueue(q, g, tmp, visited)
	}
}

func BFSDirectedGraphByQueue(g *DirectedGraph) []int {
	visited := make(map[*DirectedGraphNode]bool)
	tmp := make([]int, len(g.nodes))

	q := NewQueue()

	for _, node := range g.nodes {
		if _, isVisited := visited[node]; !isVisited {
			tmp[bfscount] = node.value
			bfscount++
			visited[node] = true
			q.EnQueue(node)
			bfsDirectedGraphByQueue(q, g, tmp, visited)
		}
	}
	bfscount = 0

	return tmp
}

func bfsDirectedGraphByQueue(q *Queue, g *DirectedGraph, tmp []int, visited map[*DirectedGraphNode]bool) {
	var next = false
	for i := 0; i < q.Len(); i++ {
		for _, node := range g.edges[q.DeQueue().(*DirectedGraphNode)] {
			next = true
			if _, isVisited := visited[node]; !isVisited {
				tmp[bfscount] = node.value
				bfscount++
				visited[node] = true
				q.EnQueue(node)
			}
		}
	}

	if next {
		bfsDirectedGraphByQueue(q, g, tmp, visited)
	}
}
