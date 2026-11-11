package main

import "testing"

func TestRomanToIntValid(t *testing.T) {
	tests := []struct {
		roman string
		want  int
	}{
		{"I", 1},
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"XIV", 14},
		{"XXVII", 27},
		{"XL", 40},
		{"CM", 900},
		{"MCMXCIV", 1994},
		{"MMMCMXCIX", 3999},
		{"xiv", 14},
		{" XIV ", 14},
	}

	for _, test := range tests {
		got, err := romanToInt(test.roman)
		if err != nil {
			t.Errorf("romanToInt(%q): неожиданная ошибка: %v", test.roman, err)
			continue
		}

		if got != test.want {
			t.Errorf("romanToInt(%q) = %d; ожидали %d", test.roman, got, test.want)
		}
	}
}

func TestRomanToIntInvalid(t *testing.T) {
	tests := []string{
		"",
		"IIII",
		"VX",
		"IC",
		"ABC",
		"MMMM",
		"IL",
		"IIV",
	}

	for _, roman := range tests {
		_, err := romanToInt(roman)
		if err == nil {
			t.Errorf("romanToInt(%q): ожидали ошибку, но её нет", roman)
		}
	}
}
