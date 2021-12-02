package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

const (
	Up Direction = "up"
	Down = "down"
	Forward = "forward"
)

type Direction string

type Movement struct {
	Direction Direction
	Value int
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
	strs := strings.Split(string(dat), "\n")
	movements := []Movement{}
	for _, i := range strs {
		substrs := strings.Split(i, " ")
		if len(substrs) != 2 {
			return 0, errors.New("Invalid input")
		}
		
		val, err := strconv.Atoi(substrs[1])
		if err != nil {
			return 0, err
		}
		movements = append(movements, Movement{Direction: Direction(substrs[0]), Value: val})
	}
	
	// Calculate
	vertical := 0
	horizontal := 0
	for _, i := range movements {
		switch i.Direction {
		case Up:
			vertical += i.Value
		case Down:
			vertical -= i.Value
		case Forward:
			horizontal += i.Value
		}
	}
	
	return -vertical * horizontal, nil
}

func b() (int, error) {
	// Read file
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	
	// Parse input
	strs := strings.Split(string(dat), "\n")
	movements := []Movement{}
	for _, i := range strs {
		substrs := strings.Split(i, " ")
		if len(substrs) != 2 {
			return 0, errors.New("Invalid input")
		}
		
		val, err := strconv.Atoi(substrs[1])
		if err != nil {
			return 0, err
		}
		movements = append(movements, Movement{Direction: Direction(substrs[0]), Value: val})
	}
	
	// Calculate
	vertical := 0
	aim := 0
	horizontal := 0
	for _, i := range movements {
		switch i.Direction {
		case Up:
			aim += i.Value
		case Down:
			aim -= i.Value
		case Forward:
			horizontal += i.Value
			vertical += i.Value * aim
		}
	}
	
	return -vertical * horizontal, nil
}
