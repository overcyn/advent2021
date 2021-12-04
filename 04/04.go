package main

import (
	"fmt"
	"os"
	// "math"
	"strings"
	"strconv"
)

type Board [][]int

func contains(a []int, b int) bool {
	for _, i := range a {
		if i == b {
			return true
		}
	}
	return false
}

func (b Board)Score(nums []int) int {
	score := 0
	for _, i := range b {
		for _, j := range i {
			if !contains(nums, j) {
				score += j
			}
		}
	}
	return score * nums[len(nums)-1]
}

func (b Board)IsWinner(nums []int) bool {
	// Check all rows
	for _, i := range b {
		matches := true
		for _, j := range i {
			if !contains(nums, j) {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}
	
	// Check all columns
	for i, _ := range b[0] {
		column := []int{}
		for _, j := range b {
			column = append(column, j[i])
		}
		
		matches := true
		for _, j := range column {
			if !contains(nums, j) {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}	
	return false
}

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
	strs := strings.Split(string(dat), "\n\n")
	nums := []int{}
	for _, i := range strings.Split(strs[0], ",") {
		val, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}
		nums = append(nums, val)
	}
	
	boards := []Board{}
	for _, i := range strs[1:] {
		board := Board{}
		for _, j := range strings.Split(i, "\n") {
			row := []int{}
			for _, k := range strings.Split(j, " ") {
				if len(k) == 0 {
					continue
				}
				val, err := strconv.Atoi(k)
				if err != nil {
					return 0, err
				}
				row = append(row, val)
			}
			board = append(board, row)
		}
		boards = append(boards, board)
	}
	
	// Calculate
	score := -1
	for i, _ := range nums {
		for _, j := range boards {		
			if j.IsWinner(nums[:i+1]) {
				score = j.Score(nums[:i+1])
				break
			}
		}
		if score != -1 {
			break
		}
	}
	return score, nil
}

func b() (int, error) {
	// Read file
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	
	// Parse input
	strs := strings.Split(string(dat), "\n\n")
	nums := []int{}
	for _, i := range strings.Split(strs[0], ",") {
		val, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}
		nums = append(nums, val)
	}
	
	boards := []Board{}
	for _, i := range strs[1:] {
		board := Board{}
		for _, j := range strings.Split(i, "\n") {
			row := []int{}
			for _, k := range strings.Split(j, " ") {
				if len(k) == 0 {
					continue
				}
				val, err := strconv.Atoi(k)
				if err != nil {
					return 0, err
				}
				row = append(row, val)
			}
			board = append(board, row)
		}
		boards = append(boards, board)
	}
	
	// Calculate
	winners := map[int]bool{}
	finalScore := -1
	for i, _ := range nums {
		for boardIdx, j := range boards {		
			if j.IsWinner(nums[:i+1]) {
				winners[boardIdx] = true
				if len(winners) == len(boards) {
					finalScore = j.Score(nums[:i+1])
					break
				}
			}
		}
		if finalScore != -1 {
			break
		}
	}
	return finalScore, nil
}
