package main

import (
	"fmt"
	"sort"
	"strings"
)

type Word struct {
	Text  string
	Count int
}

func CountWords(text string) map[string]int {
	wordMap := make(map[string]int)

	cleaned := strings.ReplaceAll(text, ".", "")
	cleaned = strings.ReplaceAll(cleaned, ",", "")
	cleaned = strings.ReplaceAll(cleaned, "'", "")
	cleaned = strings.ReplaceAll(cleaned, "`", "")

	words := strings.Fields(cleaned)

	for _, word := range words {
		wordMap[word]++
	}

	return wordMap
}

func Top5Words(wordMap map[string]int) []Word {
	var wordList []Word

	for word, count := range wordMap {
		wordList = append(wordList, Word{Text: word, Count: count})
	}

	sort.Slice(wordList, func(i, j int) bool {
		return wordList[i].Count > wordList[j].Count
	})

	if len(wordList) > 5 {
		return wordList[:5]
	}
	return wordList
}

func main() {
	fmt.Println("Word Frequency Test")

	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

	results := CountWords(text)
	MostCommon := Top5Words(results)

	fmt.Println(MostCommon)
}
