package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected string
	}{
		"Case-1":  {input: "⌘4bc2d5e", expected: "⌘⌘⌘⌘bccddddde"},
		"Case-2":  {input: "a4bc2d5e", expected: "aaaabccddddde"},
		"Case-3":  {input: "abccd", expected: "abccd"},
		"Case-4":  {input: "", expected: ""},
		"Case-5":  {input: "aaa0b", expected: "aab"},
		"Case-6":  {input: `qwe\4\5`, expected: `qwe45`},
		"Case-7":  {input: `qwe\45`, expected: `qwe44444`},
		"Case-8":  {input: `qwe\\5`, expected: `qwe\\\\\`},
		"Case-9":  {input: `qwe\\\3`, expected: `qwe\3`},
		"Case-10": {input: "d\n5abc", expected: "d\n\n\n\n\nabc"},
	}
	t.Parallel()
	for _, tc := range cases {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	cases := map[string]string{
		"Case-1": "3abc",
		"Case-2": "45",
		"Case-3": "aaa10b",
	}

	t.Parallel()
	for _, tc := range cases {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}
