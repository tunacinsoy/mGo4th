// The aim of this program is to get familiar with slices package
package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, -1, -2}
	s2 := slices.Clone(s1)
	s3 := slices.Clone(s1[2:])
	fmt.Println(s1[2], s2[2], s3[0])
	s1[2] = 0
	s1[3] = 0
	fmt.Println(s1[2], s2[2], s3[0])
	fmt.Println(s1)
	// Deletes duplicate elements within a slice
	s1 = slices.Compact(s1)
	fmt.Println("s1 (compact):", s1)
	fmt.Println(slices.Contains(s1, 2), slices.Contains(s1, -2))

	s4 := make([]int, 10, 100)
	fmt.Println(s4)
	fmt.Println("Len:", len(s4), "Cap:", cap(s4))
	s4 = slices.Clip(s4)
	fmt.Println("Len:", len(s4), "Cap:", cap(s4))

	fmt.Println("Min:", slices.Min(s1), "Max:", slices.Max(s1))

	fmt.Println("s2 before operations", s2)

	// Replace elements s2[1], s2[2] with 100 and 200 respectively, and then append 300 after s2[2]
	s2 = slices.Replace(s2, 1, 3, 100, 200, 300)
	fmt.Println("s2 (replaced):", s2)
	slices.Sort(s2)
	fmt.Println("s2 sorted:", s2)
}
