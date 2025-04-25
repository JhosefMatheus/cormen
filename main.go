package main

import (
	"cormen/algorithms"
	"cormen/models"
	"fmt"
)

func main() {
	arr := []*models.UserModel{{Id: 3, Name: "Jhosef Matheus"}, {Id: 2, Name: "Vitória"}, {Id: 1, Name: "Sérgio Verton"}}

	for i := range arr {
		fmt.Println(arr[i].ToString())
	}

	fmt.Println()

	sortableArr := make([]models.ISortable, len(arr))

	for i := range arr {
		sortableArr[i] = arr[i]
	}

	algorithms.InsertionSort(sortableArr)

	arr = make([]*models.UserModel, len(sortableArr))

	for i := range sortableArr {
		arr[i] = sortableArr[i].(*models.UserModel)
	}

	for i := range arr {
		fmt.Println(arr[i].ToString())
	}
}
