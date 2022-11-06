//Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
//
//Входные данные
//
//Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.
//
//Выходные данные
//
//Выведите максимальную цифру, которая встречается во введенной строке.
//
//Sample Input:
//
//1112221112
//Sample Output:
//
//2

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var max int
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("ошибка")
		return
	}
	for _, elem := range s {
		tmp, _ := strconv.Atoi(string(elem))
		switch {
		case max == 0:
			max = tmp
			continue
		case max < tmp:
			max = tmp
		}
	}
	fmt.Print(max)
	fmt.Print(string(getMax(s)))
}

func getMax(s string) int {
	var max int32
	for _, d := range s {
		if d > max {
			max = d
		}
	}
	return int(max)
}
