// This package is an entrypoint
// For app startup and responding
// To client requests

package main

import (
	"classifier/engine"
	"fmt"
	"log"
)

// func printSentences(sentences []engine.Sentence) {
// 	for i := range len(sentences) {
// 		fmt.Println("Context: ", sentences[i].GetMainContext())
// 		fmt.Printf("data: %s\n\n", strings.Join(sentences[i].GetData(), " "))
// 	}
// }

func main() {
	rawText := `
	`
	dataManager := engine.NewDataManager("./categories.json")
	dataset, err := dataManager.LoadDataset()
	if err != nil {
		log.Fatalf("Could not load dataset: %s", err)
	}
	classifier := engine.NewClassifier(dataset)
	sentences := classifier.ProcessInput(rawText)
	classifier.RecognizeSentences(sentences)
	fmt.Println(classifier.GetMainContext())
	unknown := engine.GetUnknownData(classifier)
	known := engine.GetKnownData(classifier)
	for _, s := range unknown {
		fmt.Printf("Sentence: %v\n\n\n", s)
	}
	fmt.Println(known)
	fmt.Println(unknown)

}
