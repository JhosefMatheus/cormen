package algorithms

import (
	"cormen/models"
	"fmt"
)

func merge(arr []models.ISortable, left, mid, right int) {
	temp := []models.ISortable{}

	i := left
	j := mid + 1

	for i <= mid && j <= right {
		if arr[i].Key() < arr[j].Key() {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}

	for j <= right {
		temp = append(temp, arr[j])
		j++
	}

	for k := 0; k < len(temp); k++ {
		arr[left+k] = temp[k]
	}
}

func MergeSort(arr []models.ISortable, left, right int) {
	if left >= right {
		return
	}

	fmt.Println(left, right)

	mid := (left + right) / 2
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}
