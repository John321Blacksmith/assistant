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
	rawText := `tire, фитн, gym, muscle, water, форма, спорт, вода, мышmuscle, график, run, exercis, schedule, упражн, питат, nutri, здоров, sport, shape, health, rest. income, cowork, friend, routine, factor, рутина`
	// rawText2 := `hour, work, finish, tommorow, график, коллег, рабо, weekday. schedule, salary, босс, будни, зарплата, collegue, collegues, офис, yesterday. director, cрок, office, шеф, завод, chief, weekdaywork, project, deadline, boss, доход`
	// rawText := `The European Central Bank must be prepared for inflation that’s been unleashed by the conflict in the Middle East but is yet to be felt, according to Chief Economist Philip Lane.
	// 		The ECB raised interest rates last week for the first time since 2023, warning that war-driven inflation is widening beyond just energy — a message repeated on Monday by President Christine Lagarde. Investors and economists still expect at least one more quarter-point hike, to 2.5%, with price gains to stay well above the 2% goal for some time.
	// `
	dataManager := engine.NewNewDataManager("./categories.json")
	dataset, err := dataManager.LoadDataset()
	if err != nil {
		log.Fatalf("Could not load dataset: %s", err)
	}
	classifier := engine.NewClassifier(dataset)
	sentences := classifier.ProcessInput(rawText)
	classifier.RecognizeSentences(sentences)
	fmt.Println(classifier.GetMainContext())
}
