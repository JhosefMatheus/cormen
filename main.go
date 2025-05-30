package main

import (
	"cormen/algorithms"
	"cormen/models"
	"fmt"
)

func main() {
	arr := []*models.UserModel{{Id: 3, Name: "Jhosef Matheus"}, {Id: 2, Name: "Vitória"}, {Id: 1, Name: "Sérgio Verton"}}

	insertionSortedArr := make([]models.ISortable, len(arr))
	mergeSortedArr := make([]models.ISortable, len(arr))

	for i := range arr {
		fmt.Println(arr[i].ToString())
		insertionSortedArr[i] = arr[i]
		mergeSortedArr[i] = arr[i]
	}

	fmt.Println()

	algorithms.InsertionSort(insertionSortedArr)

	algorithms.MergeSort(mergeSortedArr, 0, len(mergeSortedArr)-1)

	fmt.Println("Array sorted by insertion sort:")
	for _, v := range insertionSortedArr {
		user := v.(*models.UserModel)
		fmt.Println(user.ToString())
	}

	fmt.Println("\nArray sorted by merge sort:")
	for _, v := range mergeSortedArr {
		user := v.(*models.UserModel)
		fmt.Println(user.ToString())
	}
}
