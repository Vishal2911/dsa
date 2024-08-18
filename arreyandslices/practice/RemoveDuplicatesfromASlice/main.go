package main

import "fmt"

// Problem Statement 4: Remove Duplicates from a Slice
// Task:
// Write a Go function that removes duplicate elements from a slice of integers and returns a new slice with only unique elements.

// Input:
// A slice of integers, e.g., [1, 2, 2, 3, 4, 4, 5].

// Output:
// Unique Elements: [1, 2, 3, 4, 5]

func main() {

	input := []int{1, 2, 2, 3, 4, 4, 5}

	fmt.Println("output  = ", RemoveDuplicates(input))

}

func RemoveDuplicates(slice []int) []int {

	output := []int{}
	exixtingValues := make(map[int]bool)

	// for _, val := range slice { // 1      ; 2 ; 2
	// 	if exixtingValues[val] {// false; false; true
	// 		continue // move to next
	// 	} else {
	// 		output = append(output, val) // 1   ; 1,2
	// 		exixtingValues[val] = true  // 1 = true ; 1=true , 2 = true
	// 	}
	// }


	for _, val := range slice { 
		if !exixtingValues[val] {
			output = append(output, val) 
			exixtingValues[val] = true  
		} 
	}
	return output
}
