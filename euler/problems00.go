package euler

import (
	"strconv"
)

// Sum of multiples of 3 and 5 below 1000
func problem001() uint {
	// Golang doesn't have set. Use map with value of type bool and ignore the bool to emulate a set
	multiples3Or5 := make(map[uint]bool)

	for i := uint(0); i < 1000; i++ {
		if i%5 == 0 || i%3 == 0 {
			multiples3Or5[i] = true
		}
	}

	sum := uint(0)
	for multiple, _ := range multiples3Or5 {
		sum += multiple
	}

	return sum
}

// Event fibonacci numbers
func problem002() uint {
	matchedTerms := make(map[uint]bool)

	previous := uint(0)
	current := uint(1)

	sum := uint(0)

	for current <= 4000000 {
		if current%2 == 0 {
			matchedTerms[current] = true
		}

		tmp := current
		current += previous
		previous = tmp
	}

	for value, _ := range matchedTerms {
		sum += value
	}

	return sum
}

// Largest prime factor
func problem003() uint64 {
	panic("Not implemented")
}

// Largest palindrome number
func problem004() int {
	max := 0
	maxSet := false
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			product := i * j
			if IsPalindrome(strconv.Itoa(product)) {
				if product > max {
					maxSet = true
					max = product
				}
			}
		}
	}

	if maxSet {
		return max
	} else {
		panic("Palindrome number doesn't exist for the provided input")
	}
}

// Smallest multiple
func problem005() int {
	i := 20
outer:
	for {
		for j := 2; j < 20; j++ {
			if i%j != 0 {
				i++
				continue outer
			}
		}
		return i
	}
	panic("Should have found a solution by now")
}

// Sum square difference
func problem006() int {
	first, last := 1, 100
	sumOfSquares, sum := 0, 0

	for i := first; i <= last; i++ {
		sum += i
		sumOfSquares += i * i
	}

	squaresOfSum := sum * sum

	return squaresOfSum - sumOfSquares
}

// 10001st prime
func problem007() int {
	num := 1
	primeCount := 0
	for {
		if IsPrime(num) {
			primeCount++
			if primeCount == 10001 {
				return num
			}
		}
		num++
	}
	panic("Couldn't find a solution")
}
