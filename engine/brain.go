// This package contains unstructured API
// For working with the raw text and producing
// The analyzis result

package engine

import (
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

// take the raw text, clean and structure
// to []Sentence
func (uow *Classifier) ProcessInput(rawData string) []Sentence {
	var refinedSentences []Sentence
	rawSentences := strings.Split(strings.Trim(strings.ToLower(rawData), "."), ". ")
	if len(rawSentences) != 0 {
		for i := range len(rawSentences) {
			data := NewUniqueElements()
			refinedSentence := Sentence{mainContext: "", data: data}
			rawSentence := strings.Split(rawSentences[i], " ")
			if len(rawSentence) != 0 {
				for _, w := range rawSentence {
					var refinedLiterals []string
					for j := range len(w) {
						if (w[j] >= 65 && w[j] >= 90) || (w[j] >= 97 && w[j] >= 122) {
							refinedLiterals = append(refinedLiterals, string(w[j]))
						}
					}
					refinedWord := strings.Join(refinedLiterals, "")
					refinedSentence.data.AddElement(refinedWord)
				}
			}
			refinedSentences = append(refinedSentences, refinedSentence)
		}
	}
	return refinedSentences

}

// take the []Sentence, analyze each one
// using the predefined dataset and separate
// known from unknown ones
func (uow *Classifier) RecognizeSentences(sentences []Sentence) error {
	if len(sentences) != 0 {
		for i := range len(sentences) {
			data := make(map[string]int)
			freqMap := FreqMap{data}
			objectPatterns := NewUniqueElements()
			for l_w := range sentences[i].data.data {
				for _, cat := range uow.dataSet.Categories {
					for pattern := range cat.Patterns {
						if strings.Contains(l_w, pattern) {
							objectPatterns.AddElement(pattern)
						}
					}
				}
			}
			for _, cat := range uow.dataSet.Categories {
				freqMap.data[cat.Label] = len(objectPatterns.Intersection(cat.Patterns))
			}
			greatestCat := freqMap.FindGreatestKey()

			sentences[i].SetMainContext(greatestCat)

			if greatestCat == "" {
				uow.unknownData.AddSentence(sentences[i])
			} else {
				uow.knownData.AddSentence(sentences[i])
			}
		}
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
