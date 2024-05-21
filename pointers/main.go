package main

import "fmt"

func main() {
	a := 7
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 9
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	mySlice := []int{
		1, 2, 3,
	}

	for _, value := range mySlice {
		value++
	}

	fmt.Println(mySlice)

	for index := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)
}
