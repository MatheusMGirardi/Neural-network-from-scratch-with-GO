// Basic neuron
package main

import "fmt"

func main() {
	input := []float32{1, 2, 3}
	weight := []float32{0.2, 0.8, -0.5} // This is a go slice, similar to python lists
	bias := float32(2.0)

	output := (input[0]*weight[0] +
		input[1]*weight[1] +
		input[2]*weight[2] + bias)
	fmt.Printf(" Output: %f \n", output)

	//Layer of neurons
	inputs := []float32{1, 2, 3, 4}
	weight1 := []float32{0.2, 0.8, -0.5}
	weight2 := []float32{0.2, 1.1, 2.1}
	weight3 := []float32{0.2, 0.8, 3.2}

	bias1 := float32(2.0)
	bias2 := float32(8.0)
	bias3 := float32(-1.0)

	outputs := []float32{(input[0]*weight1[0] +
		inputs[1]*weight1[1] +
		inputs[2]*weight1[2] + bias1),
		(inputs[0]*weight2[0] +
			inputs[1]*weight2[1] +
			inputs[2]*weight2[2] + bias2),
		(inputs[0]*weight3[0] +
			inputs[1]*weight3[1] +
			inputs[2]*weight3[2] + bias3)}

	fmt.Printf("Neuron 1: %f, Neuron 2: %f, Neuron 3: %f\n", outputs[0], outputs[1], outputs[2])

	scalable_neurons()
}

func scalable_neurons() {
	input := []float32{1, 2, 3, 4}
	weights := [][]float32{
		{0.2, 0.8, -0.5, 1.0},
		{0.5, -0.91, 0.26, -0.5},
		{-0.26, -0.27, 0.17, 0.87},
		{1.50, -0.27, 2.6, 0.87}}

	bias := []float32{1, -1, 4, 5}

	var outputs []float32 = make([]float32, 0)

	for _, row := range weights {
		neuron_output := float32(0)
		for j, weight := range row {
			neuron_output += weight * input[j]
			neuron_output += bias[j]
		}
		outputs = append(outputs, neuron_output)
	}

	for i, output := range outputs {
		fmt.Printf("Neuron %d output: %v \n", i, output)
	}
}
