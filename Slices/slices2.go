package main

import "fmt"

func main() {
	var a1, x1, y1 *[]int
	a := []int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	a1, x1, y1 = &a, &x, &y
	fmt.Println(a, x, y, len(a), cap(a), len(x), cap(x), len(y), cap(y), a1, x1, y1)
	x = append(x, y...)
	fmt.Println(a, x, y, len(a), cap(a), len(x), cap(x), len(y), cap(y), a1, x1, y1)
	x = append(x, y...)
	fmt.Println(a, x, y, len(a), cap(a), len(x), cap(x), len(y), cap(y), a1, x1, y1)
}
