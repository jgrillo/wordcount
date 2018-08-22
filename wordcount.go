package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type words struct {
	Words []string `json:"words"`
}

type counts struct {
	Counts map[string]uint32 `json:"counts"`
}

func countWords(wds words) counts {
	cts := counts{Counts: make(map[string]uint32)}
	for _, word := range wds.Words {
		if ct, ok := cts.Counts[word]; ok {
			cts.Counts[word] = ct + 1
			continue
		}
		cts.Counts[word] = 1
	}
	return cts
}

func handleWords(writer http.ResponseWriter, request *http.Request) {
	var wds words
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&wds); err != nil {
		panic(err) // FIXME
	}
	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(countWords(wds)); err != nil {
		panic(err) // FIXME
	}
}

func main() {
	mux := http.NewServeMux()

	wordsHandler := http.HandlerFunc(handleWords)
	mux.Handle("/words", handlers.LoggingHandler(os.Stdout, wordsHandler))

	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CompressHandler(mux)))
}
