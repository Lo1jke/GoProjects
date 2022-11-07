// input
//1 149,6088607594936;1 179,0666666666666

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("%.4f", frm2digits(s))
}

func frm2digits(s string) float64 {
	s1 := strings.Split(s, ";")
	var d = []float64{}
	for _, elem := range s1 {
		var tmp = []rune(elem)
		s11 := make([]rune, 0, 0)
		for _, rn := range tmp {
			if unicode.IsDigit(rn) || rn == ',' {
				if rn == ',' {
					s11 = append(s11, '.')
					continue
				}
				s11 = append(s11, rn)
			}
		}
		tt, _ := strconv.ParseFloat(string(s11), 64)
		d = append(d, tt)
	}
	return d[0] / d[1]
}
