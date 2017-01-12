package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum := 0
	num1 := 0
	num2 := 0
	return func() int {
		sum = num1 + num2
		if sum == 0 {
			sum = 1
		}
		num1 = num2
		num2 = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
