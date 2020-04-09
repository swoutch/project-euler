package main

import "fmt"

const FIBONACCI_TERM_1 = 1
const FIBONACCI_TERM_2 = 2
const FIBONACCI_TERM_MAX = 4000000
func main() {
	var previous = FIBONACCI_TERM_1
	var current = FIBONACCI_TERM_2
	var next int
	var sum int

	for current <= FIBONACCI_TERM_MAX {
		if current % 2 == 0 {sum += current}
		next = previous + current
		previous = current
		current = next
	}
	fmt.Println(sum)
}
