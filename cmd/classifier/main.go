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
	Capital Croissant makes 10,000 pastries a week in its bakery in Ealing, West London, supplying luxury hotels and cafes across the capital with frozen croissants and pains au chocolat. Francois Bonnefoy started the company in November, calling himself “the owner and co-founder, but also the delivery driver, the packer — everything.”
	Tracking the manufacture of Bonnefoy’s viennoiserie gives an insight into the possible rise of food inflation started by the war in the Middle East. The baker sits close to the end of an international supply chain that’s come under enormous pressure over the past few months.
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
	for _, s := range unknown {
		fmt.Printf("Sentence: %v\n\n\n", s.GetData())
	}
	fmt.Println(unknown)

}
