package api

import (
	"strings"
)

// check the symbol
// if it's alphabetic
func IsAlpha(sym byte) bool {
	return (sym >= 65 && sym <= 90) ||
		(sym >= 97 && sym <= 122)
}

// receive the data
// from the external client
// and split to sentences
func SplitToParticles(data, sym string) []string {
	if len(data) != 0 {
		sentences := strings.Split(data, sym)
		return sentences
	}
	return nil
}

// check each sentence member
// if it's a proper word
func RefineMember(member string) string {
	var refinedMember string
	var ind int
	for i := range len(member) {
		letter := member[i]
		if !IsAlpha(letter) {
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
func SplitToMembers(sentence string) []string {
	members := strings.Split(sentence, " ")
	return members
}

func RefineMembers(sentence string) []string {
	var refinedMembers []string
	members := SplitToMembers(strings.TrimSpace(sentence))
	for i := range len(members) {
		refinedMember := RefineMember(members[i])
		refinedMembers = append(refinedMembers, refinedMember)
	}
	return refinedMembers
}

// check each word for
// having a non alphabetic
// symbol
func RefineSentences(sentences []string) [][]string {
	var refinedSentences [][]string
	if len(sentences) != 0 {
		for i := range len(sentences) {
			refinedMembers := RefineMembers(sentences[i])
			refinedSentences = append(refinedSentences, refinedMembers)
		}
	}
	return refinedSentences
}

// format every string
// into the sentence
// object
func AggregateSentences(refinedSentences [][]string) []map[string][]string {
	sentenceCollection := []map[string][]string{}

	for _, v := range refinedSentences {
		sentenceCollection = append(
			sentenceCollection,
			map[string][]string{"unknown": v},
		)
	}
	return sentenceCollection
}
