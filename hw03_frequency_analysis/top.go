package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[[:punct:]]`)

type countWord map[string]int

type WordCounts struct {
	words []string
	count countWord
}

func NewWordCounts() *WordCounts {
	// Конструктор для создания структуры wordCounts
	return &WordCounts{
		words: make([]string, 0),
		count: make(countWord),
	}
}

func (wC *WordCounts) getWord(word string) bool {
	// Проверка наличия слова
	_, ok := wC.count[word]
	return ok
}

func (wC *WordCounts) addWord(word string) {
	// Добавляет слово в слайс и устанавливает кол-во 1
	wC.words = append(wC.words, word)
	wC.count[word] = 1
}

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}
	compileText := re.ReplaceAllString(text, "")
	words := strings.Fields(compileText)

	wordCounter := NewWordCounts()
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		if wordCounter.getWord(lowerWord) {
			wordCounter.count[lowerWord]++
		} else {
			wordCounter.addWord(lowerWord)
		}
	}
	sort.Slice(wordCounter.words, func(i, j int) bool {
		if wordCounter.count[wordCounter.words[i]] == wordCounter.count[wordCounter.words[j]] {
			return wordCounter.words[i] < wordCounter.words[j]
		}
		return wordCounter.count[wordCounter.words[i]] > wordCounter.count[wordCounter.words[j]]
	})

	return wordCounter.words[:10]
}
