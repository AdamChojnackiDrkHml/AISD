package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func hybridSort(arr *[]int, end int, begin int, isAsc bool) (swpCounter, cpCounter int) {
	N := end - begin + 1
	if N < 40 {
		return insertSortAsc(arr, N, begin)
	} else {
		var border, swapCounter, compCounter int
		var swpAdd, cpAdd int

		border, swpAdd, cpAdd = partitionAsc(arr, end, begin)

		swapCounter += swpAdd
		compCounter += cpAdd
		swpAdd, cpAdd = hybridSort(arr, border, begin, isAsc)
		swapCounter += swpAdd
		compCounter += cpAdd
		swpAdd, cpAdd = hybridSort(arr, end, border+1, isAsc)
		swapCounter += swpAdd
		compCounter += cpAdd
		return swapCounter, compCounter
	}

}

func DPQS(array *[]int, end int, begin int) (swpCounter, cpCounter int) {

	if end <= begin {
		return 0, 0
	}
	var p, q int

	var swapCounter, compCounter int
	var swpAdd, cpAdd int

	if (*array)[begin] > (*array)[end] {
		p, q = (*array)[end], (*array)[begin]
	} else {
		q, p = (*array)[end], (*array)[begin]
	}
	//	fmt.Fprintln(os.Stderr, "comp", (*array)[j], "to ", key)
	compCounter++

	i, k, d := begin+1, end-1, 0
	j := i

	for j <= k {

		if d >= 0 {
			//	fmt.Fprintln(os.Stderr, "comp", (*array)[j], "to ", p)
			compCounter++
			if (*array)[j] < p {
				(*array)[j], (*array)[i] = (*array)[i], (*array)[j]
				//	fmt.Fprintln(os.Stderr, "swap", (*array)[j], "to ", (*array)[i])
				swapCounter++
				i++
				j++
				d++
			} else if (*array)[j] < q {
				//	fmt.Fprintln(os.Stderr, "comp", (*array)[j], "to ", q)
				compCounter++
				j++
			} else {
				//	fmt.Fprintln(os.Stderr, "comp", (*array)[j], "to ", q)
				compCounter++
				(*array)[j], (*array)[k] = (*array)[k], (*array)[j]
				//	fmt.Fprintln(os.Stderr, "swap", (*array)[j], "to ", (*array)[k])
				swapCounter++
				k--
				d--
			}

		} else {
			//	fmt.Fprintln(os.Stderr, "comp", (*array)[k], "to ", q)
			compCounter++
			if (*array)[k] > q {
				k--
				d--
			} else if (*array)[k] < p {
				//	fmt.Fprintln(os.Stderr, "comp", (*array)[k], "to ", p)
				compCounter++
				(*array)[k], (*array)[j], (*array)[i] = (*array)[j], (*array)[i], (*array)[k]
				//	fmt.Fprintln(os.Stderr, "swap", (*array)[j], "to ", (*array)[k])
				//	fmt.Fprintln(os.Stderr, "swap", (*array)[k], "to ", (*array)[i])
				swapCounter++
				swapCounter++
				i++
				d++
				j++
			} else {
				//	fmt.Fprintln(os.Stderr, "comp", (*array)[k], "to ", p)
				compCounter++
				(*array)[j], (*array)[k] = (*array)[k], (*array)[j]
				//	fmt.Fprintln(os.Stderr, "swap", (*array)[j], "to ", (*array)[k])
				swapCounter++
				j++
			}

		}
	}
	(*array)[begin], (*array)[i-1] = (*array)[i-1], p
	(*array)[end], (*array)[k+1] = (*array)[k+1], q
	//	fmt.Fprintln(os.Stderr, "swap", (*array)[begin], "to ", (*array)[i-1])
	//	fmt.Fprintln(os.Stderr, "swap", (*array)[i-1], "to ", p
	swapCounter += 2
	//	fmt.Fprintln(os.Stderr, "swap", (*array)[end], "to ", (*array)[k+1])
	//	fmt.Fprintln(os.Stderr, "swap", (*array)[k+1], "to ", q

	swapCounter += 2

	swpAdd, cpAdd = DPQS(array, i-2, begin)
	swapCounter += swpAdd
	compCounter += cpAdd
	swpAdd, cpAdd = DPQS(array, k, i)
	swapCounter += swpAdd
	compCounter += cpAdd
	swpAdd, cpAdd = DPQS(array, end, k+2)
	swapCounter += swpAdd
	compCounter += cpAdd

	return swapCounter, compCounter

}

func insertSortAsc(array *[]int, N int, start int) (swpCounter, cmpCounter int) {

	var swapCounter, compCounter int
	for i := 1 + start; i < N+start; i++ {
		key := (*array)[i]
		j := i - 1

		for j >= 0 && (*array)[j] > key {
			//	fmt.Fprintln(os.Stderr, "comp", (*array)[j], "to ", key)
			compCounter++
			(*array)[j+1] = (*array)[j]
			//	fmt.Fprintln(os.Stderr, "swap", (*array)[j], "to ", j+1)
			swapCounter++
			j--
		}
		//fmt.Fprintln(os.Stderr, "comp", (*array)[j+1], "to ", key)
		compCounter++
		//fmt.Fprintln(os.Stderr, "swap", key, "to ", j+1)
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
			//	fmt.Fprintln(os.Stderr, "comp", (*array)[j], "to ", key)
			compCounter++
			(*array)[j+1] = (*array)[j]
			//	fmt.Fprintln(os.Stderr, "swap", (*array)[j], "to ", j+1)
			swapCounter++
			j--
		}
		//	fmt.Fprintln(os.Stderr, "comp", (*array)[j+1], "to ", key)
		compCounter++
		//	fmt.Fprintln(os.Stderr, "swap", key, "to ", j+1)
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
			//		fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
			compCounter++
			j--
		}
		//	fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
		compCounter++

		for (*arr)[i] > pivot {
			//		fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
			compCounter++
			i++
		}
		//	fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
		compCounter++

		if i < j {
			//		fmt.Fprintln(os.Stderr, "swap", (*arr)[i], "to ", (*arr)[j])
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
			//		fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
			compCounter++
			j--
		}
		//	fmt.Fprintln(os.Stderr, "comp", (*arr)[j], "to ", pivot)
		compCounter++

		for (*arr)[i] < pivot {
			//		fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
			compCounter++
			i++
		}
		//	fmt.Fprintln(os.Stderr, "comp", (*arr)[i], "to ", pivot)
		compCounter++

		if i < j {
			//		fmt.Fprintln(os.Stderr, "swap", (*arr)[i], "to ", (*arr)[j])
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
		//	fmt.Fprintln(os.Stderr, "comp", arr1[i], "to ", arr2[j])
		compCounter++

		if arr1[i] <= arr2[j] {
			//		fmt.Fprintln(os.Stderr, "swap", arr1[i], "to ", k, "place")
			swapCounter++
			array[k] = arr1[i]
			i++
			continue
		}

		//	fmt.Fprintln(os.Stderr, "swap", arr2[j], "to ", k, "place")
		swapCounter++
		array[k] = arr2[j]
		j++
	}

	k++

	for i < length1 {
		//	fmt.Fprintln(os.Stderr, "swap", arr1[i], "to ", k, "place")
		swapCounter++
		array[k] = arr1[i]
		i++
		k++
	}

	for j < length2 {
		//	fmt.Fprintln(os.Stderr, "swap", arr2[j], "to ", k, "place")
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
		//	fmt.Fprintln(os.Stderr, "comp", arr1[i], "to ", arr2[j])
		compCounter++

		if arr1[i] >= arr2[j] {
			//		fmt.Fprintln(os.Stderr, "swap", arr1[i], "to ", k, "place")
			swapCounter++
			array[k] = arr1[i]
			i++
			continue
		}

		//	fmt.Fprintln(os.Stderr, "swap", arr2[j], "to ", k, "place")
		swapCounter++
		array[k] = arr2[j]
		j++
	}

	k++

	for i < length1 {
		//	fmt.Fprintln(os.Stderr, "swap", arr1[i], "to ", k, "place")
		swapCounter++
		array[k] = arr1[i]
		i++
		k++
	}

	for j < length2 {
		//	fmt.Fprintln(os.Stderr, "swap", arr2[j], "to ", k, "place")
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
		var N [100]int
		p := 100
		for i := 0; i < len(N); i++ {
			N[i] = p
			p += 100
		}
		k, _ := strconv.Atoi(argsWithoutProg[6])
		var meansComp [100][5]float32
		var meansSwap [100][5]float32
		var meansTime [100][5]int64
		for i := 0; i < k; i++ {
			for x := 0; x < len(N); x++ {
				formula := make([]int, N[x])
				for j := 0; j < N[x]; j++ {
					formula[j] = rand.Intn(N[x])
				}
				size := N[x]
				size2 := size - 1
				toSort := make([]int, size)

				//DPQS
				for index, val := range formula {
					toSort[index] = val
				}
				start := time.Now().UnixNano() / int64(time.Microsecond)
				swapCounter, compCounter := DPQS(&toSort, size2, 0)
				end := time.Now().UnixNano() / int64(time.Microsecond)
				timeStamp := end - start
				meansTime[x][0] += timeStamp
				meansComp[x][0] += float32(compCounter)
				meansSwap[x][0] += float32(swapCounter)

				//InsertionSort
				for index, val := range formula {
					toSort[index] = val
				}
				start = time.Now().UnixNano() / int64(time.Microsecond)
				swapCounter, compCounter = insertSortAsc(&toSort, size, 0)
				end = time.Now().UnixNano() / int64(time.Microsecond)
				timeStamp = end - start
				meansTime[x][1] += timeStamp
				meansComp[x][1] += float32(compCounter)
				meansSwap[x][1] += float32(swapCounter)

				//DPQS

				//MergeSort
				for index, val := range formula {
					toSort[index] = val
				}
				start = time.Now().UnixNano() / int64(time.Microsecond)
				toSort, swapCounter, compCounter = mergeSort(toSort, true)
				end = time.Now().UnixNano() / int64(time.Microsecond)
				timeStamp = end - start
				meansTime[x][2] += timeStamp
				meansComp[x][2] += float32(compCounter)
				meansSwap[x][2] += float32(swapCounter)

				//QuickSort
				for index, val := range formula {
					toSort[index] = val
				}
				start = time.Now().UnixNano() / int64(time.Microsecond)
				swapCounter, compCounter = quickSort(&toSort, size2, 0, true)
				end = time.Now().UnixNano() / int64(time.Microsecond)
				timeStamp = end - start
				meansTime[x][3] += timeStamp
				meansComp[x][3] += float32(compCounter)
				meansSwap[x][3] += float32(swapCounter)

				//Hybrid
				for index, val := range formula {
					toSort[index] = val
				}
				start = time.Now().UnixNano() / int64(time.Microsecond)
				swapCounter, compCounter = hybridSort(&toSort, size2, 0, true)
				end = time.Now().UnixNano() / int64(time.Microsecond)
				timeStamp = end - start
				meansTime[x][4] += timeStamp
				meansComp[x][4] += float32(compCounter)
				meansSwap[x][4] += float32(swapCounter)

			}

			f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(fmt.Sprintf("%d\n", i)); err != nil {
				log.Println(err)
			}
			f.Close()
		}

		for d := 0; d < 5; d++ {

			for i, _ := range meansComp {
				meansComp[i][d] /= float32(k)
				meansSwap[i][d] /= float32(k)
				meansTime[i][d] /= int64(k)
			}
		}

		for i, n := range N {
			output := fmt.Sprintf("%d", n) + ";"
			for d := 0; d < 5; d++ {
				c := meansComp[i][d]
				s := meansSwap[i][d]
				output += fmt.Sprintf("%f", c) + ";" + fmt.Sprintf("%f", s) + ";" + strconv.FormatInt(meansTime[i][d], 10) + ";" + fmt.Sprintf("%f", c/float32(n)) + ";" + fmt.Sprintf("%f", s/float32(n)) + ";"
			}
			output += "\n"
			f, err := os.OpenFile(argsWithoutProg[5], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				log.Println(err)
			}
			f.Close()
		}

	} else {
		var decide int
		comp := true
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
		case "dp":
			{
				decide += 4
				break
			}
		case "hb":
			{
				decide += 5
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
				decide += 5
				comp = false
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
		start := time.Now().UnixNano() / int64(time.Microsecond)
		switch decide {
		case 1:
			{
				swapCounter, compCounter = insertSortAsc(&toSort, N, 0)
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
		case 5:
			{
				swapCounter, compCounter = hybridSort(&toSort, N-1, 0, true)
			}
		case 4:
			{
				swapCounter, compCounter = DPQS(&toSort, N-1, 0)
			}
		case 6:
			{
				swapCounter, compCounter = insertSortDes(&toSort, N)
				break
			}
		case 7:
			{
				swapCounter, compCounter = quickSort(&toSort, N-1, 0, false)
				break
			}
		case 8:
			{
				toSort, swapCounter, compCounter = mergeSort(toSort, false)
				break
			}
		case 9:
			{
				swapCounter, compCounter = hybridSort(&toSort, N-1, 0, true)
			}
		case 10:
			{
				swapCounter, compCounter = DPQS(&toSort, N-1, 0)
			}

		}

		end := time.Now().UnixNano() / int64(time.Microsecond)

		timeStamp := end - start
		fmt.Fprintln(os.Stderr, "Total Comprasions: ", compCounter)
		fmt.Fprintln(os.Stderr, "Total Swaps: ", swapCounter)

		fmt.Fprintln(os.Stderr, "Total Time: ", timeStamp)
		timeStamp += timeStamp
		if comp {
			for i := 1; i < len(toSort); i++ {
				if toSort[i] < toSort[i-1] {
					fmt.Printf("Błąd")
					return
				}
			}
		} else {
			for i := 1; i < len(toSort); i++ {
				if toSort[i] > toSort[i-1] {
					fmt.Printf("Błąd")
					return
				}
			}
		}
		for _, n := range toSort {
			fmt.Printf("%d ", n)
		}
	}

}
