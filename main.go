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
	notHttp := flag.Bool("notHttp", false, "Run a web server")
	port := flag.String("address", "127.0.0.1:3000", "The address for the server to listen on")
	inFile := flag.String("inFile", "./words.txt", "The input file with the words in it")
	outFile := flag.String("outFile", "./outwords.txt", "The output file the words get written to")
	flag.Parse()

	if *notHttp {
		filehandler.RunFileHandler(*inFile, *outFile)
	} else {
		httphandler.RunHttpHandler(*port, *inFile, *outFile)
	}

}
