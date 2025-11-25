package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	opType, err1 := readOpType()
	if err1 != nil {
		panic(err1)
	}

	nums, err2 := readNums()
	if err2 != nil {
		panic(err2)
	}

	switch opType {
	case "avg":
		fmt.Println(average(nums))
	case "sum":
		fmt.Println(sum(nums))
	case "med":
		fmt.Println(median(nums))
	default:
		fmt.Println("Unsupported operation type:", opType)
	}
}

func readOpType() (string, error) {
	fmt.Print("Chose operation type(avg, sum, med): ")
	var operation string
	fmt.Scan(&operation)
	switch operation {
	case "avg", "sum", "med":
		return operation, nil
	default:
		return "", errors.New("Invalid operation")
	}
}

func readNums() ([]int, error) {
	fmt.Println("Write numbers in line, devide with <,>: ")
	var input string
	fmt.Scan(&input)

	input = strings.TrimSpace(input)
	numsStr := strings.Split(input, ",")
	nums := make([]int, len(numsStr))
	var err error
	for i, numStr := range numsStr {
		numStr = strings.TrimSpace(numStr)
		nums[i], err = strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
	}
	return nums, nil
}

func average(nums []int) float64 {
	sum := sum(nums)
	return float64(sum) / float64(len(nums))
}

func sum(nums []int) float64 {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum)
}

func median(nums []int) float64 {
	sort.Ints(nums)
	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	} else {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / 2
	}
}
