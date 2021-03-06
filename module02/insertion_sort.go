package module02

import (
	"fmt"
	"sort"
)

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	var sorted []int
	for _, item := range list {
		inserted := 0
		for j, val := range sorted {
			if item < val {
				inserted = 1
				sorted = append(sorted[:j], append([]int{item}, sorted[j:]...)...)
				break
			}
		}
		if inserted == 0 {
			sorted = append(sorted, item)
		}
	}
	copy(list, sorted)
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
	for _, v := range list {
		fmt.Println(v)
	}
}
