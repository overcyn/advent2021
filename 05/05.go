package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

type line struct {
	ax int
	ay int
	bx int
	by int
}

func parseLine(str string) (line, error) {
	points := strings.Split(str, " -> ")
	if len(points) != 2 {
		return line{}, errors.New("Invalid input")
	}
	
	a := strings.Split(points[0], ",")
	b := strings.Split(points[1], ",")
	if len(a) != 2 || len(b) != 2 {
		return line{}, errors.New("Invalid input")
	}
	ax, err := strconv.Atoi(a[0])
	if err != nil {
		return line{}, errors.New("Invalid input")
	}
	ay, err := strconv.Atoi(a[1])
	if err != nil {
		return line{}, errors.New("Invalid input")
	}
	bx, err := strconv.Atoi(b[0])
	if err != nil {
		return line{}, errors.New("Invalid input")
	}
	by, err := strconv.Atoi(b[1])
	if err != nil {
		return line{}, errors.New("Invalid input")
	}
	return line{ax: ax, ay: ay, bx: bx, by: by}, nil
}

type grid [][]int

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
	lines := []line{}
	for _, i := range strs {
		val, err := parseLine(i)
		if err != nil {
			return 0, err
		}
		lines = append(lines, val)
	}
	
	// Calculate
	maxX := 0
	maxY := 0
	for _, i := range lines {
		if i.ax > maxX {
			maxX = i.ax
		}
		if i.ay > maxY {
			maxY = i.ay
		}
		if i.bx > maxX {
			maxX = i.bx
		}
		if i.by > maxY {
			maxY = i.by
		}
	}
	
	grid := make([][]int, maxX+1)
	for i := 0; i <= maxX; i++ {
		grid[i] = make([]int, maxY+1)
	}
	
	for _, line := range lines {
		if line.ax == line.bx {
			// Handle y axis
			start := line.by
			end := line.ay
			if start > end {
				temp := start
				start = end
				end = temp
			}
			
			
			for i := start; i <= end; i += 1 {
				grid[line.ax][i] += 1
			}
		} else if line.ay == line.by {
			// Handle x axis
			start := line.bx
			end := line.ax
			if start > end {
				temp := start
				start = end
				end = temp
			}
			
			for i := start; i <= end; i += 1 {
				grid[i][line.ay] += 1
			}
		}
	}
	count := 0
	for _, i := range grid {
		for _, j := range i {
			if j >= 2 {
				count += 1
			}
		}
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
	lines := []line{}
	for _, i := range strs {
		val, err := parseLine(i)
		if err != nil {
			return 0, err
		}
		lines = append(lines, val)
	}
	
	// Calculate
	maxX := 0
	maxY := 0
	for _, i := range lines {
		if i.ax > maxX {
			maxX = i.ax
		}
		if i.ay > maxY {
			maxY = i.ay
		}
		if i.bx > maxX {
			maxX = i.bx
		}
		if i.by > maxY {
			maxY = i.by
		}
	}
	
	grid := make([][]int, maxX+1)
	for i := 0; i <= maxX; i++ {
		grid[i] = make([]int, maxY+1)
	}
	
	for _, line := range lines {
		if line.ax == line.bx {
			// Handle y axis
			start := line.by
			end := line.ay
			if start > end {
				temp := start
				start = end
				end = temp
			}
			
			
			for i := start; i <= end; i += 1 {
				grid[line.ax][i] += 1
			}
		} else if line.ay == line.by {
			// Handle x axis
			start := line.bx
			end := line.ax
			if start > end {
				temp := start
				start = end
				end = temp
			}
			
			for i := start; i <= end; i += 1 {
				grid[i][line.ay] += 1
			}
		} else if line.ay - line.by == line.ax - line.bx {
			// Handle positive slope
			startx := line.ax
			endx := line.bx
			starty := line.ay
			endy := line.by
			if startx > endx {
				tempx := startx
				startx = endx
				endx = tempx
				tempy := starty
				starty = endy
				endy = tempy
			}
			
			for i := 0; i <= (endx - startx); i += 1 {
				grid[startx+i][starty+i] += 1
			}
		} else if line.ay - line.by == line.bx - line.ax {
			// Handle negative slope
			startx := line.ax
			endx := line.bx
			starty := line.ay
			endy := line.by
			if startx > endx {
				tempx := startx
				startx = endx
				endx = tempx
				tempy := starty
				starty = endy
				endy = tempy
			}
			
			for i := 0; i <= (endx - startx); i += 1 {
				grid[startx+i][starty-i] += 1
			}
		}
	}
	count := 0
	for _, i := range grid {
		for _, j := range i {
			if j >= 2 {
				count += 1
			}
		}
	}
	return count, nil
}