package list

import "errors"

type DoublyLinkedList struct {
}

func (list *DoublyLinkedList) Add(val int) {
}

func (list *DoublyLinkedList) AddOnIndex(value int, index int) error {
    return errors.New("error msg")
}

func (list *DoublyLinkedList) RemoveOnIndex(index int) error {
    return errors.New("error msg")
}

func (list *DoublyLinkedList) Get(index int) (int, error) {
    return -1, errors.New("error msg")
}

func (list *DoublyLinkedList) Set(value, index int) error {
    return errors.New("error msg")
}

func (list *DoublyLinkedList) Size() int {
    return 0
}