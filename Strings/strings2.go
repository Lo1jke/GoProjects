package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))

}

func adding(s1, s2 string) int64 {
	//fmt.Println(frmstr(s1))
	//fmt.Println(frmstr(s2))
	d1, _ := strconv.ParseInt(frmstr(s1), 10, 64)
	d2, _ := strconv.ParseInt(frmstr(s2), 10, 64)
	return d1 + d2
	//return strconv.ParseInt(frmstr(s1), 10, 64) + strconv.ParseInt(frmstr(s2), 10, 64)
}

func frmstr(s string) string {
	a := []rune(s)
	s1 := make([]rune, 0, 0)
	for _, elem := range a {
		if unicode.IsDigit(elem) {
			s1 = append(s1, elem)
		}
	}
	return string(s1)
}
