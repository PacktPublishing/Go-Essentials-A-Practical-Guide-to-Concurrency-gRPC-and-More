package main

import "fmt"

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Add[T comparable, V Num](m map[T]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	ints := map[string]int {
		"first": 34,
		"second": 12,
	}

	floats := map[int]float64{
		1: 35.98,
		2: 26.99,
	}

	result := Add(ints)
	fmt.Println(result)

	result2 := Add(floats)
	fmt.Println(result2)

}