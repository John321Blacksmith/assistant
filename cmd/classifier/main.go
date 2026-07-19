// This package is an entrypoint
// For app startup and responding
// To client requests

package main

import (
	"classifier/engine"
	"fmt"
	"log"
)

func ingestData(inp string, classifier *engine.Classifier) <-chan engine.Sentence {
	outProcessed := make(chan engine.Sentence)
	go func() {
		defer close(outProcessed)
		sentences := classifier.ProcessInput(inp)
		for i := range len(sentences) {
			outProcessed <- sentences[i]
		}
	}()
	return outProcessed
}

func orchestrateData(inp <-chan engine.Sentence, classifier *engine.Classifier) (<-chan engine.Sentence, <-chan engine.Sentence) {
	outKnown := make(chan engine.Sentence)
	outUnknown := make(chan engine.Sentence)
	sentences := []engine.Sentence{}

	go func() {
		for s := range inp {
			sentences = append(sentences, s)
		}
	}()
	go func() {
		defer close(outKnown)
		defer close(outUnknown)

		classifier.RecognizeSentences(sentences)

		for _, s := range classifier.Sentences {
			switch s.MainContext {
			case "":
				outUnknown <- s
			default:
				outKnown <- s
			}
		}
	}()
	return outKnown, outUnknown
}

func main() {
	// rawText := `Yields on German benchmark 10-year bonds rose seven basis points to 3.07%, their highest in nearly a month and climbing for an eighth straight day in the longest streak since January. French debt suffered even more, also weighed down by political uncertainty following Marine Le Pen’s decision to run in next year’s presidential election. US 10-year yields were two basis points higher at 4.57%.
	// 			tire, фитн, gym, muscle, water, форма, спорт, вода, мышmuscle, график.run, exercis, schedule, упражн, питат, nutri, здоров, sport, shape, health, rest
	// 			In the final, the seeded under the ninth number of Noskova outplayed her compatriot, the tenth racket of the tournament Karolina Mukhova, with a score of 6:2, 5:7, 6:3. The meeting lasted 2 hours 28 minutes. In the second set, Mukhova played five matches. `

	rawText2 := `
	Министерство иностранных дел России уведомило недружественные страны Запада о неприемлемости размещения войск "коалиции желающих" на территории Украины. Если эти страны отправят солдат на украинские территории - ВС РФ будут рассматривать их как законные цели, препятствующие достижению задач спецоперации.
	`
	dataManager := engine.NewDataManager("./categories.json")
	dataset, err := dataManager.LoadDataset()
	if err != nil {
		log.Fatalf("Could not load dataset: %s", err)
	}
	classifier := engine.NewClassifier(dataset)
	processedData := ingestData(rawText2, classifier)
	knownData, unknownData := orchestrateData(processedData, classifier)

	fmt.Println("recognized:  ", knownData)
	fmt.Println("not recognized:  ", unknownData)

}
