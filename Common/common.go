package common

import (
	"fmt"
	"strings"
)

type InputWordLists struct {
	ListA []string
	ListB []string
}

func JoinWords(words InputWordLists) string {
	outWords := make([]string, 0)
	fmt.Print(words)

	for _, wordA := range words.ListA {
		for _, wordB := range words.ListB {
			outWords = append(outWords, wordA+" "+wordB)
		}
	}

	outString := strings.Join(outWords, "\n")

	return outString
}
