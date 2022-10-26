package main

import "fmt"

func main() {
	var n int
	fmt.Println("Последовательность Фибаначчи от числа ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n < 1 {
		fmt.Println("Error:", n, "число слишком мало или отридцательное")
		return
	} else if n > int(^uint(0)>>1) {
		fmt.Println("Error:", n, "Превышена размерность переменной")
		return
	}

	fmt.Println("Последовательность Фибаначчи:", n)
	fmt.Print("0, 1, ")

	p := 0
	q := 1
	for p+q <= n {
		if p < q {
			p += q
			fmt.Printf("%d", p)
		} else {
			q += p
			fmt.Printf("%d", q)
		}

		if p+q <= n {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}
}