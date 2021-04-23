package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func binarySearch(arr *[]int, p int, q int, key int) (int, int) {
	var compCounter, cpAdd int

	if q-p < 0 {
		return 0, compCounter
	}
	pivot := (*arr)[(p+q)/2]
	compCounter++
	if pivot == key {
		return 1, compCounter
	}
	if pivot > key {
		compCounter++
		ans, cpAdd := binarySearch(arr, p, ((p+q)/2)-1, key)
		compCounter += cpAdd
		return ans, compCounter
	}
	compCounter++
	ans, cpAdd := binarySearch(arr, ((p+q)/2)+1, q, key)
	compCounter += cpAdd
	return ans, compCounter

}
func main() {
	const nsize = 1000
	const m = 100
	const expCt = 1
	var N [nsize]int

	p := 100
	for i := 0; i < len(N); i++ {
		N[i] = p
		p += 100
	}

	var totalComp [m][nsize][expCt]int
	var totalTime [m][nsize][expCt]int64
	sum := 0
	S := make([]int, expCt)
	for i := 0; i < m; i++ {

		for x := 0; x < len(N); x++ {

			toSearch := make([]int, N[x])
			size := N[x] - 1
			for p := range toSearch {
				toSearch[p] = (p + 1) * 2
			}
			S = []int{toSearch[rand.Intn(N[x])%N[x]]}
			//S = []int{toSearch[N[x]/100], toSearch[N[x]-(N[x]/100)], toSearch[(N[x]/100)*99], toSearch[1+((N[x]/100)*80)]}
			for w, s := range S {
				start := time.Now().UnixNano()
				z, compCounter := binarySearch(&toSearch, 0, size, s)
				end := time.Now().UnixNano()
				timeStamp := end - start
				totalComp[i][x][w] = compCounter
				totalTime[i][x][w] = timeStamp
				sum += z
			}
		}
	}
	var avgComp [nsize][expCt]float64
	var avgTime [nsize][expCt]float64
	for i := 0; i < len(S); i++ {

		for k := 0; k < nsize; k++ {
			var sumsComp, sumsTime int64

			for j := 0; j < m; j++ {
				sumsComp += int64(totalComp[j][k][i])
				sumsTime += totalTime[j][k][i]

			}

			avgComp[k][i] = float64(sumsComp) / float64(m)
			avgTime[k][i] = float64(sumsTime) / float64(m)
		}
	}
	for i, val := range N {

		output := strconv.Itoa(val) + ";"

		for d, _ := range S {
			output += fmt.Sprintf("%f", avgComp[i][d]) + ";"
		}
		output += "\n"
		f, err := os.OpenFile("Data/DataR/rAvgComp.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
			output += fmt.Sprintf("%f", avgTime[i][d]) + ";"
		}
		output += "\n"
		f, err := os.OpenFile("Data/DataR/rAvgTime.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
