package test

import (
	"fmt"
	"testing"

	"github.com/WuShaoQiang/data_structure_and_algorithm_analysis/datastructure"
)

func TestUndirectedGraph(t *testing.T) {

	fmt.Println("==================TestUndirectedGraph==================")

	g := datastructure.NewUndirectedGraph()
	n1 := datastructure.NewUndirectedGraphNode(1)
	n2 := datastructure.NewUndirectedGraphNode(2)
	n3 := datastructure.NewUndirectedGraphNode(3)
	n4 := datastructure.NewUndirectedGraphNode(4)
	n5 := datastructure.NewUndirectedGraphNode(5)

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)

	g.AddEdge(n3, n4)
	g.AddEdge(n4, n5)
	g.AddEdge(n1, n2)
	g.AddEdge(n1, n5)
	g.AddEdge(n2, n3)
	g.AddEdge(n2, n4)
	g.AddEdge(n2, n5)

	g.DeleteEdge(n1, n5)
	g.DeleteEdge(n1, n2)
	g.DeleteEdge(n3, n4)

	g.String()

	fmt.Println("")
}

func TestDirectedGraph(t *testing.T) {
	g := datastructure.NewDirectedGraph()

	n1 := datastructure.NewDirectedGraphNode(1)
	n2 := datastructure.NewDirectedGraphNode(2)
	n3 := datastructure.NewDirectedGraphNode(3)
	n4 := datastructure.NewDirectedGraphNode(4)
	n5 := datastructure.NewDirectedGraphNode(5)

	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)

	g.AddEdge(n3, n4, 10)
	g.AddEdge(n4, n5, 20)
	g.AddEdge(n1, n2, 15)
	g.AddEdge(n1, n5, 10)
	g.AddEdge(n2, n3, 5)
	g.AddEdge(n2, n4, 30)
	g.AddEdge(n2, n5, 20)

	// g.DeleteEdge(n1, n5)
	// g.DeleteEdge(n1, n2)
	// g.DeleteEdge(n3, n4)

	g.String()
}
