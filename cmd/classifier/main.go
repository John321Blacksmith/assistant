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
	rawText := `This is the first sent*ence. This is the sec&ond sentence. Th#is is the third sentence`
	dataManager := engine.NewDataManager("./test_dataset.json")
	dataset, err := dataManager.LoadDataset()
	if err != nil {
		log.Fatalf("Could not load dataset: %s", err)
	}
	classifier := engine.NewClassifier(dataset)
	sentences := classifier.ProcessInput(rawText)
	classifier.RecognizeSentences(sentences)
	fmt.Println(classifier.GetMainContext())
	unknown := engine.GetUnknownData(classifier)
	for _, s := range unknown {
		fmt.Printf("Sentence: %v\n\n\n", s.GetData())
	}
	fmt.Println(unknown)

}
