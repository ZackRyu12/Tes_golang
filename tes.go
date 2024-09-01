package main

import (
	"fmt"
	"math"
)

// Fungsi untuk memeriksa apakah suatu bilangan adalah bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fungsi untuk menghasilkan bilangan prima hingga batas tertentu
func generatePrimes(limit int) []int {
	primes := []int{}
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	// Menghasilkan bilangan prima hingga 50
	primes := generatePrimes(50)
	fmt.Println("Bilangan prima hingga 50:", primes)
}
