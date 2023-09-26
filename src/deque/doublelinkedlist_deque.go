package deque

import "errors"

type DoublyLinkedListDeque struct {
    front *Node
    rear  *Node
    size  int
}

type Node struct {
    prev *Node
    val  int
    next *Node
}

func (deque *DoublyLinkedListDeque) EnqueueRear(val int) {
    newNode := Node{nil, val, nil}
    if deque.size == 0 {
        deque.front = &newNode
    } else {
        newNode.prev = deque.rear
        deque.rear.next = &newNode
    }
    deque.rear = &newNode
    deque.size++
}

func (deque *DoublyLinkedListDeque) EnqueueFront(val int) {
    newNode := Node{nil, val, nil}
    if deque.size == 0 {
        deque.rear = &newNode
    } else {
        newNode.next = deque.front
        deque.front.prev = &newNode
    }
    deque.front = &newNode
    deque.size++
}

func (deque *DoublyLinkedListDeque) DequeueFront() (int, error) {
    if deque.size == 0 {
        return -1, errors.New("Can't dequeue from empty deque")
    } else {
        lastFront := deque.front.val
        if deque.size == 1 {
            deque.front = nil
            deque.rear = nil
        } else {
            deque.front = deque.front.next
            deque.front.prev = nil
        }
        deque.size--
        return lastFront, nil
    }
}

func (deque *DoublyLinkedListDeque) DequeueRear() (int, error) {
    if deque.size == 0 {
        return -1, errors.New("Can't dequeue from empty deque")
    } else {
        lastRear := deque.rear.val
        if deque.size == 1 {
            deque.front = nil
            deque.rear = nil
        } else {
            deque.rear = deque.rear.prev
            deque.rear.next = nil
        }
        deque.size--
        return lastRear, nil
    }
}

func (deque *DoublyLinkedListDeque) Front() (int, error) {
    if deque.size == 0 {
        return -1, errors.New("Empty deque has no front")
    } else {
        return deque.front.val, nil
    }
}

func (deque *DoublyLinkedListDeque) Rear() (int, error) {
    if deque.size == 0 {
        return -1, errors.New("Empty deque has no rear")
    } else {
        return deque.rear.val, nil
    }
}

func (deque *DoublyLinkedListDeque) IsEmpty() bool {
    return deque.size == 0
}

func (deque *DoublyLinkedListDeque) Size() int {
    return deque.size
}