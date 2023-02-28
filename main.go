package main

import "fmt"

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s

}

func main() {
	ints := map[string]int64{
		"first":  19,
		"second": 88,
	}
	floats := map[string]float64{
		"first":  22.34,
		"second": 76.33,
	}
	fmt.Printf("Non Generic Sum: %v and %v \n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v \n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums type inferred: %v and %v \n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

}
