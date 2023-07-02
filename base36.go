package base36

import (
	"errors"
	"strings"
)

// ErrNumberTooLarge is returned when the value represented by a string exceeds
// the max int64 value.
var ErrNumberTooLarge = errors.New("value represented by string exceed max int64 value")

// ErrInvalidCharacter is returned when a string contains a character not in the
// encoding's 36 character encoder.
var ErrInvalidCharacter = errors.New("invalid character in string")

var powMap = map[int]int64{
	0:  1,
	1:  36,
	2:  1296,
	3:  46656,
	4:  1679616,
	5:  60466176,
	6:  2176782336,
	7:  78364164096,
	8:  2821109907456,
	9:  101559956668416,
	10: 3656158440062976,
	11: 131621703842267136,
	12: 4738381338321616896,
}

// Encoding is a radix 36 encoding/decoding scheme for int64 values.
type Encoding struct {
	code       []byte
	encoderMap map[byte]int
}

// NewEncoder creates a new base36 encoder using the provided code.
// The code must be 36 characters long and contain no duplicates.
func NewEncoder(encoder string) *Encoding {
	if len(encoder) != 36 {
		panic("encoder must be 36 characters long")
	}

	encoding := &Encoding{
		code:       []byte(encoder),
		encoderMap: make(map[byte]int, 32),
	}

	for i, char := range encoding.code {
		if _, ok := encoding.encoderMap[char]; ok {
			panic("encoder must not contain duplicate characters")
		}

		encoding.encoderMap[char] = i
	}

	return encoding
}

// StdEncoding is the default base36 encoder using the following code: "0123456789abcdefghijklmnopqrstuvwxyz"
var StdEncoding = NewEncoder("0123456789abcdefghijklmnopqrstuvwxyz")

// Encode encodes an int64 value to a base36 string.
// If the input value is negative, the output string will have a '-' prefix.
func (e *Encoding) Encode(input int64) string {
	out := ""
	negative := false

	if input < 0 {
		negative = true
		input *= -1
	}

	if input == 0 {
		return "0"
	}

	for {
		if input == 0 {
			break
		}

		mod := input % 36
		input /= 36

		out = string(e.code[mod]) + out
	}

	if negative {
		return "-" + out
	}

	return out
}

// Decode decodes a base36 string and returns the corresponding int64 value.
// If the input string is too large to decode, it returns an ErrStringTooLarge error.
// If the value represented by the input string exceeds the max int64 value, it returns an ErrNumberTooLarge error.
func (e *Encoding) Decode(input string) (int64, error) {
	negative := false
	input = strings.ToLower(input)

	if input[0] == '-' {
		input = input[1:]
		negative = true
	}

	if len(input) > 13 {
		return 0, ErrNumberTooLarge
	}

	var value int64 = 0

	for n := len(input); n > 0; n-- {
		char := input[len(input)-n]
		index, ok := e.encoderMap[char]
		if !ok {
			return 0, ErrInvalidCharacter
		}

		value += int64(index) * powMap[n-1]
	}

	// If we get a negative number, we've exceeded the max int64 size
	if value < 0 {
		return 0, ErrNumberTooLarge
	}

	if negative {
		return -value, nil
	}

	return value, nil
}
