package main

import (
	"reflect"
	"testing"
)

// TestFindAnagrams - функция для тестирования findAnagrams.
func TestFindAnagrams(t *testing.T) {
	// testCases содержит набор тестовых сценариев.
	testCases := []struct {
		name     string              // Имя теста для удобства идентификации.
		input    []string            // Входные данные: срез слов.
		expected map[string][]string // Ожидаемый результат: карта анаграмм.
	}{
		{
			name:     "Пустой ввод",
			input:    []string{},
			expected: map[string][]string{},
		},
		{
			name:     "Нет анаграмм",
			input:    []string{"стол", "стул", "дом"},
			expected: map[string][]string{},
		},
		{
			name:  "Базовый случай с анаграммами",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			expected: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:  "Разный регистр букв",
			input: []string{"Пятак", "пятка", "Тяпка"},
			expected: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name:  "Несколько групп анаграмм",
			input: []string{"апельсин", "спаниель", "актёр", "катер", "клоун", "кулон", "уклон"},
			expected: map[string][]string{
				"апельсин": {"апельсин", "спаниель"},
				"актёр":    {"актёр", "катер"},
				"клоун":    {"клоун", "кулон", "уклон"},
			},
		},
	}

	// Проходим по всем тестовым сценариям.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Вызываем тестируемую функцию.
			result := findAnagrams(tc.input)
			// Сравниваем полученный результат с ожидаемым.
			// reflect.DeepEqual используется для глубокого сравнения карт.
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("findAnagrams() = %v, ожидалось %v", result, tc.expected)
			}
		})
	}
}
