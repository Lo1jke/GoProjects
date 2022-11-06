//На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
//
//Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
//
//Sample Input:
//
//9119
//Sample Output:
//
//811181

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	for fmt.Scan(&n); n > 0; n /= 10 {
		defer fmt.Print(int(math.Pow(float64(n%10), 2)))
	}
}

//package main
//
//import (
//	"fmt"
//	"strconv"
//)
//func main(){
//	var s string
//	fmt.Scan(&s)
//	for _,elem := range s{
//		tmp, _ := strconv.Atoi(string(elem))
//		fmt.Print(tmp*tmp)
//	}
//}
