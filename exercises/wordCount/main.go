// Implement WordCount. It should return a map of the counts of each “word” in the string s.

package main

import (
	"fmt"
)

func checkWord(w string, m map[string]int) {
	_, ok := m[w]
	if ok {
		m[w] = m[w] + 1
	} else {
		m[w] = 1
	}
}

func WordCount(s string) map[string]int {
	m := map[string]int{}
	prevIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			word := s[prevIndex:i]
			checkWord(word, m)
			prevIndex = i + 1
		}
	}
	word := s[prevIndex:]
	checkWord(word, m)
	fmt.Println(m)
	return m
}

func main() {
	s := "I ate a donut. Then I ate another donut."
	WordCount(s)
}
