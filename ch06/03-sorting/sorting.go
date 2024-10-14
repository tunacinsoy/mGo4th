// The aim of this program is to get familiar with functions that accept function as input argument, like sort.Slice
package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {

	data := []Grades{{"J.", "Lewis", 10}, {"M.", "Tsoukalos", 7},
		{"D.", "Tsoukalos", 8}, {"J.", "Lewis", 9}}

	isSorted := sort.SliceIsSorted(data, func(i int, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	if isSorted {
		fmt.Println("Sorted")
	} else {
		fmt.Println("Not sorted, sorting now...")

		sort.Slice(data, func(i int, j int) bool {
			if data[i].Grade < data[j].Grade {
				return true
			} else {
				return false
			}
		})

	}

	fmt.Println("By grade:", data)

}
