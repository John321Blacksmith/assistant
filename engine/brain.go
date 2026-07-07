// This package contains unstructured API
// For working with the raw text and producing
// The analyzis result

package engine

import (
	"errors"
	"strings"
)

// Classification actor behaviour
type TextClassification interface {
	ProcessInput(rawData string) []string
	RecognizeSentences(sentences []Sentence) error
	GetMainConext() string
}

// Classification implementation
type Classifier struct {
	dataSet     *RefinedDataSet
	mainContext string
	knownData   KnownData
	unknownData UnknownData
}

// instantiate a new Classifier object
// with a dataset provided
func NewClassifier(dataSet *RefinedDataSet) *Classifier {
	return &Classifier{dataSet: dataSet}
}

func refineWord(w string) []string {
	var refinedWord []string
	if len(w) < 3 {
		return nil
	}
	for i := range len(w) {
		if (w[i] >= 65 && w[i] >= 90) || (w[i] >= 97 && w[i] >= 122) {
			refinedWord = append(refinedWord, string(w[i]))
		}
	}
	return refinedWord
}

// take the raw text, clean and structure
// to []Sentence
func (uow *Classifier) ProcessInput(rawData string) []Sentence {
	var refinedSentences []Sentence
	trimmedData := strings.Trim(strings.ToLower(rawData), ".")
	rawSentences := strings.Split(trimmedData, ". ")
	if len(rawSentences) == 0 {
		return []Sentence{}
	}
	for i := range len(rawSentences) {
		refinedSentence := Sentence{mainContext: "", data: NewUniqueElements()}
		if len(rawSentences[i]) == 0 {
			return []Sentence{}
		}
		rawSentence := strings.Split(rawSentences[i], " ")
		if len(rawSentence) == 0 {
			return []Sentence{}
		}
		for _, w := range rawSentence {
			refinedWord := refineWord(w)
			refinedSentence.data.AddElement(strings.Join(refinedWord, ""))
		}
		refinedSentences = append(refinedSentences, refinedSentence)
	}
	return refinedSentences
}

// take the []Sentence, analyze each one
// using the predefined dataset and separate
// known from unknown ones
func (uow *Classifier) RecognizeSentences(sentences []Sentence) error {
	if len(sentences) == 0 {
		return errors.New("No sentences taken from input")
	}
	for i := range len(sentences) {
		data := make(map[string]int)
		freqMap := FreqMap{data}
		objectPatterns := NewUniqueElements()
		for l_w := range sentences[i].data.data {
			for _, cat := range uow.dataSet.Categories {
				for pattern := range cat.Patterns.data {
					if strings.Contains(l_w, pattern) {
						objectPatterns.AddElement(pattern)
					}
				}
			}
		}
		for _, cat := range uow.dataSet.Categories {
			freqMap.data[cat.Label] = objectPatterns.Intersection(cat.Patterns).Card()
		}
		greatestCat := freqMap.FindGreatestKey()
		if greatestCat == "" {
			uow.unknownData.AddSentence(sentences[i])
		}
		uow.knownData.AddSentence(sentences[i])
	}
	return nil
}

// take the KnownData and get a main context
func (uow *Classifier) GetMainContext() string {
	return uow.knownData.GetMainConext()
}

func GetUnknownData(classifier *Classifier) []Sentence {
	return classifier.unknownData.sentences
}

func GetKnownData(classifier *Classifier) []Sentence {
	return classifier.knownData.sentences
}
