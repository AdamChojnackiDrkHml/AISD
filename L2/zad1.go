package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func insertSortAsc(array *[]int, N int) (swpCounter, cmpCounter int) {

	var swapCounter, compCounter int
	for i := 1; i < N; i++ {
		key := (*array)[i]
		j := i - 1

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

	}
	return swapCounter, compCounter
}

func insertSortDes(array *[]int, N int) (swpCounter, cmpCounter int) {
	var swapCounter, compCounter int
	for i := 1; i < N; i++ {
		key := (*array)[i]
		j := i - 1

		for j >= 0 && (*array)[j] < key {
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

	}

	return swapCounter, compCounter
}

func quickSort(arr *[]int, end int, begin int, isAsc bool) (swpCounter, cpCounter int) {
	var border, swapCounter, compCounter int
	var swpAdd, cpAdd int
	if begin < end {
		if isAsc {
			border, swpAdd, cpAdd = partitionAsc(arr, end, begin)
		} else {
			border, swpAdd, cpAdd = partitionDes(arr, end, begin)
		}
		swapCounter += swpAdd
		compCounter += cpAdd
		swpAdd, cpAdd = quickSort(arr, border, begin, isAsc)
		swapCounter += swpAdd
		compCounter += cpAdd
		swpAdd, cpAdd = quickSort(arr, end, border+1, isAsc)
		swapCounter += swpAdd
		compCounter += cpAdd
	}

	return swapCounter, compCounter
}

func partitionDes(arr *[]int, end int, begin int) (border, swpCounter, cpCounter int) {
	pivot := (*arr)[(end+begin)/2]
	var swapCounter, compCounter int
	i, j := begin, end

	for {

		for (*arr)[j] < pivot {
			fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
			compCounter++
			j--
		}
		fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
		compCounter++

		for (*arr)[i] > pivot {
			fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
			compCounter++
			i++
		}
		fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
		compCounter++

		if i < j {
			fmt.Fprintln(os.Stderr, "swap", (*arr)[i], "to ", (*arr)[j])
			swapCounter++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
			j--
		} else if i >= j {
			break
		}
	}
	return j, swapCounter, compCounter
}

func partitionAsc(arr *[]int, end int, begin int) (border, swpCounter, cpCounter int) {
	pivot := (*arr)[(end+begin)/2]
	var swapCounter, compCounter int
	i, j := begin, end

	for {
		for (*arr)[j] > pivot {
			fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
			compCounter++
			j--
		}
		fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
		compCounter++

		for (*arr)[i] < pivot {
			fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
			compCounter++
			i++
		}
		fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
		compCounter++

		if i < j {
			fmt.Fprintln(os.Stderr, "swap", (*arr)[i], "to ", (*arr)[j])
			swapCounter++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
			j--
		} else if i >= j {
			break
		}
	}
	return j, swapCounter, compCounter
}

func mergeSort(arr []int, isAsc bool) (sorted []int, swpCounter, cpCounter int) {
	length := len(arr)
	var swapCounter, compCounter int

	if length > 1 {

		middle := length / 2

		var swpAdd, cpAdd int

		arr1, swpAdd, cpAdd := mergeSort(arr[:middle], isAsc)
		swapCounter += swpAdd
		compCounter += cpAdd

		arr2, swpAdd, cpAdd := mergeSort(arr[middle:], isAsc)
		swapCounter += swpAdd
		compCounter += cpAdd

		if isAsc {
			arr, swpAdd, cpAdd = mergeAsc(arr1, arr2)
		}
		if !isAsc {
			arr, swpAdd, cpAdd = mergeDes(arr1, arr2)
		}
		swapCounter += swpAdd
		compCounter += cpAdd

	}

	return arr, swapCounter, compCounter
}

func mergeAsc(arr1 []int, arr2 []int) (arr []int, swpCounter, cpCounter int) {

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

func mergeDes(arr1 []int, arr2 []int) (arr []int, swpCounter, cpCounter int) {

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

		if arr1[i] >= arr2[j] {
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

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 4 {
		var N [101]int
		p := 100
		for i := 0; i < len(N); i++ {
			N[i] = p
			p += 100
		}
		k, _ := strconv.Atoi(argsWithoutProg[6])

		for i := 0; i < k; i++ {
			for _, n := range N {
				formula := make([]int, n)
				for j := 0; j < n; j++ {
					formula[j] = rand.Intn(n)
				}

				//InsertionSort
				toSort := formula
				start := time.Now().UnixNano() / int64(time.Millisecond)
				swapCounter, compCounter := insertSortAsc(&toSort, n)
				end := time.Now().UnixNano() / int64(time.Millisecond)
				timeStamp := end - start
				output := "insert;" + strconv.Itoa(n) + ";" + strconv.Itoa(compCounter) + ";" + strconv.Itoa(swapCounter) + ";" + strconv.FormatInt(timeStamp, 10) + "\n"
				f, err := os.OpenFile(argsWithoutProg[5], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Println(err)
				}
				defer f.Close()
				if _, err := f.WriteString(output); err != nil {
					log.Println(err)
				}

				//MergeSort
				toSort = formula
				start = time.Now().UnixNano() / int64(time.Millisecond)
				_, swapCounter, compCounter = mergeSort(toSort, true)
				end = time.Now().UnixNano() / int64(time.Millisecond)
				timeStamp = end - start
				output = "merge;" + strconv.Itoa(n) + ";" + strconv.Itoa(compCounter) + ";" + strconv.Itoa(swapCounter) + ";" + strconv.FormatInt(timeStamp, 10) + "\n"
				f, err = os.OpenFile(argsWithoutProg[5], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Println(err)
				}
				defer f.Close()
				if _, err := f.WriteString(output); err != nil {
					log.Println(err)
				}

				//QuickSort
				toSort = formula
				start = time.Now().UnixNano() / int64(time.Millisecond)
				swapCounter, compCounter = quickSort(&toSort, n-1, 0, true)
				end = time.Now().UnixNano() / int64(time.Millisecond)
				timeStamp = end - start
				output = "quick;" + strconv.Itoa(n) + ";" + strconv.Itoa(compCounter) + ";" + strconv.Itoa(swapCounter) + ";" + strconv.FormatInt(timeStamp, 10) + "\n"
				f, err = os.OpenFile(argsWithoutProg[5], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					log.Println(err)
				}
				defer f.Close()
				if _, err := f.WriteString(output); err != nil {
					log.Println(err)
				}
			}
		}
	} else {
		var decide int
		switch argsWithoutProg[1] {
		case "insert":
			{
				decide += 1
				break
			}
		case "quick":
			{
				decide += 2
				break
			}
		case "merge":
			{
				decide += 3
				break
			}
		}

		switch argsWithoutProg[3] {
		case ">=":
			{
				decide += 0
				break
			}
		case "<=":
			{
				decide += 3
				break
			}
		}

		var N int
		fmt.Scanf("%d", &N)
		var toSort = make([]int, N)

		for i, _ := range toSort {
			fmt.Scanf("%d", &(toSort[i]))
		}

		var swapCounter, compCounter int
		start := time.Now()
		switch decide {
		case 1:
			{
				swapCounter, compCounter = insertSortAsc(&toSort, N)
				break
			}
		case 2:
			{
				swapCounter, compCounter = quickSort(&toSort, N-1, 0, true)
				break
			}
		case 3:
			{
				toSort, swapCounter, compCounter = mergeSort(toSort, true)
				break
			}
		case 4:
			{
				swapCounter, compCounter = insertSortDes(&toSort, N)
				break
			}
		case 5:
			{
				swapCounter, compCounter = quickSort(&toSort, N-1, 0, false)
				break
			}
		case 6:
			{
				toSort, swapCounter, compCounter = mergeSort(toSort, false)
				break
			}

		}

		end := time.Now()

		time := end.Sub(start)
		//time += time
		fmt.Fprintln(os.Stderr, "Total Comprasions: ", compCounter)
		fmt.Fprintln(os.Stderr, "Total Swaps: ", swapCounter)

		fmt.Println(time)
		/*for i := range toSort {
			fmt.Println(toSort[i])
		} */
	}
}
