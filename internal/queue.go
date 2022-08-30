package internal

type Queue struct {
	arr []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(i int) {
	q.arr = append(q.arr, i)
}

func (q *Queue) Pop() int {
	if len(q.arr) == 0 {
		return -1
	}

	head := q.arr[0]
	q.arr = q.arr[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}
