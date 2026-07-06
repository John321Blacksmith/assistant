package engine

import (
	"fmt"
	"log/slog"
	"testing"
)

func BenchmarkDatasetLoadPerformance(b *testing.B) {
	datasetPath := "../test_dataset.json"
	dataManager := NewDataManager(datasetPath)

	for b.Loop() {
		refinedDataSet, _ := dataManager.LoadDataset()
		slog.Info(fmt.Sprintf("Lenght of dataset: %d", len(refinedDataSet.Categories)))
	}
}

func BenchmarkInputProcessingPerformance(b *testing.B) {
	var sentences []Sentence
	datasetPath := "../test_dataset.json"
	dataManager := NewDataManager(datasetPath)
	refinedDataSet, err := dataManager.LoadDataset()
	if err != nil {
		b.Error("DataSet not loaded. Aborting...")
	}
	testInput := `This is the first sent*ence. This is the sec&ond sentence. Th#is is the third sentence`
	classifier := NewClassifier(refinedDataSet)
	b.StartTimer()
	for b.Loop() {
		sentences = classifier.ProcessInput(testInput)
	}
	b.StopTimer()
	fmt.Println(len(sentences))
}

func BenchmarkClassificationPerformance(b *testing.B) {
	var err error
	refinedSentences := []Sentence{
		{data: &UniqueElements{data: map[string]bool{"This": true, "first": true, "sentence": true}}, mainContext: ""},
		{data: &UniqueElements{data: map[string]bool{"This": true, "third": true, "sentence": true}}, mainContext: ""},
		{data: &UniqueElements{data: map[string]bool{"This": true, "second": true, "sentence": true}}, mainContext: ""},
	}
	dataManager := NewDataManager("../test_dataset.json")
	dataSet, err := dataManager.LoadDataset()
	if err != nil {
		b.Error("DataSet not loaded. Aborting...")
	}
	classifier := NewClassifier(dataSet)
	b.StartTimer()
	for b.Loop() {
		err = classifier.RefactoredRecognizeSentences(refinedSentences)
	}
	b.StopTimer()
	if err != nil {
		b.Error("RecognizeSentences() -	FAILED")
	}
}
