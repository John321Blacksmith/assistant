// This package contains unstructured API
// For working with the raw text and producing
// The analyzis result

package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Product abstractions
type DataSet struct {
	Categories []Category `json:"categories"`
}

func (entity *DataSet) Refine() *RefinedDataSet {
	var categories []RefinedCategory
	for _, cat := range entity.Categories {
		var uniquePatterns map[string]bool = map[string]bool{}
		for _, pattern := range cat.Patterns {
			uniquePatterns[pattern] = true
		}
		categories = append(categories, RefinedCategory{Label: cat.Label, Patterns: uniquePatterns})
	}
	return &RefinedDataSet{Categories: categories}

}

type Category struct {
	Label    string   `json:"label"`
	Patterns []string `json:"patterns"`
}

type RefinedDataSet struct {
	Categories []RefinedCategory
}

type RefinedCategory struct {
	Label    string
	Patterns map[string]bool
}

type Sentence struct {
	data        *UniqueElements
	mainContext string
}

func (entity *Sentence) GetData() map[string]bool {
	return entity.data.data
}

func (entity *Sentence) GetMainContext() string {
	return entity.mainContext
}

func (entity *Sentence) SetMainContext(context string) {
	entity.mainContext = context
}

type KnownData struct {
	sentences []Sentence
}

func (entity *KnownData) AddSentence(sentence Sentence) {
	entity.sentences = append(entity.sentences, sentence)
}

// Find the main context of the known
// []Sentence
func (entity *KnownData) GetMainConext() string {
	var contexts []string
	var mainContext string
	for _, s := range entity.sentences {
		contexts = append(contexts, s.mainContext)
	}
	var freqMap map[string]int = map[string]int{}

	for _, ctx := range contexts {
		if _, exists := freqMap[ctx]; !exists {
			freqMap[ctx] = 1
		} else {
			freqMap[ctx] += 1
		}
	}

	var greatestCxt string
	var greatestFreq int
	for ctx, freq := range freqMap {
		if freq > greatestFreq {
			greatestCxt = ctx
			greatestFreq = freq
		}
	}

	mainContext = greatestCxt

	return mainContext
}

type UnknownData struct {
	sentences []Sentence
}

// Define central Behaviors

// Classification actor behaviour
type TextClassification interface {
	ProcessInput(rawData string) []string
	RecognizeSentences(sentences []Sentence) error
}

// Data management behaviour
type DataManagement interface {
	LoadDataset() (DataSet, error)
	UpdateDataset() error
}

// Define real actors from
// from the Behaviours

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
			refinedSentence := Sentence{mainContext: "unknown", data: data}
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
					refinedSentence.data.AddLiteral(refinedWord)
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
	// KnownData
	// UnknownData
	// TODO:
	//	1. Take a Sentence array and check if it's not empty
	//	2. Iterate through the []Sentence and evaluate Sentence.data
	// 		against all categories from the provided dataset
	//		- Form an array of the patterns from the Category.Patterns, occurred in a Sentence.data
	// 		- Find and intersection between the formed array of patterns and each Category.Patterns
	//  	- Form a frequency map with categories and a cardinality number of the intersection
	//		- Find the category with the highest cardinal
	//		- Assign that category to the Sentence.mainContext
	//	3. Separate []Sentence to known and uknown
	//		- Iterate though the []Sentence
	//			- Put the Sentence to the KnownData.sentences if Sentence.mainContext in not "unknown"
	//			- Othervice, put the one to the UnknownData.sentences
	if len(sentences) != 0 {
		for i := range len(sentences) {
			var freqMap map[string]int = map[string]int{}
			objectPatterns := NewUniqueElements()
			// uow.dataSet.categories[n].label (compare) sentences[i].data
			for l_w := range sentences[i].data.data {
				for _, cat := range uow.dataSet.Categories {
					for pattern := range cat.Patterns {
						// - Populate an array of the patterns from the Category.Patterns, occurred in a Sentence.data
						if strings.Contains(l_w, pattern) {
							objectPatterns.AddLiteral(pattern)
						}
					}
				}
			}
			for _, cat := range uow.dataSet.Categories {
				freqMap[cat.Label] = len(objectPatterns.Intersection(cat.Patterns))
			}
			fmt.Printf("freqMap for the sentence #%d: %v\n\n\n", i, freqMap)
			// Find the category with the highest cardinal
			var greatestCat string
			var greatestCard int
			for cat, cardinal := range freqMap {
				if cardinal > greatestCard {
					greatestCat = cat
					greatestCard = cardinal
				}
			}
			sentences[i].SetMainContext(greatestCat)

			//3. Separate []Sentence to known and uknown
			if sentences[i].GetMainContext() != "unknown" {
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

// DataManagement implementation
type DataManager struct {
	datasetPath string
}

// instantiate a new DataManager object
func NewNewDataManager(datasetPath string) *DataManager {
	return &DataManager{datasetPath: datasetPath}
}

// load the dataset from the storage using
// the provided path
func (uow *DataManager) LoadDataset() (*RefinedDataSet, error) {
	var dataSet DataSet

	binary, err := os.ReadFile(uow.datasetPath)
	if err != nil {
		return &RefinedDataSet{}, err
	}
	err = json.Unmarshal(binary, &dataSet)

	return dataSet.Refine(), nil
}

// update an existing dataset
func (uow *DataManager) UpdateDatatset() error {
	return nil
}
