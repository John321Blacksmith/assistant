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
	// rawText := `Yields on German benchmark 10-year bonds rose seven basis points to 3.07%, their highest in nearly a month and climbing for an eighth straight day in the longest streak since January. French debt suffered even more, also weighed down by political uncertainty following Marine Le Pen’s decision to run in next year’s presidential election. US 10-year yields were two basis points higher at 4.57%.`
	// rawText2 := `tire, фитн, gym, muscle, water, форма, спорт, вода, мышmuscle, график, run, exercis, schedule, упражн, питат, nutri, здоров, sport, shape, health, rest`
	rawText := `In the final, the seeded under the ninth number of Noskova outplayed her compatriot, the tenth racket of the tournament Karolina Mukhova, with a score of 6:2, 5:7, 6:3. The meeting lasted 2 hours 28 minutes. In the second set, Mukhova played five matches. `
	dataManager := engine.NewDataManager("./categories.json")
	dataset, err := dataManager.LoadDataset()
	if err != nil {
		log.Fatalf("Could not load dataset: %s", err)
	}
	classifier := engine.NewClassifier(dataset)
	sentences := classifier.ProcessInput(rawText)
	fmt.Println(sentences)
	classifier.RecognizeSentences(sentences)

}
