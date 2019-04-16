package test

import (
	"fmt"
	"testing"

	"github.com/WuShaoQiang/data_structure_and_algorithm_analysis/datastructure"
)

func TestQueue(t *testing.T) {

	fmt.Println("==================TestQueue==================")

	q := datastructure.NewQueue()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.DeQueue())

	q.EnQueue(1)
	q.EnQueue(2)

	fmt.Println(q.GetHead())
	fmt.Println(q.DeQueue())

	fmt.Println(q)

	fmt.Println(q.Len())
}
