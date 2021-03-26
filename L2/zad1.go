package main

import (
	"fmt"
	"os"
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

func main() {

	xd := []int{6, 75, 24, 86, 234, 7532, 78, 3}

	insertSort(&xd, len(xd))

	for i := range xd {
		fmt.Printf("%d ", xd[i])
	}
}
