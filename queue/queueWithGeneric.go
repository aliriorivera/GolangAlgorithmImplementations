package queue

type QueueWithGeneric[T any] struct {
	values []*T
}

func NewQueueWithGeneric[T any]() *QueueWithGeneric[T] {
	return &QueueWithGeneric[T]{
		values: []*T{},
	}
}

func (q *QueueWithGeneric[T]) Queue(val *T) {
	q.values = append(q.values, val)
}

func (q *QueueWithGeneric[T]) DeQueue() *T {
	if q.IsEmpty() {
		return nil
	}
	first := q.values[0]
	q.values = q.values[1:]
	return first
}

func (q *QueueWithGeneric[T]) IsEmpty() bool {
	return len(q.values) == 0
}
