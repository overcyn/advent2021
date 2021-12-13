/*
--- Day 9: Smoke Basin ---

These caves seem to be lava tubes. Parts are even still volcanically active; small hydrothermal vents release smoke into the caves that slowly settles like rain.

If you can model how the smoke flows through the caves, you might be able to avoid it and be that much safer. The submarine generates a heightmap of the floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the following heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678
Each number corresponds to the height of a particular location, where 9 is the highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent locations. Most locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1 and a 0), one is in the third row (a 5), and one is in the bottom row (also a 5). All other locations on the heightmap have some lower adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low points on your heightmap?
*/

package main

import (
	"fmt"
	"os"
	// "math"
	"strings"
	"strconv"
)

func main() {
	rlt, err := a()
	if err != nil {
		panic(err)
	}
	fmt.Println("Result A:", rlt)
	
	// rlt, err = b()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Result B:", rlt)
}

func riskLevel(grid [][]int, x int, y int) int {
	val := grid[x][y]
	if x > 0 && grid[x-1][y] <= val {
		return 0
	}
	if x < len(grid)-1 && grid[x+1][y] <= val {
		return 0
	}
	if y > 0 && grid[x][y-1] <= val {
		return 0
	}
	if y < len(grid[0])-1 && grid[x][y+1] <= val {
		return 0
	}
	return val + 1
}

func a() (int, error) {
	// Read file
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	
	// Parse input
	grid := [][]int{}
	for _, i := range strings.Split(string(dat), "\n") {
		line := []int{}
		for _, j := range strings.Split(i, "") {
			val, err := strconv.Atoi(j)
			if err != nil {
				return 0, err
			}
			line = append(line, val)
		}
		grid = append(grid, line)
	}
	
	// Calculate
	total := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			total += riskLevel(grid, x, y)
		}
	}
    return total, nil
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