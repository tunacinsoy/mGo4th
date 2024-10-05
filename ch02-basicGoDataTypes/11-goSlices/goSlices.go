// The aim of this program is to get familiar with slices
package main

import "fmt"

func main() {
	aSlice := []float64{}

	fmt.Println(aSlice, len(aSlice), cap(aSlice)) // [] 0 0
	aSlice = append(aSlice, -1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, len(aSlice), cap(aSlice)) // [-1234.56 -34] 2 2

	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4

	t = append(t, -5)
	fmt.Println(t)

	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	// The outer for loop actually gets the rows of twoD matrix and assigns it into variable i
	// Then variable i becomes []int obviously, and then variable k traverses each of the elements of these arrays

	for _, i := range twoD { // outer loop traverses rows
		fmt.Printf("%T, %v \n", i, i) // []int, [1 2 3] then [int], [4 5 6]
		for _, k := range i {         // inner loop traverses columns
			fmt.Printf("Val: %T, %v ", k, k) // Val: int, 1 Val: int, 2 Val: int, 3 AND Val: int, 4 Val: int, 5 Val: int, 6
		}
		fmt.Println()
	}

	// Either definition is correct
	make2D := make([][]int, 2)
	var make2DwithVar [][]int = make([][]int, 2) // 2 rows, and we do not know the count of columns

	fmt.Println(make2D)
	fmt.Println(make2DwithVar)

	make2D[0] = []int{1, 2, 3, 4} // initialize first row
	make2D[1] = []int{5, 6, 7, 8} // initialize second row
	fmt.Println(make2D)

	for _, i := range make2D {
		for _, k := range i {
			fmt.Printf("%d ", k)
		}
		fmt.Println()
	}

}
