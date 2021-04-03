package main

import (
	"flag"
	"fmt"

	"github.com/catbuttes/word-combinator/filehandler"
	"github.com/catbuttes/word-combinator/httphandler"
)

func usage() {
	fmt.Println("Word-Combinator")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	http := flag.Bool("http", false, "Run a web server")
	port := flag.Int("port", 3000, "The port for the server to listen on")
	inFile := flag.String("inFile", "./words.txt", "The input file with the words in it")
	outFile := flag.String("outFile", "./outwords.txt", "The output file the words get written to")
	flag.Parse()

	if !*http {
		filehandler.RunFileHandler(*inFile, *outFile)
	} else {
		httphandler.RunHttpHandler(*port, *inFile, *outFile)
	}

}
