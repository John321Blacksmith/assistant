package engine

import (
	"testing"
)

func TestDatasetRefinement(t *testing.T) {
	datasetPath := "../test_dataset.json"
	dataManager := NewDataManager(datasetPath)

	TestDataSet := RefinedDataSet{
		Categories: []RefinedCategory{
			{Label: "test", Patterns: map[string]bool{"sentence": true, "third": true, "this": true, "first": true}},
		},
	}

	want := func(dataSet1 RefinedDataSet, dataSet2 *RefinedDataSet) bool {
		var result bool
		if len(dataSet1.Categories) == len(dataSet2.Categories) {
			for i, cat := range dataSet1.Categories {
				if cat.Label == dataSet2.Categories[i].Label {
					if len(cat.Patterns) == len(dataSet2.Categories[i].Patterns) {
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

func TestUniqueElementsIntersection(t *testing.T) {}

func TestFreqDictGreatestContext(t *testing.T) {}

func TestKnownDataStructWork(t *testing.T) {}

func TestUnknownDataStructWork(t *testing.T) {}
