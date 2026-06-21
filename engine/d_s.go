package engine

type UniqueElements struct {
	data map[string]bool
}

func NewUniqueElements() *UniqueElements {
	return &UniqueElements{data: map[string]bool{}}
}

func (ds *UniqueElements) AddElement(items ...string) {
	for _, item := range items {
		ds.data[item] = true
	}
}

func (ds *UniqueElements) Intersection(comparableArr map[string]bool) map[string]bool {
	var intersection map[string]bool = map[string]bool{}

	if len(ds.data) > len(comparableArr) {
		for item := range comparableArr {
			if ds.data[item] {
				intersection[item] = true
			}
		}
	} else {
		for item := range ds.data {
			if comparableArr[item] {
				intersection[item] = true
			}
		}
	}

	return intersection
}

type DataSet struct {
	Categories []Category `json:"categories"`
}

func (ds *DataSet) Refine() *RefinedDataSet {
	var categories []RefinedCategory
	for _, cat := range ds.Categories {
		var uniquePatterns map[string]bool = map[string]bool{}
		for _, pattern := range cat.Patterns {
			uniquePatterns[pattern] = true
		}
		categories = append(
			categories,
			RefinedCategory{
				Label:    cat.Label,
				Patterns: uniquePatterns,
			})
	}
	return &RefinedDataSet{Categories: categories}
}

type RefinedDataSet struct {
	Categories []RefinedCategory
}

type KnownData struct {
	sentences []Sentence
}

func (ds *KnownData) AddSentence(sentence Sentence) {
	ds.sentences = append(ds.sentences, sentence)
}

// Find the main context of the known
// []Sentence
func (ds *KnownData) GetMainConext() string {
	var contexts []string
	for _, s := range ds.sentences {
		contexts = append(contexts, s.mainContext)
	}
	data := make(map[string]int)
	freqMap := FreqMap{data}

	for _, ctx := range contexts {
		if _, exists := freqMap.data[ctx]; !exists {
			freqMap.data[ctx] = 1
		} else {
			freqMap.data[ctx] += 1
		}
	}

	return freqMap.FindGreatestKey()
}

func (ds *KnownData) PresentData() []Sentence {
	return ds.sentences
}

type UnknownData struct {
	sentences []Sentence
}

func (ds *UnknownData) AddSentence(sentence Sentence) {
	ds.sentences = append(ds.sentences, sentence)
}

func (ds *UnknownData) PresentData() []Sentence {
	return ds.sentences
}

type FreqMap struct {
	data map[string]int
}

func (ds FreqMap) FindGreatestKey() string {
	var greatestKey string
	var greatestVal int

	for k, v := range ds.data {
		if v > greatestVal {
			greatestKey = k
			greatestVal = v
		}
	}
	return greatestKey
}
