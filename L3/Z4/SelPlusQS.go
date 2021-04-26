package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Select(arr *[]int, p int, q int, i int) (answer, swpCounter, cpCounter int) {
	var swapCounter, compCounter int
	var swpAdd, cpAdd int
	compCounter++
	if p == q {
		return (*arr)[p], swapCounter, compCounter
	}

	pivot, swpAdd, cpAdd := MedianOfMedians(arr, p, q)
	compCounter += cpAdd
	swapCounter += swpAdd

	r, swpAdd, cpAdd := partition(arr, p, q, pivot)
	compCounter += cpAdd
	swapCounter += swpAdd

	k := r - p + 1

	if i == k {
		compCounter++
		return (*arr)[r], swapCounter, compCounter
	} else if i < k {
		compCounter += 2
		ans, swpAdd, cpAdd := Select(arr, p, r-1, i)
		compCounter += cpAdd
		swapCounter += swpAdd
		return ans, swapCounter, compCounter
	} else {
		ans, swpAdd, cpAdd := Select(arr, r+1, q, i-k)
		compCounter += cpAdd
		swapCounter += swpAdd
		return ans, swapCounter, compCounter
	}

}

func MedianOfMedians(arr *[]int, p int, q int) (MedianOfMedians, swpCounter, cpCounter int) {
	var swapCounter, compCounter int
	if q-p < 5 {
		return medianOfSmallArr(arr, p, q)
	}

	medianArray := make([]int, 0)

	for i := p; i <= q; i = i + 5 {
		subRight := i + 4

		if subRight > q {
			subRight = q
		}
		median, swpAdd, cpAdd := medianOfSmallArr(arr, i, subRight)
		compCounter += cpAdd
		swapCounter += swpAdd
		medianArray = append(medianArray, median)
	}
	arrLen := len(medianArray)
	return Select(&medianArray, 0, arrLen-1, (arrLen+1)/2)
}

func medianOfSmallArr(arr *[]int, p int, q int) (median, swpCounter, cpCounter int) {
	var swapCounter, compCounter int
	if p == q {
		return (*arr)[p], swapCounter, compCounter
	}

	for i := p + 1; i < q+1; i++ {
		j := i - 1

		key := (*arr)[i]
		compCounter++
		for j >= p && (*arr)[j] > key {
			compCounter++
			(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			swapCounter++
			j--
		}
		swapCounter++
		(*arr)[j+1] = key
	}

	return (*arr)[(p+q)/2], swapCounter, compCounter
}

func partition(arr *[]int, left int, right int, pivot int) (pivotIndex, swpCounter, cpCounter int) {
	i := left
	var pivotPos int
	var swapCounter, compCounter int

	for k := left; k <= right; k++ {

		if (*arr)[k] == pivot {
			pivotPos = k
			break
		}
	}
	swapCounter++
	(*arr)[pivotPos], (*arr)[right] = (*arr)[right], (*arr)[pivotPos]

	for j := left; j < right; j++ {
		compCounter++
		if (*arr)[j] <= pivot {
			swapCounter++
			(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
			i++
		}
	}
	swapCounter++
	(*arr)[right], (*arr)[i] = (*arr)[i], (*arr)[right]

	return i, swapCounter, compCounter

}

func quickSort(arr *[]int, end int, begin int) (swpCounter, cpCounter int) {
	var border, swapCounter, compCounter int
	var swpAdd, cpAdd int
	if begin < end {

		pivot, _, _ := MedianOfMedians(arr, begin, end)
		border, swpAdd, cpAdd = partition(arr, begin, end, pivot)
		swapCounter += swpAdd
		compCounter += cpAdd
		swpAdd, cpAdd = quickSort(arr, border, begin)
		swapCounter += swpAdd
		compCounter += cpAdd
		swpAdd, cpAdd = quickSort(arr, end, border+1)
		swapCounter += swpAdd
		compCounter += cpAdd
	}

	return swapCounter, compCounter
}

func main() {

	formula := make([]int, 1000)
	for i := range formula {
		formula[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(formula), func(i, j int) { formula[i], formula[j] = formula[j], formula[i] })

	quickSort(&formula, len(formula)-1, 0)

	fmt.Println(formula)
}
