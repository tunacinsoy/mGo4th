// The aim of this program is to define custom interface and implement it with defined struct type
package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	Radius float64
}

// type circle implements Shape2D interface by defining Perimeter() function
func (c circle) Perimeter() float64 {
	var perimeter float64 = 2 * math.Pi * c.Radius
	return perimeter
}

func main() {

	circle1 := circle{Radius: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f \n", circle1.Radius, circle1.Perimeter()) // R 1.50 -> Perimeter 9.425

	// Checks whether circle1 is shape2d or not
	_, ok := interface{}(circle1).(Shape2D)

	if ok {
		fmt.Println("a is a shape2d!") // a is a shape2d!
	}

	i := 12

	_, ok = interface{}(i).(int)

	if ok {
		fmt.Println("i is an int!") // i is an int!
	}

	j := 12.4
	_, ok = interface{}(j).(float64)

	if ok {
		fmt.Println("j is a float64!") // j is a float64!
	}

}
