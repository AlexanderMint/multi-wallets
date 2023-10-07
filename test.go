package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1,3,2}
	sort.Ints(a)
	fmt.Println(a[(len(a) - 2):len(a)])
}
