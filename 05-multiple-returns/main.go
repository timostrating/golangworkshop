package main

import "fmt"

func main() {
	a, b := "Go", "Oh I like"
	x, y := swap(a, b)
	fmt.Println(x, y)
}

func swap(a, b string) (c, d string) {
  tmp := b
  b = a
  a = tmp
  return a, b
}

// Implement a function function that takes two strings and returns them swapped
