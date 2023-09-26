package list

type Node struct {
    data interface{}
    next *Node
}

type LinkedList struct {
    head *Node
    size int
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (ll *LinkedList) Add(data interface{}) {
    newNode := &Node{data: data, next: nil}
    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
    ll.size++
}

func (ll *LinkedList) AddOnIndex(index int, data interface{}) {
    if index < 0 || index > ll.size {
        fmt.Println("Índice inválido.")
        return
    }

    newNode := &Node{data: data, next: nil}

    if index == 0 {
        newNode.next = ll.head
        ll.head = newNode
    } else {
        current := ll.head
        for i := 0; i < index-1; i++ {
            current = current.next
        }
        newNode.next = current.next
        current.next = newNode
    }
    ll.size++
}

func (ll *LinkedList) RemoveOnIndex(index int) {
    if index < 0 || index >= ll.size {
        fmt.Println("Índice inválido.")
        return
    }

    if index == 0 {
        ll.head = ll.head.next
    } else {
        current := ll.head
        for i := 0; i < index-1; i++ {
            current = current.next
        }
        current.next = current.next.next
    }
    ll.size--
}

func (ll *LinkedList) Get(index int) interface{} {
    if index < 0 || index >= ll.size {
        fmt.Println("Índice inválido.")
        return nil
    }

    current := ll.head
    for i := 0; i < index; i++ {
        current = current.next
    }
    return current.data
}

func (ll *LinkedList) Set(index int, data interface{}) {
    if index < 0 || index >= ll.size {
        fmt.Println("Índice inválido.")
        return
    }

    current := ll.head
    for i := 0; i < index; i++ {
        current = current.next
    }
    current.data = data
}

func (ll *LinkedList) Size() int {
    return ll.size
}

func main() {
    linkedList := NewLinkedList()

    linkedList.Add(1)
    linkedList.Add(2)
    linkedList.Add(3)

    fmt.Println("Size:", linkedList.Size()) // Size: 3

    linkedList.AddOnIndex(1, 4)

    fmt.Println("Get(1):", linkedList.Get(1)) // Get(1): 4

    linkedList.Set(2, 5)

    fmt.Println("Get(2):", linkedList.Get(2)) // Get(2): 5

    linkedList.RemoveOnIndex(0)

    fmt.Println("Size:", linkedList.Size()) // Size: 3
}