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

var numlen = 12

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
	nums := [][]bool{}
	for _, i := range strs {
		substrs := strings.Split(i, "")
		if len(substrs) != numlen {
			return 0, errors.New("Invalid input")
		}
		
		num := []bool{}
		for _, i := range substrs {
			switch i {
			case "0":
				num = append(num, false)
			case "1":
				num = append(num, true)
			}
		}
		
		nums = append(nums, num)
	}
	
	// Calculate
	count := []int{}
	for i := 0; i < numlen; i++ {
		count = append(count, 0)
	}
	
	for _, i := range nums {
		for idx, j := range i {
			if j {
				count[idx] += 1
			}
		}
	}
	
	gammastr := ""
	epsilonstr := ""
	for i := 0; i < numlen; i++ {
		if count[i] < len(nums)/2 {
			gammastr += "0"
			epsilonstr += "1"
		} else {
			gammastr += "1"
			epsilonstr += "0"
		}
	}
	
	gamma, err := strconv.ParseInt(gammastr, 2, 64)
	if err != nil {
		return 0, err
	}
	epsilon, err := strconv.ParseInt(epsilonstr, 2, 64)
	if err != nil {
		return 0, err
	}
		
	return int(gamma * epsilon), nil
}

func b() (int, error) {
	// Read file
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	
	// Parse input
	strs := strings.Split(string(dat), "\n")
	nums := [][]bool{}
	for _, i := range strs {
		substrs := strings.Split(i, "")
		if len(substrs) != numlen {
			return 0, errors.New("Invalid input")
		}
		
		num := []bool{}
		for _, i := range substrs {
			switch i {
			case "0":
				num = append(num, false)
			case "1":
				num = append(num, true)
			}
		}
		
		nums = append(nums, num)
	}
	
	// Calculate
	var oxygen int64
	{
		filter := nums
		for i := 0; i < numlen; i++ {		
			// Get most common value at filter[i]
			count := 0
			for _, j := range filter {
				if j[i] == true {
					count += 1
				}
			}
			var mostcommon bool
			if count == len(filter) {
				mostcommon = true
			} else if count == 0 {
				mostcommon = false
			} else if len(filter) % 2 == 0 && count == len(filter) / 2 {
				mostcommon = true
			} else {
				mostcommon = count > len(filter)/2
			}
			
			// filter out values 
			copy := [][]bool{}
			for _, j := range filter {
				if j[i] == mostcommon {
					copy = append(copy, j)
				}
			}
			filter = copy
		}
		
		str := ""
		for _, i := range filter[0] {
			if !i {
				str += "0"
			} else {
				str += "1"
			}
		}
		
		oxygen, err = strconv.ParseInt(str, 2, 64)
	}
	
	var co2 int64
	{
		filter := nums
		for i := 0; i < numlen; i++ {
			// Get least common value at filter[i]
			count := 0
			for _, j := range filter {
				if j[i] == true {
					count += 1
				}
			}
			
			var leastcommon bool
			if count == len(filter) {
				leastcommon = true
			} else if count == 0 {
				leastcommon = false
			} else if len(filter) % 2 == 0 && count == len(filter) / 2 {
				leastcommon = false
			} else {
				leastcommon = count <= len(filter)/2
			}
			
			// filter out values 
			copy := [][]bool{}
			for _, j := range filter {
				if j[i] == leastcommon {
					copy = append(copy, j)
				}
			}
			filter = copy
		}
		
		str := ""
		for _, i := range filter[0] {
			if !i {
				str += "0"
			} else {
				str += "1"
			}
		}
		
		co2, err = strconv.ParseInt(str, 2, 64)
	}
	
	return int(oxygen*co2), nil
}
