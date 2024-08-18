package main

import "fmt"

// Problem Statement 6: Rotate an Array
// Task:
// Write a Go function that rotates the elements of an array to the right by a given number of steps.

// Input:
// An array of integers, e.g., [1, 2, 3, 4, 5] and a number of steps, e.g., 2.

// Output:
// Rotated Array: [4, 5, 1, 2, 3]
func main() {
slice:= []int{1, 2, 3, 4, 5}

fmt.Println("rotated result = ", rotate(slice , 2))
}

func rotate(slice []int, index int) []int {
	result := []int{}
	for i := index + 1; i < len(slice); i++ {
		result = append(result, slice[i])
	}

	for i := 0; i <= index; i++ {
		result = append(result, slice[i])
	}

	return result
}
