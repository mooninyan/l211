package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// Открываем файл со словарем. Имя файла "dictionary" жестко задано.
	wordsFile, err := os.Open("dictionary.txt")
	if err != nil {
		log.Printf("Ошибка при открытии файла: %v\n", err)
		return
	}
	// Гарантируем закрытие файла при выходе из функции.
	defer func() {
		if err = wordsFile.Close(); err != nil {
			log.Printf("Ошибка при закрытии файла: %v\n", err)
		}
	}()

	// Создаем сканер для чтения файла построчно.
	scanner := bufio.NewScanner(wordsFile)

	// Считываем все слова из файла в срез (slice).
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Ищем анаграммы в списке слов.
	anagrams := findAnagrams(words)
	// Выводим результат в лог.
	log.Println(anagrams)
}

// findAnagrams принимает срез слов и возвращает карту (map),
// где ключ - первое слово из группы анаграмм, а значение - срез всех анаграмм этой группы.
func findAnagrams(words []string) map[string][]string {
	// preDict - временная карта для группировки слов по ключу-анаграмме.
	// Ключ - отсортированные буквы слова, значение - указатель на срез слов.
	preDict := make(map[string]*[]string)
	for _, word := range words {
		// Приводим слово к нижнему регистру для унификации.
		word = strings.ToLower(word)
		// Создаем ключ для слова (отсортированные буквы).
		key := createKey(word)

		// Если ключ уже есть в карте, добавляем слово в существующий срез.
		if list, ok := preDict[key]; ok {
			*list = append(*list, word)
		} else {
			// Иначе создаем новый срез с текущим словом.
			preDict[key] = &[]string{word}
		}
	}

	// result - итоговая карта для анаграмм.
	result := make(map[string][]string)
	for _, it := range preDict {
		// Если в группе только одно слово, это не анаграмма, пропускаем.
		if len(*it) < 2 {
			continue
		}
		// В качестве ключа для результата берем первое слово из группы.
		result[(*it)[0]] = *it
	}
	return result
}

// createKey создает уникальный ключ для слова, который идентифицирует его группу анаграмм.
// Ключ - это строка, состоящая из отсортированных по алфавиту букв слова.
func createKey(word string) string {
	// Преобразуем слово в срез рун для корректной работы с Unicode.
	runes := []rune(strings.ToLower(word))

	// Заменяем букву 'ё' на 'е' для унификации.
	for i, r := range runes {
		if r == 'ё' {
			runes[i] = 'е'
		}
	}

	// Сортируем срез рун в алфавитном порядке.
	sort.Slice(runes, func(o1, o2 int) bool {
		return runes[o1] < runes[o2]
	})

	// Преобразуем отсортированный срез рун обратно в строку.
	key := string(runes)
	return key
}
