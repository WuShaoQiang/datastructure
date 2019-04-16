package test

import (
	"fmt"
	"testing"

	"github.com/WuShaoQiang/data_structure_and_algorithm_analysis/datastructure"
)

func TestBFSUndirectedGraphByQueue(t *testing.T) {
	fmt.Println("==================TestBFSUndirectedGraphByQueue==================")

	g := getUndirectedGraph()
	result := datastructure.BFSUndirectedGraphByQueue(g)
	fmt.Println("result : ", result)
}

func TestBFSDirectedGraphByQueue(t *testing.T) {
	fmt.Println("==================TestBFSUndirectedGraphByQueue==================")

	g := getDirectedGraph()
	result := datastructure.BFSDirectedGraphByQueue(g)
	fmt.Println("result : ", result)
}
