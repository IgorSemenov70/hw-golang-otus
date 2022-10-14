package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

const limitWords = 10

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

func (wC *WordCounts) getMostWords() []string {
	// Сортирует и возвращает 10 наиболее частых слов
	sort.Slice(wC.words, func(i, j int) bool {
		if wC.count[wC.words[i]] == wC.count[wC.words[j]] {
			return wC.words[i] < wC.words[j]
		}
		return wC.count[wC.words[i]] > wC.count[wC.words[j]]
	})
	if len(wC.words) > limitWords {
		return wC.words[:10]
	}
	return wC.words
}

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}
	wordCounter := NewWordCounts()
	compileText := re.ReplaceAllString(text, "")

	for _, word := range strings.Fields(compileText) {
		lowerWord := strings.ToLower(word)

		if wordCounter.getWord(lowerWord) {
			wordCounter.count[lowerWord]++
		} else {
			wordCounter.addWord(lowerWord)
		}
	}
	return wordCounter.getMostWords()
}
