package main

import (
	"fmt"
)

func main() {

	s1 := make(map[float32]int)

	s1[3.14] = 10
	s1[5.11] = 0
	s1[5.21] = 100

	fmt.Println(s1)

	var s2 []map[int]int //Декларирование без присваивания. Пустой слайс мар

	var s3 = []map[int]string{} //Декларирование и присваивание если = тогда ставим фигурные скобки

	fmt.Println(len(s2), len(s3))

	fmt.Println(s2)
	s2 = append(s2, map[int]int{10: 10})
	s2 = append(s2, map[int]int{1: 1})
	s2 = append(s2, map[int]int{20: 230})
	s2 = append(s2, map[int]int{10: 3312})
	fmt.Println(s2)

	fmt.Println()
	fmt.Println(s3)

	s3 = append(s3, map[int]string{10: "test"})

	fmt.Println(s3)
}
