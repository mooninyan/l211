package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	wordsFile, err := os.Open("dictionary")
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err = wordsFile.Close(); err != nil {
			log.Println(err)
		}
	}()

	scanner := bufio.NewScanner(wordsFile)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	anagrams := findAnagrams(words)
	log.Println(anagrams)
}

func findAnagrams(words []string) map[string][]string {
	preDict := make(map[string]*[]string)
	for _, word := range words {
		word = strings.ToLower(word)
		var key string
		key = createKey(word)
		if list, ok := preDict[key]; ok {
			*list = append(*list, word)
		} else {
			preDict[key] = &[]string{word}
		}
	}
	result := make(map[string][]string)
	for _, it := range preDict {
		if len(*it) == 1 {
			continue
		}
		result[(*it)[0]] = *it
	}
	return result
}

func createKey(word string) string {
	runes := []rune(strings.ToLower(word))

	for i, r := range runes {
		if r == 'ё' {
			runes[i] = 'е'
		}
	}

	sort.Slice(runes, func(o1, o2 int) bool {
		return runes[o1] < runes[o2]
	})

	key := string(runes)
	return key
}
