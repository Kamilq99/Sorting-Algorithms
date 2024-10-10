package main

import (
	"fmt"
)

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Printf("Sorted array by Bubble Sort: %v", arr)
}

func main() {

	arr := []int{34, 12, 5, 99, 23, 47, 8, 64, 28, 1}

	fmt.Printf("Unsorted array: %v", arr)

	var choice int

	fmt.Println("Choose a sroting algorithm: ")
	fmt.Println("1 - Bubble sort")
	fmt.Println("2 - Selection Sort")
	fmt.Println("3 - Insertion Sort")
	fmt.Println("4 - Quick sort")
	fmt.Println("5 - Merge sort")
	fmt.Println("6 - Bucket Sort")
	fmt.Println("7 - Heap sort")

	fmt.Scanf("%d", &choice)

	switch choice {

	case 1:
		BubbleSort(arr)
	case 2:
		SelectionSort(arr)
	/*
		case 3:
			InsertionSort(arr)
		case 4:
			QuickSort(arr, 0, len(arr)-1)
		case 5:
			MergeSort(arr, 0, len(arr)-1)
		case 6:
			BucketSort(arr)
		case 7:
			HeapSort(arr)

	*/
	default:
		fmt.Println("Invalid choice")
	}
}
