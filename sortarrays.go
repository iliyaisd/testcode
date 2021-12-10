package main

import "log"

const (
	MinUint uint = 0                 // binary: all zeroes
	MaxUint      = ^MinUint          // binary: all ones
	MaxInt       = int(MaxUint >> 1) // binary: all ones except high bit
	MinInt       = ^MaxInt           // binary: all zeroes except high bit
)

func main() {
	log.Printf("%#v", sortArrays([]int{1, 2, 3}, []int{2, 4}, []int{1, 5, 7, 8}))
	log.Printf("%#v", sortArrays([]int{}, []int{}))
	log.Printf("%#v", sortArrays([]int{1, 2, 3}, []int{}, []int{1, 44}))
	log.Printf("%#v", sortArrays([]int{1, 2, 3}, []int{-5, 8}, []int{0, 1, 8}))
}

func sortArrays(inputs ...[]int) []int {
	var counters []int
	var totalEl int

	for _, inp := range inputs {
		counters = append(counters, 0)
		totalEl += len(inp)
	}

	var result []int
	var indexesSmallestNum []int

	for {
		var curSmallest = MaxInt
		var countersExhausted int
		var foundSmallest bool
		for n, curInput := range inputs {
			if counters[n] == len(curInput) {
				countersExhausted++
				continue
			}
			if curSmallest > curInput[counters[n]] {
				curSmallest = curInput[counters[n]]
				indexesSmallestNum = []int{n}
				foundSmallest = true
				continue
			}
			if curSmallest == curInput[counters[n]] {
				indexesSmallestNum = append(indexesSmallestNum, n)
			}
		}
		if !foundSmallest {
			break
		}
		for _, index := range indexesSmallestNum {
			result = append(result, inputs[index][counters[index]])
			if counters[index] < len(inputs[index]) {
				counters[index]++
			}
		}
		if countersExhausted == len(inputs) {
			break
		}
	}

	return result
}
