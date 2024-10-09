// The aim of this program is to understand implemeting sort.Interface to create our own sort functions
package main

import (
	"fmt"
	"sort"
)

type Size struct {
	F1 int
	F2 string
	F3 int
}

type Person struct {
	F1 int
	F2 string
	F3 Size
}

// In Go, when we say "implements sort.Interface,"
// it means that a type provides the required methods
// (Len(), Less(), and Swap()) defined by the sort.Interface.
// This allows the sort.Sort() function to sort that type according to the
// logic defined in those methods.
type Personslice []Person

func (a Personslice) Len() int {
	return len(a)
}

func (a Personslice) Less(i int, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func (a Personslice) Swap(i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func main() {

	data := []Person{
		{1, "One", Size{1, "Person_1", 10}},
		{2, "Two", Size{2, "Person_2", 20}},
		{-1, "Two", Size{-1, "Person_3", -20}},
	}

	fmt.Println("Before:", data)
	sort.Sort(Personslice(data))
	fmt.Println("After:", data)

	a := []int{2, 3, 4, 5, 6, 1}
	fmt.Println("a", a)
	sort.Ints(a)
	fmt.Println("a", a)
	sort.Sort(sort.Reverse(Personslice(data)))
	fmt.Println("Reverse:", data)
}
