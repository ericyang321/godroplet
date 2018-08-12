package main

import (
	"github.com/ericyang321/godroplet/src/algo/linkedlist"
	"github.com/kr/pretty"
)

func main() {
	l := linkedlist.List{}
	l.Insert(0)
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)

	l.Insert(6)
	l.Print()

	l.Pop()
	l.Print()

	pretty.Println(l.Get(3))
	pretty.Println(l.GetIndex(3))
}
