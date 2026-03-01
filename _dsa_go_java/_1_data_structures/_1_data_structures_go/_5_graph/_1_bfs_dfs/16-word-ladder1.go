// https://leetcode.com/problems/word-ladder/

package main

import "fmt"

type Pair struct {
	word  string
	steps int
}

type Solution struct{}

func (s *Solution) wordLadderLength(startWord string, targetWord string, wordList []string) int {

	queue := []Pair{{startWord, 1}}

	// create set
	st := make(map[string]bool)
	for _, w := range wordList {
		st[w] = true
	}

	delete(st, startWord)

	for len(queue) > 0 {

		cur := queue[0]
		queue = queue[1:]

		word := cur.word
		steps := cur.steps

		if word == targetWord {
			return steps
		}

		wordBytes := []byte(word)

		for i := 0; i < len(wordBytes); i++ {

			original := wordBytes[i]

			for ch := byte('a'); ch <= byte('z'); ch++ {

				wordBytes[i] = ch
				newWord := string(wordBytes)

				if st[newWord] {
					delete(st, newWord)
					queue = append(queue, Pair{newWord, steps + 1})
				}
			}

			wordBytes[i] = original
		}
	}

	return 0
}

func main() {

	wordList := []string{"des", "der", "dfr", "dgt", "dfs"}
	startWord := "der"
	targetWord := "dfs"

	obj := Solution{}
	ans := obj.wordLadderLength(startWord, targetWord, wordList)

	fmt.Println(ans)
}