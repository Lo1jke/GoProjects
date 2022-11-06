package main

import "fmt"

func main() {
	reverseIntLoop()
	reverseString()
}

func reverseString() {
	s := "Hello World!"
	for _, elem := range s {
		defer fmt.Println(string(elem))
	}
}

func reverseIntLoop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
