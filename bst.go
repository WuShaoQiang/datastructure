package datastructure

import (
	"fmt"
)

//左小右大
type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type Tree struct {
	Head *Node
	Size int
}

func NewNode(i int) *Node {
	return &Node{Value: i}
}
func (n *Node) Compare(m *Node) int {
	if n.Value > m.Value {
		return 1
	}
	if n.Value < m.Value {
		return -1
	}
	return 0
}
func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	}
	return &Tree{Head: n, Size: 1}
}
func (t *Tree) Insert(i int) {
	n := NewNode(i)
	h := t.Head

	if h == nil {
		t.Head = n
		t.Size++
		// t.Head.Parent = nil
		return
	}

	for {
		//n>h right
		if n.Compare(h) == 1 {
			if h.Right == nil {
				h.Right = n
				// n.Parent = h
				t.Size++
				break
			} else {
				h = h.Right
			}
		} else { //n<=h left
			if h.Left == nil {
				h.Left = n
				// n.Parent = h
				t.Size++
				break
			} else {
				h = h.Left
			}
		}
	}
}

func (t *Tree) Search(i int) *Node {
	h := t.Head
	n := NewNode(i)

	for h != nil {
		switch n.Compare(h) {
		case 1:
			h = h.Right
		case -1:
			h = h.Left
		case 0:
			return h
		}
	}
	return nil
}

func (t *Tree) Delete(i int) bool {
	var parent *Node
	h := t.Head
	n := &Node{Value: i}

	for h != nil {
		switch n.Compare(h) {
		case 1:
			parent = h
			h = h.Right
		case -1:
			parent = h
			h = h.Left
		case 0:
			fmt.Println(h)
			fmt.Println(parent)
			fmt.Println(parent.Left)
			//双孩子
			if h.Left != nil && h.Right != nil {
				n := findMin(h.Right)
				if n.Parent != nil {
					n.Parent.Left = n.Right
				}

				h.Value = n.Value
				n.Right = nil
				n.Parent = nil

				return true
			} else if h.Left != nil && h.Right == nil { //左子树
				h.Value = h.Left.Value
				h.Left = h.Left.Left
				h.Right = h.Left.Right

				t.Size--
				return true
			} else if h.Right != nil && h.Left == nil { //右子树
				h.Value = h.Right.Value
				h.Right = h.Right.Right
				h.Left = h.Right.Left

				t.Size--
				return true
			} else if parent == nil { //根节点
				t.Head = nil
				t.Size--
				return true
			} else if parent.Left == h { //叶子节点
				parent.Left = nil
				t.Size--
				return true
			} else if parent.Right == h {
				parent.Right = nil
				t.Size--
				return true
			}
			return false
		}
	}
	return false
}

func findMin(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}

	return n
}
