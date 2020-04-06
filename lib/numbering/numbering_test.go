package numbering

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name           string
		inputX         int
		inputY         int
		expectedOutput int
	}{
		{
			name:           "Regular",
			inputX:         3,
			inputY:         4,
			expectedOutput: 7,
		},
		{
			name:           "WithZero",
			inputX:         3,
			inputY:         0,
			expectedOutput: 3,
		},
		{
			name:           "PositiveNegative",
			inputX:         3,
			inputY:         -4,
			expectedOutput: -1,
		},
		{
			name:           "DoubleNegative",
			inputX:         -3,
			inputY:         -4,
			expectedOutput: -7,
		},
		{
			name:           "WithUndefined",
			inputX:         3,
			expectedOutput: 3,
		},
	}

	num := NewNumber()

	for _, test := range cases {
		t.Run(test.name, func(tester *testing.T) {
			output := num.Sum(test.inputX, test.inputY)

			assert.Equal(t, test.expectedOutput, output)
		})
	}
}

func TestMultiply(t *testing.T) {
	cases := []struct {
		name           string
		inputX         int
		inputY         int
		expectedOutput int
	}{
		{
			name:           "Regular",
			inputX:         3,
			inputY:         4,
			expectedOutput: 12,
		},
		{
			name:           "WithZero",
			inputX:         3,
			inputY:         0,
			expectedOutput: 0,
		},
		{
			name:           "PositiveNegative",
			inputX:         3,
			inputY:         -4,
			expectedOutput: -12,
		},
		{
			name:           "DoubleNegative",
			inputX:         -3,
			inputY:         -4,
			expectedOutput: 12,
		},
		{
			name:           "WithUndefined",
			inputX:         3,
			expectedOutput: 0,
		},
	}

	num := NewNumber()

	for _, test := range cases {
		t.Run(test.name, func(tester *testing.T) {
			output := num.Multiply(test.inputX, test.inputY)

			assert.Equal(t, test.expectedOutput, output)
		})
	}
}

func TestGetPrimes(t *testing.T) {
	cases := []struct {
		name           string
		input          int
		expectedOutput []int
	}{
		{
			name:           "FourPrime",
			input:          4,
			expectedOutput: []int{2, 3, 5, 7},
		},
		{
			name:           "TenPrimes",
			input:          10,
			expectedOutput: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
		{
			name:  "TwentyFivePrimes",
			input: 25,
			expectedOutput: []int{
				2,
				3,
				5,
				7,
				11,
				13,
				17,
				19,
				23,
				29,
				31,
				37,
				41,
				43,
				47,
				53,
				59,
				61,
				67,
				71,
				73,
				79,
				83,
				89,
				97,
			},
		},
		{
			name:           "ZeroPrime",
			input:          0,
			expectedOutput: nil,
		},
		{
			name:           "NegativeInput",
			input:          -3,
			expectedOutput: nil,
		},
		{
			name:           "UndefinedInput",
			expectedOutput: nil,
		},
	}

	num := NewNumber()

	for _, test := range cases {
		t.Run(test.name, func(tester *testing.T) {
			output := num.GetPrimes(test.input)

			assert.Equal(t, len(test.expectedOutput), len(output))
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}

func TestGetFibonnacy(t *testing.T) {
	cases := []struct {
		name           string
		input          int
		expectedOutput []int
	}{
		{
			name:           "FourFibs",
			input:          4,
			expectedOutput: []int{0, 1, 1, 2},
		},
		{
			name:           "TenFibss",
			input:          10,
			expectedOutput: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
		{
			name:  "TwentyFiveFibs",
			input: 25,
			expectedOutput: []int{
				0,
				1,
				1,
				2,
				3,
				5,
				8,
				13,
				21,
				34,
				55,
				89,
				144,
				233,
				377,
				610,
				987,
				1597,
				2584,
				4181,
				6765,
				10946,
				17711,
				28657,
				46368,
			},
		},
		{
			name:           "ZeroFibs",
			input:          0,
			expectedOutput: nil,
		},
		{
			name:           "NegativeInput",
			input:          -3,
			expectedOutput: nil,
		},
		{
			name:           "UndefinedInput",
			expectedOutput: nil,
		},
	}

	num := NewNumber()

	for _, test := range cases {
		t.Run(test.name, func(tester *testing.T) {
			output := num.GetFibonnacy(test.input)

			assert.Equal(t, len(test.expectedOutput), len(output))
			assert.Equal(t, test.expectedOutput, output)
		})
	}
}
