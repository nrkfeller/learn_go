package main

import "fmt"

func main() {
	myll := newll()
	myll.pushfront(12)
	myll.pushfront(1)
	myll.pushfront(4)
	myll.pushfront(3)
	myll.pushfront(2)
	myll.removeValue(4)
	myll.insertAfter(7, 1)
	traverse(myll)
}

type linkedlist struct {
	root *element
	len  int
}

type element struct {
	value      interface{}
	next, prev *element
}

func (l *linkedlist) removeValue(v interface{}) {
	if l.len == 0 {
		return
	}
	if l.len == 1 {
		l.len--
		l.root = nil
	}
	current := l.root
	previous := l.root
	for {
		if current.value == v {
			previous.next = current.next
			//newelement.next = current.next
			l.len--
			return
		}
		previous = current
		current = current.next
	}
}

func (l *linkedlist) insertAfter(v interface{}, after interface{}) {
	if l.len == 0 {
		return
	}
	current := l.root
	newelement := &element{v, nil, nil}
	for {
		if current.value == after {
			newelement.next = current.next
			current.next = newelement
			l.len++
			return
		}
		current = current.next
	}
}

func traverse(l *linkedlist) {
	current := l.root
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
func newll() *linkedlist {
	return &linkedlist{root: nil, len: 0}
}

func (l *linkedlist) pushfront(v interface{}) {
	newel := &element{v, l.root, nil}
	if l.len == 0 {
		l.root = newel
	} else {
		l.root.prev = newel
		l.root = newel
	}
	l.len++
}
