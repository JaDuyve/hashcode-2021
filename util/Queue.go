package util

import (
	"container/list"
	"log"
)

type Queue struct {
	queue *list.List
}

func NewQueue() Queue {
	return Queue{
		queue: list.New(),
	}
}

func (q *Queue) Enqueue(value interface{}) {
	q.queue.PushBack(value)
}

func (q *Queue) Dequeue() {
	if q.Empty() {
		log.Fatal("Dequeue error: Queue was empty")
	}

	ele := q.queue.Front()
	q.queue.Remove(ele)
}

func (q Queue) Front() interface{} {
	if q.Empty() {
		log.Fatal("Front error: Queue was empty")
	}

	return q.queue.Front().Value
}

func (q Queue) size() int {
	return q.queue.Len()
}

func (q Queue) Empty() bool {
	return q.queue.Len() == 0
}
