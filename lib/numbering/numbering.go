package numbering

// Number ...
type Number interface {
	Sum(x, y int) (int, error)
	Multiply(x, y int) (int, error)
	GetPrimes(n int) ([]int, error)
	GetFibonnacy(n int) ([]int, error)
}
