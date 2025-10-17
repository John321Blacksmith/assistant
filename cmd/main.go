package main

import (
	"fmt"
	"tesla_go/pkg/classifier"
)

func main() {
	input := `Little/ did they/ know the ice cream was invented in Russia.
			  This kind/ of food was liked by almost everyone/ in the world.
			  because of it's taste/ and unique, nutrient/ features/.`

	inputManager := classifier.NewInputManager()
	rawSentences := inputManager.SplitToParticles(input)
	refinedSentences := inputManager.RefineSentences(rawSentences)
	preparedSentences := inputManager.AggregateSentences(refinedSentences)
	for ind, v := range preparedSentences {
		fmt.Println(ind, v)
	}
}
