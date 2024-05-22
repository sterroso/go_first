package main

import (
	"fmt"
	"math"
)

const (
	PI = 3.141592653589
)

type Circle struct {
	Radius float64
}

func (c Circle) GetPerimeter() float64 {
	return 2 * c.Radius * PI
}

func (c Circle) GetArea() float64 {
	return PI * math.Pow(c.Radius, 2)
}

func main() {
	myCircle := Circle{
		Radius: 10.11,
	}

	fmt.Printf("My circle radius: %.2f\n", myCircle.Radius)
	fmt.Printf("My circle perimeter: %.4f\n", myCircle.GetPerimeter())
	fmt.Printf("My circle area: %.4f\n", myCircle.GetArea())
}
