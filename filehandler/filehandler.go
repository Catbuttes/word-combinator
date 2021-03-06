package filehandler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/catbuttes/word-combinator/common"
)

func RunFileHandler(inFile string, outFile string) {
	words := loadWordFile(inFile)
	outWords := common.JoinWords(words, "")

	if inFile == outFile {
		outWords = "\n#\n" + outWords
	}

	writeOutput(outWords, outFile)
}

func loadWordFile(file string) common.InputWordLists {
	_ = file
	contents := common.InputWordLists{
		ListA: make([]string, 0),
		ListB: make([]string, 0),
	}

	wordsBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		fmt.Println("")
		os.Exit(-1)
	}

	wordsString := string(wordsBytes)
	words := strings.Split(wordsString, "\n")

	first := true

	for _, word := range words {
		if strings.HasPrefix(string(word), "#") {
			first = false
			continue
		}

		if first {
			contents.ListA = append(contents.ListA, word)
		} else {
			contents.ListB = append(contents.ListB, word)
		}
	}

	return contents
}

func writeOutput(outWords string, outFile string) {
	fs, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fs, err = os.OpenFile(outFile, os.O_RDWR|os.O_APPEND|os.O_EXCL, 0666)
			if err != nil {
				fmt.Print(err)
				fmt.Println("")
				os.Exit(-1)
			}
		} else {
			fmt.Print(err)
			fmt.Println("")
			os.Exit(-1)
		}
	}
	defer fs.Close()

	_, err = fs.WriteString(outWords)
	if err != nil {
		fmt.Print(err)
		fmt.Println("")
		os.Exit(-1)
	}
}
