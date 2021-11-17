package findwords

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

type WordList struct {
	theList []string
}

var words WordList

func init() {
	// We need the directory of the package because that's where we keep the word list
	// Yes, this is a bit of a hack. Good enough for our use-cases, though
	_, file, _, _ := runtime.Caller(0)
	thePath := path.Dir(file)

	// Yes, this only works for Unix-like OSes right now
	f, err := os.Open(thePath + "/dict_merged_00.txt")
	if err != nil {
		fmt.Println("fail!")
		panic(`Cant't open dictionary`)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		words.theList = append(words.theList, input.Text())
	}
	f.Close()
}

func NewWordList() *WordList {
	return &words
}

func (*WordList) AllWords() []string {
	returnWords := make([]string, len(words.theList))
	copy(returnWords, words.theList)
	return returnWords
}

func StringToMap(s string) map[string]int {
	theMap := make(map[string]int)

	characters := strings.Split(s, "")

	for _, val := range characters {
		fmt.Println(val)
		theMap[val]++
	}

	return theMap
}

func WordCanBeMadeFromLetters(word string, pool map[string]int) bool {
	workingPool := make(map[string]int, len(pool))

	// we don't want to mutate the argument
	for k, v := range pool {
		workingPool[k] = v
	}

	for _, key := range strings.Split(word, "") {
		val, ok := workingPool[key]
		if ok {
			if val > 0 {
				workingPool[key] = workingPool[key] - 1
			}
		} else {
			return false
		}
	}

	fmt.Println(workingPool)
	return true
}
