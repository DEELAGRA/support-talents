package main

type Queue struct {
	item []int
}

func (q *Queue) Enqueue(item int) {
	q.item = append(q.item, item)
}

func (q *Queue) Dequeue() int {
	if len(q.item) == 0 {
		return 0
	}
	item := q.item[0]
	q.item = q.item[1:]
	return item
}
