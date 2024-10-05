package main

import "fmt"

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("The type of c1 is: %T\n", c1)
	fmt.Printf("The type of c2 is: %T\n", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3", c3)
	fmt.Printf("Type of c3 is: %T\n", c3)
	cZero := c3 - c3
	fmt.Println("cZero:", cZero)
	fmt.Printf("Type of cZero is: %T\n", cZero)
	fmt.Println("----------------------")
	x := 12
	k := 5

	fmt.Printf("Type of x is: %T, val is: %d\n", x, x)
	div := x / k
	fmt.Println("div: ", div)
	fmt.Println("----------------------")

	var m, n float64
	m = 1.223
	fmt.Println("m and n is: ", m, n)

	y := 4 / 2.3
	fmt.Println("y is: ", y)
	divFloat := float64(x) / float64(k)
	fmt.Println("divFloat: ", divFloat)
	fmt.Printf("Type of divFloat is: %T\n", divFloat)
}
