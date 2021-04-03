package httphandler

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/catbuttes/word-combinator/common"
)

// Embed the entire directory.
var (
	//go:embed templates
	indexHTML embed.FS
)

type UiModel struct {
	ListA  string
	ListB  string
	Output string
}

func RunHttpHandler(port string, inFile string, outFile string) {
	http.HandleFunc("/", Page)
	fmt.Printf("Listening on http://%s\n", port)
	http.ListenAndServe(port, nil)
}

func Page(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		postPage(w, req)
	default:
		getPage(w, req)
	}
}

func getPage(w http.ResponseWriter, req *http.Request) {
	indexTemplate, err := template.ParseFS(indexHTML, "templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	blank := UiModel{
		ListA:  "",
		ListB:  "",
		Output: "",
	}

	indexTemplate.Execute(w, blank)
}

func postPage(w http.ResponseWriter, req *http.Request) {
	indexTemplate, err := template.ParseFS(indexHTML, "templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	req.ParseForm()
	defer req.Body.Close()

	words := processInput(req.Form.Get("A"), req.Form.Get("B"))

	uiData := UiModel{
		ListA:  req.Form.Get("A"),
		ListB:  req.Form.Get("B"),
		Output: common.JoinWords(words),
	}

	indexTemplate.Execute(w, uiData)
}

func processInput(inputA string, inputB string) common.InputWordLists {
	contents := common.InputWordLists{
		ListA: make([]string, 0),
		ListB: make([]string, 0),
	}

	wordsA := strings.Split(strings.ReplaceAll(inputA, "\r\n", "\n"), "\n")
	wordsB := strings.Split(strings.ReplaceAll(inputB, "\r\n", "\n"), "\n")

	contents.ListA = append(contents.ListA, wordsA...)
	contents.ListB = append(contents.ListB, wordsB...)

	return contents
}
