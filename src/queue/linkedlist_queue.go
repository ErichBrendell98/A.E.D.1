package queue

import "errors"

type LinkedListQueue struct {
    front *Node
    rear  *Node
    size  int
}

type Node struct {
    val  int
    next *Node
}

func (queue *LinkedListQueue) Enqueue(val int) {
    newNode := Node{val, nil}
    if queue.size == 0 {
        queue.front = &newNode
    } else {
        queue.rear.next = &newNode
    }
    queue.rear = &newNode
    queue.size++
}

func (queue *LinkedListQueue) Dequeue() (int, error) {
    if queue.size == 0 {
        return -1, errors.New("Can't dequeue from empty queue")
    } else {
        lastFront := queue.front.val
        if queue.size == 1 {
            queue.front = nil
            queue.rear = nil
        } else {
            queue.front = queue.front.next
        }
        queue.size--
        return lastFront, nil
    }
}

func (queue *LinkedListQueue) Front() (int, error) {
    if queue.size == 0 {
        return -1, errors.New("Empty queue has no front")
    } else {
        return queue.front.val, nil
    }
}

func (queue *LinkedListQueue) IsEmpty() bool {
    return queue.size == 0
}

func (queue *LinkedListQueue) Size() int {
    return queue.size
}