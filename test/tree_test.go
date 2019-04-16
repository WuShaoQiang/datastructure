package test

import (
	"fmt"
	"testing"

	"github.com/WuShaoQiang/data_structure_and_algorithm_analysis/datastructure"
)

func TestTree(t *testing.T) {

	fmt.Println("==================TestTree==================")


	tree := datastructure.NewTree(nil)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(1)

	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(7)
	tree.Insert(6)
	tree.Insert(9)
	tree.Insert(15)
	tree.Insert(13)
	tree.Insert(14)
	tree.Insert(17)
	tree.Insert(16)
	fmt.Println(tree.Head.Left.Value)
	fmt.Println(tree.Head.Left.Left.Value)

	fmt.Println(tree.Search(10).Value)

	fmt.Println(tree.Delete(1))
	fmt.Println(tree.Delete(2))
	fmt.Println(tree.Delete(10))
	fmt.Println(tree.Head.Right.Left.Value)

}
