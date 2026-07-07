package engine

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestDatasetLoad(t *testing.T) {
	datasetPath := "../test_dataset.json"
	dataManager := NewDataManager(datasetPath)
	patterns := NewUniqueElements()
	patterns.AddElement("sentence", "third", "this", "first")
	TestDataSet := RefinedDataSet{
		Categories: []RefinedCategory{
			{Label: "test", Patterns: patterns},
		},
	}

	want := func(dataSet1 RefinedDataSet, dataSet2 *RefinedDataSet) bool {
		var result bool
		if len(dataSet1.Categories) == len(dataSet2.Categories) {
			for i, cat := range dataSet1.Categories {
				if cat.Label == dataSet2.Categories[i].Label {
					if cat.Patterns.Card() == dataSet2.Categories[i].Patterns.Card() {
						result = true
						continue
					}
				}
			}
		}
		return result
	}

	wantErr := "Data set load error"

	t.Run("test validity of dataset", func(t *testing.T) {
		slog.Info("TestDatasetRefinement")
		refinedDataSet, err := dataManager.LoadDataset()
		if err != nil {
			if wantErr != err.Error() {
				t.Errorf("The function LoadDataSet() returns an unexpected function: want - %s, returned - %s", wantErr, err.Error())
			}
		}
		if want(TestDataSet, refinedDataSet) == false {
			t.Error("The function does not load the categories as expected")
		}
	})

}

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
