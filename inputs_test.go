package main

import "testing"

func TestCleanInput(t *testing.T) {
cases := []struct{
	input string
	expected []string
}{
	{
		input: "  Hello World  ",
		expected: []string{"hello", "world"},
	},
	{
		input: "  PeePee Poo Poo  ",
		expected: []string{"peepee", "poo", "poo"},
	},
}

for _, c := range cases{
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected){
		t.Errorf("Unmatched length. %d != %d",len(actual), len(c.expected))
	}
	for i := range actual{
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord{
			t.Errorf("Unmatched word. %s != %s",word,expectedWord)
		}
	}
}

}