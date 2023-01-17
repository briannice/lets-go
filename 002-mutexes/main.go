package main

import (
	"fmt"
	"sync"
)

const (
	MAX = 1000
)

// Package level WaitGroup.
var wg sync.WaitGroup

// Package level Mutex.
var mtx sync.Mutex

// Check if a given integer is a prime number. If it is a prime number, add it
// to the given list of prime numbers.
func isPrime(i int, primes *[]int) {
	defer wg.Done()

	if i < 2 {
		return
	}

	for j := 2; j <= i/2; j++ {
		if i%j == 0 {
			return
		}
	}

	// Use a Mutex to lock the access to the prime number list. This way, no
	// more than one Go routine can access the prime number list. After adding
	// the prime number to the list, the Mutex can be unlocked.
	mtx.Lock()
	*primes = append(*primes, i)
	mtx.Unlock()
}

// Print a list of prime numbers to the console. If there are no prime numbers,
// a messsage will be printed indicating there are no prime numbers.
func printPrimes(primes *[]int) {
	result := "Prime numbers:"

	if len(*primes) == 0 {
		result = fmt.Sprintf("%s no primes found...", result)
	} else {
		for _, prime := range *primes {
			result = fmt.Sprintf("%s %d", result, prime)
		}
	}

	fmt.Println(result)
}

func main() {
	primes := []int{}

	for i := 0; i < MAX; i++ {
		wg.Add(1)
		go isPrime(i, &primes)
	}
	wg.Wait()

	printPrimes(&primes)
}
