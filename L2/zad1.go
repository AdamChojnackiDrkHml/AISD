package main

import (
	"fmt"
	"os"
	"time"
)

func insertSort(array *[]int, N int) {

	for i := 1; i < N; i++ {
		key := (*array)[i]
		j := i - 1
		var swapCounter, compCounter int
		for j >= 0 && (*array)[j] > key {
			fmt.Fprintln(os.Stderr, "comp", (*array)[j], "to ", key)
			compCounter++
			(*array)[j+1] = (*array)[j]
			fmt.Fprintln(os.Stderr, "swap", (*array)[j], "to ", j+1)
			swapCounter++
			j--
		}
		fmt.Fprintln(os.Stderr, "comp", (*array)[j+1], "to ", key)
		compCounter++
		fmt.Fprintln(os.Stderr, "swap", key, "to ", j+1)
		swapCounter++
		(*array)[j+1] = key

		fmt.Fprintln(os.Stderr, "Total Comprasions: ", compCounter)
		fmt.Fprintln(os.Stderr, "Total Swaps: ", swapCounter)

	}

}

func quickSort(arr *[]int, end int, begin int) (swpCounter, cpCounter int) {
	j := end + 1
	i := begin - 1
	pivot := (*arr)[(end+begin)/2]
	var swapCounter, compCounter int
	for {
		i++
		j--

		fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
		compCounter++
		for (*arr)[i] < pivot {
			fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
			compCounter++
			i++
		}

		fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
		compCounter++
		for (*arr)[j] > pivot {
			fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
			compCounter++
			j--
		}

		if i < j {
			fmt.Fprintln(os.Stderr, "Total Swaps: ", swapCounter)
			swapCounter++
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[j]
			(*arr)[j] = temp
			continue
		}

		if i > j {
			break
		}

	}
	j--
	var swpAdd, cpAdd int
	if j > begin {
		swpAdd, cpAdd = quickSort(arr, j, begin)
	}
	swapCounter += swpAdd
	compCounter += cpAdd
	if i < end {
		swpAdd, cpAdd = quickSort(arr, end, i)
	}
	swapCounter += swpAdd
	compCounter += cpAdd

	return swapCounter, compCounter
}

func mergeSort(arr []int) (sorted []int, swpCounter, cpCounter int) {
	length := len(arr)
	var swapCounter, compCounter int

	if length > 1 {

		middle := length / 2

		var swpAdd, cpAdd int

		arr1, swpAdd, cpAdd := mergeSort(arr[:middle])
		swapCounter += swpAdd
		compCounter += cpAdd

		arr2, swpAdd, cpAdd := mergeSort(arr[middle:])
		swapCounter += swpAdd
		compCounter += cpAdd

		arr, swpAdd, cpAdd = merge(arr1, arr2)
		swapCounter += swpAdd
		compCounter += cpAdd

	}

	return arr, swapCounter, compCounter
}

func merge(arr1 []int, arr2 []int) (arr []int, swpCounter, cpCounter int) {

	var swapCounter, compCounter int

	length1, length2 := len(arr1), len(arr2)
	length := length1 + length2

	var array = make([]int, length)

	var i, j int
	k := -1
	for i < length1 && j < length2 {
		k++
		fmt.Fprintln(os.Stderr, "comp", arr1[i], "to ", arr2[j])
		compCounter++

		if arr1[i] <= arr2[j] {
			fmt.Fprintln(os.Stderr, "swap", arr1[i], "to ", k, "place")
			swapCounter++
			array[k] = arr1[i]
			i++
			continue
		}

		fmt.Fprintln(os.Stderr, "swap", arr2[j], "to ", k, "place")
		swapCounter++
		array[k] = arr2[j]
		j++
	}

	k++

	for i < length1 {
		fmt.Fprintln(os.Stderr, "swap", arr1[i], "to ", k, "place")
		swapCounter++
		array[k] = arr1[i]
		i++
		k++
	}

	for j < length2 {
		fmt.Fprintln(os.Stderr, "swap", arr2[j], "to ", k, "place")
		swapCounter++
		array[k] = arr2[j]
		j++
		k++
	}

	return array, swapCounter, compCounter
}
func main() {

	xd := []int{2, 8, 5, 2, 64, 12, 7, 86, 132, 8}

	start := time.Now()
	xd, swapCounter, compCounter := mergeSort(xd)
	end := time.Now()

	time := end.Sub(start)

	fmt.Fprintln(os.Stderr, "Total Comprasions: ", compCounter)
	fmt.Fprintln(os.Stderr, "Total Swaps: ", swapCounter)

	fmt.Printf("%s\n", time)
	for i := range xd {
		fmt.Printf("%d ", xd[i])
	}
}
