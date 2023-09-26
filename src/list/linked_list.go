package list

import "errors"

type LinkedList struct {
}

func (list *LinkedList) Add(val int) {
}

func (list *LinkedList) AddOnIndex(value int, index int) error {
    return errors.New("error msg")
}

func (list *LinkedList) RemoveOnIndex(index int) error {
    return errors.New("error msg")
}

func (list *LinkedList) Get(index int) (int, error) {
    return -1, errors.New("error msg")
}

func (list *LinkedList) Set(value, index int) error {
    return errors.New("error msg")
}

func (list *LinkedList) Size() int {
    return 0
}