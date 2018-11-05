package queue

type Any interface{}

// Queue is makeshift queue imitation using array.
// don't use it in production
type Queue struct {
	items []Any
}

// Empty shows if queue has items in it
func (q *Queue) Empty() bool {
	return len(q.items) == 0
}

// Push adds new item in queue
func (q *Queue) Push(item Any) {
	q.items = append(q.items, item)
}

// Pop removes last in item in queue
func (q *Queue) Pop() Any {
	first := q.items[0]
	rest := q.items[1:]
	q.items = rest
	return first
}
