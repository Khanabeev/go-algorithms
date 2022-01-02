package main

import (
	"fmt"

	"github.com/Khanabeev/go-algorithms/selectionsort"
)

func main() {
	arr := []int{1, 10, 3, 2, 5, 4, 6, 8, 9, 7}

	fmt.Println(selectionsort.Sort(arr))
}
