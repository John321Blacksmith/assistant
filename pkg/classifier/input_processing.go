package classifier

import (
	"osint_agent/internal/domain"
	"strings"
)

// this actor performs
// input processing
// work and prepares
// sentences for further
// classification
type InputManager struct{}

// implement new
// input manager
func NewInputManager() *InputManager {
	return &InputManager{}
}

// check the symbol
// if it's alphabetic
func isAlpha(sym byte) bool {
	return (sym >= 65 && sym <= 90) ||
		(sym >= 97 && sym <= 122)
}

// receive the data
// from the external client
// and split to sentences
func (m *InputManager) SplitToParticles(data string) []string {
	if len(data) != 0 {
		sentences := strings.Split(data, ".")
		return sentences
	}
	return nil
}

// check each sentence member
// if it's a proper word
func refineMember(member string) string {
	var refinedMember string
	var ind int
	for i := range len(member) {
		letter := member[i]
		if !isAlpha(letter) {
			ind = strings.Index(member, string(letter))
		}
	}
	if ind == (len(member) - 1) {
		refinedMember = member[:ind]
	} else {
		refinedMember = member
	}
	return refinedMember
}

// split a sentence to
// members
func splitToMembers(sentence string) []string {
	members := strings.Split(sentence, " ")
	return members
}

func refineMembers(sentence string) []string {
	var refinedMembers []string
	members := splitToMembers(strings.TrimSpace(sentence))
	for i := range len(members) {
		refinedMember := refineMember(members[i])
		refinedMembers = append(refinedMembers, refinedMember)
	}
	return refinedMembers
}

// check each word for
// having a non alphabetic
// symbol
func (m *InputManager) RefineSentences(sentences []string) [][]string {
	var refinedSentences [][]string
	if len(sentences) != 0 {
		for i := range len(sentences) {
			refinedMembers := refineMembers(sentences[i])
			refinedSentences = append(refinedSentences, refinedMembers)
		}
	}
	return refinedSentences
}

// format every string
// into the sentence
// object
func (m *InputManager) AggregateSentences(refinedSentences [][]string) []domain.Sentence {
	sentenceCollection := []domain.Sentence{}

	for _, v := range refinedSentences {
		sentenceCollection = append(
			sentenceCollection,
			domain.Sentence{
				Category:    "unknown",
				LiteralData: v,
			},
		)
	}
	return sentenceCollection
}
