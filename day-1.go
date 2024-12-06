package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() (slice []string, err error) {
	file, err := os.Open("input")

	if err != nil {
		return
	}

	defer file.Close()

	r := bufio.NewReader(file)

	var s []string

	// Section 2
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			s = append(s, string(line))
		}
		if err != nil {
			break
		}
	}

	return s, err
}

func diff(a, b int64) int64 {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	input, err := getInput()

	if err != nil {
		panic(err)
	}

	var totalDistance int64
	var left []int64
	var right []int64

	for _, element := range input {

		var numbers = strings.Split(element, "   ")

		a, err := strconv.ParseInt(numbers[0], 10, 64)
		if err != nil {
			panic(err)
		}
		left = append(left, a)
		b, err := strconv.ParseInt(numbers[1], 10, 64)
		if err != nil {
			panic(err)
		}
		right = append(right, b)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	for index, number := range left {
		totalDistance += diff(number, right[index])
	}

	fmt.Printf(string(totalDistance))

}
