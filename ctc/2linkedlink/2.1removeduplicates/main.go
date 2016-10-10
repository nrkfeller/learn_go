package main

import (
	"container/list"
	"fmt"
)

func (l *list.List) traverse() {
	fmt.Println(l.Len())
}

func main() {
	myll := list.New()
	myll.PushBack(10)
	myll.traverse()
	fmt.Printf("%T", myll)
}
