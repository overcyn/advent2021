package main

import (
	"fmt"
	"os"
	"math"
	"strings"
	"strconv"
)

func main() {
	rlt, err := a()
	if err != nil {
		panic(err)
	}
	fmt.Println("Result A:", rlt)
	
	rlt, err = b()
	if err != nil {
		panic(err)
	}
	fmt.Println("Result B:", rlt)
}

func a() (int, error) {
	// Read file
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	
	// Parse input
	strs := strings.Split(string(dat), "\n")
	ints := []int{}
	for _, i := range strs {
		val, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}
		ints = append(ints, val)
	}
	
	// Calculate
	count := 0
	prev := math.MaxInt64
	for _, i := range ints {
		if i > prev {
			count += 1
		}
		prev = i
	}
    return count, nil
}

func b() (int, error) {
	// Read file
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	
	// Parse input
	strs := strings.Split(string(dat), "\n")
	ints := []int{}
	for _, i := range strs {
		val, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}
		ints = append(ints, val)
	}
	
	// Calculate
	count := 0
	for i := 0; i < len(ints)-3; i++ {
		if ints[i] < ints[i+3] {
			count += 1
		}
	}
    return count, nil
}