package tests

import (
	"testing"

	"github.com/MrElementary/Pokedex/utils"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}

	cases := []testCase{
		{
			// test to check lowercase
			input:    "Charmander Bulbasaur SQUIRTLE",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			// test to clear trailing whitespace
			input:    "  Cyndaquil Totodile CHIKORITA     ",
			expected: []string{"cyndaquil", "totodile", "chikorita"},
		},
		{
			// test with multiple white space between words. Fails if using Split instead of Fields.
			input:    "Torchic     Mudkip Treecko",
			expected: []string{"torchic", "mudkip", "treecko"},
		},
		{
			// test for no input
			input:    "  ",
			expected: []string{},
		},
	}

	// Compare length of output vs expected arrays
	for _, c := range cases {
		actual := utils.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths not matching = %v, should be %v", actual, c.expected)
		}

		// compare each value in arrays after length matches.
		for i := range actual {
			input_pokename := actual[i]
			exp_pokename := c.expected[i]
			if input_pokename != exp_pokename {
				t.Errorf("cleanInput(%v) == %v, should be %v", c.input, actual, c.expected)
			}
		}
	}
}
