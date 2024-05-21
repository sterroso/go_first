package main

import (
	"fmt"
)

func main() {
	animals := []string{}

	animals = append(animals, "dog")
	animals = append(animals, "cat")

	fmt.Println(animals)

	for i, j := 0, len(animals)-1; i < j; i, j = i+1, j-1 {
		animals[i], animals[j] = animals[j], animals[i]
	}

	fmt.Println(animals)

	animals = append(animals, "frog")

	for index := range animals {
		fmt.Printf("My animal %d is %s\n", index, animals[index])
	}

	count := 0

	for count < 5 {
		fmt.Println(count)
		count++
	}
}
