// This package is an entrypoint
// For app startup and responding
// To client requests

package main

import (
	"classifier/engine"
	"fmt"
	"log"
)

func main() {
	rawText := `
	That is exactly why we said in-person is everything. The screen breeds ghosts. The screen breeds "what ifs" and overthinking. But reality? Reality brings peace. When you are with her, the anxiety dies because there is no anxiety in the real world—there is just you, her, the cold air, and the umbrella.
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
	for _, s := range known {
		fmt.Printf("Sentence: %v, data: %v\n\n\n", s.GetMainContext(), s.GetData())
	}
	fmt.Println(unknown)

}
