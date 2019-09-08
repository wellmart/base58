package base58

import (
	"fmt"
	"math/big"
)

var (
	zero  = big.NewInt(0)
	radix = big.NewInt(58)

	alphabet = []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "A",
		"B", "C", "D", "E", "F", "G", "H", "J", "K", "L",
		"M", "N", "P", "Q", "R", "S", "T", "U", "V", "W",
		"X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z",
	}

	table = [256]byte{
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 0, 1, 2, 3, 4, 5, 6,
		7, 8, 255, 255, 255, 255, 255, 255,
		255, 9, 10, 11, 12, 13, 14, 15,
		16, 255, 17, 18, 19, 20, 21, 255,
		22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 255, 255, 255, 255, 255,
		255, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 255, 44, 45, 46,
		47, 48, 49, 50, 51, 52, 53, 54,
		55, 56, 57, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255, 255, 255,
	}
)

// EncodeToString returns the base64 encoded string of bytes.
func EncodeToString(input []byte) (output string) {
	numbers := new(big.Int).SetBytes(input)

	for numbers.Cmp(zero) > 0 {
		module := new(big.Int)
		numbers.DivMod(numbers, radix, module)

		output = alphabet[module.Int64()] + output
	}

	return
}

// DecodeString returns the bytes by the base64 string.
func DecodeString(input string) (output []byte, e error) {
	var i int

	number := new(big.Int)
	multiplier := big.NewInt(1)
	numbers := big.NewInt(0)

	for i = len(input) - 1; i >= 0; i-- {
		value := table[input[i]]

		if value == 255 {
			e = fmt.Errorf("Invalid character \"%s\", position %d", string(input[i]), i)
			return
		}

		number.SetInt64(int64(value))
		number.Mul(multiplier, number)

		numbers.Add(numbers, number)
		multiplier.Mul(multiplier, radix)
	}

	for i = 0; i < len(input); i++ {
		if input[i] != '1' {
			break
		}
	}

	bytes := numbers.Bytes()
	output = make([]byte, i+len(bytes))

	copy(output[i:], bytes)
	return
}
