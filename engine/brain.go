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

type Category struct {
	Label    string   `json:"label"`
	Patterns []string `json:"patterns"`
}

type Sentence struct {
	data        []string
	mainContext string
}

func (entity *Sentence) GetData() ([]string, error) {
	if len(entity.data) != 0 {
		return entity.data, nil
	}
	return nil, fmt.Errorf("No content found")
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

// Find the main context of the known
// []Sentence
func (entity *KnownData) GetMainConext() string {
	return ""
}

type UnknownData struct {
	sentences []Sentence
}

type UniqueElements struct {
	data []string
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
	dataSet     DataSet
	mainContext string
	knownData   KnownData
	unknownData UnknownData
}

// instantiate a new Classifier object
// with a dataset provided
func NewClassifier(dataSet DataSet) *Classifier {
	return &Classifier{dataSet: dataSet}
}

// take the raw text, clean and structure
// to []Sentence
func (uow *Classifier) ProcessInput(rawData string) []Sentence {
	var refinedSentences []Sentence
	rawSentences := strings.Split(strings.Trim(strings.ToLower(rawData), "."), ". ")
	if len(rawSentences) != 0 {
		for i := range len(rawSentences) {
			refinedSentence := Sentence{mainContext: "unknown"}
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
					refinedSentence.data = append(refinedSentence.data, refinedWord)
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
			var frequencyDict map[string]int
			// uow.dataSet.categories[n].label (compare) sentences[i].data
			for _, cat := range uow.dataSet.Categories {
				for _, l_w := range sentences[i].data {
					for _, pattern := cat.Patterns {
						// - Form an array of the patterns from the Category.Patterns, occurred in a Sentence.data
					}
				}
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
func (uow *DataManager) LoadDataset() (DataSet, error) {
	var dataSet DataSet

	binary, err := os.ReadFile(uow.datasetPath)
	if err != nil {
		return DataSet{}, err
	}
	err = json.Unmarshal(binary, &dataSet)
	return dataSet, nil
}

// update an existing dataset
func (uow *DataManager) UpdateDatatset() error {
	return nil
}
