package algorithms

import "cormen/models"

func InsertionSort(arr []models.ISortable) {
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		j := i - 1

		for j >= 0 && arr[j].Key() > curr.Key() {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = curr
	}
}
