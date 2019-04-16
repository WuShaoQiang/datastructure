package datastructure

import (
	"log"
	"sync"
)

type Stack struct {
	stack []interface{}
	len   int
	lock  sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]interface{}, 0),
		len:   0,
	}
}

func (s *Stack) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.len
}

func (s *Stack) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return (s.len == 0)
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	var el interface{}
	el, s.stack = s.stack[s.len-1], s.stack[:s.len-1]
	s.len--
	return el
}

func (s *Stack) Push(el interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stack = append(s.stack, el)
	s.len++
}

func (s *Stack) GetTop() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	if s.stack == nil || s.len == 0 {
		log.Println("Nothing in stack")
		return nil
	}

	return s.stack[s.len-1]

}

func (s *Stack) ClearStack() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.len = 0
	s.stack = make([]interface{}, 0)
}

func (s *Stack) DestoryStack() {
	s = nil
}
