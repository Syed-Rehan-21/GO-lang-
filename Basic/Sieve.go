package main

import "fmt"

func main() {
	var n int
	fmt.Print("Till which digit you want the prime numbers to be printed: ")
	fmt.Scan(&n)
	primes := make([]bool, 40)
	prime(n, 2, 4, primes)
	fmt.Print("-->")
	printingPrimes(n, 2, primes)
}
func prime(size, r, c int, primes []bool) {
	if r*r > size {
		return
	}
	if !primes[r] {
		if c <= size {
			primes[c] = true
			prime(size, r, c+r, primes)
		} else {
			prime(size, r+1, (r+1)*2, primes)
		}
	} else {
		prime(size, r+1, (r+1)*2, primes)
	}
}
func printingPrimes(size, r int, primes []bool) {
	if r > size {
		return
	}
	if !primes[r] {
		fmt.Print(r, "  ")
		printingPrimes(size, r+1, primes)
	} else {
		printingPrimes(size, r+1, primes)
	}
}
