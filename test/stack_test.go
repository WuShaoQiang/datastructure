package test

import (
	"fmt"
	"testing"

	"github.com/WuShaoQiang/data_structure_and_algorithm_analysis/datastructure"
)

func TestStack(t *testing.T) {

	fmt.Println("==================TestStack==================")


	s := datastructure.NewStack()
	fmt.Println(s.Len())
	fmt.Println(s.IsEmpty())

	s.Push(1)
	s.Push(2)
	fmt.Println(s.GetTop())
	fmt.Println(s.Pop())

	s.ClearStack()
	fmt.Println(s.GetTop())
	fmt.Println(s)
	s.DestoryStack()
	fmt.Println(s.GetTop())
}
