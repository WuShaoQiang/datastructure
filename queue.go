package datastructure

import (
	"log"
	"sync"
)

type Queue struct {
	queue []interface{}
	len   int
	lock  sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]interface{}, 0),
		len:   0,
	}
}

func (q *Queue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.len
}

func (q *Queue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return (q.len == 0)
}

func (q *Queue) DeQueue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	var el interface{}
	if q.len >= 2 {
		el, q.queue = q.queue[0], q.queue[1:q.len]
	} else if q.len == 1 {
		el, q.queue = q.queue[0], q.queue[0:0]
	} else {
		log.Println("Nothing in queue")
		return nil
	}
	q.len--
	return el
}

func (q *Queue) EnQueue(el interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.queue = append(q.queue, el)
	q.len++
}

func (q *Queue) GetHead() interface{} {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if q.queue == nil || q.len == 0 {
		log.Println("Nothing in queue")
		return nil
	}

	return q.queue[0]

}
