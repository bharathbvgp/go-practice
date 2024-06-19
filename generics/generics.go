package main

import "fmt"

func mapKeys[K comparable , V any](m map[K]V) []K{
	keys := make([]K , 0 , len(m))
	for k := range m {
		keys = append(keys , k);
	}
	return keys
}


type element[T any] struct {
	next *element[T]
	val T
}

type List[T any] struct {
	head , tail *element[T]
}

func (l *List[T]) Push(val T) {
	if l.tail == nil {
		l.head = &element[T]{ val : val }
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{ val : val }
		l.tail = l.tail.next
	}
}

func (l *List[T]) GetAll() []T {
	var elems []T
	for i := l.head ; i != nil ; i = i.next {
		elems = append(elems, i.val);
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("Keys" , mapKeys(m))
	lst := List[int]{}
	lst.Push(10);
	lst.Push(100);
	lst.Push(200);
	all := lst.GetAll();
	fmt.Println(all);
}