package stack

import "errors"

type LinkedListStack struct {
}

func (stack *LinkedListStack) Push(val int) {
}

func (stack *LinkedListStack) Pop() (int, error) {
    return -1, errors.New("error msg")
}

func (stack *LinkedListStack) Peek() (int, error) {
    return -1, errors.New("error msg")
}

func (stack *LinkedListStack) IsEmpty() bool {
    return false
}

func (stack *LinkedListStack) Size() int {
    return 0
}