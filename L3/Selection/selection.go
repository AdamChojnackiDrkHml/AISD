package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
	//	"os"
)

func RandomSelect(arr *[]int, p int, q int, i int) (answer, swpCounter, cpCounter int) {
	var swapCounter, compCounter int
	var swpAdd, cpAdd int
	compCounter++
	if p == q {
		return (*arr)[p], swapCounter, compCounter
	}

	r, swpAdd, cpAdd := randomPartition(arr, p, q)
	swapCounter += swpAdd
	compCounter += cpAdd

	k := r - p + 1

	compCounter++
	if i == k {
		return (*arr)[r], swapCounter, compCounter
	}

	compCounter++
	if i < k {
		ans, cpAdd, swpAdd := RandomSelect(arr, p, r-1, i)
		compCounter += cpAdd
		swapCounter += swpAdd
		return ans, swapCounter, compCounter
	}

	ans, cpAdd, swpAdd := RandomSelect(arr, r+1, q, i-k)
	compCounter += cpAdd
	swapCounter += swpAdd
	return ans, swapCounter, compCounter
}

func randomPartition(arr *[]int, left int, right int) (border, swpCounter, cpCounter int) {
	pivotPos := rand.Intn(right-left) + left
	i := left
	pivot := (*arr)[pivotPos]
	var swapCounter, compCounter int

	(*arr)[pivotPos], (*arr)[right] = (*arr)[right], (*arr)[pivotPos]
	swapCounter++

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

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		var N [100]int
		p := 100
		for i := 0; i < len(N); i++ {
			N[i] = p
			p += 100
		}
		m := 100
		var S [6]int
		var totalComp [100][100][6]int
		var totalSwap [100][100][6]int
		for i := 0; i < m; i++ {
			for x := 0; x < len(N); x++ {

				formula := make([]int, N[x])
				/*
					for i := range formula {
						formula[i] = rand.Int()
					}
				*/
				for i := range formula {
					formula[i] = i + 1
				}
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(formula), func(i, j int) { formula[i], formula[j] = formula[j], formula[i] })
				size := N[x] - 1

				toSort := make([]int, N[x])

				//
				//
				//RandomSelect K 30
				for index, val := range formula {
					toSort[index] = val
				}
				_, swapCounter, compCounter := RandomSelect(&toSort, 0, size, 30)
				totalComp[i][x][0] = compCounter
				totalSwap[i][x][0] = swapCounter

				//
				//
				//RandomSelect K N/2
				for index, val := range formula {
					toSort[index] = val
				}
				_, swapCounter, compCounter = RandomSelect(&toSort, 0, size, N[x]/2)
				totalComp[i][x][1] = compCounter
				totalSwap[i][x][1] = swapCounter

				//
				//
				//RandomSelect K N - 50
				for index, val := range formula {
					toSort[index] = val
				}
				_, swapCounter, compCounter = RandomSelect(&toSort, 0, size, N[x]-50)
				totalComp[i][x][2] = compCounter
				totalSwap[i][x][2] = swapCounter

				//
				//
				//Select
				for index, val := range formula {
					toSort[index] = val
				}
				_, swapCounter, compCounter = Select(&toSort, 0, size, 30)
				totalComp[i][x][3] = compCounter
				totalSwap[i][x][3] = swapCounter

				//
				//
				//Select K N/2
				for index, val := range formula {
					toSort[index] = val
				}
				_, swapCounter, compCounter = Select(&toSort, 0, size, N[x]/2)
				totalComp[i][x][4] = compCounter
				totalSwap[i][x][4] = swapCounter

				//
				//
				//Select K N - 50
				for index, val := range formula {
					toSort[index] = val
				}
				_, swapCounter, compCounter = Select(&toSort, 0, size, N[x]-50)
				totalComp[i][x][5] = compCounter
				totalSwap[i][x][5] = swapCounter

			}

		}

		var avgComp [100][6]int64
		var avgSwap [100][6]int64

		var stdDevComp [100][6]float64
		var stdDevSwap [100][6]float64

		var maxSwap [100][6]int
		var minSwap [100][6]int

		var maxComp [100][6]int
		var minComp [100][6]int
		for i := 0; i < 6; i++ {

			for k := 0; k < 100; k++ {
				var sumsComp, sumsSwap int64
				var sumsOfSquareComp, sumsOfSquareSwap int64

				var tempMaxComp, tempMaxSwap int

				tempMinComp, tempMinSwap := math.MaxInt32, math.MaxInt32

				for j := 0; j < 100; j++ {
					sumsComp += int64(totalComp[j][k][i])
					sumsSwap += int64(totalSwap[j][k][i])

					if totalComp[j][k][i] > tempMaxComp {
						tempMaxComp = totalComp[j][k][i]
					}
					if totalComp[j][k][i] < tempMinComp {
						tempMinComp = totalComp[j][k][i]
					}
					if totalSwap[j][k][i] > tempMaxSwap {
						tempMaxSwap = totalSwap[j][k][i]
					}
					if totalSwap[j][k][i] < tempMinComp {
						tempMinSwap = totalSwap[j][k][i]
					}
					sumsOfSquareComp += int64(totalComp[j][k][i]) * int64(totalComp[j][k][i])
					sumsOfSquareSwap += int64(totalSwap[j][k][i]) * int64(totalSwap[j][k][i])
				}

				maxSwap[k][i] = tempMaxSwap
				maxComp[k][i] = tempMaxComp
				minSwap[k][i] = tempMinSwap
				minComp[k][i] = tempMinComp
				avgComp[k][i] = sumsComp / 100
				avgSwap[k][i] = sumsSwap / 100

				stdDevComp[k][i] = math.Sqrt(float64((sumsOfSquareComp - (2 * avgComp[k][i] * sumsComp) + 100*(avgComp[k][i]*avgComp[k][i])) / int64(m)))
				stdDevSwap[k][i] = math.Sqrt(float64((sumsOfSquareSwap - (2 * avgSwap[k][i] * sumsSwap) + 100*(avgSwap[k][i]*avgSwap[k][i])) / int64(m)))
			}
		}

		for i, val := range N {

			output := strconv.Itoa(val) + ";"

			for d, _ := range S {
				output += fmt.Sprintf("%f", stdDevComp[i][d]) + ";"
			}
			output += "\n"
			f, err := os.OpenFile("Data/DataP/pCompDev.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				log.Println(err)
			}
			f.Close()
		}

		for i, val := range N {

			output := strconv.Itoa(val) + ";"

			for d, _ := range S {
				output += fmt.Sprintf("%f", stdDevSwap[i][d]) + ";"
			}
			output += "\n"
			f, err := os.OpenFile("Data/DataP/pSwapDev.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				log.Println(err)
			}
			f.Close()
		}

		for i, val := range N {

			output := strconv.Itoa(val) + ";"

			for d, _ := range S {
				output += strconv.FormatInt(avgComp[i][d], 10) + ";"
			}
			output += "\n"
			f, err := os.OpenFile("Data/DataP/pCompAvg.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				log.Println(err)
			}
			f.Close()
		}

		for i, val := range N {

			output := strconv.Itoa(val) + ";"

			for d, _ := range S {
				output += strconv.FormatInt(avgSwap[i][d], 10) + ";"
			}
			output += "\n"
			f, err := os.OpenFile("Data/DataP/pSwapAvg.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				log.Println(err)
			}
			f.Close()
		}

		for i, val := range N {

			output := strconv.Itoa(val) + ";"

			for d, _ := range S {
				output += strconv.Itoa(maxComp[i][d]) + ";"
			}
			output += "\n"
			f, err := os.OpenFile("Data/DataP/pCompMax.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				log.Println(err)
			}
			f.Close()
		}

		for i, val := range N {

			output := strconv.Itoa(val) + ";"

			for d, _ := range S {
				output += strconv.Itoa(minSwap[i][d]) + ";"
			}
			output += "\n"

			f, err := os.OpenFile("Data/DataP/pSwapMin.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				log.Println(err)
			}
			f.Close()
		}
		for i, val := range N {

			output := strconv.Itoa(val) + ";"

			for d, _ := range S {
				output += strconv.Itoa(maxSwap[i][d]) + ";"
			}
			output += "\n"

			f, err := os.OpenFile("Data/DataP/pSwapMax.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(output); err != nil {
				log.Println(err)
			}
			f.Close()
		}
		for i, val := range N {

			output := strconv.Itoa(val) + ";"

			for d, _ := range S {
				output += strconv.Itoa(minComp[i][d]) + ";"
			}
			output += "\n"
			f, err := os.OpenFile("Data/DataP/pCompMin.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		var N, K int
		//N = 11
		//K = 8
		fmt.Scanf("%d", &N)
		fmt.Scanf("%d", &K)

		array := make([]int, N)

		switch argsWithoutProg[0] {
		case "-p":
			{
				for i := range array {
					array[i] = i + 1
				}
				rand.Seed(time.Now().UnixNano())
				rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
			}
		case "-r":
			{
				for i := range array {
					array[i] = rand.Intn(1000)
				}
			}
		}

		//	array := make([]int, 6)
		//	array = []int{3, 23, 12, 1, 6, 4}

		w, _, _ := Select(&array, 0, N-1, K)

		//fmt.Println(array)

		fmt.Println(array)
		fmt.Println(strconv.Itoa(K) + "   " + strconv.Itoa(w))

	}
}
