package findwords

import (
	"fmt"
)

type WordList struct {
	theList []string
}

var words WordList

func NewWordList() *WordList {
	words.theList = append(words.theList, "apple", "boy", "cat", "damage")

	return &words
}

func (*WordList) AllWords() []string {
	returnWords := make([]string, len(words.theList))
	copy(returnWords, words.theList)
	fmt.Println(returnWords)
	return returnWords
}
