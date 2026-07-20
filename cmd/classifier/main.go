// This package is an entrypoint
// For app startup and responding
// To client requests

package main

import (
	"classifier/engine"
	"fmt"
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

// ingestData takes rawData
// from client and returns
// a channel with []Sentence.
func ingestData(inp string) chan<- engine.Sentence

// classifyData takes a channel
// with []Sentence and classifies
// each Sentence. Returns a channel
// with []ClassifiedSentence.
func classifyData(inp chan<- engine.Sentence) chan<- engine.Sentence

// orchestrateData takes a channel
// with ClassifiedSentence, delivers
// the classified ones to both Known/Unknown
// data channels and returns these ones.
func orchestrateData(inp chan<- engine.Sentence) (chan<- engine.Sentence, chan<- engine.Sentence)

// findMainContext takes a channel
// with KnownData and finds an overall
// topic of the whole text. Returns
// a string of the main context.
func findMainContext(inp chan<- engine.Sentence) string

// processUnknownData takes a channel
// with UnknownData and clusterizes the
// the literals collection. Returns a
// a probabalistic dataset
func processUnknownData(inp chan<- engine.Sentence) engine.DataSet

func main() {
	fmt.Println("hello, world")
}
