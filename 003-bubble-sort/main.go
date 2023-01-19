package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(arr *[]int) {
	length := len(*arr)

	var first, second int
	var sorted bool

	for i := 0; i < length; i++ {
		sorted = true
		for j := 1; j < length-i; j++ {
			first = (*arr)[j-1]
			second = (*arr)[j]
			if first > second {
				(*arr)[j-1] = first
				(*arr)[j] = second
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}

func main() {
	var arr []int

	for len(arr) < 1000 {
		random := rand.Intn(10000)
		arr = append(arr, random)
	}

	bubbleSort(&arr)

	fmt.Println(arr)
}
