package selectionsort

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex
}

func Sort(arr []int) []int {
	var newArr []int
	var smallestIndex int

	for len(arr) != 0 {
		smallestIndex = findSmallest(arr)
		newArr = append(newArr, arr[smallestIndex])
		arr = removeIndex(arr, smallestIndex)
	}

	return newArr

}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
