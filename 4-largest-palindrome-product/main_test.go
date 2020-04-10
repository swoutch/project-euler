package main

import "testing"

func TestFindLargestPalindrome(t *testing.T) {
	digits := 1
	actual := FindLargestPalindrome(digits)
	expected := 9
	if actual != expected {
		t.Error("The largest palindrome product of numbers of", digits, "digits is", expected, ", not", actual)
	}

	digits = 2
	actual = FindLargestPalindrome(digits)
	expected = 9009
	if actual != expected {
		t.Error("The largest palindrome product of numbers of", digits, "digits is", expected, ", not", actual)
	}
}

func TestIsPalindrome(t *testing.T) {

	type TestCase struct {
		input int
		expected bool
	}

	testCases := []TestCase{
		{input: 1, expected: true},
		{input: 22, expected: true},
		{input: 909, expected: true},
		{input: 19, expected: false},
		{input: 908, expected: false},
	}

	for _, testCase := range testCases {
		actual := IsPalindrome(testCase.input)
		if actual != testCase.expected {
			t.Errorf("(expected) %t vs %t (actual)", testCase.expected, actual)
		}
	}
}
