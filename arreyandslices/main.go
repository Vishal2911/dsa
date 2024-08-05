package main

import "fmt"

func main() {
	// arrey()
	// slice()
	min, max := minmax([]int{5, 9, 8, 7, 4, 6, 5, 66, 5, 5, -54, -845, -545, -5, -4, 4})

	fmt.Printf("Min = %v , max = %d ,  summ = %d  , reverse = %v\n", min, max, sum([]int{1, 2, 3, 4, 5}), reverse([]int{1, 2, 3, 4, 5}))
}

func arrey() {
	var arr [5]int

	arr = [5]int{1, 2, 3, 4}

	// Insertion
	arr[0] = 10

	// Traversal
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Deletion (set to zero)
	arr[1] = 0
	fmt.Println("After deletion:", arr)
}

func slice() {
	slice := []int{1, 2, 3, 4, 5}

	// Insertion
	slice = append(slice, 6)
	slice = append(slice, 6)
	slice = append(slice, 6)
	slice = append(slice, 6)
	slice = append(slice, 6)
	// Traversal
	for i, v := range slice {
		fmt.Println("indel =", i, " val = ", v)
	}
	a := slice[5]
	fmt.Println("a = ", a)
	// Deletion
	i := 2
	slice = append(slice[:i], slice[i+1:]...)
	fmt.Println("After deletion:", slice)
}

// 5,4,8,3,-5,-15
func minmax(arr []int) (int, int) {

	if len(arr) == 0 {
		return 0, 0
	}
	min, max := arr[0], arr[0]

	for _, val := range arr {

		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
	}

	return min, max
}

func sum(arr []int) int {

	sum := 0

	for _, val := range arr {

		sum = sum + val
	}

	return sum
}

// 5,4,8,3,-5,-15
func reverse(arr []int) []int {

	if len(arr) == 0 {
		return nil
	}
	reverse := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		reverse = append(reverse, arr[i])
	}

	return reverse
}
