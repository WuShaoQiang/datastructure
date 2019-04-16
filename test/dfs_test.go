package test

import (
	"fmt"
	"testing"

	"github.com/WuShaoQiang/data_structure_and_algorithm_analysis/datastructure"
)

func getUndirectedGraph() *datastructure.UndirectedGraph {

	g := datastructure.NewUndirectedGraph()
	n1 := datastructure.NewUndirectedGraphNode(1)
	n2 := datastructure.NewUndirectedGraphNode(2)
	n3 := datastructure.NewUndirectedGraphNode(3)
	n4 := datastructure.NewUndirectedGraphNode(4)
	n5 := datastructure.NewUndirectedGraphNode(5)

	n6 := datastructure.NewUndirectedGraphNode(6)
	n7 := datastructure.NewUndirectedGraphNode(7)
	n8 := datastructure.NewUndirectedGraphNode(8)
	n9 := datastructure.NewUndirectedGraphNode(9)
	n10 := datastructure.NewUndirectedGraphNode(10)

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	g.AddNode(n7)
	g.AddNode(n8)
	g.AddNode(n9)
	g.AddNode(n10)

	g.AddEdge(n1, n2)
	g.AddEdge(n1, n5)
	g.AddEdge(n2, n3)
	g.AddEdge(n2, n4)
	g.AddEdge(n2, n5)
	g.AddEdge(n3, n4)
	g.AddEdge(n4, n5)
	g.AddEdge(n10, n9)
	g.AddEdge(n8, n9)

	return g
}

func getDirectedGraph() *datastructure.DirectedGraph {

	g := datastructure.NewDirectedGraph()
	n1 := datastructure.NewDirectedGraphNode(1)
	n2 := datastructure.NewDirectedGraphNode(2)
	n3 := datastructure.NewDirectedGraphNode(3)
	n4 := datastructure.NewDirectedGraphNode(4)
	n5 := datastructure.NewDirectedGraphNode(5)

	n6 := datastructure.NewDirectedGraphNode(6)
	n7 := datastructure.NewDirectedGraphNode(7)
	n8 := datastructure.NewDirectedGraphNode(8)
	n9 := datastructure.NewDirectedGraphNode(9)
	n10 := datastructure.NewDirectedGraphNode(10)

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	g.AddNode(n7)
	g.AddNode(n8)
	g.AddNode(n9)
	g.AddNode(n10)

	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n5, 5)
	g.AddEdge(n1, n10, 7)
	g.AddEdge(n2, n3, 8)
	g.AddEdge(n2, n4, 4)
	g.AddEdge(n2, n5, 3)
	g.AddEdge(n3, n4, 6)
	g.AddEdge(n4, n5, 7)
	g.AddEdge(n6, n8, 1)
	g.AddEdge(n7, n6, 13)
	g.AddEdge(n8, n9, 2)
	g.AddEdge(n10, n7, 10)
	g.AddEdge(n10, n9, 10)

	return g
}

func TestDFSUndirectedGraph(t *testing.T) {

	fmt.Println("==================TestDFSUndirectedGraph==================")

	g := getUndirectedGraph()
	result := datastructure.DFSUndirectedGraph(g)
	fmt.Println("result : ", result)
}

func TestDFSDirectedGraph(t *testing.T) {
	fmt.Println("==================TestDFSDirectedGraph==================")

	g := getDirectedGraph()
	result := datastructure.DFSDirectedGraph(g)
	fmt.Println("result : ", result)

}

func TestDFSDirectedGraphByStack(t *testing.T) {
	fmt.Println("==================TestDFSDirectedGraphByStack==================")

	g := getDirectedGraph()
	result := datastructure.DFSDirectedGraphByStack(g)
	fmt.Println("result : ", result)

}
