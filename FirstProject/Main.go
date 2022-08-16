package main

import (
	"fmt"
)

func main() {
	grades := []int{1, 2, 3, 4, 5}
	b := make([]int, 3, 100)
	fmt.Println(cap(grades) + len(b))
}
