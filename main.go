package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func (l *List) Insert(d int) {
	list := &ListNode{Val: d, Next: nil}
	if l.Head == nil {
		l.Head = list
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = list
	}
}
func (l *List) Show() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("-> %v ", curr.Val)
		curr = curr.Next
	}
	fmt.Print("\n")
}

func (l *List) Count() int {
	curr := l.Head
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}

	return count
}

func (l *List) ExtractMiddle() *ListNode {
	count := l.Count()
	if count == 0 {
		return nil
	}
	middleIndex := count / 2

	found := l.Head
	currInd := 0
	for found != nil && currInd < middleIndex {
		currInd++
		found = found.Next
	}

	return found
}

func (l *List) Invert() *List {
	if l == nil {
		return nil
	}

	curr := l.Head.Next
	prev := &ListNode{
		Val: l.Head.Val,
	}
	for curr != nil {
		prev = &ListNode{
			Val:  curr.Val,
			Next: prev,
		}
		curr = curr.Next
	}

	return &List{
		Head: prev,
	}
}

func main() {
	sl := List{}
	for i := 0; i < 5; i++ {
		sl.Insert(rand.Intn(100))
	}
	sl.Show()
	fmt.Println(sl.Count())

	sl2 := List{}
	for _, item := range []int{1, 2, 3, 4, 5} {
		sl2.Insert(item)
	}
	fmt.Println(sl2.ExtractMiddle().Val)

	sl3 := List{}
	for _, item := range []int{1, 2, 3, 4, 5, 6} {
		sl3.Insert(item)
	}
	fmt.Println(sl3.ExtractMiddle().Val)

	sl3.Show()
	sl4 := sl3.Invert()
	sl4.Show()
}
