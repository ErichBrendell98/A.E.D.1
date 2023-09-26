package list

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) Add(data interface{}) {
	newNode := &Node{data: data, next: nil, prev: dll.tail}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		dll.tail = newNode
	}
	dll.size++
}

func (dll *DoublyLinkedList) AddOnIndex(index int, data interface{}) {
	if index < 0 || index > dll.size {
		fmt.Println("Índice inválido.")
		return
	}

	newNode := &Node{data: data, next: nil, prev: nil}

	if index == 0 {
		newNode.next = dll.head
		if dll.head != nil {
			dll.head.prev = newNode
		}
		dll.head = newNode
		if dll.tail == nil {
			dll.tail = newNode
		}
	} else if index == dll.size {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	} else {
		current := dll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		newNode.prev = current
		current.next.prev = newNode
		current.next = newNode
	}
	dll.size++
}

func (dll *DoublyLinkedList) RemoveOnIndex(index int) {
	if index < 0 || index >= dll.size {
		fmt.Println("Índice inválido.")
		return
	}

	if index == 0 {
		dll.head = dll.head.next
		if dll.head != nil {
			dll.head.prev = nil
		}
		if dll.head == nil {
			dll.tail = nil
		}
	} else if index == dll.size-1 {
		dll.tail = dll.tail.prev
		if dll.tail != nil {
			dll.tail.next = nil
		}
		if dll.tail == nil {
			dll.head = nil
		}
	} else {
		current := dll.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		current.prev.next = current.next
		current.next.prev = current.prev
	}
	dll.size--
}

func (dll *DoublyLinkedList) Get(index int) interface{} {
	if index < 0 || index >= dll.size {
		fmt.Println("Índice inválido.")
		return nil
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data
}

func (dll *DoublyLinkedList) Set(index int, data interface{}) {
	if index < 0 || index >= dll.size {
		fmt.Println("Índice inválido.")
		return
	}

	current := dll.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.data = data
}

func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

func main() {
	dll := NewDoublyLinkedList()

	dll.Add(1)
	dll.Add(2)
	dll.Add(3)

	fmt.Println("Size:", dll.Size()) // Size: 3

	dll.AddOnIndex(1, 4)

	fmt.Println("Get(1):", dll.Get(1)) // Get(1): 4

	dll.Set(2, 5)

	fmt.Println("Get(2):", dll.Get(2)) // Get(2): 5

	dll.RemoveOnIndex(0)

	fmt.Println("Size:", dll.Size()) // Size: 3
}