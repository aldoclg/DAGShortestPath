package queue

import (
	"reflect"
	"sort"
)

type PriorityQueue[T IComparable[T]] struct {
	queue []T
}

type IComparable[T any] interface {
	Compare(e T) bool
}

func NewPriorityQueue[T IComparable[T]]() PriorityQueue[T] {
	return PriorityQueue[T]{queue: make([]T, 0)}
}

func (q *PriorityQueue[T]) Poll() T {
	if len(q.queue) > 0 {
		item := q.queue[0]
		q.queue = q.queue[1:]
		return item
	}
	return q.queue[0]
}

func (q *PriorityQueue[T]) Peek() T {
	return q.queue[0]
}

func (q *PriorityQueue[T]) Offer(e T) {
	q.queue = append(q.queue, e)
	q.Update()
}

func (q *PriorityQueue[T]) Update() {
	sort.Slice(q.queue, func(i, j int) bool {
		return q.queue[i].Compare(q.queue[j])
	})
}

func (q *PriorityQueue[T]) IsNotEmpty() bool {
	return len(q.queue) != 0
}

func (q *PriorityQueue[T]) Remove(e T) {
	for i, v := range q.queue {
		if reflect.DeepEqual(v, e) {
			q.queue = append(q.queue[:i], q.queue[i+1:]...)
			return
		}

	}
}
