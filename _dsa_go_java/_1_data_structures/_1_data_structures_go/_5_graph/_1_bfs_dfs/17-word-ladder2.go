// https://leetcode.com/problems/word-ladder-ii/

package main

import (
	"fmt"
	"sort"
)

type Solution struct {
	mpp map[string]int
	ans [][]string
	b   string
}

// ---------- DFS BACKTRACK ----------
func (s *Solution) dfs(word string, seq *[]string) {

	if word == s.b {
		// reverse copy
		tmp := make([]string, len(*seq))
		for i := range *seq {
			tmp[i] = (*seq)[len(*seq)-1-i]
		}
		s.ans = append(s.ans, tmp)
		return
	}

	steps := s.mpp[word]
	wordBytes := []byte(word)

	for i := 0; i < len(wordBytes); i++ {
		original := wordBytes[i]

		for ch := byte('a'); ch <= byte('z'); ch++ {
			wordBytes[i] = ch
			newWord := string(wordBytes)

			if val, ok := s.mpp[newWord]; ok && val+1 == steps {
				*seq = append(*seq, newWord)
				s.dfs(newWord, seq)
				*seq = (*seq)[:len(*seq)-1]
			}
		}

		wordBytes[i] = original
	}
}

// ---------- MAIN FUNCTION ----------
func (s *Solution) findLadders(beginWord string, endWord string, wordList []string) [][]string {

	s.mpp = make(map[string]int)
	s.ans = [][]string{}
	s.b = beginWord

	st := make(map[string]bool)
	for _, w := range wordList {
		st[w] = true
	}

	queue := []string{beginWord}
	s.mpp[beginWord] = 1
	delete(st, beginWord)

	wordLen := len(beginWord)

	// ---------- BFS ----------
	for len(queue) > 0 {

		word := queue[0]
		queue = queue[1:]

		steps := s.mpp[word]

		if word == endWord {
			break
		}

		wordBytes := []byte(word)

		for i := 0; i < wordLen; i++ {

			original := wordBytes[i]

			for ch := byte('a'); ch <= byte('z'); ch++ {

				wordBytes[i] = ch
				newWord := string(wordBytes)

				if st[newWord] {
					queue = append(queue, newWord)
					delete(st, newWord)
					s.mpp[newWord] = steps + 1
				}
			}

			wordBytes[i] = original
		}
	}

	// ---------- DFS BUILD PATHS ----------
	if _, ok := s.mpp[endWord]; ok {
		seq := []string{endWord}
		s.dfs(endWord, &seq)
	}

	return s.ans
}

func main() {

	wordList := []string{"des", "der", "dfr", "dgt", "dfs"}
	startWord := "der"
	targetWord := "dfs"

	obj := Solution{}
	ans := obj.findLadders(startWord, targetWord, wordList)

	if len(ans) == 0 {
		fmt.Println(-1)
		return
	}

	sort.Slice(ans, func(i, j int) bool {
		a := ""
		b := ""
		for _, w := range ans[i] {
			a += w
		}
		for _, w := range ans[j] {
			b += w
		}
		return a < b
	})

	for _, seq := range ans {
		for _, w := range seq {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}