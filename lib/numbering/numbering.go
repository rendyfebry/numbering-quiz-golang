package numbering

// Number ...
type Number interface {
	Sum(x, y int) int
	Multiply(x, y int) int
	GetPrimes(n int) []int
	GetFibonnacy(n int) []int
}

// NumberSvc ...
type NumberSvc struct{}

// NewNumber ...
func NewNumber() Number {
	return &NumberSvc{}
}

// Sum will return result of x + y
func (*NumberSvc) Sum(x, y int) int {
	return x + y
}

// Multiply will return result of x * y
func (*NumberSvc) Multiply(x, y int) int {
	return x * y
}

func isPrime(num int) bool {
	// Fail fast, 0 & 1 is not prime
	if num < 2 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func fibonacci(number int) int {
	if number < 1 {
		return 0
	}
	if number <= 2 {
		return 1
	}

	return fibonacci(number-1) + fibonacci(number-2)
}

// GetPrimes will return n number of primes
func (*NumberSvc) GetPrimes(n int) []int {
	var primes []int
	var num = 2 // start with 2

	for len(primes) < n {
		if isPrime(num) {
			primes = append(primes, num)
		}

		num++
	}

	return primes
}

// GetFibonnacy will return n number of fibonaccy sequence
func (*NumberSvc) GetFibonnacy(n int) []int {
	var fibs []int

	for i := 0; i < n; i++ {
		fibs = append(fibs, fibonacci(i))
	}

	return fibs
}
