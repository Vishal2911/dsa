package main

import "fmt"

// Problem Statement 5: Find the Intersection of Two Arrays
// Task:
// Write a Go function that takes two arrays and returns a new array containing the elements that appear in both arrays.

// Input:
// Two arrays of integers, e.g., [1, 2, 3, 4, 5] and [3, 4, 5, 6, 7].

// Output:
// Intersection: [3, 4, 5]

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	fmt.Printf("final output = %v", intersection(slice1 , slice2))

}

func intersection(slice1, slice2 []int) []int {

	result := []int{}
	m := make(map[int]bool)

	for _ , val := range slice1 {
		m[val] = true
	}

	//m [1= true , 2 = true , .... 5 = true]

	for _ , val := range slice2 {
		if m[val] {
			result = append(result,  val)
		}
	}

	return result
}
