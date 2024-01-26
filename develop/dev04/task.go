package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func findAnagrams(list []string) map[string][]string {
	anagramSets := make(map[string][]string)
	letters := make(map[string]string)
	uniqueWords := map[string]string{}

	for _, word := range list {
		_, ok := uniqueWords[word]
		if ok {
			continue
		}
		uniqueWords[word] = word
		sortedLetters := sortLetters(strings.ToLower(word))
		var key string
		key, ok = letters[sortedLetters]
		if !ok {
			letters[sortedLetters] = word
			key = word
		}
		anagramSets[key] = append(anagramSets[key], word)
	}

	for key, value := range anagramSets {
		if len(value) < 2 {
			delete(anagramSets, key)
		} else {
			sort.Strings(value)
		}
	}

	return anagramSets
}

func sortLetters(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func main() {
	dictionary := []string{"пятак", "тяпка", "пятка", "листок", "пятка", "слиток", "столик"}
	anagrams := findAnagrams(dictionary)

	for key, value := range anagrams {
		fmt.Printf("%s: %v\n", key, value)
	}
}
