package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Select(arr *[]int, p int, q int, i int, partSize int) (answer, swpCounter, cpCounter int) {
	var swapCounter, compCounter int
	var swpAdd, cpAdd int

	if p == q {
		return (*arr)[p], swapCounter, compCounter
	}

	pivot, swpAdd, cpAdd := MedianOfMedians(arr, p, q, partSize)
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
		ans, swpAdd, cpAdd := Select(arr, p, r-1, i, partSize)
		compCounter += cpAdd
		swapCounter += swpAdd
		return ans, swapCounter, compCounter
	} else {
		compCounter += 2
		ans, swpAdd, cpAdd := Select(arr, r+1, q, i-k, partSize)
		compCounter += cpAdd
		swapCounter += swpAdd
		return ans, swapCounter, compCounter
	}

}

func MedianOfMedians(arr *[]int, p int, q int, partSize int) (MedianOfMedians, swpCounter, cpCounter int) {
	var swapCounter, compCounter int
	if q-p < partSize {
		return medianOfSmallArr(arr, p, q)
	}

	medianArray := make([]int, 0)

	for i := p; i <= q; i = i + partSize {
		subRight := i + partSize - 1

		if subRight > q {
			subRight = q
		}
		median, swpAdd, cpAdd := medianOfSmallArr(arr, i, subRight)
		compCounter += cpAdd
		swapCounter += swpAdd
		medianArray = append(medianArray, median)
	}
	arrLen := len(medianArray)
	return Select(&medianArray, 0, arrLen-1, (arrLen+1)/2, partSize)
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

	var N [100]int
	var S [12]int
	p := 100
	for i := 0; i < len(N); i++ {
		N[i] = p
		p += 100
	}
	for i, _ := range S {
		S[i] = 3 + (2 * i)
	}
	m := 100

	var totalComp [100][100][12]int
	var totalSwap [100][100][12]int
	var totalTime [100][100][12]int64
	for i := 0; i < m; i++ {
		for x := 0; x < len(N); x++ {

			formula := make([]int, N[x])
			size := N[x] - 1
			toSort := make([]int, N[x])
			/*			for p := range formula {
							formula[p] = p + 1
						}

						rand.Seed(time.Now().UnixNano())
						rand.Shuffle(len(formula), func(i, j int) { formula[i], formula[j] = formula[j], formula[i] })
			*/

			for l := range formula {
				formula[l] = rand.Int()
			}
			for w, s := range S {
				for index, val := range formula {
					toSort[index] = val
				}
				start := time.Now().UnixNano() / int64(time.Microsecond)
				_, swapCounter, compCounter := Select(&toSort, 0, size, 10, s)
				end := time.Now().UnixNano() / int64(time.Microsecond)
				timeStamp := end - start
				totalComp[i][x][w] = compCounter
				totalSwap[i][x][w] = swapCounter
				totalTime[i][x][w] = timeStamp
			}
		}
	}
	var avgComp [100][12]int64
	var avgSwap [100][12]int64
	var avgTime [100][12]int64

	var stdDevComp [100][12]float64
	var stdDevSwap [100][12]float64
	var stdDevTime [100][12]float64

	var maxSwap [100][12]int
	var minSwap [100][12]int

	var maxComp [100][12]int
	var minComp [100][12]int

	var maxTime [100][12]int64
	var minTime [100][12]int64
	for i := 0; i < 12; i++ {

		for k := 0; k < 100; k++ {
			var sumsComp, sumsSwap, sumsTime int64
			var sumsOfSquareComp, sumsOfSquareSwap, sumsOfSquareTime int64

			var tempMaxComp, tempMaxSwap int
			var tempMaxTime int64
			tempMinComp, tempMinSwap := math.MaxInt32, math.MaxInt32
			tempMinTime := int64(math.MaxInt64)

			for j := 0; j < 100; j++ {
				sumsComp += int64(totalComp[j][k][i])
				sumsSwap += int64(totalSwap[j][k][i])
				sumsTime += totalTime[j][k][i]

				if totalComp[j][k][i] > tempMaxComp {
					tempMaxComp = totalComp[j][k][i]
				}
				if totalComp[j][k][i] < tempMinComp {
					tempMinComp = totalComp[j][k][i]
				}
				if totalSwap[j][k][i] > tempMaxSwap {
					tempMaxSwap = totalSwap[j][k][i]
				}
				if totalSwap[j][k][i] < tempMinSwap {
					tempMinSwap = totalSwap[j][k][i]
				}
				if totalTime[j][k][i] > tempMaxTime {
					tempMaxTime = totalTime[j][k][i]
				}
				if totalTime[j][k][i] < tempMinTime {
					tempMinTime = totalTime[j][k][i]
				}
				sumsOfSquareComp += int64(totalComp[j][k][i]) * int64(totalComp[j][k][i])
				sumsOfSquareSwap += int64(totalSwap[j][k][i]) * int64(totalSwap[j][k][i])
				sumsOfSquareTime += totalTime[j][k][i] * totalTime[j][k][i]
			}

			maxSwap[k][i] = tempMaxSwap
			maxComp[k][i] = tempMaxComp
			maxTime[k][i] = tempMaxTime
			minSwap[k][i] = tempMinSwap
			minComp[k][i] = tempMinComp
			minTime[k][i] = tempMinTime

			avgComp[k][i] = sumsComp / 100
			avgSwap[k][i] = sumsSwap / 100
			avgTime[k][i] = sumsTime / 100

			stdDevComp[k][i] = math.Sqrt(float64((sumsOfSquareComp - (2 * avgComp[k][i] * sumsComp) + int64(m)*(avgComp[k][i]*avgComp[k][i])) / int64(m)))
			stdDevSwap[k][i] = math.Sqrt(float64((sumsOfSquareSwap - (2 * avgSwap[k][i] * sumsSwap) + int64(m)*(avgSwap[k][i]*avgSwap[k][i])) / int64(m)))
			stdDevTime[k][i] = math.Sqrt(float64((sumsOfSquareTime - (2 * avgTime[k][i] * sumsTime) + int64(m)*(avgTime[k][i]*avgTime[k][i])) / int64(m)))
		}
	}

	for i, val := range N {

		output := strconv.Itoa(val) + ";"

		for d, _ := range S {
			output += fmt.Sprintf("%f", stdDevTime[i][d]) + ";"
		}
		output += "\n"
		f, err := os.OpenFile("DataBeg/DataR/rTimeDev.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
			output += fmt.Sprintf("%f", stdDevComp[i][d]) + ";"
		}
		output += "\n"
		f, err := os.OpenFile("DataBeg/DataR/rCompDev.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		f, err := os.OpenFile("DataBeg/DataR/rSwapDev.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
			output += strconv.FormatInt(avgTime[i][d], 10) + ";"
		}
		output += "\n"
		f, err := os.OpenFile("DataBeg/DataR/rTimeAvg.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		f, err := os.OpenFile("DataBeg/DataR/rCompAvg.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		f, err := os.OpenFile("DataBeg/DataR/rSwapAvg.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
			output += strconv.FormatInt(maxTime[i][d], 10) + ";"
		}
		output += "\n"
		f, err := os.OpenFile("DataBeg/DataR/rTimeMax.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		f, err := os.OpenFile("DataBeg/DataR/rCompMax.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
			output += strconv.FormatInt(minTime[i][d], 10) + ";"
		}
		output += "\n"
		f, err := os.OpenFile("DataBeg/DataR/rTimeMin.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

		f, err := os.OpenFile("DataBeg/DataR/rSwapMin.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

		f, err := os.OpenFile("DataBeg/DataR/rSwapMax.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
		f, err := os.OpenFile("DataBeg/DataR/rCompMin.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(output); err != nil {
			log.Println(err)
		}
		f.Close()
	}
}
