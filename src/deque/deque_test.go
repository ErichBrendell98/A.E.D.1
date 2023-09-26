package deque

import (
    "testing"
)

var size int

var deques [1]IDeque

//var deques [2]IDeque

func createDeques(size int) {
    //array_deque := &ArrayDeque{}
    doublylinkedlist_deque := &DoublyLinkedListDeque{}
    //(*array_deque).Init(size)
    deques = [1]IDeque{doublylinkedlist_deque}
    //deques = [2]IDeque{array_deque, doublylinkedlist_deque}
}

func deleteDeques() {
    //in this case, createQueues alone solves the problem
    //however, I let the template here to be used in other tests
    deques[0] = nil
    //deques[1] = nil
}

func setupTest() func() {
    //before each test
    size = 10
    createDeques(size)

    //after each test
    return func() {
        deleteDeques()
    }
}

func TestEnqueueRear(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        for i := 0; i < 2*size; i++ {
            deque.EnqueueRear(i)
            if deque.Size() != i+1 {
                t.Errorf("%T size is %d, but we expected it to be %d", deque, deque.Size(), i+1)
            }
        }
    }
}

func TestEnqueueFront(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        for i := 0; i < 2*size; i++ {
            deque.EnqueueFront(i)
            if deque.Size() != i+1 {
                t.Errorf("%T size is %d, but we expected it to be %d", deque, deque.Size(), i+1)
            }
        }
    }
}

func TestDequeueFront(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        for i := 0; i < size; i++ {
            deque.EnqueueRear(i)
        }
        for i := 0; i < size; i++ {
            val, err := deque.DequeueFront()
            //front ==> 0, 1, 2, 3, ..., 9 ==> rear
            if val != i {
                t.Errorf("Value dequeued from %T is %d, but we expected it to be %d", deque, val, i)
            }
            if err != nil {
                t.Errorf(err.Error())
            }
            if deque.Size() != size-i-1 {
                t.Errorf("%T size is %d, but we expected it to be %d", deque, deque.Size(), size-i-1)
            }
        }
    }
}

func TestDequeueRear(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        for i := 0; i < size; i++ {
            deque.EnqueueRear(i)
        }
        for i := 0; i < size; i++ {
            val, err := deque.DequeueRear()
            //front ==> 0, 1, 2, 3, ..., 9 ==> rear
            if val != size-i-1 {
                t.Errorf("Value dequeued from %T is %d, but we expected it to be %d", deque, val, size-i-1)
            }
            if err != nil {
                t.Errorf(err.Error())
            }
            if deque.Size() != size-i-1 {
                t.Errorf("%T size is %d, but we expected it to be %d", deque, deque.Size(), size-i-1)
            }
        }
    }
}

func TestDequeueRearEmptyQueue(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        _, err := deque.DequeueRear()
        if err == nil {
            t.Errorf(err.Error())
        }
    }
}

func TestDequeueFrontEmptyQueue(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        _, err := deque.DequeueFront()
        if err == nil {
            t.Errorf(err.Error())
        }
    }
}

func TestFrontAfterEnqueueRear(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        for i := 0; i < size; i++ {
            deque.EnqueueRear(i)
            val, err := deque.Front()
            if val != 0 {
                t.Errorf("Value in front of from %T is %d, but we expected it to be %d", deque, val, 0)
            }
            if err != nil {
                t.Errorf(err.Error())
            }
        }
    }
}

func TestFrontAfterEnqueueFront(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        for i := 0; i < size; i++ {
            deque.EnqueueFront(i)
            val, err := deque.Front()
            if val != i {
                t.Errorf("Value in front of from %T is %d, but we expected it to be %d", deque, val, i)
            }
            if err != nil {
                t.Errorf(err.Error())
            }
        }
    }
}

func TestFrontEmptyQueue(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        _, err := deque.Front()
        if err == nil {
            t.Errorf("Empty %T should return error when asked the front element", deque)
        }
    }
}

func TestIsEmptyAfterEnqueueRear(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        deque.EnqueueRear(0)
        empty := deque.IsEmpty()
        if empty {
            t.Errorf("Non Empty %T should return false for IsEmpty operation", deque)
        }
    }
}

func TestIsEmptyAfterEnqueueFront(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        deque.EnqueueFront(0)
        empty := deque.IsEmpty()
        if empty {
            t.Errorf("Non Empty %T should return false for IsEmpty operation", deque)
        }
    }
}

func TestIsEmptyOnEmptyQueue(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        empty := deque.IsEmpty()
        if !empty {
            t.Errorf("Empty %T should return true for IsEmpty operation", deque)
        }
    }
}

func TestSizeAfterEnqueueRear(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        deque.EnqueueRear(0)
        size := deque.Size()
        if size != 1 {
            t.Errorf("%T size is %d, but we expected it to be %d", deque, size, 1)
        }
    }
}

func TestSizeAfterEnqueueFront(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        deque.EnqueueFront(0)
        size := deque.Size()
        if size != 1 {
            t.Errorf("%T size is %d, but we expected it to be %d", deque, size, 1)
        }
    }
}

func TestSizeEmptyQueue(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        size := deque.Size()
        if size != 0 {
            t.Errorf("%T size is %d, but we expected it to be %d", deque, size, 0)
        }
    }
}

func TestEnqueueCircularRight(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        for i := 0; i < size; i++ {
            deque.EnqueueRear(i)
        }
        for i := 0; i < size-2; i++ {
            deque.DequeueFront()
        }
        //front ==> 8, 9 <== rear
        for i := 10; i < 16; i++ {
            deque.EnqueueRear(i)
        }

        //front ==> 8, 9, 10, 11, 12, 13, 14, 15 <== rear
        //se for vetor, 10 a 15 foram inseridos de forma circular, no início

        for i := 8; i < 16; i++ {
            val, err := deque.Front()
            if val != i {
                t.Errorf("Value dequeued from %T is %d, but we expected it to be %d", deque, val, i)
            }
            if err != nil {
                t.Errorf(err.Error())
            }
            deque.DequeueFront()
        }

        size := deque.Size()
        if size != 0 {
            t.Errorf("%T size is %d, but we expected it to be %d", deque, size, 0)
        }
    }
}

func TestEnqueueCircularLeft(t *testing.T) {
    defer setupTest()()
    for _, deque := range deques {
        for i := 0; i < size; i++ {
            deque.EnqueueRear(i)
        }
        for i := 0; i < 2; i++ {
            deque.DequeueRear()
        }
        //front ==> 0, 1, ..., 7 <== rear
        for i := -1; i > -3; i-- {
            deque.EnqueueFront(i)
        }

        //front ==> -2, -1, 0, 1, ..., 7 <== rear
        //se for vetor, -2 e -1 foram inseridos de forma circular, no final

        for i := -2; i < 8; i++ {
            val, err := deque.Front()
            if val != i {
                t.Errorf("Value dequeued from %T is %d, but we expected it to be %d", deque, val, i)
            }
            if err != nil {
                t.Errorf(err.Error())
            }
            deque.DequeueFront()
        }

        size := deque.Size()
        if size != 0 {
            t.Errorf("%T size is %d, but we expected it to be %d", deque, size, 0)
        }
    }
}