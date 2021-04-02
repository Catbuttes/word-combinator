package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type WordFileContents struct {
	listA []string
	listB []string
}

func usage() {
	fmt.Println("Word-Combinator")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	//http := flag.Bool("http", false, "Run a web server")
	//port := flag.Int("port", 3000, "The port for the server to listen on")
	inFile := flag.String("inFile", "./words.txt", "The input file with the words in it")
	outFile := flag.String("outFile", "./outwords.txt", "The output file the words get written to")
	flag.Parse()

	//_ = http
	//_ = port

	//if !*http {
	runFileHandler(*inFile, *outFile)
	//}

}

func runFileHandler(inFile string, outFile string) {
	words := loadWordFile(inFile)
	outWords := make([]string, 0)

	if inFile == outFile {
		outWords = append(outWords, "\n#")
	}

	for _, wordA := range words.listA {
		for _, wordB := range words.listB {
			outWords = append(outWords, wordA+" "+wordB)
		}
	}

	outString := strings.Join(outWords, "\n")

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

	_, err = fs.WriteString(outString)
	if err != nil {
		fmt.Print(err)
		fmt.Println("")
		os.Exit(-1)
	}

}

func loadWordFile(file string) WordFileContents {
	_ = file
	contents := WordFileContents{
		listA: make([]string, 0),
		listB: make([]string, 0),
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
			contents.listA = append(contents.listA, word)
		} else {
			contents.listB = append(contents.listB, word)
		}
	}

	return contents
}
