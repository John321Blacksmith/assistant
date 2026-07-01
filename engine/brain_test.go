package engine

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestInputProcess(t *testing.T) {
	rawData := `This is the first sent*ence. This is the sec&ond sentence. Th#is is the third sentence`
	dataManager := NewDataManager("../test_dataset.json")
	dataSet, err := dataManager.LoadDataset()

	if err != nil {
		fmt.Println("Error loading dataset")
	}

	classifier := NewClassifier(dataSet)
	result := classifier.ProcessInput(rawData)

	tests := []struct {
		name string
		want any
	}{
		{name: "sentence objects length", want: 3},
	}
	t.Run(tests[0].name, func(t *testing.T) {
		slog.Info("TestInputProcess")
		if len(result) != tests[0].want {
			t.Errorf("The function ProcessInput returns un enexpected amount of sentences. Returned: %d, want: %d", len(result), tests[0].want)
		}
	})
}

func TestRecognizeSentences(t *testing.T) {
	refinedSentences := []Sentence{
		{data: &UniqueElements{data: map[string]bool{"This": true, "first": true, "sentence": true}}, mainContext: ""},
		{data: &UniqueElements{data: map[string]bool{"This": true, "third": true, "sentence": true}}, mainContext: ""},
		{data: &UniqueElements{data: map[string]bool{"This": true, "second": true, "sentence": true}}, mainContext: ""},
	}
	dataManager := NewDataManager("../test_dataset.json")
	dataSet, err := dataManager.LoadDataset()

	if err != nil {
		fmt.Println("Error loading dataset")
	}

	classifier := NewClassifier(dataSet)
	classifier.RecognizeSentences(refinedSentences)

	tests := []struct {
		name string
		want bool
	}{
		{name: "known data not empty", want: true},
		{name: "unknown data empty", want: true},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		slog.Info("TestRecognizeSentences")
		if (len(classifier.knownData.sentences) > 0) != tests[0].want {
			t.Errorf("The function RecognizeSentences doesn't recognize sentences even though some of them are recognizable")
		}
	})
}

func TestGetMainContext(t *testing.T) {}
