// This package is an entrypoint
// For app startup and responding
// To client requests

package main

import (
	"classifier/engine"
	"fmt"
	"log"
	"strings"
)

func printSentences(sentences []engine.Sentence) {
	for i := range len(sentences) {
		data, err := sentences[i].GetData()
		if err != nil {
			fmt.Printf("Empty sentence: %v", i)
		}
		fmt.Println("Context: ", sentences[i].GetMainContext())
		fmt.Printf("data: %s\n\n", strings.Join(data, " "))
	}
}

func printKnownSentences(classifier *engine.Classifier) {
}

func main() {
	rawText := `As the gigantic space exploration startup barreled toward a
				public listing after 24 years as a private company, the urgency
				was apparent. The acquisition of a $250 billion artificial
				intelligence company in the middle of the process?
				Don’t let it slow you down. Why have everyone wait late
				into the evening for the final terms? Instead,
				close the deal during market hours.
				`

	dataManager := engine.NewNewDataManager("./categories.json")
	dataset, err := dataManager.LoadDataset()
	if err != nil {
		log.Fatalf("Could not load dataset: %s", err)
	}
	classifier := engine.NewClassifier(dataset)
	sentences := classifier.ProcessInput(rawText)
	classifier.RecognizeSentences(sentences)
	printKnownSentences(classifier)
}
