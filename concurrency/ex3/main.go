package main

import (
	"fmt"
	"math"
	"sync"
)

func BuildMap() map[int]float64 {
	nums := make(map[int]float64, 100_000)
	for key := 0; key < 100_000; key++ {
		nums[key] = math.Sqrt(float64(key))
	}
	return nums
}

var squareRootMapCache = sync.OnceValue(BuildMap)

func main() {
	for i := 0; i < 100_000; i += 1_000 {
		fmt.Println(i, squareRootMapCache()[i])
	}
}
