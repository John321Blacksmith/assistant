package engine

import (
	"fmt"
	"testing"
)

func TestDatasetRefinement(t *testing.T) {
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

func TestUniqueElementsIntersection(t *testing.T) {
	set1 := NewUniqueElements()
	set1.AddElement("1", "2", "3", "4")

	set2 := NewUniqueElements()
	set2.AddElement("3", "4", "5", "6", "1")
	t.Run("test intersection done well", func(t *testing.T) {
		intersection := set1.Intersection(set2)
		if len(intersection.data) != 0 {
			result, err := intersection.Contains("3", "4")
			if err != nil {
				fmt.Println("The functon Intersection returned an error", err)
			}
			if result == false {
				t.Error("The function Intersection does not return a valid result")
			}
		}
	})
}

func TestFreqDictGreatestContext(t *testing.T) {}

func TestKnownDataStructWork(t *testing.T) {}

func TestUnknownDataStructWork(t *testing.T) {}
