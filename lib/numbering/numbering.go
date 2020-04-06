package numbering

// Number ...
type Number interface {
	Sum(x, y int) int
	Multiply(x, y int) int
	GetPrimes(n int) ([]int, error)
	GetFibonnacy(n int) ([]int, error)
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

// GetPrimes will return n number of primes
func (*NumberSvc) GetPrimes(n int) ([]int, error) {
	return []int{}, nil
}

// GetFibonnacy will return n number of fibonaccy sequence
func (*NumberSvc) GetFibonnacy(n int) ([]int, error) {
	return []int{}, nil
}
