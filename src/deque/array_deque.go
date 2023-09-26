package deque

import "fmt"

type Deque struct {
	array     []interface{}
	capacity  int
	front     int
	rear      int
	size      int
}

func NewDeque(capacity int) *Deque {
	deque := &Deque{
		array:    make([]interface{}, capacity),
		capacity: capacity,
		front:    -1,
		rear:     -1,
		size:     0,
	}
	return deque
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) IsFull() bool {
	return d.size == d.capacity
}

func (d *Deque) Front() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.array[d.front]
}

func (d *Deque) Back() interface{} {
	if d.IsEmpty() {
		return nil
	}
	return d.array[d.rear]
}

func (d *Deque) PushFront(item interface{}) {
	if d.IsFull() {
		fmt.Println("O deque está cheio. Não é possível inserir.")
		return
	}

	if d.front == -1 {
		d.front = 0
		d.rear = 0
	} else if d.front == 0 {
		d.front = d.capacity - 1
	} else {
		d.front--
	}

	d.array[d.front] = item
	d.size++
}

func (d *Deque) PushBack(item interface{}) {
	if d.IsFull() {
		fmt.Println("O deque está cheio. Não é possível inserir.")
		return
	}

	if d.front == -1 {
		d.front = 0
		d.rear = 0
	} else if d.rear == d.capacity-1 {
		d.rear = 0
	} else {
		d.rear++
	}

	d.array[d.rear] = item
	d.size++
}

func (d *Deque) PopFront() interface{} {
	if d.IsEmpty() {
		fmt.Println("O deque está vazio. Não é possível remover.")
		return nil
	}

	item := d.array[d.front]

	if d.front == d.rear {
		d.front = -1
		d.rear = -1
	} else if d.front == d.capacity-1 {
		d.front = 0
	} else {
		d.front++
	}

	d.size--
	return item
}

func (d *Deque) PopBack() interface{} {
	if d.IsEmpty() {
		fmt.Println("O deque está vazio. Não é possível remover.")
		return nil
	}

	item := d.array[d.rear]

	if d.front == d.rear {
		d.front = -1
		d.rear = -1
	} else if d.rear == 0 {
		d.rear = d.capacity - 1
	} else {
		d.rear--
	}

	d.size--
	return item
}

func main() {
	deque := NewDeque(5)

	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushFront(0)

	fmt.Println("Front:", deque.Front()) // Front: 0
	fmt.Println("Back:", deque.Back())   // Back: 2

	fmt.Println("PopFront:", deque.PopFront()) // PopFront: 0
	fmt.Println("PopBack:", deque.PopBack())   // PopBack: 2
}