package utils

// Abs returns the absolute value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GCD returns the greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Sum returns the sum of all integers in the slice
func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Product returns the product of all integers in the slice
func Product(nums []int) int {
	result := 1
	for _, n := range nums {
		result *= n
	}
	return result
}
