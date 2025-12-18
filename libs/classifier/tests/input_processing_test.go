package tests

import (
	"osint_agent/libs/classifier/api"
	"testing"
)

func TestUtilFunctions(t *testing.T) {
	sentence := "Little/ did they/ know the ice cream was invented in Russia."
	normalWord := "word"
	incorrectWord := "wor@d"

	refinedSentence := api.RefineMembers(sentence)
	if len(refinedSentence) == 0 {
		t.Errorf("length of result from api.RefineMembers(sentence) = %v; want >0", len(refinedSentence))
	}

	intactWord := api.RefineMember(normalWord)
	if len(intactWord) == 0 {
		t.Errorf("length of result from api.RefineMember(normalWord) = %v; want >0", len(intactWord))
	}

	refinedWord := api.RefineMember(incorrectWord)
	if len(refinedWord) == len(incorrectWord) {
		t.Errorf("length of result from api.RefineMember(normalWord) = %v; want len(refinedWord) < len(incorrectWord)", len(refinedWord))
	}
}
