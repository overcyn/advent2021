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
    strs := strings.Split(string(dat), ",")
    counts := make([]int, 9)
    for _, i := range strs {
        val, err := strconv.Atoi(i)
        if err != nil {
            return 0, err
        }
        counts[val] += 1
    }
    fmt.Println("str:", counts)
    
    // Calculate
    for i := 0; i < 80; i++ {
        zeroCount := counts[0]
        for j, count := range counts {
            if j == 0 {
                continue
            }
            counts[j-1] = count
            counts[j] = 0
        }
        counts[8] += zeroCount
        counts[6] += zeroCount
    }
    
    count := 0 
    for _, i := range counts {
        count += i
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
    strs := strings.Split(string(dat), ",")
    counts := make([]int, 9)
    for _, i := range strs {
        val, err := strconv.Atoi(i)
        if err != nil {
            return 0, err
        }
        counts[val] += 1
    }
    fmt.Println("str:", counts)
    
    // Calculate
    for i := 0; i < 256; i++ {
        zeroCount := counts[0]
        for j, count := range counts {
            if j == 0 {
                continue
            }
            counts[j-1] = count
            counts[j] = 0
        }
        counts[8] += zeroCount
        counts[6] += zeroCount
    }
    
    count := 0 
    for _, i := range counts {
        count += i
    }
    return count, nil
}