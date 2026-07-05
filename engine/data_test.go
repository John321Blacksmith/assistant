package engine

import (
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
