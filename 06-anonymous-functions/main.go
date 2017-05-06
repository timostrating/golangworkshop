package main

import "fmt"

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	squaredItems := applyOneEach(func(a int)int{return a*a }, items...)
	cubedItems := applyOneEach(func(a int)int{return a*a*a }, items...)

	fmt.Println("Items", items)
	fmt.Println("Squares", squaredItems)
	fmt.Println("Cubes", cubedItems)
}

// This will be explained later... TODO maybe remove?
func applyOneEach(op func(int) int, items ...int) []int {
	var results []int
	for _, value := range items {
		results = append(results, op(value))
	}
	return results
}
