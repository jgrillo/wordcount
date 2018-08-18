package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type words struct {
	words []string
}

type counts struct {
	counts map[string]uint32
}

func countWords(wds words) counts {
	cts := make(map[string]uint32)

	for _, word := range wds.words {
		ct, prs := cts[word]

		if prs {
			cts[word] = ct + 1
		} else {
			cts[word] = 1
		}
	}

	return counts{cts}
}

func handleWords(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var wds words
	decodeErr := decoder.Decode(&wds)
	if decodeErr != nil {
		panic(decodeErr) // FIXME
	}

	encoder := json.NewEncoder(writer)

	encodeErr := encoder.Encode(countWords(wds))
	if encodeErr != nil {
		panic(encodeErr) // FIXME
	}
}

func main() {
	mux := http.NewServeMux()

	wordsHandler := http.HandlerFunc(handleWords)
	mux.Handle("/words", handlers.LoggingHandler(os.Stdout, wordsHandler))

	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CompressHandler(mux)))
}
