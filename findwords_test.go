package findwords

import (
	"testing"
)

// TestSingletonBehavior calls NewWordList, checking that it returns the same, un-mutated object each time.
func TestSingletonBehavior(t *testing.T) {
	wl1 := NewWordList()
	wl2 := NewWordList()
	if wl1 != wl2 {
		t.Fatal(`NewWordList() doesn't return the same object when called multiple times`)
	}
}

// Test WordList.theList is fetched only one time
func TestWordListFetchedOneTime(t *testing.T) {
	wl1 := NewWordList()
	startLen := len(wl1.AllWords())
	wl2 := NewWordList()
	endLen := len(wl2.AllWords())

	if startLen != endLen {
		t.Fatalf(`NewWordList() mutates the length of theList expected %v got %v`, startLen, endLen)
	}

}

func TestStringToMap(t *testing.T) {
	theMap := StringToMap("abcdefghbb")

	if theMap["a"] != 1 {
		t.Fatalf(`wrong count for letter "a"`)
	}

	if theMap["b"] != 3 {
		t.Fatalf(`wrong count for letter "a"`)
	}
}

func TestWordCanBeMadeFromLetters(t *testing.T) {
	if WordCanBeMadeFromLetters("dog", StringToMap("abcdefg")) {
		t.Fatalf(`"dog" not composable with provided letters`)
	}

	if !WordCanBeMadeFromLetters("cab", StringToMap("abcdefg")) {
		t.Fatalf(`"cab" is composable with provided letters`)
	}
}
