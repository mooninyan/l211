package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:     "Empty input",
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "No anagrams",
			input:    []string{"стол", "стул", "дом"},
			expected: map[string][]string{},
		},
		{
			name:  "Basic anagrams",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:  "Mixed case",
			input: []string{"Пятак", "пятка", "Тяпка"},
			expected: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:  "Multiple anagram groups",
			input: []string{"апельсин", "спаниель", "актёр", "катер", "клоун", "кулон", "уклон"},
			expected: map[string][]string{
				"апельсин": {"апельсин", "спаниель"},
				"актёр":    {"актёр", "катер"},
				"клоун":    {"клоун", "кулон", "уклон"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findAnagrams(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("findAnagrams() = %v, want %v", result, tc.expected)
			}
		})
	}
}
