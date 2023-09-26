package deque

import "errors"

type ArrayDeque struct {
}

func (deque *ArrayDeque) Init(size int) {
}

func (deque *ArrayDeque) EnqueueRear(val int) {
}

func (deque *ArrayDeque) EnqueueFront(val int) {
}

func (deque *ArrayDeque) DequeueFront() (int, error) {
    return -1, errors.New("error msg")
}

func (deque *ArrayDeque) DequeueRear() (int, error) {
    return -1, errors.New("error msg")
}

func (deque *ArrayDeque) Front() (int, error) {
    return -1, errors.New("error msg")
}

func (deque *ArrayDeque) Rear() (int, error) {
    return -1, errors.New("error msg")
}

func (deque *ArrayDeque) IsEmpty() bool {
    return false
}

func (deque *ArrayDeque) Size() int {
    return 0
}