package main

import "fmt"

func main() {
	vector1 := []int{1, 2, 3, 4}
	vector2 := []int{5, 6, 7, 8}
	sum := 0

	for int i = 0; i < len(vector1); i++{
		sum += vector1[i] * vector2[i]		
	}
	fmt.PrintF()
}
