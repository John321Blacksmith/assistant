package engine

import (
	"log/slog"
	"testing"
)

func BenchmarkUniqueElementsIntersection(b *testing.B) {
	slog.Info("BenchmarkUniqueElementsIntersection")
	set1 := NewUniqueElements()
	set1.AddElement("1", "2", "3", "4")

	set2 := NewUniqueElements()
	set2.AddElement("3", "4", "5", "6", "1")

	for b.Loop() {
		result := set1.Intersection(set2)
		slog.Debug("Intersection()", "card of intersect", result)
	}
}

func BenchmarkFreqMapGreatestContext(b *testing.B) {
	slog.Info("BenchmarkFreqMapGreatestContext")
	testFreqMap := FreqMap{
		data: map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
		},
	}
	for b.Loop() {
		result := testFreqMap.FindGreatestKey()
		slog.Debug("FreqMapGreatestContext()", "greatest pair", result)
	}
}

func BenchmarkKnownDataStructMainContext(b *testing.B) {
	slog.Info("BenchmarkKnownDataStructMainContext")
	knownSentences := []Sentence{
		{data: &UniqueElements{data: map[string]bool{}}, mainContext: ""},
		{data: &UniqueElements{data: map[string]bool{"three": true, "four": true}}, mainContext: "four"},
		{data: &UniqueElements{data: map[string]bool{"one": true, "two": true}}, mainContext: "two"},
		{data: &UniqueElements{data: map[string]bool{"three": true, "four": true}}, mainContext: "four"},
		{data: &UniqueElements{data: map[string]bool{"five": true, "four": true}}, mainContext: "five"},
	}
	knownData := KnownData{sentences: knownSentences}

	for b.Loop() {
		result := knownData.GetMainConext()
		slog.Debug("GetMainConext", "result", result)
	}
}

func BenchmarkUnknownDataStructWork(b *testing.B) {
	slog.Info("BenchmarkUnknownDataStructWork")
}
