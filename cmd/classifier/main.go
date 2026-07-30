// This package is an entrypoint
// For app startup and responding
// To client requests

package main

import (
	"classifier/engine"
	"fmt"
	"strings"
)

// STAGES
// 													| -> findMainContext
// ingestData -> classifyData -> orchestrateData -> |
// 													| -> processUnknownData

// DATA PIPELINE
// 							| -> KnownData
// rawData -> []Sentence -> |
//							| -> UnknownData -> ClusteredData

// PROCESS
//																								 | -> KnownData -> findMainContext()
// rawData -> ingestData() -> []Sentence -> classifyData() -> []Sentence -> orchestrateData() -> |
//																								 | -> UnknownData -> processUnknownData() -> ClusteredData

func refineWord(w string) []string {
	var refinedWord []string
	if len(w) < 3 {
		return nil
	}
	for i := range len(w) {
		if (w[i] >= 65 && w[i] >= 90) || (w[i] >= 97 && w[i] >= 122) {
			refinedWord = append(refinedWord, string(w[i]))
		}
	}
	return refinedWord
}

// processSentence goroutine takes
// each sentence and makes a processing
func processSentence(rawSentence string, refinedSentences chan<- engine.Sentence) {
	refinedSentence := engine.Sentence{MainContext: "", Data: engine.NewUniqueElements()}
	splittedSentence := strings.SplitSeq(rawSentence, " ")

	for w := range splittedSentence {
		refinedWord := refineWord(w)
		refinedSentence.Data.AddElement(strings.Join(refinedWord, ""))
	}
	refinedSentences <- refinedSentence
}

// ingestData takes rawData
// from client and returns
// a channel with []Sentence.
func ingestData(inp string) (<-chan engine.Sentence, <-chan error) {
	out := make(chan engine.Sentence)
	errChan := make(chan error)
	trimmedData := strings.Trim(strings.ToLower(inp), ".")
	rawSentences := strings.Split(trimmedData, ". ")

	defer close(out)
	for i := range len(rawSentences) {
		go processSentence(rawSentences[i], out)
	}

	return out, errChan
}

// classifyData takes a channel
// with []Sentence and classifies
// each Sentence. Returns a channel
// with []ClassifiedSentence.
func classifyData(inp chan<- engine.Sentence) chan<- engine.Sentence {
	out := make(chan engine.Sentence)
	return out
}

// orchestrateData takes a channel
// with ClassifiedSentence, delivers
// the classified ones to both Known/Unknown
// data channels and returns these ones.
// // func orchestrateData(inp chan<- engine.Sentence) (chan<- engine.Sentence, chan<- engine.Sentence)

// findMainContext takes a channel
// with KnownData and finds an overall
// topic of the whole text. Returns
// a string of the main context.
func findMainContext(inp chan<- engine.Sentence) string {
	var mainContext string
	return mainContext
}

// processUnknownData takes a channel
// with UnknownData and clusterizes the
// the literals collection. Returns a
// a probabalistic dataset
func processUnknownData(inp chan<- engine.Sentence) engine.DataSet {
	var dataset engine.DataSet
	return dataset
}

func main() {
	text1 := ` processUnknownData takes a channel
			   with UnknownData and clusterizes the
			   the literals collection. Returns a
			   a probabalistic dataset`
	result, _ := ingestData(text1)
	fmt.Println(<-result)
}
