package queue

// La Cola Queue es FIFO First In First Out
// entrada 1 2 3 .. salida 1 2 3

import "errors"

type Queue struct {
	values []int
}

func NewQueue() *Queue {
	return &Queue{
		values: []int{},
	}
}

func (q *Queue) Queue(newValue int) {
	q.values = append(q.values, newValue)
}

func (q *Queue) Deque() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty!! ")
	}

	valueToReturn := q.values[0]
	q.values = q.values[1:]
	return valueToReturn, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}
