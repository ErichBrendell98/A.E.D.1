package main

import (
	"fmt"
)

type List interface {
	AddLast(value int)
	AddPos(value int, pos int)
	Update(value int, pos int)
	RemoveLast()
	Remove(pos int)
	Get(pos int) int
	Size() int
}

type ArrayList struct {
	values []int
	inserted int
}

func Init(size int) ArrayList {
	return ArrayList{values: make([]int, 10), inserted: 0}
}

func (list *ArrayList) Add(val int) {

}

func main() {
	fmt.Println("Hello, World!")
}
