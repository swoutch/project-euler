package main

import "fmt"

func IsDividable(n int, dividers []int) bool {
	dividerFound := false
	for i:=0; i<len(dividers) && !dividerFound; i++{
		if n%dividers[i] == 0 {
			dividerFound = true
		}
	}
	return dividerFound
}

func LargestPrimeFactor(n int) int {
	primes := []int{2}
	remaining := n
	divider := 2
	for remaining != 1 {
		for remaining%divider == 0 {
			remaining = remaining/divider
		}
		if remaining != 1 {
			for IsDividable(divider, primes) {
				divider++
			}
			primes = append(primes, divider)
		}
	}
	return primes[len(primes)-1]
}

func main() {
	fmt.Println(LargestPrimeFactor(600851475143))
}
