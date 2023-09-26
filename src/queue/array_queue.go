package queue

import "errors"

type ArrayQueue struct {
  v []int;
  front int;
  rear int;
}

func (queue *ArrayQueue) Init(size int) {
}

func (queue *ArrayQueue) Enqueue(val int) {
}

func (queue *ArrayQueue) Dequeue() (int, error) {
    return -1, errors.New("error msg")
}

func (queue *ArrayQueue) Front() (int, error) {
    return -1, errors.New("error msg")
}

func (queue *ArrayQueue) IsEmpty() bool {
    return false
}

func (queue *ArrayQueue) Size() int {
    return 0
}