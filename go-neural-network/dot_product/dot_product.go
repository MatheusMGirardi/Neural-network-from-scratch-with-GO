package main

import "fmt"

func dotproduct(vector1 []float64, vector2 []float64) float64 {
	output := 0.0
	for k := 0; k < len(vector1); k++ {
		output += vector1[k] * vector2[k]
	}
	return output
}

func main() {
	vector1 := []float64{0.5, 1.5, 2.5, 3.5}
	vector2 := []float64{4.5, 5.5, 6.5, 7.5}
	bias := 2.0

	result := dotproduct(vector1, vector2) + bias
	fmt.Printf("Dot product result: %f", result)
}
