package base36

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncoder_Encode(t *testing.T) {
	testCases := map[string]struct {
		Input  int64
		Output string
	}{
		"zero":                {Input: 0, Output: "0"},
		"first alpha":         {Input: 10, Output: "a"},
		"last alpha":          {Input: 35, Output: "z"},
		"double digits":       {Input: 36, Output: "10"},
		"big number":          {Input: 1812367, Output: "12ufj"},
		"readme number":       {Input: 1234567890, Output: "kf12oi"},
		"negative number":     {Input: -1812367, Output: "-12ufj"},
		"max number":          {Input: math.MaxInt64, Output: "1y2p0ij32e8e7"},
		"negative max number": {Input: -math.MaxInt64, Output: "-1y2p0ij32e8e7"},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res := StdEncoding.Encode(tc.Input)

			require.Equal(t, tc.Output, res)
			reverse, err := StdEncoding.Decode(res)
			require.NoError(t, err)
			require.Equal(t, tc.Input, reverse)
		})
	}
}

func TestNewEncoder(t *testing.T) {
	encoder := NewEncoder("abcdefghijklmnopqrstuvwxyz0123456789")

	require.Equal(t, "upbcys", encoder.Encode(1234567890))

	value, err := encoder.Decode("upbcys")
	require.NoError(t, err)
	require.Equal(t, int64(1234567890), value)
}

func TestEncoder_Decode_TooLarge(t *testing.T) {
	_, err := StdEncoding.Decode("1y2p0ij32e8e8")
	require.ErrorIs(t, err, ErrNumberTooLarge)

	_, err = StdEncoding.Decode("-1y2p0ij32e8e8")
	require.ErrorIs(t, err, ErrNumberTooLarge)
}

func TestEncoder_Decode_InvalidChars(t *testing.T) {
	_, err := StdEncoding.Decode("abc123!")
	require.ErrorIs(t, err, ErrInvalidCharacter)
}

func TestEncoder_Decode_InputTooLarge(t *testing.T) {
	_, err := StdEncoding.Decode("11y2p0ij32e8e8")
	require.ErrorIs(t, err, ErrNumberTooLarge)
}
