package module02

import "sort"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	length := len(list)
	for j := 0; j < length-1; j++ {
		swapped := 0
		for i := 0; i < length-1-j; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = 1
			}
		}
		if swapped == 0{
			break
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	length := list.Len()
	for j := 0; j < length-1; j++ {
		swapped := 0
		for i := 0; i < length-1-j; i++ {
			if list.Less(i+1, i) {
				list.Swap(i, i+1)
				swapped = 1
			}
		}
		if swapped == 0 {
			break
		}
	}
}
