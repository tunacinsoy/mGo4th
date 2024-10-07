// The main package is defined, indicating this is an executable program
package main

// The fmt package is imported for outputting values to the console.
import "fmt"

// A generic node is defined for the linked list.
// The type parameter [T any] allows the list to store values of any type.
type node[T any] struct {
	// The field Data stores a value of type T.
	// Since T is a generic type parameter,
	// Data can hold any type of value (e.g., int, string, custom structs, etc.).
	Data T
	// Points to the next node in the list, making this a singly linked list
	next *node[T]
}

// Lot less confusing node declaration
type node1 struct {
	data float64
	next *node1
}

// defines a generic list
type list[T any] struct {
	// pointer to the first node (the head) of the list.
	start *node[T]
}

// This method is tied to a generic list (list[T]), allowing to add elements of type T to the list.
func (l *list[T]) add(data T) {
	// A new node n is created with the data provided and next set to nil.
	n := node[T]{
		Data: data,
		next: nil,
	}

	// If the list is empty (l.start == nil), the newly created node becomes the first node in the list (l.start).
	if l.start == nil {
		l.start = &n
		return
	}

	// If there’s only one node in the list (i.e., l.start.next == nil), the new node is added as the next node in the list.
	if l.start.next == nil {
		l.start.next = &n
		return
	}

	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
}

func main() {
	var myList list[int]
	fmt.Println(myList)
	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)
	fmt.Println(myList)

	cur := myList.start

	for {

		if cur == nil {
			break
		}
		fmt.Println("*", cur)
		cur = cur.next
	}
}
