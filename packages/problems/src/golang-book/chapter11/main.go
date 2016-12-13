package main

import "fmt"
import "golang-book/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	min := math.Min(xs)
	max := math.Max(xs)
	fmt.Println(avg)
	fmt.Printf("The min is: %f\n", min)
	fmt.Printf("The max is: %f\n", max)
}
