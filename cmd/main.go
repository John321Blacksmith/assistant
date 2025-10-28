package main

import (
	"fmt"
	"osint_agent/internal/usecase"
	// "osint_agent/internal/domain"
	// "osint_agent/pkg/classifier"
)

// func recognizeSentence(sentence []string, categories ) string {
// 	return ""
// }

func main() {
	// input := `Little/ did they/ know the ice cream was invented in Russia.
	// 		This kind/ of food was liked by almost everyone/ in the world.
	// 		  because of it's taste/ and unique, nutrient/ features/.`

	// inputManager := classifier.NewInputManager()
	// rawSentences := inputManager.SplitToParticles(input)
	// refinedSentences := inputManager.RefineSentences(rawSentences)
	// preparedSentences := inputManager.AggregateSentences(refinedSentences)

	// for ind, v := range preparedSentences {
	// 	fmt.Println(ind, v)
	// }

	usecase := usecase.NewDataUseCase()
	dataset, err := usecase.LoadDataset("categories.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataset)

	// for i := range len(preparedSentences) {
	// 	category := recognizeSentence(preparedSentences[i])

	// }
}
