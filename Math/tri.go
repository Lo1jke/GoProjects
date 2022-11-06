//На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
//
//Sample Input:
//
//6 8
//Sample Output:
//
//10

package main

import (
	"fmt"
	. "math"
)

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Print("ошибка")
	}
	c = cLength(a, b)
	fmt.Print(c)
}

func cLength(a, b int) int {
	return int(
		Sqrt(
			Pow(float64(a), 2) + Pow(float64(b), 2)))

}
