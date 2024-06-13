package main

type Queue struct {
	values []int8
}

func newQueue(size int8) *Queue {
	return &Queue{make([]int8, 0, size)}
}

func (queue *Queue) enqueue(data int8) {
	queue.values = append(queue.values, data)
}

func (queue *Queue) dequeue() {
	if len(queue.values) == 0 {
		return
	}
	queue.values = queue.values[1:]
}

func (queue *Queue) front() (int8, bool) {
	if len(queue.values) == 0 {
		return 0, false
	}
	return queue.values[0], true

}

func (queue *Queue) back() (int8, bool) {
	if len(queue.values) == 0 {
		return 0, false
	}
	return queue.values[len(queue.values)-1], true
}

func main() {}
