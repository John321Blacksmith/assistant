package engine

import (
	"fmt"
	"log/slog"
	"testing"
)

func TestUniqueElementsIntersection(t *testing.T) {
	set1 := NewUniqueElements()
	set1.AddElement("1", "2", "3", "4")

	set2 := NewUniqueElements()
	set2.AddElement("3", "4", "5", "6", "1")
	t.Run("test intersection done well", func(t *testing.T) {
		slog.Info("TestUniqueElementsIntersection")
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

func TestFreqMapGreatestContext(t *testing.T) {
	testFreqMap := FreqMap{
		data: map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
		},
	}
	want := "five"
	t.Run("testing finding greatest context", func(t *testing.T) {
		slog.Info("TestFreqMapGreatestContext")
		result := testFreqMap.FindGreatestKey()
		if result != want {
			t.Errorf("The function FreqMap.FindGreatestKey() does not retrun a top context: want - %s, returned - %s", want, result)
		}
	})

}

func TestKnownDataStructWork(t *testing.T) {
	knownSentences := []Sentence{
		{data: &UniqueElements{data: map[string]bool{}}, mainContext: ""},
		{data: &UniqueElements{data: map[string]bool{"three": true, "four": true}}, mainContext: "four"},
		{data: &UniqueElements{data: map[string]bool{"one": true, "two": true}}, mainContext: "two"},
		{data: &UniqueElements{data: map[string]bool{"three": true, "four": true}}, mainContext: "four"},
		{data: &UniqueElements{data: map[string]bool{"five": true, "four": true}}, mainContext: "five"},
	}
	knownData := KnownData{sentences: knownSentences}
	want := "four"

	t.Run("testing getting a document-level main context", func(t *testing.T) {
		slog.Info("TestKnownDataStructWork")
		result := knownData.GetMainConext()
		if result != want {
			t.Errorf("The function KownData.GetMainContext() returns a wrong context: want - %s, returned - %s", want, result)
		}
	})

}

func TestUnknownDataStructWork(t *testing.T) {}
